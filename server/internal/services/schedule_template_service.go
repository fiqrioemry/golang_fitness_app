package services

import (
	"encoding/json"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
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
		DayOfWeeks:   utils.IntSliceToJSON([]int{req.DayOfWeek}),
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

	today := time.Now().Truncate(24 * time.Hour)

	for _, template := range templates {
		// Ambil rule jika recurring
		rule, err := s.templateRepo.GetRecurrenceRuleByTemplateID(template.ID.String())
		if err != nil || rule.Frequency != "recurring" {
			continue
		}

		// Parse DayOfWeeks
		var days []int
		if err := json.Unmarshal(template.DayOfWeeks, &days); err != nil {
			continue
		}

		// Hitung end date
		endDate := today.AddDate(0, 1, 0)
		if rule.EndType == "until" && rule.EndDate != nil {
			if rule.EndDate.Before(today) {
				continue
			}
			endDate = rule.EndDate.Truncate(24 * time.Hour)
		}

		// Loop semua hari dari today sampai endDate
		for d := today; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			if !utils.IsDayMatched(int(d.Weekday()), days) {
				continue
			}

			// Cek duplicate
			existing, err := s.classScheduleRepo.GetClassSchedulesWithFilter(dto.ClassScheduleQueryParam{
				StartDate: d.Format("2006-01-02"),
				EndDate:   d.Format("2006-01-02"),
			})
			if err != nil {
				continue
			}

			conflict := false
			for _, e := range existing {
				if e.ClassID == template.ClassID &&
					e.StartHour == template.StartHour &&
					e.StartMinute == template.StartMinute {
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
				Date:         d,
				StartHour:    template.StartHour,
				StartMinute:  template.StartMinute,
			}
			_ = s.classScheduleRepo.CreateClassSchedule(&schedule)
		}
	}

	return nil
}

func (s *scheduleTemplateService) CreateRecurringScheduleTemplate(req dto.CreateRecurringScheduleTemplateRequest) error {
	template := models.ScheduleTemplate{
		ID:           uuid.New(),
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		DayOfWeeks:   utils.IntSliceToJSON([]int{req.DayOfWeek}),
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
		DayOfWeeks:   utils.IntSliceToJSON([]int{req.DayOfWeek}),
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
