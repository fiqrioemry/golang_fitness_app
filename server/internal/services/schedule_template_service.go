package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ScheduleTemplateService interface {
	CreateRecurringScheduleTemplate(req dto.CreateRecurringScheduleTemplateRequest) error
	CreateTemplate(req dto.CreateScheduleTemplateRequest) error
	AutoGenerateSchedules() error
	DeleteTemplate(id string) error
	UpdateTemplate(id string, req dto.CreateScheduleTemplateRequest) error
}

type scheduleTemplateService struct {
	templateRepo      repositories.ScheduleTemplateRepository
	classRepo         repositories.ClassRepository
	classScheduleRepo repositories.ClassScheduleRepository
}

func NewScheduleTemplateService(
	templateRepo repositories.ScheduleTemplateRepository,
	classRepo repositories.ClassRepository,
	classScheduleRepo repositories.ClassScheduleRepository,
) ScheduleTemplateService {
	return &scheduleTemplateService{templateRepo, classRepo, classScheduleRepo}
}

func (s *scheduleTemplateService) CreateTemplate(req dto.CreateScheduleTemplateRequest) error {
	template := models.ScheduleTemplate{
		ID:           uuid.New(),
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		DayOfWeeks:   req.DayOfWeeks,
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
		Capacity:     req.Capacity,
		IsActive:     true,
	}

	return s.templateRepo.CreateTemplate(&template)
}

func (s *scheduleTemplateService) AutoGenerateSchedules() error {
	templates, err := s.templateRepo.GetActiveTemplates()
	if err != nil {
		return err
	}

	today := time.Now()

	for _, template := range templates {
		// Cek apakah template memiliki recurrence rule
		rule, err := s.templateRepo.GetRecurrenceRuleByTemplateID(template.ID.String())
		if err != nil {
			continue
		}

		// Skip jika rule bukan recurring
		if rule.Frequency != "recurring" {
			continue
		}

		// Hitung tanggal untuk jadwal berikutnya
		dayOffset := (int(template.DayOfWeek) - int(today.Weekday()) + 7) % 7
		scheduleDate := today.AddDate(0, 0, dayOffset)
		scheduleDate = time.Date(scheduleDate.Year(), scheduleDate.Month(), scheduleDate.Day(), 0, 0, 0, 0, time.UTC)

		// Cek endType dan endDate jika ada
		if rule.EndType == "until" && rule.EndDate != nil && scheduleDate.After(*rule.EndDate) {
			continue
		}

		// Cek apakah sudah ada jadwal yang sama (prevent duplicate)
		existings, err := s.classScheduleRepo.GetClassSchedules()
		if err != nil {
			continue
		}
		conflict := false
		for _, s := range existings {
			if s.ClassID == template.ClassID && s.Date.Equal(scheduleDate) && s.StartHour == template.StartHour && s.StartMinute == template.StartMinute {
				conflict = true
				break
			}
		}
		if conflict {
			continue
		}

		schedule := models.ClassSchedule{
			ID:           uuid.New(),
			ClassID:      template.ClassID,
			InstructorID: template.InstructorID,
			Capacity:     template.Capacity,
			IsActive:     true,
			Date:         scheduleDate,
			StartHour:    template.StartHour,
			StartMinute:  template.StartMinute,
		}
		s.classScheduleRepo.CreateClassSchedule(&schedule)
	}

	return nil
}
func nextWeekday(t time.Time, weekday time.Weekday) time.Time {
	daysUntil := (int(weekday) - int(t.Weekday()) + 7) % 7
	if daysUntil == 0 {
		daysUntil = 7
	}
	return t.AddDate(0, 0, daysUntil)
}

func (s *scheduleTemplateService) CreateRecurringScheduleTemplate(req dto.CreateRecurringScheduleTemplateRequest) error {
	template := models.ScheduleTemplate{
		ID:           uuid.New(),
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		DayOfWeek:    req.DayOfWeek,
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
		Capacity:     req.Capacity,
		IsActive:     true,
	}

	if err := s.templateRepo.CreateTemplate(&template); err != nil {
		return err
	}

	if req.Frequency == "recurring" {
		rule := models.RecurrenceRule{
			TemplateID: template.ID,
			Frequency:  req.Frequency,
			EndType:    req.EndType,
			EndDate:    req.EndDate,
		}
		return s.templateRepo.CreateRecurrenceRule(&rule)
	}

	return nil
}

// schedule_template_service.go
func (s *scheduleTemplateService) UpdateTemplate(id string, req dto.CreateScheduleTemplateRequest) error {
	templateID, _ := uuid.Parse(id)
	template := models.ScheduleTemplate{
		ID:           templateID,
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		DayOfWeek:    req.DayOfWeek,
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
		Capacity:     req.Capacity,
		IsActive:     true,
	}
	return s.templateRepo.UpdateTemplate(&template)
}

func (s *scheduleTemplateService) DeleteTemplate(id string) error {
	return s.templateRepo.DeleteTemplate(id)
}
