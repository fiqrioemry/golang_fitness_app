package services

import (
	"errors"
	"fmt"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(userID, packageID, scheduleID string) error
	MarkAbsentBookings() error
	EnterClassSchedule(userID, bookingID string) (*dto.QRCodeAttendanceResponse, error)
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

	userPackage.RemainingCredit -= 1
	if err := s.userPackageRepo.UpdateUserPackage(&userPackage); err != nil {
		return err
	}

	if err := s.classScheduleRepo.IncrementBooked(schedule.ID); err != nil {
		return err
	}

	// TODO: Use RabbitMQ to emit "payment_success" event for async email delivery (only in production with EDA)
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
		attendanceStatus := ""
		if b.Attendance.ID != uuid.Nil {
			attendanceStatus = b.Attendance.Status
		}

		participantCount, _ := s.bookingRepo.CountBookingBySchedule(schedule.ID.String())

		result = append(result, dto.BookingResponse{
			ID:               b.ID.String(),
			BookingStatus:    b.Status,
			AttendanceStatus: attendanceStatus,
			ClassID:          schedule.ClassID.String(),
			ClassName:        schedule.ClassName,
			ClassImage:       schedule.ClassImage,
			InstructorName:   schedule.InstructorName,
			Duration:         schedule.Duration,
			StartHour:        schedule.StartHour,
			StartMinute:      schedule.StartMinute,
			Location:         schedule.Location,
			Participant:      int(participantCount),
			Date:             schedule.Date.Format("2006-01-02"),
			BookedAt:         b.CreatedAt.Format("2006-01-02 15:04:05"),
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

func (s *bookingService) EnterClassSchedule(userID, bookingID string) (*dto.QRCodeAttendanceResponse, error) {
	log.Printf("📥 [EnterClass] User %s attempting to enter booking %s\n", userID, bookingID)

	booking, err := s.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		log.Printf("❌ [EnterClass] Booking not found: %v\n", err)
		return nil, errors.New("booking not found")
	}
	log.Printf("✅ [EnterClass] Booking retrieved. Class: %s\n", booking.ClassSchedule.ClassName)

	if booking.UserID.String() != userID {
		log.Printf("⛔ [EnterClass] Unauthorized access. Booking owned by %s\n", booking.UserID.String())
		return nil, errors.New("unauthorized access")
	}

	schedule := booking.ClassSchedule
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	startTime := time.Date(
		schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
		schedule.StartHour, schedule.StartMinute, 0, 0, loc,
	)

	if now.Before(startTime.Add(-15*time.Minute)) || now.After(startTime.Add(30*time.Minute)) {
		log.Printf("⏱️ [EnterClass] Access denied. Current time %v is outside attendance window for %s\n", now, schedule.ClassName)
		return nil, errors.New("attendance window closed")
	}

	log.Println("🟢 [EnterClass] Attendance time is valid. Marking attendance...")
	if err := s.bookingRepo.MarkAsAttendance(userID, bookingID); err != nil {
		log.Printf("❌ [EnterClass] Failed to mark attendance: %v\n", err)
		return nil, errors.New("failed to mark attendance")
	}

	booking, err = s.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		log.Printf("❌ [EnterClass] Failed to fetch updated booking: %v\n", err)
		return nil, errors.New("failed to fetch updated booking")
	}

	if booking.Attendance.ID == uuid.Nil {
		log.Println("❌ [EnterClass] Attendance not created")
		return nil, errors.New("attendance not created")
	}
	log.Printf("✅ [EnterClass] Attendance created successfully: %s\n", booking.Attendance.ID)

	qr := utils.GenerateBase64QR(booking.Attendance.ID.String())
	log.Printf("🔐 [EnterClass] QR code generated for attendance ID: %s\n", booking.Attendance.ID)

	response := dto.QRCodeAttendanceResponse{
		QR:             qr,
		ClassName:      booking.ClassSchedule.ClassName,
		InstructorName: booking.ClassSchedule.InstructorName,
		Date:           booking.ClassSchedule.Date.Format("2006-01-02"),
		StartTime:      fmt.Sprintf("%02d:%02d", booking.ClassSchedule.StartHour, booking.ClassSchedule.StartMinute),
	}

	log.Printf("📤 [EnterClass] Returning QR attendance response for user %s\n", userID)
	return &response, nil
}

func (s *bookingService) GetQRCode(userID, bookingID string) (*dto.QRCodeAttendanceResponse, error) {
	booking, err := s.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		return nil, errors.New("failed to fetch updated booking")
	}

	qr := utils.GenerateBase64QR(booking.Attendance.ID.String())

	response := dto.QRCodeAttendanceResponse{
		QR:             qr,
		ClassName:      booking.ClassSchedule.ClassName,
		InstructorName: booking.ClassSchedule.InstructorName,
		Date:           booking.ClassSchedule.Date.Format("2006-01-02"),
		StartTime:      fmt.Sprintf("%02d:%02d", booking.ClassSchedule.StartHour, booking.ClassSchedule.StartMinute),
	}

	return &response, nil
}

// ** buat cron job
func (s *bookingService) MarkAbsentBookings() error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

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
			schedule.StartHour, schedule.StartMinute, 0, 0, loc,
		)
		endTime := startTime.Add(time.Duration(schedule.Duration) * time.Minute)

		if now.After(endTime) {
			exists, err := s.bookingRepo.CheckAttendanceExists(b.ID)
			if err != nil {
				log.Printf("Failed to check attendance for booking %s: %v\n", b.ID, err)
				continue
			}

			if !exists {
				attendance := &models.Attendance{
					ID:        uuid.New(),
					BookingID: b.ID,
					Status:    "absent",
					CreatedAt: now,
				}

				if err := s.bookingRepo.CreateAttendance(attendance); err != nil {
					log.Printf("Failed to create absent attendance for booking %s: %v\n", b.ID, err)
				} else {
					totalMarked++
				}
			}
		}
	}

	log.Printf("Marked %d users as absent\n", totalMarked)
	return nil
}
