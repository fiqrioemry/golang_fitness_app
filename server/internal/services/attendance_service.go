package services

import (
	"errors"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
)

type AttendanceService interface {
	GetBookingInfo(bookingID string) (*models.Booking, error)
	ValidateQRCodeData(qr string) (*dto.AttendanceResponse, error)
	GetAllAttendances(userID string) ([]dto.AttendanceResponse, error)
	CheckinAttendance(userID string, bookingID string) (string, error)
	GetQRCode(userID, bookingID string) (string, *models.Booking, error)
	GetAttendanceDetail(scheduleID string) ([]dto.AttendanceDetailResponse, error)
	MarkAbsentBookings() error
}

type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
	bookingRepo    repositories.BookingRepository
	reviewRepo     repositories.ReviewRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository, bookingRepo repositories.BookingRepository, reviewRepo repositories.ReviewRepository) AttendanceService {
	return &attendanceService{attendanceRepo, bookingRepo, reviewRepo}
}

func (s *attendanceService) GetBookingInfo(bookingID string) (*models.Booking, error) {
	return s.bookingRepo.GetBookingByID(bookingID)
}

func (s *attendanceService) GetAllAttendances(userID string) ([]dto.AttendanceResponse, error) {
	attendances, err := s.attendanceRepo.GetAllAttendancesByUser(userID)
	if err != nil {
		return nil, err
	}
	var result []dto.AttendanceResponse
	for _, a := range attendances {

		classID := a.ClassSchedule.Class.ID.String()

		reviews, err := s.reviewRepo.GetUserReviewStatus(userID, classID)
		if err != nil {
			return nil, err
		}

		reviewed := reviews != nil

		checkedAt := ""
		if a.CheckedAt != nil {
			checkedAt = a.CheckedAt.Format(time.RFC3339)
		}
		result = append(result, dto.AttendanceResponse{
			ID:             a.ID.String(),
			ScheduleID:     a.ClassSchedule.ID.String(),
			ClassName:      a.ClassSchedule.ClassName,
			ClassImage:     a.ClassSchedule.ClassImage,
			InstructorID:   a.ClassSchedule.InstructorID.String(),
			InstructorName: a.ClassSchedule.InstructorName,
			Date:           a.ClassSchedule.Date.Format(time.RFC3339),
			StartHour:      a.ClassSchedule.StartHour,
			StartMinute:    a.ClassSchedule.StartMinute,
			Status:         a.Status,
			Reviewed:       reviewed,
			CheckedAt:      checkedAt,
		})
	}
	return result, nil
}

func (s *attendanceService) CheckinAttendance(userID, bookingID string) (string, error) {
	booking, err := s.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		return "", errors.New("booking not found")
	}

	if booking.UserID.String() != userID {
		return "", errors.New("unauthorized access")
	}

	schedule := booking.ClassSchedule

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	startTime := time.Date(
		schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
		schedule.StartHour, schedule.StartMinute, 0, 0, loc,
	)

	if now.Before(startTime.Add(-15*time.Minute)) || now.After(startTime.Add(30*time.Minute)) {
		return "", errors.New("attendance window closed")
	}

	attendance, err := s.attendanceRepo.MarkAsAttendance(userID, bookingID)
	if err != nil {
		return "", err
	}
	if attendance.Status != "attended" {
		return "", errors.New("attendance not marked properly")
	}
	booking.Status = "checked_in"
	if err := s.bookingRepo.UpdateBookingStatus(booking.ID, booking.Status); err != nil {
		return "", errors.New("failed to update booking status")
	}

	qr := utils.GenerateBase64QR(attendance.ID.String())
	return qr, nil
}

func (s *attendanceService) GetAttendanceDetail(scheduleID string) ([]dto.AttendanceDetailResponse, error) {
	attendances, err := s.attendanceRepo.GetClassAttendance(scheduleID)
	if err != nil {
		return nil, err
	}

	var result []dto.AttendanceDetailResponse
	for _, a := range attendances {
		checkedAt := ""
		if a.CheckedAt != nil {
			checkedAt = a.CheckedAt.Format(time.RFC3339)
		}
		result = append(result, dto.AttendanceDetailResponse{
			ID:        a.ID.String(),
			Fullname:  a.User.Profile.Fullname,
			Avatar:    a.User.Profile.Avatar,
			CheckedAt: checkedAt,
		})
	}

	return result, nil
}

func (s *attendanceService) GetQRCode(userID, bookingID string) (string, *models.Booking, error) {
	attendance, err := s.attendanceRepo.FindByUserBooking(userID, bookingID)
	if err != nil {
		return "", nil, err
	}
	if attendance.Status != "attended" {
		return "", nil, errors.New("attendance not marked yet")
	}
	qr := utils.GenerateBase64QR(attendance.ID.String())

	booking, err := s.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		return "", nil, err
	}

	return qr, booking, nil
}
func (s *attendanceService) ValidateQRCodeData(qr string) (*dto.AttendanceResponse, error) {
	data, err := utils.ParseQRPayload(qr)
	if err != nil {
		return nil, errors.New("invalid QR format")
	}

	attendance, err := s.attendanceRepo.FindByUserBooking(data.UserID, data.BookingID)
	if err != nil {
		return nil, errors.New("attendance record not found")
	}

	if !attendance.Verified {
		now := time.Now()
		attendance.Verified = true
		attendance.VerifiedAt = &now
		if err := s.attendanceRepo.UpdateAttendance(attendance); err != nil {
			return nil, errors.New("failed to mark as verified")
		}
	}

	return &dto.AttendanceResponse{
		ID:             attendance.ID.String(),
		ScheduleID:     attendance.ClassSchedule.ID.String(),
		ClassName:      attendance.ClassSchedule.ClassName,
		ClassImage:     attendance.ClassSchedule.ClassImage,
		InstructorID:   attendance.ClassSchedule.InstructorID.String(),
		InstructorName: attendance.ClassSchedule.InstructorName,
		Date:           attendance.ClassSchedule.Date.Format(time.RFC1123),
		StartHour:      attendance.ClassSchedule.StartHour,
		StartMinute:    attendance.ClassSchedule.StartMinute,
		Status:         attendance.Status,
		Verified:       attendance.Verified,
		CheckedAt:      attendance.CheckedAt.Format(time.RFC1123),
	}, nil

}

// ** buat cron job
func (s *attendanceService) MarkAbsentBookings() error {
	now := time.Now()

	bookings, err := s.bookingRepo.GetAllBookedWithScheduleAndClass()
	if err != nil {
		return err
	}

	var totalMarked int
	for _, b := range bookings {
		schedule := b.ClassSchedule
		class := schedule.Class

		loc, _ := time.LoadLocation("Asia/Jakarta")

		startTime := time.Date(
			schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
			schedule.StartHour, schedule.StartMinute, 0, 0, loc,
		)
		endTime := startTime.Add(time.Duration(class.Duration) * time.Minute)

		if now.After(endTime) {
			exists, err := s.attendanceRepo.CheckAttendanceExists(b.UserID, b.ID)
			if err != nil {
				continue
			}
			if !exists {
				attendance := &models.Attendance{
					ID:              uuid.New(),
					UserID:          b.UserID,
					ClassScheduleID: schedule.ID,
					Status:          "absent",
				}
				err := s.attendanceRepo.CreateAttendance(attendance)
				if err == nil {
					totalMarked++
				}
			}
		}
	}

	log.Printf(" Marked %d users as absent\n", totalMarked)
	return nil
}
