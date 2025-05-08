package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	CountBookingBySchedule(scheduleID string) (int64, error)
	GetBookingByID(id string) (*models.Booking, error)
	IsUserBookedSchedule(userID, scheduleID string) (bool, error)
	GetBookingsByUserID(userID string) ([]models.Booking, error)
	FindByUserAndSchedule(userID, scheduleID string) (*models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) GetBookingsByUserID(userID string) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.
		Preload("ClassSchedule").
		Preload("ClassSchedule.Class.Location").
		Preload("ClassSchedule.Instructor.User.Profile").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&bookings).Error

	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingByID(id string) (*models.Booking, error) {
	var booking models.Booking
	if err := r.db.Preload("ClassSchedule.Instructor.User.Profile").First(&booking, "id = ?", id).Error; err != nil {
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
