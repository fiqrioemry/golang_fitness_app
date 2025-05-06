package services

import (
	"errors"
	"fmt"
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
	// ValidateQRCode(attendanceID string) error
	ValidateQRCodeData(qr string) (*dto.AttendanceInfoResponse, error)

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

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	startTime := time.Date(
		schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
		schedule.StartHour, schedule.StartMinute, 0, 0, loc,
	)

	fmt.Println("=== Debug Waktu Check-in ===")
	fmt.Println("schedule.Date       :", schedule.Date)
	fmt.Println("schedule.StartHour  :", schedule.StartHour)
	fmt.Println("schedule.StartMinute:", schedule.StartMinute)
	fmt.Println("startTime           :", startTime)
	fmt.Println("now                 :", now)
	fmt.Println("startTime -15m      :", startTime.Add(-15*time.Minute))
	fmt.Println("startTime +30m      :", startTime.Add(30*time.Minute))

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

// func (s *attendanceService) ValidateQRCode(attendanceID string) error {
// 	attendance, err := s.attendanceRepo.GetByID(attendanceID)
// 	if err != nil {
// 		return errors.New("attendance not found")
// 	}

// 	if attendance.Verified {
// 		return errors.New("already verified")
// 	}

// 	now := time.Now()
// 	attendance.Verified = true
// 	attendance.VerifiedAt = &now

// 	err = s.attendanceRepo.UpdateAttendance(attendance)
// 	if err != nil {
// 		return errors.New("failed to verify attendance")
// 	}

// 	return nil
// }

// func (s *attendanceService) ValidateQRCodeAttendance(payload string) (*dto.AttendanceValidationResponse, error) {
// 	var data struct {
// 		UserID     string `json:"userId"`
// 		BookingID  string `json:"bookingId"`
// 		ScheduleID string `json:"scheduleId"`
// 	}

// 	err := json.Unmarshal([]byte(payload), &data)
// 	if err != nil {
// 		return nil, errors.New("invalid QR payload")
// 	}

// 	attendance, err := s.attendanceRepo.FindByUserBooking(data.UserID, data.BookingID)
// 	if err != nil {
// 		return nil, errors.New("attendance not found")
// 	}

// 	if !attendance.Verified {
// 		now := time.Now()
// 		attendance.Verified = true
// 		attendance.VerifiedAt = &now
// 		s.attendanceRepo.UpdateAttendance(attendance)
// 	}

// 	// response
// 	response := &dto.AttendanceValidationResponse{
// 		Fullname:   attendance.User.Profile.Fullname,
// 		ClassName:  attendance.ClassSchedule.Class.Title,
// 		StartTime:  fmt.Sprintf("%02d:%02d", attendance.ClassSchedule.StartHour, attendance.ClassSchedule.StartMinute),
// 		Status:     attendance.Status,
// 		Verified:   attendance.Verified,
// 		VerifiedAt: attendance.VerifiedAt,
// 	}
// 	return response, nil
// }

func (s *attendanceService) ValidateQRCodeData(qr string) (*dto.AttendanceInfoResponse, error) {
	data, err := utils.ParseQRPayload(qr)
	if err != nil {
		return nil, errors.New("invalid QR format")
	}

	attendance, err := s.attendanceRepo.FindByUserBooking(data.UserID, data.BookingID)
	if err != nil {
		return nil, errors.New("attendance record not found")
	}

	// Optional: verifikasi dan update attendance jika belum verified
	if !attendance.Verified {
		now := time.Now()
		attendance.Verified = true
		attendance.VerifiedAt = &now
		if err := s.attendanceRepo.UpdateAttendance(attendance); err != nil {
			return nil, errors.New("failed to mark as verified")
		}
	}

	return &dto.AttendanceInfoResponse{
		ID:          attendance.ID.String(),
		ClassName:   attendance.ClassSchedule.Class.Title,
		Date:        attendance.ClassSchedule.Date.Format("2006-01-02"),
		StartHour:   attendance.ClassSchedule.StartHour,
		StartMinute: attendance.ClassSchedule.StartMinute,
		Fullname:    attendance.User.Profile.Fullname,
		Status:      attendance.Status,
		Verified:    attendance.Verified,
	}, nil
}
