package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type ScheduleTemplateRepository interface {
	CreateTemplate(template *models.ScheduleTemplate) error
	GetActiveTemplates() ([]models.ScheduleTemplate, error)
}

type scheduleTemplateRepository struct {
	db *gorm.DB
}

func NewScheduleTemplateRepository(db *gorm.DB) ScheduleTemplateRepository {
	return &scheduleTemplateRepository{db}
}

func (r *scheduleTemplateRepository) CreateTemplate(template *models.ScheduleTemplate) error {
	return r.db.Create(template).Error
}

func (r *scheduleTemplateRepository) GetActiveTemplates() ([]models.ScheduleTemplate, error) {
	var templates []models.ScheduleTemplate
	if err := r.db.Where("is_active = ?", true).Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}
