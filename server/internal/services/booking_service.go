package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(userID, packageID, scheduleID string) error
	GetUserBookings(userID string) ([]dto.BookingResponse, error)
}

type bookingService struct {
	bookingRepo       repositories.BookingRepository
	classScheduleRepo repositories.ClassScheduleRepository
	userPackageRepo   repositories.UserPackageRepository
	packageRepo       repositories.PackageRepository
}

func NewBookingService(bookingRepo repositories.BookingRepository, classScheduleRepo repositories.ClassScheduleRepository, userPackageRepo repositories.UserPackageRepository, packageRepo repositories.PackageRepository) BookingService {
	return &bookingService{
		bookingRepo:       bookingRepo,
		classScheduleRepo: classScheduleRepo,
		userPackageRepo:   userPackageRepo,
		packageRepo:       packageRepo,
	}
}

func (s *bookingService) CreateBooking(userID, packageID, scheduleID string) error {

	pkg, err := s.packageRepo.GetPackageByID(packageID)
	if err != nil {
		return fmt.Errorf("package not found")
	}
	if !pkg.IsActive {
		return fmt.Errorf("package is not active")
	}

	// Validasi schedule
	schedule, err := s.classScheduleRepo.GetClassScheduleByID(scheduleID)
	if err != nil {
		return fmt.Errorf("class schedule not found")
	}

	// Validasi userPackage aktif sesuai package
	var userPackage models.UserPackage
	err = s.userPackageRepo.FindActiveByUserAndPackage(userID, packageID, &userPackage)
	if err != nil {
		return fmt.Errorf("you don’t have an active package for this class")
	}
	if userPackage.RemainingCredit <= 0 {
		return fmt.Errorf("not enough credit")
	}

	// Cek kapasitas
	count, err := s.bookingRepo.CountBookingBySchedule(schedule.ID.String())
	if err != nil {
		return fmt.Errorf("failed to count schedule bookings")
	}
	if int(count) >= schedule.Capacity {
		return fmt.Errorf("class schedule is full")
	}

	// Buat booking
	booking := models.Booking{
		ID:              uuid.New(),
		UserID:          uuid.MustParse(userID),
		ClassScheduleID: schedule.ID,
		Status:          "booked",
	}
	if err := s.bookingRepo.CreateBooking(&booking); err != nil {
		return err
	}

	// Kurangi kredit
	userPackage.RemainingCredit -= 1
	if err := s.userPackageRepo.UpdateUserPackage(&userPackage); err != nil {
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
		schedule := b.ClassSchedule
		class := schedule.Class
		location := class.Location
		instructor := schedule.Instructor
		instructorName := instructor.User.Profile.Fullname

		participantCount, err := s.bookingRepo.CountBookingBySchedule(schedule.ID.String())
		if err != nil {
			participantCount = 0
		}

		result = append(result, dto.BookingResponse{
			ID:          b.ID.String(),
			Status:      b.Status,
			BookedAt:    b.CreatedAt.Format("2006-01-02 15:04:05"),
			ClassID:     class.ID.String(),
			ClassTitle:  class.Title,
			ClassImage:  class.Image,
			Duration:    class.Duration,
			StartHour:   schedule.StartHour,
			StartMinute: schedule.StartMinute,
			Location:    location.Name,
			Instructor:  instructorName,
			Participant: int(participantCount),
		})
	}

	return result, nil
}
