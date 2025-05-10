package repositories

import (
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	UpdateAttendance(att *models.Attendance) error
	MarkAbsentIfNotCheckedIn(scheduleID uuid.UUID) error
	FindAllSchedulesBefore(t time.Time) ([]models.ClassSchedule, error)
	GetClassAttendance(scheduleID string) ([]models.Attendance, error)
	GetAllAttendancesByUser(userID string) ([]models.Attendance, error)
	MarkAsAttendance(userID string, bookingID string) (*models.Attendance, error)
	FindByUserBooking(userID string, bookingID string) (*models.Attendance, error)
}
type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) GetAllAttendancesByUser(userID string) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.
		Preload("ClassSchedule.Class").
		Preload("ClassSchedule.Instructor.User.Profile").
		Preload("User.Profile").
		Where("user_id = ?", userID).
		Find(&attendances).Error
	return attendances, err
}

func (r *attendanceRepository) FindByUserBooking(userID string, bookingID string) (*models.Attendance, error) {
	var booking models.Booking
	if err := r.db.Where("id = ? AND user_id = ?", bookingID, userID).First(&booking).Error; err != nil {
		return nil, err
	}

	var attendance models.Attendance
	err := r.db.Where("user_id = ? AND class_schedule_id = ?", userID, booking.ClassScheduleID).First(&attendance).Error
	return &attendance, err
}

func (r *attendanceRepository) GetClassAttendance(scheduleID string) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.Preload("User.Profile").
		Where("class_schedule_id = ?", scheduleID).
		Find(&attendances).Error
	return attendances, err
}

func (r *attendanceRepository) FindAllSchedulesBefore(t time.Time) ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	err := r.db.
		Where("date + INTERVAL start_hour HOUR + INTERVAL start_minute MINUTE <= ?", t).
		Find(&schedules).Error
	return schedules, err
}

func (r *attendanceRepository) UpdateAttendance(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *attendanceRepository) MarkAsAttendance(userID string, bookingID string) (*models.Attendance, error) {
	var booking models.Booking
	if err := r.db.Where("id = ? AND user_id = ?", bookingID, userID).First(&booking).Error; err != nil {
		return nil, err
	}

	var attendance models.Attendance
	err := r.db.Where("user_id = ? AND class_schedule_id = ?", userID, booking.ClassScheduleID.String()).First(&attendance).Error

	now := time.Now()
	if err == nil {
		attendance.Status = "attended"
		attendance.CheckedAt = &now
		if err := r.db.Save(&attendance).Error; err != nil {
			return nil, err
		}
		return &attendance, nil
	}

	newAttendance := models.Attendance{
		UserID:          uuid.MustParse(userID),
		ClassScheduleID: booking.ClassScheduleID,
		Status:          "attended",
		CheckedAt:       &now,
	}
	if err := r.db.Create(&newAttendance).Error; err != nil {
		return nil, err
	}
	return &newAttendance, nil
}

func (r *attendanceRepository) MarkAbsentIfNotCheckedIn(scheduleID uuid.UUID) error {
	var bookings []models.Booking
	err := r.db.Where("class_schedule_id = ?", scheduleID).
		Where("status != ?", "checked_in").
		Find(&bookings).Error
	if err != nil {
		return err
	}

	for _, booking := range bookings {
		var existing models.Attendance
		err := r.db.Where("user_id = ? AND class_schedule_id = ?", booking.UserID, booking.ClassScheduleID).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			attendance := models.Attendance{
				UserID:          booking.UserID,
				ClassScheduleID: booking.ClassScheduleID,
				Status:          "absent",
			}
			if err := r.db.Create(&attendance).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
