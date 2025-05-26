package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	ExpireOldPendingPayments() (int64, error)
	CreatePayment(payment *models.Payment) error
	UpdatePayment(payment *models.Payment) error
	GetPaymentByID(id string) (*models.Payment, error)
	GetPaymentByOrderID(orderID string) (*models.Payment, error)
	GetAllUserPayments(params dto.PaymentQueryParam) ([]models.Payment, int64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetPaymentByID(id string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetPaymentByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, "id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) GetAllUserPayments(params dto.PaymentQueryParam) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	db := r.db.Model(&models.Payment{})
	// search
	if params.Search != "" {
		likeQuery := "%" + params.Search + "%"
		db = db.Where("email LIKE ? OR fullname LIKE ?", likeQuery, likeQuery)
	}

	// status filter
	if params.Status != "" && params.Status != "all" {
		db = db.Where("payments.status = ?", params.Status)
	}

	// date range filter
	if params.StartDate != "" && params.EndDate != "" {
		db = db.Where("paid_at BETWEEN ? AND ?", params.StartDate, params.EndDate)
	}

	// sorting
	switch params.Sort {
	case "paid_at_asc":
		db = db.Order("paid_at asc")
	case "paid_at_desc":
		db = db.Order("paid_at desc")
	case "name_asc":
		db = db.Order("fullname asc")
	case "name_desc":
		db = db.Order("fullname desc")
	case "email_asc":
		db = db.Order("email asc")
	case "email_desc":
		db = db.Order("email desc")
	default:
		db = db.Order("paid_at desc")
	}

	// pagination
	page := params.Page
	if page <= 0 {
		page = 1
	}
	// limit per page
	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

// ** khusus cron job update status ke failed
func (r *paymentRepository) ExpireOldPendingPayments() (int64, error) {
	threshold := time.Now().Add(-24 * time.Hour)

	result := r.db.Model(&models.Payment{}).
		Where("status = ? AND paid_at <= ?", "pending", threshold).
		Update("status", "failed")

	return result.RowsAffected, result.Error
}
