package repositories

import (
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers(q, role, sort string, limit, offset int) ([]models.User, int64, error)
	FindUserByID(id string) (*models.User, error)
	GetUserStats() (total, customers, instructors, admins, newThisMonth int64, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAllUsers(q, role, sort string, limit, offset int) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	db := r.db.Model(&models.User{}).Joins("JOIN profiles ON users.id = profiles.user_id").Preload("Profile")

	if q != "" {
		db = db.Where("users.email LIKE ? OR profiles.fullname LIKE ?", "%"+q+"%", "%"+q+"%")
	}
	if role != "" {
		db = db.Where("users.role = ?", role)
	}

	switch sort {
	case "oldest":
		db = db.Order("users.created_at asc")
	case "name_asc":
		db = db.Order("profiles.fullname asc")
	case "name_desc":
		db = db.Order("profiles.fullname desc")
	default:
		db = db.Order("users.created_at desc")
	}

	db.Count(&count)
	if err := db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
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
