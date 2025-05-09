package repositories

import (
	"errors"
	"server/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
	GetReviewsByClassID(classID string) ([]models.Review, error)
	GetUserReviewStatus(userID, classID string) (*models.Review, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) GetUserReviewStatus(userID, classID string) (*models.Review, error) {
	var review models.Review
	err := r.db.Where("class_id = ? AND user_id = ?", classID, userID).First(&review).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &review, err
}

func (r *reviewRepository) GetReviewsByClassID(classID string) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Preload("User.Profile").Preload("Class").
		Where("class_id = ?", classID).
		Order("created_at desc").
		Find(&reviews).Error
	return reviews, err
}
