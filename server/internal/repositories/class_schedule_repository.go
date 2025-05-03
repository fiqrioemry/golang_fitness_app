package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type ClassScheduleRepository interface {
	DeleteClassSchedule(id string) error
	GetClassSchedules() ([]models.ClassSchedule, error)
	CreateClassSchedule(schedule *models.ClassSchedule) error
	UpdateClassSchedule(schedule *models.ClassSchedule) error
	GetClassScheduleByID(id string) (*models.ClassSchedule, error)
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
	if err := r.db.Preload("Class").Preload("Instructor.User.Profile").First(&schedule, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *classScheduleRepository) GetClassSchedules() ([]models.ClassSchedule, error) {
	var schedules []models.ClassSchedule
	if err := r.db.
		Preload("Class").
		Preload("Instructor.User.Profile").
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
		Preload("Class.Category").
		Preload("Instructor.User.Profile").
		Order("start_hour asc").
		Order("start_minute asc")

	if filter.StartDate != "" {
		if date, err := time.Parse("2006-01-02", filter.StartDate); err == nil {
			db = db.Where("start_time >= ?", date)
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
