package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	CreateAttendance(attendance *models.Attendance) error
	GetAttendanceByBooking(userID, scheduleID string) (*models.Attendance, error)
	UpdateAttendance(attendance *models.Attendance) error
	GetAllAttendances() ([]models.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) CreateAttendance(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepository) GetAttendanceByBooking(userID, scheduleID string) (*models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.Where("user_id = ? AND class_schedule_id = ?", userID, scheduleID).First(&attendance).Error
	return &attendance, err
}

func (r *attendanceRepository) UpdateAttendance(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *attendanceRepository) GetAllAttendances() ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.
		Preload("User.Profile").
		Preload("ClassSchedule.Class").
		Find(&attendances).Error
	return attendances, err
}
