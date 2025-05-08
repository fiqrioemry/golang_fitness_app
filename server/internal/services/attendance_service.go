package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"
)

type AttendanceService interface {
	MarkAbsentAttendances() error
	GetBookingInfo(bookingID string) (*models.Booking, error)
	ValidateQRCodeData(qr string) (*dto.AttendanceResponse, error)
	CheckinAttendance(userID string, bookingID string) (string, error)
	GetAllAttendances(userID string) ([]dto.AttendanceResponse, error)
	GetQRCode(userID, bookingID string) (string, *models.Booking, error)
	GetAttendanceDetail(scheduleID string) ([]dto.AttendanceDetailResponse, error)
}

type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
	bookingRepo    repositories.BookingRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository, bookingRepo repositories.BookingRepository) AttendanceService {
	return &attendanceService{attendanceRepo, bookingRepo}
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
		checkedAt := ""
		if a.CheckedAt != nil {
			checkedAt = a.CheckedAt.Format(time.RFC3339)
		}
		result = append(result, dto.AttendanceResponse{
			ID: a.ID.String(),
			Class: dto.ClassBrief{
				ID:       a.ClassSchedule.Class.ID.String(),
				Title:    a.ClassSchedule.Class.Title,
				Image:    a.ClassSchedule.Class.Image,
				Duration: a.ClassSchedule.Class.Duration,
			},
			Instructor: dto.InstructorBrief{
				ID:       a.ClassSchedule.Instructor.ID.String(),
				Fullname: a.ClassSchedule.Instructor.User.Profile.Fullname,
				Rating:   a.ClassSchedule.Instructor.Rating,
			},
			Fullname:    a.User.Profile.Fullname,
			Date:        a.ClassSchedule.Date.Format(time.RFC3339),
			StartHour:   a.ClassSchedule.StartHour,
			StartMinute: a.ClassSchedule.StartMinute,
			Status:      a.Status,
			CheckedAt:   checkedAt,
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

	// schedule := booking.ClassSchedule

	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// now := time.Now().In(loc)

	// startTime := time.Date(
	// 	schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
	// 	schedule.StartHour, schedule.StartMinute, 0, 0, loc,
	// )

	// if now.Before(startTime.Add(-15*time.Minute)) || now.After(startTime.Add(30*time.Minute)) {
	// 	return "", errors.New("attendance window closed")
	// }

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

func (s *attendanceService) MarkAbsentAttendances() error {
	now := time.Now()
	schedules, err := s.attendanceRepo.FindAllSchedulesBefore(now.Add(-30 * time.Minute))
	if err != nil {
		return err
	}
	for _, sched := range schedules {
		err := s.attendanceRepo.MarkAbsentIfNotCheckedIn(sched.ID)
		if err != nil {
			return err
		}
	}
	return nil
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
		ID: attendance.ID.String(),
		Class: dto.ClassBrief{
			ID:       attendance.ClassSchedule.Class.ID.String(),
			Title:    attendance.ClassSchedule.Class.Title,
			Image:    attendance.ClassSchedule.Class.Image,
			Duration: attendance.ClassSchedule.Class.Duration,
		},
		Instructor: dto.InstructorBrief{
			ID:       attendance.ClassSchedule.Instructor.ID.String(),
			Fullname: attendance.ClassSchedule.Instructor.User.Profile.Fullname,
			Rating:   attendance.ClassSchedule.Instructor.Rating,
		},
		StartHour:   attendance.ClassSchedule.StartHour,
		StartMinute: attendance.ClassSchedule.StartMinute,
		Fullname:    attendance.User.Profile.Fullname,
		Status:      attendance.Status,
		Verified:    attendance.Verified,
		CheckedAt:   attendance.CheckedAt.Format(time.RFC1123),
	}, nil

}
