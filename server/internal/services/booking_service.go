package services

import (
	"errors"
	"fmt"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(userID, packageID, scheduleID string) error
	MarkAbsentBookings() error
	CheckedInClassSchedule(userID, bookingID string) error
	CheckoutClassSchedule(userID, bookingID string, req dto.ValidateCheckoutRequest) error
	GetBookingDetail(userID, bookingID string) (*dto.BookingDetailResponse, error)

	GetBookingByUser(userID string, params dto.BookingQueryParam) ([]dto.BookingResponse, *dto.PaginationResponse, error)
}

type bookingService struct {
	bookingRepo         repositories.BookingRepository
	classScheduleRepo   repositories.ClassScheduleRepository
	userPackageRepo     repositories.UserPackageRepository
	packageRepo         repositories.PackageRepository
	notificationService NotificationService
}

func NewBookingService(bookingRepo repositories.BookingRepository, classScheduleRepo repositories.ClassScheduleRepository, userPackageRepo repositories.UserPackageRepository, packageRepo repositories.PackageRepository, notificationService NotificationService) BookingService {
	return &bookingService{
		bookingRepo:         bookingRepo,
		classScheduleRepo:   classScheduleRepo,
		userPackageRepo:     userPackageRepo,
		packageRepo:         packageRepo,
		notificationService: notificationService,
	}
}

func (s *bookingService) CreateBooking(userID, packageID, scheduleID string) error {

	schedule, err := s.classScheduleRepo.GetClassScheduleByID(scheduleID)
	if err != nil {
		return fmt.Errorf("class schedule not found")
	}

	var userPackage models.UserPackage
	err = s.userPackageRepo.GetActiveUserPackages(userID, packageID, &userPackage)
	if err != nil {
		return fmt.Errorf("you don’t have an active package for this class")
	}
	if userPackage.RemainingCredit <= 0 {
		return fmt.Errorf("not enough credit")
	}

	count, err := s.bookingRepo.CountBookingBySchedule(schedule.ID.String())
	if err != nil {
		return fmt.Errorf("failed to count schedule bookings")
	}
	if int(count) >= schedule.Capacity {
		return fmt.Errorf("class schedule is full")
	}

	booking := models.Booking{
		UserID:          uuid.MustParse(userID),
		ClassScheduleID: schedule.ID,
		Status:          "booked",
	}
	if err := s.bookingRepo.CreateBooking(&booking); err != nil {
		return err
	}

	attendance := models.Attendance{
		ID:        uuid.New(),
		BookingID: booking.ID,
		Status:    "none",
	}

	if err := s.bookingRepo.CreateAttendance(&attendance); err != nil {
		return fmt.Errorf("failed to create attendance record: %v", err)
	}

	userPackage.RemainingCredit -= 1
	if err := s.userPackageRepo.UpdateUserPackage(&userPackage); err != nil {
		return err
	}

	if err := s.classScheduleRepo.IncrementBooked(schedule.ID); err != nil {
		return err
	}

	// TODO: Use RabbitMQ to emit "booking_success" event for async email delivery (only in production with EDA)
	payload := dto.NotificationEvent{
		UserID: booking.UserID.String(),
		Type:   "system_message",
		Title:  "Class Booked Successfully",
		Message: fmt.Sprintf(
			"You have successfully booked the class \"%s\" on %s at %02d:%02d. 1 credit has been deducted from your package.",
			schedule.ClassName,
			schedule.Date.Format("January 2, 2006"),
			schedule.StartHour,
			schedule.StartMinute,
		),
	}

	if err := s.notificationService.SendToUser(payload); err != nil {
		log.Printf("failed sending notification to user %s: %v\n", payload.UserID, err)
	}
	// TODO: Use RabbitMQ to emit "payment_success" event for async email delivery (only in production with EDA)

	return nil
}

