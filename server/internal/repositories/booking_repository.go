package repositories

import (
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	GetBookingByID(userID, bookingID string) (*models.Booking, error)
	CountBookingBySchedule(scheduleID string) (int64, error)
	UpdateBookingStatus(bookingID uuid.UUID, status string) error
	IsUserBookedSchedule(userID, scheduleID string) (bool, error)
	FindByUserAndSchedule(userID, scheduleID string) (*models.Booking, error)
	CheckAttendanceExists(bookingID uuid.UUID) (bool, error)
	CreateAttendance(attendance *models.Attendance) error
	UpdateAttendanceStatus(bookingID, status string) error
	GetBookingsByUserID(userID string, params dto.BookingQueryParam) ([]models.Booking, int64, error)
	GetAttendanceByBookingID(bookingID uuid.UUID) (*models.Attendance, error)
	// ** cron job
	GetAllBookedWithScheduleAndClass() ([]models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) UpdateBookingStatus(bookingID uuid.UUID, status string) error {
	return r.db.Model(&models.Booking{}).
		Where("id = ?", bookingID).
		Update("status", status).Error
}

func (r *bookingRepository) GetBookingsByUserID(userID string, params dto.BookingQueryParam) ([]models.Booking, int64, error) {
	var bookings []models.Booking
	var count int64

	db := r.db.Model(&models.Booking{}).
		Preload("ClassSchedule").
		Where("user_id = ?", userID).
		Joins("JOIN class_schedules ON class_schedules.id = bookings.class_schedule_id")

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("class_schedules.class_name LIKE ?", like)
	}

	if params.Status != "" && params.Status != "all" {
		db = db.Where("bookings.status = ?", params.Status)
	}

	switch params.Sort {
	case "name_asc":
		db = db.Order("class_schedules.class_name ASC")
	case "name_desc":
		db = db.Order("class_schedules.class_name DESC")
	case "date_asc":
		db = db.Order("class_schedules.date ASC")
	case "date_desc":
		db = db.Order("class_schedules.date DESC")
	default:
		db = db.Order("bookings.created_at DESC")
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.Limit
	if err := db.Limit(params.Limit).Offset(offset).Find(&bookings).Error; err != nil {
		return nil, 0, err
	}

	return bookings, count, nil
}

func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingByID(userID, bookingID string) (*models.Booking, error) {
	var booking models.Booking
	if err := r.db.Preload("ClassSchedule").Where("user_id = ?", userID).First(&booking, "id = ?", bookingID).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) CountBookingBySchedule(scheduleID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).Where("class_schedule_id = ?", scheduleID).Count(&count).Error
	return count, err
}

func (r *bookingRepository) IsUserBookedSchedule(userID, scheduleID string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).
		Where("user_id = ? AND class_schedule_id = ?", userID, scheduleID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *bookingRepository) FindByUserAndSchedule(userID, scheduleID string) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.
		Preload("ClassSchedule").
		Where("user_id = ? AND class_schedule_id = ?", userID, scheduleID).
		First(&booking).Error
	return &booking, err
}

// ** cron job
func (r *bookingRepository) GetAllBookedWithScheduleAndClass() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.
		Preload("ClassSchedule").
		Where("status = ?", "booked").
		Find(&bookings).Error

	return bookings, err
}

func (r *bookingRepository) CheckAttendanceExists(bookingID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.
		Model(&models.Attendance{}).
		Where("booking_id = ?", bookingID).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r bookingRepository) CreateAttendance(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r bookingRepository) UpdateAttendanceStatus(bookingID, status string) error {
	return r.db.Model(&models.Attendance{}).
		Where("booking_id = ?", bookingID).
		Update("status", status).Error
}

func (r *bookingRepository) GetAttendanceByBookingID(bookingID uuid.UUID) (*models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.
		Preload("Booking.ClassSchedule").
		First(&attendance, "booking_id = ?", bookingID).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}
