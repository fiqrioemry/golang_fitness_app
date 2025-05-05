package repositories

import (
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	GetAllAttendances() ([]models.Attendance, error)
	MarkAbsentIfNotCheckedIn(scheduleID uuid.UUID) error
	UpdateAttendance(attendance *models.Attendance) error
	CreateAttendance(attendance *models.Attendance) error
	FindAllSchedulesBefore(t time.Time) ([]models.ClassSchedule, error)
	MarkAsAttendance(userID string, bookingID string) (*models.Attendance, error)
	FindByUserBooking(userID string, bookingID string) (*models.Attendance, error)
	GetAttendanceByBooking(userID, scheduleID string) (*models.Attendance, error)
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

func (r *attendanceRepository) MarkAsAttendance(userID string, bookingID string) (*models.Attendance, error) {
	var booking models.Booking
	if err := r.db.Where("id = ? AND user_id = ?", bookingID, userID).First(&booking).Error; err != nil {
		return nil, err
	}

	var exist models.Attendance
	err := r.db.Where("user_id = ? AND class_schedule_id = ?", userID, booking.ClassScheduleID.String()).First(&exist).Error
	if err == nil {
		return &exist, nil
	}

	now := time.Now()
	attendance := models.Attendance{
		ID:              uuid.New(),
		UserID:          uuid.MustParse(userID),
		ClassScheduleID: booking.ClassScheduleID,
		Status:          "attended",
		CheckedAt:       &now,
	}

	err = r.db.Create(&attendance).Error
	return &attendance, err
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

func (r *attendanceRepository) FindAllSchedulesBefore(t time.Time) ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	err := r.db.
		Where("date + INTERVAL start_hour HOUR + INTERVAL start_minute MINUTE <= ?", t).
		Find(&schedules).Error
	return schedules, err
}

func (r *attendanceRepository) MarkAbsentIfNotCheckedIn(scheduleID uuid.UUID) error {
	var bookings []models.Booking
	err := r.db.Where("class_schedule_id = ?", scheduleID).Find(&bookings).Error
	if err != nil {
		return err
	}

	for _, booking := range bookings {
		var existing models.Attendance
		err := r.db.Where("user_id = ? AND class_schedule_id = ?", booking.UserID, booking.ClassScheduleID).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			attendance := models.Attendance{
				ID:              uuid.New(),
				UserID:          booking.UserID,
				ClassScheduleID: booking.ClassScheduleID,
				Status:          "absent",
				CreatedAt:       time.Now(),
			}
			if err := r.db.Create(&attendance).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
