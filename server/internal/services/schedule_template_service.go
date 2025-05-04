package services

import (
	"encoding/json"
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ScheduleTemplateService interface {
	AutoGenerateSchedules() error
	RunTemplate(id string) error
	StopTemplate(id string) error
	DeleteTemplate(id string) error
	GetAllTemplates() ([]dto.ScheduleTemplateResponse, error)
	CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) error
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

func (s *scheduleTemplateService) GetAllTemplates() ([]dto.ScheduleTemplateResponse, error) {
	templates, err := s.templateRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.ScheduleTemplateResponse
	for _, t := range templates {
		var days []int
		_ = json.Unmarshal(t.DayOfWeeks, &days)
		rule, _ := s.templateRepo.GetRecurrenceRuleByTemplateID(t.ID.String())
		resp := dto.ScheduleTemplateResponse{
			ID:           t.ID,
			ClassID:      t.ClassID,
			ClassName:    t.Class.Title,
			InstructorID: t.InstructorID,
			Instructor:   t.Instructor.User.Profile.Fullname,
			DayOfWeeks:   days,
			StartHour:    t.StartHour,
			StartMinute:  t.StartMinute,
			Capacity:     t.Capacity,
			IsActive:     t.IsActive,
			Frequency:    rule.Frequency,
			EndType:      rule.EndType,
			EndDate:      rule.EndDate,
		}
		result = append(result, resp)
	}
	return result, nil
}

func (s *scheduleTemplateService) CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) error {
	template := models.ScheduleTemplate{
		ID:           uuid.New(),
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		DayOfWeeks:   utils.IntSliceToJSON([]int{req.DayOfWeek}),
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
		Capacity:     req.Capacity,
		IsActive:     false,
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

func (s *scheduleTemplateService) UpdateTemplate(id string, req dto.CreateScheduleTemplateRequest) error {

	template, err := s.templateRepo.GetTemplateByID(id)
	if err != nil {
		return err
	}

	if template.IsActive {
		return fmt.Errorf("cannot update an active template, please stop it first")
	}

	template.ClassID = uuid.MustParse(req.ClassID)
	template.InstructorID = uuid.MustParse(req.InstructorID)
	template.DayOfWeeks = utils.IntSliceToJSON([]int{req.DayOfWeek})
	template.StartHour = req.StartHour
	template.StartMinute = req.StartMinute
	template.Capacity = req.Capacity

	return s.templateRepo.UpdateTemplate(template)
}

func (s *scheduleTemplateService) DeleteTemplate(id string) error {
	return s.templateRepo.DeleteTemplate(id)
}

func (s *scheduleTemplateService) AutoGenerateSchedules() error {
	templates, err := s.templateRepo.GetActiveTemplates()
	if err != nil {
		return err
	}

	today := time.Now().Truncate(24 * time.Hour)
	oneMonthAhead := today.AddDate(0, 1, 0)

	var anySuccess bool
	var errs []string

	for _, template := range templates {
		rule, err := s.templateRepo.GetRecurrenceRuleByTemplateID(template.ID.String())
		if err != nil || rule.Frequency != "recurring" {
			continue
		}

		var days []int
		if err := json.Unmarshal(template.DayOfWeeks, &days); err != nil {
			errs = append(errs, fmt.Sprintf("template %s: failed to unmarshal days", template.ID))
			continue
		}

		generateUntil := oneMonthAhead
		if rule.EndType == "until" && rule.EndDate != nil && rule.EndDate.Before(oneMonthAhead) {
			if rule.EndDate.Before(today) {
				continue
			}
			generateUntil = rule.EndDate.Truncate(24 * time.Hour)
		}

		for d := today; !d.After(generateUntil); d = d.AddDate(0, 0, 1) {
			if !utils.IsDayMatched(int(d.Weekday()), days) {
				continue
			}

			existing, err := s.classScheduleRepo.GetClassSchedulesWithFilter(dto.ClassScheduleQueryParam{
				StartDate: d.Format("2006-01-02"),
				EndDate:   d.Format("2006-01-02"),
			})
			if err != nil {
				errs = append(errs, fmt.Sprintf("template %s: failed to get schedules on %s", template.ID, d))
				continue
			}

			conflict := false
			newStart := time.Date(d.Year(), d.Month(), d.Day(), template.StartHour, template.StartMinute, 0, 0, time.Local)
			newEnd := newStart.Add(time.Hour)

			for _, e := range existing {
				if e.InstructorID != template.InstructorID {
					continue
				}
				existStart := time.Date(e.Date.Year(), e.Date.Month(), e.Date.Day(), e.StartHour, e.StartMinute, 0, 0, time.Local)
				existEnd := existStart.Add(time.Hour)

				if newStart.Before(existEnd) && existStart.Before(newEnd) {
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

			if err := s.classScheduleRepo.CreateClassSchedule(&schedule); err != nil {
				errs = append(errs, fmt.Sprintf("template %s: failed to create schedule on %s", template.ID, d))
				continue
			}

			anySuccess = true
		}
	}

	if !anySuccess {
		return fmt.Errorf("no schedules were successfully generated: %v", errs)
	}

	if len(errs) > 0 {
		return fmt.Errorf("partial success with some errors: %v", errs)
	}

	return nil
}

func (s *scheduleTemplateService) RunTemplate(id string) error {
	template, err := s.templateRepo.GetTemplateByID(id)
	if err != nil {
		return err
	}
	if template.IsActive {
		return fmt.Errorf("template is already active")
	}
	template.IsActive = true
	return s.templateRepo.UpdateTemplate(template)
}

func (s *scheduleTemplateService) StopTemplate(id string) error {
	template, err := s.templateRepo.GetTemplateByID(id)
	if err != nil {
		return err
	}
	if !template.IsActive {
		return fmt.Errorf("template is already inactive")
	}
	template.IsActive = false
	return s.templateRepo.UpdateTemplate(template)
}
