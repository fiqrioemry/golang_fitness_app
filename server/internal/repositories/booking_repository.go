package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	GetBookingsByUserID(userID string) ([]models.Booking, error)
	CountBookingBySchedule(scheduleID string) (int64, error)
	GetBookingByID(id string) (*models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingsByUserID(userID string) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Preload("ClassSchedule.Class.Location").
		Preload("ClassSchedule.Instructor.User.Profile").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&bookings).Error

	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *bookingRepository) CountBookingBySchedule(scheduleID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).Where("class_schedule_id = ?", scheduleID).Count(&count).Error
	return count, err
}

func (r *bookingRepository) GetBookingByID(id string) (*models.Booking, error) {
	var booking models.Booking
	if err := r.db.Preload("ClassSchedule.Instructor.User.Profile").First(&booking, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
