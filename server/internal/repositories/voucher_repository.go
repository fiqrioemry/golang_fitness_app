package repositories

import (
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	Create(v *models.Voucher) error
	GetAll() ([]models.Voucher, error)
	UpdateVoucher(v *models.Voucher) error
	GetByCode(code string) (*models.Voucher, error)
	CheckVoucherUsed(userID, voucherID uuid.UUID) (bool, error)
	InsertUsedVoucher(userID, voucherID uuid.UUID) error
	GetValidVoucherByCode(code string) (*models.Voucher, error)
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

func (r *voucherRepository) UpdateVoucher(v *models.Voucher) error {
	return r.db.Save(v).Error
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

func (r *voucherRepository) GetValidVoucherByCode(code string) (*models.Voucher, error) {
	var v models.Voucher
	err := r.db.
		Where("code = ? AND expired_at > NOW() AND quota > 0", code).
		First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *voucherRepository) CheckVoucherUsed(userID, voucherID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.UsedVoucher{}).
		Where("user_id = ? AND voucher_id = ?", userID, voucherID).
		Count(&count).Error
	return count > 0, err
}

func (r *voucherRepository) InsertUsedVoucher(userID, voucherID uuid.UUID) error {
	return r.db.Create(&models.UsedVoucher{
		ID:        uuid.New(),
		UserID:    userID,
		VoucherID: voucherID,
		UsedAt:    time.Now(),
	}).Error
}
