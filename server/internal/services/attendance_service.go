package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type AttendanceService interface {
	MarkAttendance(userID string, req dto.MarkAttendanceRequest) error
	GetAllAttendances() ([]dto.AttendanceResponse, error)
	ExportAttendancesToExcel() (*excelize.File, error)
}

type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
	bookingRepo    repositories.BookingRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository, bookingRepo repositories.BookingRepository) AttendanceService {
	return &attendanceService{attendanceRepo, bookingRepo}
}

func (s *attendanceService) MarkAttendance(userID string, req dto.MarkAttendanceRequest) error {
	// Fetch Booking
	booking, err := s.bookingRepo.GetBookingByID(req.BookingID)
	if err != nil {
		return errors.New("booking not found")
	}

	// Check if already marked
	exist, _ := s.attendanceRepo.GetAttendanceByBooking(userID, booking.ClassScheduleID.String())
	if exist != nil && exist.ID != uuid.Nil {
		// If already exist, update
		exist.Status = req.Status
		now := time.Now()
		exist.CheckedAt = &now
		return s.attendanceRepo.UpdateAttendance(exist)
	}

	// New Attendance
	now := time.Now()
	attendance := models.Attendance{
		ID:              uuid.New(),
		UserID:          uuid.MustParse(userID),
		ClassScheduleID: booking.ClassScheduleID,
		Status:          req.Status,
		CheckedAt:       &now,
	}
	return s.attendanceRepo.CreateAttendance(&attendance)
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
			ID:           a.ID.String(),
			ClassTitle:   a.ClassSchedule.Class.Title,
			UserFullname: a.User.Profile.Fullname,
			StartTime:    a.ClassSchedule.StartTime.Format(time.RFC3339),
			EndTime:      a.ClassSchedule.EndTime.Format(time.RFC3339),
			Status:       a.Status,
			CheckedAt:    checkedAt,
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

	// Header
	headers := []string{"No", "User Name", "Class Title", "Start Time", "End Time", "Status"}
	for idx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(idx+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// Isi data
	for i, a := range attendances {
		row := i + 2
		startTime := a.ClassSchedule.StartTime.Format("2006-01-02 15:04")
		endTime := a.ClassSchedule.EndTime.Format("2006-01-02 15:04")

		values := []interface{}{
			i + 1,
			a.User.Profile.Fullname,
			a.ClassSchedule.Class.Title,
			startTime,
			endTime,
			a.Status,
		}

		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
	}

	return f, nil
}
