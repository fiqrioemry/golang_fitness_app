package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/xuri/excelize/v2"
)

type AttendanceService interface {
	MarkAbsentAttendances() error
	ExportAttendancesToExcel() (*excelize.File, error)
	GetAllAttendances() ([]dto.AttendanceResponse, error)
	GetQRCode(userID, bookingID string) (string, error)
	CheckinAttendance(userID string, bookingID string) (string, error)
}

type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
	bookingRepo    repositories.BookingRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository, bookingRepo repositories.BookingRepository) AttendanceService {
	return &attendanceService{attendanceRepo, bookingRepo}
}

func (s *attendanceService) GetAllAttendances() ([]dto.AttendanceResponse, error) {
	attendances, err := s.attendanceRepo.GetAllAttendances()
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
			ID:          a.ID.String(),
			ClassName:   a.ClassSchedule.Class.Title,
			Fullname:    a.User.Profile.Fullname,
			StartHour:   a.ClassSchedule.StartHour,
			StartMinute: a.ClassSchedule.StartMinute,
			Status:      a.Status,
			CheckedAt:   checkedAt,
		})
	}
	return result, nil
}

func (s *attendanceService) ExportAttendancesToExcel() (*excelize.File, error) {
	attendances, err := s.attendanceRepo.GetAllAttendances()
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Attendances"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{"No", "User Name", "Class Title", "Start Time", "End Time", "Status"}
	for idx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(idx+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	for i, a := range attendances {
		row := i + 2
		startHour := a.ClassSchedule.StartHour
		startMinute := a.ClassSchedule.StartMinute

		values := []interface{}{
			i + 1,
			a.User.Profile.Fullname,
			a.ClassSchedule.Class.Title,
			startHour,
			startMinute,
			a.Status,
		}

		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
	}

	return f, nil
}

func (s *attendanceService) CheckinAttendance(userID string, bookingID string) (string, error) {
	booking, err := s.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		return "", errors.New("booking not found")
	}
	if booking.UserID.String() != userID {
		return "", errors.New("unauthorized")
	}

	schedule := booking.ClassSchedule
	startTime := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(), schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC)
	now := time.Now().UTC()

	// attendance window: 15 menit sebelum sampai 30 menit setelah
	if now.Before(startTime.Add(-15*time.Minute)) || now.After(startTime.Add(30*time.Minute)) {
		return "", errors.New("attendance window closed")
	}

	attendance, err := s.attendanceRepo.MarkAsAttendance(userID, bookingID)
	if err != nil {
		return "", err
	}

	// generate QR hanya jika status "attended"
	if attendance.Status != "attended" {
		return "", errors.New("attendance not marked properly")
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

func (s *attendanceService) GetQRCode(userID, bookingID string) (string, error) {
	attendance, err := s.attendanceRepo.FindByUserBooking(userID, bookingID)
	if err != nil {
		return "", err
	}
	if attendance.Status != "attended" {
		return "", errors.New("attendance not marked yet")
	}
	qr := utils.GenerateBase64QR(attendance.ID.String())
	return qr, nil
}
