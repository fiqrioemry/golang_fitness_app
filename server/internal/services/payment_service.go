package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"gorm.io/gorm"
)

type PaymentService interface {
	ExpireOldPendingPayments() error
	StripeWebhookNotification(event stripe.Event) error
	GetAllUserPayments(params dto.PaymentQueryParam) ([]dto.PaymentListResponse, *dto.PaginationResponse, error)
	CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error)
}
type paymentService struct {
	paymentRepo         repositories.PaymentRepository
	packageRepo         repositories.PackageRepository
	userPackageRepo     repositories.UserPackageRepository
	authRepo            repositories.AuthRepository
	voucherService      VoucherService
	notificationService NotificationService
}

func NewPaymentService(
	paymentRepo repositories.PaymentRepository,
	packageRepo repositories.PackageRepository,
	userPackageRepo repositories.UserPackageRepository,
	authRepo repositories.AuthRepository,
	voucherService VoucherService,
	notificationService NotificationService,
) PaymentService {
	return &paymentService{
		paymentRepo:         paymentRepo,
		packageRepo:         packageRepo,
		userPackageRepo:     userPackageRepo,
		authRepo:            authRepo,
		voucherService:      voucherService,
		notificationService: notificationService,
	}
}

func buildLineItem(name string, unitAmount float64, quantity int64) *stripe.CheckoutSessionLineItemParams {
	if unitAmount < 0 {
		unitAmount = 0
	}
	return &stripe.CheckoutSessionLineItemParams{
		PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String("idr"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name: stripe.String(name),
			},
			UnitAmount: stripe.Int64(int64(unitAmount * 100)),
		},
		Quantity: stripe.Int64(quantity),
	}
}

func (s *paymentService) CreatePayment(userID string, req dto.CreatePaymentRequest) (*dto.CreatePaymentResponse, error) {
	uid := uuid.MustParse(userID)

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
	invoice := utils.GenerateInvoiceNumber(paymentID)

	// Stripe Line Items
	lineItems := []*stripe.CheckoutSessionLineItemParams{
		buildLineItem(pkg.Name, base, 1),
	}
	if tax > 0 {
		lineItems = append(lineItems, buildLineItem("Tax (PPN) 10% ", tax, 1))
	}

	var successURL, cancelURL string
	if os.Getenv("NODE_ENV") == "production" {
		successURL = os.Getenv("STRIPE_SUCCESS_URL_PROD")
		cancelURL = os.Getenv("STRIPE_CANCEL_URL_PROD")
	} else {
		successURL = os.Getenv("STRIPE_SUCCESS_URL_DEV")
		cancelURL = os.Getenv("STRIPE_CANCEL_URL_DEV")
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String(successURL),
		CancelURL:          stripe.String(cancelURL),
		ClientReferenceID:  stripe.String(paymentID.String()),
		Metadata: map[string]string{
			"order_id":   paymentID.String(),
			"user_id":    userID,
			"package_id": pkg.ID.String(),
		},
	}

	// 🔁 Buat Stripe session sebelum menyimpan ke DB
	sess, err := session.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create stripe session: %w", err)
	}

	payment := models.Payment{
		ID:              paymentID,
		PackageID:       pkg.ID,
		PackageName:     pkg.Name,
		InvoiceNumber:   invoice,
		Fullname:        user.Profile.Fullname,
		Email:           user.Email,
		PaymentLink:     sess.URL,
		UserID:          uid,
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

	if payment.VoucherCode != nil && *payment.VoucherCode != "" {
		if err := s.voucherService.DecreaseQuota(payment.UserID, *payment.VoucherCode); err != nil {
			return nil, fmt.Errorf("failed to decrease voucher quota: %w", err)
		}
	}

	payload := dto.NotificationEvent{
		UserID:  user.ID.String(),
		Title:   "Pending Payments",
		Type:    "system_message",
		Message: fmt.Sprintf("Thank you %s, your purchasement with invoice no. %s is created. Please complete your payment", user.Profile.Fullname, invoice),
	}

	if err := s.notificationService.SendToUser(payload); err != nil {
		log.Printf("failed sending notification to user %s: %v\n", payload.UserID, err)
	}

	return &dto.CreatePaymentResponse{
		PaymentID: paymentID.String(),
		SnapURL:   sess.URL,
		SessionID: sess.ID,
	}, nil
}

func (s *paymentService) StripeWebhookNotification(event stripe.Event) error {
	if event.Type != "checkout.session.completed" {
		return fmt.Errorf("%s is not a valid event", event.Type)
	}

	var session stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
		return fmt.Errorf("invalid session data")
	}

	orderID, ok := session.Metadata["order_id"]
	if !ok || orderID == "" {
		return fmt.Errorf("missing order_id in Stripe metadata")
	}

	payment, err := s.paymentRepo.GetPaymentByOrderID(orderID)
	if err != nil {
		return fmt.Errorf("payment not found: %w", err)
	}

	if payment.Status == "success" {
		return nil
	}
	payment.PaymentMethod = "card"
	payment.Status = "success"
	payment.PaidAt = time.Now().UTC()

	if err := s.paymentRepo.UpdatePayment(payment); err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	// TODO: Use RabbitMQ to emit "payment_success" event for async email delivery (only in production with EDA)
	payload := dto.NotificationEvent{
		UserID: payment.UserID.String(),
		Type:   "system_message",
		Title:  "Payment Successful & Package Activated",
		Message: fmt.Sprintf(
			"Hi %s, your payment for %q (Invoice: %s) was successful. Your package has been activated and is now ready to use.",
			payment.Fullname,
			payment.PackageName,
			payment.InvoiceNumber,
		),
	}

	if err := s.notificationService.SendToUser(payload); err != nil {
		log.Printf("failed sending notification to user %s: %v\n", payload.UserID, err)
	}
	// TODO: Use RabbitMQ to emit "payment_success" event for async email delivery (only in production with EDA)

	pkg, err := s.packageRepo.GetPackageByID(payment.PackageID.String())
	if err != nil {
		return fmt.Errorf("package not found: %w", err)
	}

	var existing models.UserPackage
	now := time.Now().UTC()

	err = s.userPackageRepo.GetActiveUserPackages(payment.UserID.String(), payment.PackageID.String(), &existing)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed checking existing user package: %w", err)
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
	}

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
			InvoiceNumber: p.InvoiceNumber,
			Email:         p.Email,
			Fullname:      p.Fullname,
			PackageID:     p.PackageID.String(),
			PackageName:   p.PackageName,
			Total:         p.Total,
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
