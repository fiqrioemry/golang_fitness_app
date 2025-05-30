package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers(params dto.UserQueryParam) ([]models.User, int64, error)
	FindUserByID(id string) (*models.User, error)
	GetUserStats() (total, customers, instructors, admins, newThisMonth int64, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAllUsers(params dto.UserQueryParam) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	db := r.db.Model(&models.User{}).Joins("JOIN profiles ON users.id = profiles.user_id").Preload("Profile")

	if params.Q != "" {
		db = db.Where("users.email LIKE ? OR profiles.fullname LIKE ?", "%"+params.Q+"%", "%"+params.Q+"%")
	}
	if params.Role != "" && params.Role != "all" {
		db = db.Where("users.role = ?", params.Role)
	}

	switch params.Sort {
	case "joined_asc":
		db = db.Order("users.created_at asc")
	case "joined_desc":
		db = db.Order("users.created_at desc")
	case "email_asc":
		db = db.Order("users.email asc")
	case "email_desc":
		db = db.Order("users.email desc")
	case "name_asc":
		db = db.Order("profiles.fullname asc")
	case "name_desc":
		db = db.Order("profiles.fullname desc")
	default:
		db = db.Order("users.created_at desc")
	}

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(params.Limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (r *userRepository) FindUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Profile").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserStats() (int64, int64, int64, int64, int64, error) {
	var total, customers, instructors, admins, newThisMonth int64
	var err error
	db := r.db.Model(&models.User{})

	db.Count(&total)

	db.Where("role = ?", "customer").Count(&customers)

	db.Where("role = ?", "instructor").Count(&instructors)

	db.Where("role = ?", "admin").Count(&admins)

	db.Where("created_at >= ?", time.Now().AddDate(0, -1, 0)).Count(&newThisMonth)

	return total, customers, instructors, admins, newThisMonth, err
}
