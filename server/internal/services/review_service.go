package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ReviewService interface {
	GetReviewsByClassID(classID string) ([]dto.ReviewResponse, error)
	CreateReview(userID string, bookingID string, req dto.CreateReviewRequest) error
}

type reviewService struct {
	repo           repositories.ReviewRepository
	bookingRepo    repositories.BookingRepository
	instructorRepo repositories.InstructorRepository
}

func NewReviewService(repo repositories.ReviewRepository, bookingRepo repositories.BookingRepository, instructorRepo repositories.InstructorRepository) ReviewService {
	return &reviewService{repo, bookingRepo, instructorRepo}
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

func (s *reviewService) CreateReview(userID string, bookingID string, req dto.CreateReviewRequest) error {

	booking, err := s.bookingRepo.GetBookingByID(userID, bookingID)
	if err != nil {
		return err
	}

	attendance := booking.Attendance
	if attendance.IsReviewed {
		return errors.New("you already submitted a review")
	}

	schedule := booking.ClassSchedule
	review := models.Review{
		UserID:  uuid.MustParse(userID),
		ClassID: schedule.ClassID,
		Rating:  req.Rating,
		Comment: req.Comment,
	}

	if err := s.repo.CreateReview(&review); err != nil {
		return err
	}

	attendance.IsReviewed = true

	if err := s.bookingRepo.UpdateAttendance(&attendance); err != nil {
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
