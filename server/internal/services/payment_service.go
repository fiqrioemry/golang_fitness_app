package services

import (
	"errors"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error)
	HandlePaymentNotification(req dto.MidtransNotificationRequest) error
}

type paymentService struct {
	paymentRepo     repositories.PaymentRepository
	packageRepo     repositories.PackageRepository
	userPackageRepo repositories.UserPackageRepository
}

func NewPaymentService(paymentRepo repositories.PaymentRepository, packageRepo repositories.PackageRepository, userPackageRepo repositories.UserPackageRepository) PaymentService {
	return &paymentService{
		paymentRepo:     paymentRepo,
		packageRepo:     packageRepo,
		userPackageRepo: userPackageRepo,
	}
}

func (s *paymentService) CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error) {
	// Fetch package info
	pkg, err := s.packageRepo.GetPackageByID(req.PackageID)
	if err != nil {
		return nil, errors.New("package not found")
	}

	// Generate Payment ID
	paymentID := uuid.New()

	// Create Payment record
	payment := models.Payment{
		ID:            paymentID,
		PackageID:     pkg.ID,
		UserID:        uuid.MustParse(userID),
		PaymentMethod: "midtrans",
		Status:        "pending",
		PaidAt:        time.Now(),
	}

	if err := s.paymentRepo.CreatePayment(&payment); err != nil {
		return nil, err
	}

	// Create Midtrans Snap request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentID.String(),
			GrossAmt: int64(pkg.Price),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: "customer.email@example.com",
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

	if req.TransactionStatus == "settlement" || (req.TransactionStatus == "capture" && req.FraudStatus == "accept") {
		payment.Status = "success"
	} else if req.TransactionStatus == "pending" {
		payment.Status = "pending"
	} else {
		payment.Status = "failed"
	}

	if err := s.paymentRepo.UpdatePayment(payment); err != nil {
		return err
	}

	if payment.Status == "success" {
		pkg, err := s.packageRepo.GetPackageByID(payment.PackageID.String())
		if err != nil {
			return err
		}

		expired := time.Now().AddDate(0, 0, *pkg.Expired)

		userPackage := models.UserPackage{
			ID:              uuid.New(),
			UserID:          payment.UserID,
			PackageID:       payment.PackageID,
			RemainingCredit: pkg.Credit,
			ExpiredAt:       &expired,
			PurchasedAt:     time.Now(),
		}

		if err := s.userPackageRepo.CreateUserPackage(&userPackage); err != nil {
			return err
		}
	}

	return nil
}
