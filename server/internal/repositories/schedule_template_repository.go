package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type ScheduleTemplateRepository interface {
	DeleteTemplate(id string) error
	CreateTemplate(template *models.ScheduleTemplate) error
	GetActiveTemplates() ([]models.ScheduleTemplate, error)
	CreateRecurrenceRule(rule *models.RecurrenceRule) error
	UpdateTemplate(template *models.ScheduleTemplate) error
	GetRecurrenceRuleByTemplateID(templateID string) (*models.RecurrenceRule, error)
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

func (r *scheduleTemplateRepository) CreateRecurrenceRule(rule *models.RecurrenceRule) error {
	return r.db.Create(rule).Error
}

func (r *scheduleTemplateRepository) UpdateTemplate(template *models.ScheduleTemplate) error {
	return r.db.Save(template).Error
}

func (r *scheduleTemplateRepository) DeleteTemplate(id string) error {
	return r.db.Delete(&models.ScheduleTemplate{}, "id = ?", id).Error
}

func (r *scheduleTemplateRepository) GetRecurrenceRuleByTemplateID(templateID string) (*models.RecurrenceRule, error) {
	var rule models.RecurrenceRule
	err := r.db.Where("template_id = ?", templateID).First(&rule).Error
	return &rule, err
}
