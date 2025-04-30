package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type UserPackageRepository interface {
	CreateUserPackage(userPackage *models.UserPackage) error
	GetUserPackagesByUserID(userID string) ([]models.UserPackage, error)
	GetActiveUserPackage(userID string) (*models.UserPackage, error)
	UpdateUserPackage(userPackage *models.UserPackage) error
}

type userPackageRepository struct {
	db *gorm.DB
}

func NewUserPackageRepository(db *gorm.DB) UserPackageRepository {
	return &userPackageRepository{db}
}

func (r *userPackageRepository) CreateUserPackage(userPackage *models.UserPackage) error {
	return r.db.Create(userPackage).Error
}

func (r *userPackageRepository) GetUserPackagesByUserID(userID string) ([]models.UserPackage, error) {
	var userPackages []models.UserPackage
	if err := r.db.Where("user_id = ?", userID).Order("purchased_at desc").Preload("Package").Preload("User").Find(&userPackages).Error; err != nil {
		return nil, err
	}
	return userPackages, nil
}

func (r *userPackageRepository) GetActiveUserPackage(userID string) (*models.UserPackage, error) {
	var userPackage models.UserPackage
	if err := r.db.
		Where("user_id = ? AND remaining_credit > 0", userID).
		Order("purchased_at desc").
		First(&userPackage).Error; err != nil {
		return nil, err
	}
	return &userPackage, nil
}

func (r *userPackageRepository) UpdateUserPackage(userPackage *models.UserPackage) error {
	return r.db.Save(userPackage).Error
}
