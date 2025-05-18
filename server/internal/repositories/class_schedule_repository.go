package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassScheduleRepository interface {
	DeleteClassSchedule(id string) error
	IncrementBooked(scheduleID uuid.UUID) error
	GetClassSchedules() ([]models.ClassSchedule, error)
	GetClassByID(id uuid.UUID) (*models.Class, error)
	HasActiveBooking(scheduleID uuid.UUID) (bool, error)
	CreateClassSchedule(schedule *models.ClassSchedule) error
	UpdateClassSchedule(schedule *models.ClassSchedule) error
	GetClassScheduleByID(id string) (*models.ClassSchedule, error)
	GetInstructorWithProfileByID(id uuid.UUID) (*models.Instructor, error)
	GetClassSchedulesWithFilter(filter dto.ClassScheduleQueryParam) ([]models.ClassSchedule, error)
}

type classScheduleRepository struct {
	db *gorm.DB
}

func NewClassScheduleRepository(db *gorm.DB) ClassScheduleRepository {
	return &classScheduleRepository{db}
}

func (r *classScheduleRepository) CreateClassSchedule(schedule *models.ClassSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *classScheduleRepository) UpdateClassSchedule(schedule *models.ClassSchedule) error {
	return r.db.Save(schedule).Error
}

func (r *classScheduleRepository) DeleteClassSchedule(id string) error {
	return r.db.Delete(&models.ClassSchedule{}, "id = ?", id).Error
}

func (r *classScheduleRepository) GetClassScheduleByID(id string) (*models.ClassSchedule, error) {
	var schedule models.ClassSchedule
	if err := r.db.Preload("Class", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).
		First(&schedule, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *classScheduleRepository) GetClassSchedules() ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	if err := r.db.
		Preload("Class", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Order("date asc").
		Order("start_hour asc").
		Order("start_minute asc").
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *classScheduleRepository) GetClassSchedulesWithFilter(filter dto.ClassScheduleQueryParam) ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	db := r.db.
		Preload("Class", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Order("start_hour asc").
		Order("start_minute asc")

	if filter.StartDate != "" {
		if start, err := time.Parse("2006-01-02", filter.StartDate); err == nil {
			db = db.Where("date >= ?", start)
		}
	}
	if filter.EndDate != "" {
		if end, err := time.Parse("2006-01-02", filter.EndDate); err == nil {
			db = db.Where("date <= ?", end)
		}
	}
	if filter.CategoryID != "" {
		db = db.Joins("JOIN classes ON classes.id = class_schedules.class_id").
			Where("classes.category_id = ?", filter.CategoryID)
	}

	if err := db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *classScheduleRepository) IncrementBooked(scheduleID uuid.UUID) error {
	return r.db.Model(&models.ClassSchedule{}).
		Where("id = ?", scheduleID).
		Update("booked", gorm.Expr("booked + 1")).Error
}
func (r *classScheduleRepository) HasActiveBooking(scheduleID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Booking{}).
		Where("class_schedule_id = ? AND status = ?", scheduleID, "booked").
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *classScheduleRepository) GetClassByID(id uuid.UUID) (*models.Class, error) {
	var class models.Class
	if err := r.db.Unscoped().First(&class, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *classScheduleRepository) GetInstructorWithProfileByID(id uuid.UUID) (*models.Instructor, error) {
	var instructor models.Instructor
	if err := r.db.
		Preload("User.Profile", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		First(&instructor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &instructor, nil
}
