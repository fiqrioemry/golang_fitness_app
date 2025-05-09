package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type VoucherRepository interface {
	Create(v *models.Voucher) error
	GetAll() ([]models.Voucher, error)
	GetByCode(code string) (*models.Voucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db}
}

func (r *voucherRepository) Create(v *models.Voucher) error {
	return r.db.Create(v).Error
}

func (r *voucherRepository) GetAll() ([]models.Voucher, error) {
	var vouchers []models.Voucher
	err := r.db.Order("created_at desc").Find(&vouchers).Error
	return vouchers, err
}

func (r *voucherRepository) GetByCode(code string) (*models.Voucher, error) {
	var v models.Voucher
	err := r.db.Where("code = ?", code).First(&v).Error
	return &v, err
}
