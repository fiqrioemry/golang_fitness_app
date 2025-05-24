package services

import (
	"errors"
	"fmt"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type PaymentService interface {
	ExpireOldPendingPayments() error
	HandlePaymentNotification(req dto.MidtransNotificationRequest) error
	GetAllUserPayments(params dto.PaymentQueryParam) ([]dto.PaymentListResponse, *dto.PaginationResponse, error)
	CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error)
}
type paymentService struct {
	paymentRepo     repositories.PaymentRepository
	packageRepo     repositories.PackageRepository
	userPackageRepo repositories.UserPackageRepository
	authRepo        repositories.AuthRepository
	voucherService  VoucherService
}

func NewPaymentService(
	paymentRepo repositories.PaymentRepository,
	packageRepo repositories.PackageRepository,
	userPackageRepo repositories.UserPackageRepository,
	authRepo repositories.AuthRepository,
	voucherService VoucherService,
) PaymentService {
	return &paymentService{
		paymentRepo:     paymentRepo,
		packageRepo:     packageRepo,
		userPackageRepo: userPackageRepo,
		authRepo:        authRepo,
		voucherService:  voucherService,
	}
}

func (s *paymentService) CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error) {
	pkg, err := s.packageRepo.GetPackageByID(req.PackageID)
	if err != nil {
		return nil, fmt.Errorf("package not found: %w", err)
	}

	user, err := s.authRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	taxRate := utils.GetTaxRate()
	discounted := pkg.Price * (1 - pkg.Discount/100)
	base := discounted
	var voucherCode *string
	var voucherDiscount float64

	if req.VoucherCode != nil {
		apply, err := s.voucherService.ApplyVoucher(dto.ApplyVoucherRequest{
			Code:  *req.VoucherCode,
			Total: base,
		})
		if err == nil {
			base = apply.FinalTotal
			voucherCode = &apply.Code
			voucherDiscount = apply.DiscountValue
		}
	}

	tax := base * taxRate
	total := base + tax
	paymentID := uuid.New()

	payment := models.Payment{
		ID:              paymentID,
		PackageID:       pkg.ID,
		PackageName:     pkg.Name,
		UserID:          uuid.MustParse(userID),
		PaymentMethod:   "-",
		Status:          "pending",
		BasePrice:       base,
		Tax:             tax,
		Total:           total,
		VoucherCode:     voucherCode,
		VoucherDiscount: voucherDiscount,
	}

	if err := s.paymentRepo.CreatePayment(&payment); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentID.String(),
			GrossAmt: int64(total),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Profile.Fullname,
			Email: user.Email,
			Phone: user.Profile.Phone,
		},
	}

	snapResp, err := config.SnapClient.CreateTransaction(snapReq)

	return &dto.CreatePaymentResponse{
		PaymentID: paymentID.String(),
		SnapToken: snapResp.Token,
		SnapURL:   snapResp.RedirectURL,
	}, nil
}

func (s *paymentService) HandlePaymentNotification(req dto.MidtransNotificationRequest) error {
	payment, err := s.paymentRepo.GetPaymentByOrderID(req.OrderID)
	if err != nil {
		return err
	}

	if payment.Status == "success" {
		payment.PaidAt = time.Now()
	}

	payment.PaymentMethod = req.PaymentType

	switch req.TransactionStatus {
	case "settlement", "capture":
		if req.FraudStatus == "accept" || req.FraudStatus == "" {
			payment.Status = "success"
		}
	case "pending":
		payment.Status = "pending"
	default:
		payment.Status = "failed"
	}

	if err := s.paymentRepo.UpdatePayment(payment); err != nil {
		return err
	}

	if payment.Status == "success" {

		if payment.VoucherCode != nil && *payment.VoucherCode != "" {
			if err := s.voucherService.DecreaseQuota(payment.UserID, *payment.VoucherCode); err != nil {
				return err
			}
		}

		pkg, err := s.packageRepo.GetPackageByID(payment.PackageID.String())
		if err != nil {
			return err
		}

		var existing models.UserPackage
		now := time.Now()
		err = s.userPackageRepo.
			FindActiveByUserAndPackage(payment.UserID.String(), payment.PackageID.String(), &existing)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if existing.ID == uuid.Nil {
			expired := now.AddDate(0, 0, pkg.Expired)
			newUP := models.UserPackage{
				ID:              uuid.New(),
				UserID:          payment.UserID,
				PackageID:       payment.PackageID,
				PackageName:     payment.PackageName,
				RemainingCredit: pkg.Credit,
				ExpiredAt:       &expired,
				PurchasedAt:     now,
			}
			return s.userPackageRepo.CreateUserPackage(&newUP)
		} else {
			existing.RemainingCredit += pkg.Credit
			if existing.ExpiredAt != nil {
				*existing.ExpiredAt = existing.ExpiredAt.AddDate(0, 0, pkg.Expired)
			} else {
				exp := now.AddDate(0, 0, pkg.Expired)
				existing.ExpiredAt = &exp
			}
			existing.PurchasedAt = now
			return s.userPackageRepo.UpdateUserPackage(&existing)
		}
	}

	return nil
}

func (s *paymentService) GetAllUserPayments(params dto.PaymentQueryParam) ([]dto.PaymentListResponse, *dto.PaginationResponse, error) {

	payments, total, err := s.paymentRepo.GetAllUserPayments(params)
	if err != nil {
		return nil, nil, err
	}

	var results []dto.PaymentListResponse
	for _, p := range payments {
		results = append(results, dto.PaymentListResponse{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			UserEmail:     p.User.Email,
			Fullname:      p.User.Profile.Fullname,
			PackageID:     p.PackageID.String(),
			PackageName:   p.PackageName,
			Price:         p.Total,
			PaymentMethod: p.PaymentMethod,
			Status:        p.Status,
			PaidAt:        p.PaidAt.Format("2006-01-02"),
		})
	}
	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))

	pagination := &dto.PaginationResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}
	return results, pagination, nil
}

// ** khusus cron job update status to failed
func (s *paymentService) ExpireOldPendingPayments() error {
	rows, err := s.paymentRepo.ExpireOldPendingPayments()
	if err != nil {
		return fmt.Errorf("failed to expire pending payments: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("no expired pending payments found")
	}
	fmt.Printf("✅ %d pending payments marked as failed\n", rows)
	return nil
}
