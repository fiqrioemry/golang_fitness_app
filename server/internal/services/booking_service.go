package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(userID string, req dto.CreateBookingRequest) error
	GetUserBookings(userID string) ([]dto.BookingResponse, error)
}

type bookingService struct {
	bookingRepo       repositories.BookingRepository
	classScheduleRepo repositories.ClassScheduleRepository
	userPackageRepo   repositories.UserPackageRepository
}

func NewBookingService(bookingRepo repositories.BookingRepository, classScheduleRepo repositories.ClassScheduleRepository, userPackageRepo repositories.UserPackageRepository) BookingService {
	return &bookingService{
		bookingRepo:       bookingRepo,
		classScheduleRepo: classScheduleRepo,
		userPackageRepo:   userPackageRepo,
	}
}

func (s *bookingService) CreateBooking(userID string, req dto.CreateBookingRequest) error {
	// Cek UserPackage
	userPackage, err := s.userPackageRepo.GetActiveUserPackage(userID)
	if err != nil {
		return errors.New("user has no active package")
	}
	if userPackage.RemainingCredit <= 0 {
		return errors.New("no remaining credit available")
	}
	if userPackage.ExpiredAt != nil && time.Now().After(*userPackage.ExpiredAt) {
		return errors.New("your package has expired")
	}

	// Cek ClassSchedule
	schedule, err := s.classScheduleRepo.GetClassScheduleByID(req.ClassScheduleID)
	if err != nil {
		return errors.New("class schedule not found")
	}

	// Cek quota
	count, err := s.bookingRepo.CountBookingBySchedule(schedule.ID.String())
	if err != nil {
		return errors.New("failed to count schedule bookings")
	}
	if int(count) >= schedule.Capacity {
		return errors.New("class schedule is full")
	}

	// Create booking
	booking := models.Booking{
		ID:              uuid.New(),
		UserID:          uuid.MustParse(userID),
		ClassScheduleID: schedule.ID,
		Status:          "booked",
	}
	if err := s.bookingRepo.CreateBooking(&booking); err != nil {
		return err
	}

	// Update remaining credit
	userPackage.RemainingCredit -= 1
	if err := s.userPackageRepo.UpdateUserPackage(userPackage); err != nil {
		return err
	}

	return nil
}

func (s *bookingService) GetUserBookings(userID string) ([]dto.BookingResponse, error) {
	bookings, err := s.bookingRepo.GetBookingsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var result []dto.BookingResponse
	for _, b := range bookings {
		result = append(result, dto.BookingResponse{
			ID:         b.ID.String(),
			ClassTitle: b.ClassSchedule.Class.Title,
			StartTime:  b.ClassSchedule.StartTime.Format(time.RFC3339),
			EndTime:    b.ClassSchedule.EndTime.Format(time.RFC3339),
			Status:     b.Status,
		})
	}
	return result, nil
}
