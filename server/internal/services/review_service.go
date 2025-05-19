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
	repo           repositories.ReviewRepository
	scheduleRepo   repositories.ClassScheduleRepository
	instructorRepo repositories.InstructorRepository
}

func NewReviewService(repo repositories.ReviewRepository, scheduleRepo repositories.ClassScheduleRepository, instructorRepo repositories.InstructorRepository) ReviewService {
	return &reviewService{repo, scheduleRepo, instructorRepo}
}

func (s *reviewService) CreateReview(userID string, req dto.CreateReviewRequest) error {

	schedule, err := s.scheduleRepo.GetClassScheduleByID(req.ClassScheduleID)
	if err != nil {
		return err
	}

	review := models.Review{
		ID:        uuid.New(),
		UserID:    uuid.MustParse(userID),
		ClassID:   schedule.ClassID,
		Rating:    req.Rating,
		Comment:   req.Comment,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateReview(&review); err != nil {
		return err
	}

	avgRating, err := s.repo.GetAverageRatingByInstructorID(schedule.InstructorID)
	if err != nil {
		return err
	}

	if err := s.instructorRepo.UpdateRating(schedule.InstructorID, avgRating); err != nil {
		return err
	}

	return nil
}

func (s *reviewService) GetReviewsByClassID(classID string) ([]dto.ReviewResponse, error) {
	reviews, err := s.repo.GetReviewsByClassID(classID)
	if err != nil {
		return nil, err
	}

	var result []dto.ReviewResponse
	for _, r := range reviews {
		result = append(result, dto.ReviewResponse{
			ID:        r.ID.String(),
			UserName:  r.User.Profile.Fullname,
			ClassName: r.Class.Title,
			Rating:    r.Rating,
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
		})
	}
	return result, nil
}
