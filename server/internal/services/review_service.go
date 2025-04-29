package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ReviewService interface {
	CreateReview(userID string, req dto.CreateReviewRequest) error
	GetReviewsByClassID(classID string) ([]dto.ReviewResponse, error)
}

type reviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{repo}
}

func (s *reviewService) CreateReview(userID string, req dto.CreateReviewRequest) error {
	review := models.Review{
		ID:      uuid.New(),
		UserID:  uuid.MustParse(userID),
		ClassID: uuid.MustParse(req.ClassID),
		Rating:  req.Rating,
		Comment: req.Comment,
	}
	return s.repo.CreateReview(&review)
}

func (s *reviewService) GetReviewsByClassID(classID string) ([]dto.ReviewResponse, error) {
	reviews, err := s.repo.GetReviewsByClassID(classID)
	if err != nil {
		return nil, err
	}

	var result []dto.ReviewResponse
	for _, r := range reviews {
		result = append(result, dto.ReviewResponse{
			ID:         r.ID.String(),
			UserName:   r.User.Profile.Fullname,
			ClassTitle: r.Class.Title,
			Rating:     r.Rating,
			Comment:    r.Comment,
			CreatedAt:  r.CreatedAt.Format(time.RFC3339),
		})
	}
	return result, nil
}