func (s *bookingService) GetBookingByUser(userID string, params dto.BookingQueryParam) ([]dto.BookingResponse, *dto.PaginationResponse, error) {
	bookings, total, err := s.bookingRepo.GetBookingsByUserID(userID, params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.BookingResponse
	for _, b := range bookings {
		schedule := b.ClassSchedule

		result = append(result, dto.BookingResponse{
			ID:             b.ID.String(),
			BookingStatus:  b.Status,
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorName: schedule.InstructorName,
			Date:           schedule.Date.Format("2006-01-02"),
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Duration:       schedule.Duration,
			Location:       schedule.Location,
			BookedAt:       b.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))
	pagination := &dto.PaginationResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil
}

func (s *bookingService) GetBookingDetail(userID, bookingID string) (*dto.BookingDetailResponse, error) {
	booking, err := s.bookingRepo.GetBookingByID(userID, bookingID)
	if err != nil {
		return nil, errors.New("booking not found")
	}

	attendance := booking.Attendance
	schedule := booking.ClassSchedule

	res := &dto.BookingDetailResponse{
		ID:               booking.ID.String(),
		ClassID:          schedule.ClassID.String(),
		ClassName:        schedule.ClassName,
		ClassImage:       schedule.ClassImage,
		InstructorName:   schedule.InstructorName,
		Date:             schedule.Date.Format("2006-01-02"),
		StartHour:        schedule.StartHour,
		StartMinute:      schedule.StartMinute,
		Duration:         schedule.Duration,
		CheckedIn:        attendance.CheckedIn,
		CheckedOut:       attendance.CheckedOut,
		IsOpened:         schedule.IsOpened,
		IsReviewed:       attendance.IsReviewed,
		AttendanceStatus: attendance.Status,
		CheckedAt:        "",
		VerifiedAt:       "",
	}

	if schedule.ZoomLink != nil {
		res.ZoomLink = *schedule.ZoomLink
	}

	if attendance.CheckedIn && attendance.CheckedAt != nil {
		res.CheckedAt = attendance.CheckedAt.Format("2006-01-02 15:04:05")
	}
	if attendance.CheckedOut && attendance.VerifiedAt != nil {
		res.VerifiedAt = attendance.VerifiedAt.Format("2006-01-02 15:04:05")
	}

	return res, nil
}

func (s *bookingService) CheckedInClassSchedule(userID, bookingID string) error {
	booking, err := s.bookingRepo.GetBookingByID(userID, bookingID)
	if err != nil {
		return errors.New("booking not found")
	}

	if !booking.ClassSchedule.IsOpened {
		return errors.New("class schedule is not opened yet")
	}

	attendance := booking.Attendance

	if attendance.CheckedIn {
		return errors.New("you have already checked in to this class")
	}

	now := time.Now().UTC()
	attendance.CheckedIn = true
	attendance.Status = "entered"
	attendance.CheckedAt = &now

	return s.bookingRepo.UpdateAttendance(&attendance)
}

func (s *bookingService) CheckoutClassSchedule(userID, bookingID string, req dto.ValidateCheckoutRequest) error {
	booking, err := s.bookingRepo.GetBookingByID(userID, bookingID)
	if err != nil {
		return errors.New("booking not found")
	}
	attendance := booking.Attendance

	if attendance.CheckedOut {
		return errors.New("you have already checkedout from this class")
	}

	if booking.ClassSchedule.VerificationCode == nil || *booking.ClassSchedule.VerificationCode != req.VerificationCode {
		return errors.New("invalid verification code")
	}

	now := time.Now().UTC()
	attendance.CheckedOut = true
	attendance.Status = "attended"
	attendance.VerifiedAt = &now

	return s.bookingRepo.UpdateAttendance(&attendance)
}

// ** buat cron job
func (s *bookingService) MarkAbsentBookings() error {
	now := time.Now().UTC()

	bookings, err := s.bookingRepo.GetAllBookedWithScheduleAndClass()
	if err != nil {
		return err
	}

	var totalMarked int
	for _, b := range bookings {
		if b.Status != "booked" {
			continue
		}

		schedule := b.ClassSchedule

		startTime := time.Date(
			schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
			schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC,
		)
		endTime := startTime.Add(time.Duration(schedule.Duration) * time.Minute)

		if now.After(endTime) {
			exists, err := s.bookingRepo.CheckAttendanceExists(b.ID)
			if err != nil {
				log.Printf("❌ Failed to check attendance for booking %s: %v\n", b.ID, err)
				continue
			}

			if !exists {
				attendance := &models.Attendance{
					ID:        uuid.New(),
					BookingID: b.ID,
					Status:    "absent",
				}

				if err := s.bookingRepo.CreateAttendance(attendance); err != nil {
					log.Printf("❌ Failed to create absent attendance for booking %s: %v\n", b.ID, err)
				} else {
					totalMarked++
				}
			}
		}
	}

	log.Printf("✅ Marked %d users as absent\n", totalMarked)
	return nil
}
