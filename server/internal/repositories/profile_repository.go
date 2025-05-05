package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(user *models.User) error
	GetUserTransactions(userID string, limit, offset int) ([]models.Payment, int64, error)
	GetUserPackages(userID string, limit, offset int) ([]models.UserPackage, int64, error)
	GetUserPackagesByClassID(userID, classID string) ([]models.UserPackage, error)
	GetUserBookings(userID string, limit, offset int) ([]models.Booking, int64, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}

func (r *profileRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Profile").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *profileRepository) UpdateUser(user *models.User) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(user).Error
}

func (r *profileRepository) GetUserTransactions(userID string, limit, offset int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	query := r.db.
		Model(&models.Payment{}).
		Preload("Package").
		Where("user_id = ?", userID)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("paid_at DESC").Limit(limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

func (r *profileRepository) GetUserPackages(userID string, limit, offset int) ([]models.UserPackage, int64, error) {
	var data []models.UserPackage
	var count int64

	query := r.db.Model(&models.UserPackage{}).
		Preload("Package").
		Where("user_id = ?", userID)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("purchased_at DESC").Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

func (r *profileRepository) GetUserBookings(userID string, limit, offset int) ([]models.Booking, int64, error) {
	var data []models.Booking
	var count int64

	query := r.db.Model(&models.Booking{}).
		Preload("ClassSchedule.Class.Location").
		Preload("ClassSchedule.Instructor.User.Profile").
		Where("user_id = ?", userID)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

func (r *profileRepository) GetUserPackagesByClassID(userID, classID string) ([]models.UserPackage, error) {
	var userPackages []models.UserPackage

	err := r.db.
		Joins("JOIN package_classes pc ON pc.package_id = user_packages.package_id").
		Where("user_packages.user_id = ? AND pc.class_id = ?", userID, classID).
		Preload("Package.Classes").
		Find(&userPackages).Error

	return userPackages, err
}
