package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentByID(id string) (*models.Payment, error)
	UpdatePayment(payment *models.Payment) error
	GetPaymentByOrderID(orderID string) (*models.Payment, error)
	GetAllUserPayments(query string, limit, offset int) ([]models.Payment, int64, error)
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

func (r *paymentRepository) GetAllUserPayments(query string, limit, offset int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	db := r.db.Model(&models.Payment{}).
		Preload("Package", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("User.Profile")

	if query != "" {
		db = db.Joins("JOIN users ON users.id = payments.user_id").
			Where("users.email LIKE ? OR users.id LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Order("paid_at DESC").Limit(limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}
