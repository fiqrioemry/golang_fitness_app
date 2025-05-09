package services

import (
	"errors"
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
	HandlePaymentNotification(req dto.MidtransNotificationRequest) error
	GetAllUserPayments(query string, page, limit int) (*dto.AdminPaymentListResponse, error)
	CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error)
}

type paymentService struct {
	paymentRepo     repositories.PaymentRepository
	packageRepo     repositories.PackageRepository
	userPackageRepo repositories.UserPackageRepository
	authRepo        repositories.AuthRepository
}

func NewPaymentService(paymentRepo repositories.PaymentRepository, packageRepo repositories.PackageRepository, userPackageRepo repositories.UserPackageRepository, authRepo repositories.AuthRepository) PaymentService {
	return &paymentService{
		paymentRepo:     paymentRepo,
		packageRepo:     packageRepo,
		userPackageRepo: userPackageRepo,
		authRepo:        authRepo,
	}
}

func (s *paymentService) CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error) {
	pkg, err := s.packageRepo.GetPackageByID(req.PackageID)
	if err != nil {
		return nil, errors.New("package not found")
	}

	user, err := s.authRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	taxRate := utils.GetTaxRate()
	discounted := pkg.Price * (1 - pkg.Discount/100)
	base := discounted
	tax := base * taxRate
	total := base + tax

	paymentID := uuid.New()
	payment := models.Payment{
		ID:            paymentID,
		PackageID:     pkg.ID,
		UserID:        uuid.MustParse(userID),
		PaymentMethod: "-",
		Status:        "pending",
		PaidAt:        time.Now(),
		BasePrice:     base,
		Tax:           tax,
		Total:         total,
	}

	if err := s.paymentRepo.CreatePayment(&payment); err != nil {
		return nil, err
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentID.String(),
			GrossAmt: int64(total),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
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
		return nil
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

	// On success, handle UserPackage logic
	if payment.Status == "success" {
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
			// Kondisi 1: belum punya, buat baru
			expired := now.AddDate(0, 0, pkg.Expired)
			newUP := models.UserPackage{
				ID:              uuid.New(),
				UserID:          payment.UserID,
				PackageID:       payment.PackageID,
				RemainingCredit: pkg.Credit,
				ExpiredAt:       &expired,
				PurchasedAt:     now,
			}
			return s.userPackageRepo.CreateUserPackage(&newUP)
		} else {
			// Kondisi 2: sudah punya dan masih aktif
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

func (s *paymentService) GetAllUserPayments(query string, page, limit int) (*dto.AdminPaymentListResponse, error) {
	offset := (page - 1) * limit
	payments, total, err := s.paymentRepo.GetAllUserPayments(query, limit, offset)
	if err != nil {
		return nil, err
	}

	var results []dto.AdminPaymentResponse
	for _, p := range payments {
		results = append(results, dto.AdminPaymentResponse{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			UserEmail:     p.User.Email,
			Fullname:      p.User.Profile.Fullname,
			PackageID:     p.PackageID.String(),
			PackageName:   p.Package.Name,
			Price:         p.Package.Price,
			PaymentMethod: p.PaymentMethod,
			Status:        p.Status,
			PaidAt:        p.PaidAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &dto.AdminPaymentListResponse{
		Payments: results,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
