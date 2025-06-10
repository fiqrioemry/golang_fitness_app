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
	GenerateScheduleByTemplateID(templateID string) error
	GetAllTemplates() ([]dto.ScheduleTemplateResponse, error)
	CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) (string, error)
	UpdateScheduleTemplate(id string, req dto.UpdateScheduleTemplateRequest) error
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
		resp := dto.ScheduleTemplateResponse{
			ID:             t.ID.String(),
			ClassID:        t.ClassID.String(),
			ClassName:      t.Class.Title,
			InstructorID:   t.InstructorID.String(),
			InstructorName: t.Instructor.User.Profile.Fullname,
			DayOfWeeks:     days,
			StartHour:      t.StartHour,
			StartMinute:    t.StartMinute,
			Capacity:       t.Capacity,
			IsActive:       t.IsActive,
			EndDate:        t.EndDate.Format("2006-01-02"),
			CreatedAt:      t.CreatedAt.Format("2006-01-02"),
		}
		result = append(result, resp)
	}
	return result, nil
}

func containsInt(list []int, target int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func (s *scheduleTemplateService) CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) (string, error) {
	now := time.Now().UTC()
	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return "", err
	}

	if endDate.Before(now) {
		return "", fmt.Errorf("end date must be in the future")
	}

	instructorID := uuid.MustParse(req.InstructorID)

	existingSchedules, err := s.classScheduleRepo.GetClassSchedules()
	if err != nil {
		return "", err
	}

	existingTemplates, err := s.templateRepo.GetAllTemplates()
	if err != nil {
		return "", err
	}

	simulationEnd := endDate

	for date := now; !date.After(simulationEnd); date = date.AddDate(0, 0, 1) {
		weekday := int(date.Weekday())
		if !utils.IsDayMatched(weekday, req.DayOfWeeks) {
			continue
		}

		simulatedStart := time.Date(date.Year(), date.Month(), date.Day(), req.StartHour, req.StartMinute, 0, 0, time.UTC)
		simulatedEnd := simulatedStart.Add(time.Hour)

		for _, schedule := range existingSchedules {
			if schedule.InstructorID != instructorID {
				continue
			}
			if schedule.Date.Year() != date.Year() || schedule.Date.Month() != date.Month() || schedule.Date.Day() != date.Day() {
				continue
			}

			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(), schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC)
			existEnd := existStart.Add(time.Hour)

			if simulatedStart.Before(existEnd) && existStart.Before(simulatedEnd) {
				return "", fmt.Errorf("instructor %s is already booked on %s at %02d:%02d (actual schedule)", schedule.InstructorName, date.Format("2006-01-02"), req.StartHour, req.StartMinute)
			}
		}

		for _, tpl := range existingTemplates {
			if tpl.InstructorID != instructorID {
				continue
			}
			var tplDays []int
			_ = json.Unmarshal(tpl.DayOfWeeks, &tplDays)
			if !utils.IsDayMatched(weekday, tplDays) {
				continue
			}

			tplStart := time.Date(date.Year(), date.Month(), date.Day(), tpl.StartHour, tpl.StartMinute, 0, 0, time.UTC)
			tplEnd := tplStart.Add(time.Hour)

			if simulatedStart.Before(tplEnd) && tplStart.Before(simulatedEnd) {
				return "", fmt.Errorf("instructor %s is already booked on %s at %02d:%02d (template)", tpl.InstructorName, date.Format("2006-01-02"), req.StartHour, req.StartMinute)
			}
		}
	}

	class, err := s.classScheduleRepo.GetClassByID(uuid.MustParse(req.ClassID))
	if err != nil {
		return "", fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.classScheduleRepo.GetInstructorWithProfileByID(instructorID)
	if err != nil {
		return "", fmt.Errorf("instructor not found: %w", err)
	}

	template := models.ScheduleTemplate{
		ID:             uuid.New(),
		ClassID:        class.ID,
		ClassName:      class.Title,
		ClassImage:     class.Image,
		InstructorID:   instructor.ID,
		Location:       class.Location.Name,
		InstructorName: instructor.User.Profile.Fullname,
		DayOfWeeks:     utils.IntSliceToJSON(req.DayOfWeeks),
		StartHour:      req.StartHour,
		StartMinute:    req.StartMinute,
		Capacity:       req.Capacity,
		IsActive:       false,
		Color:          req.Color,
		EndDate:        endDate,
	}

	err = s.templateRepo.CreateTemplate(&template)
	if err != nil {
		return "", fmt.Errorf("failed to create template: %w", err)
	}

	return template.ID.String(), nil
}

func (s *scheduleTemplateService) UpdateScheduleTemplate(id string, req dto.UpdateScheduleTemplateRequest) error {
	template, err := s.templateRepo.GetTemplateByID(id)
	if err != nil {
		return err
	}

	if template.IsActive {
		return fmt.Errorf("cannot update an active template, please stop it first")
	}

	now := time.Now().UTC()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), req.StartHour, req.StartMinute, 0, 0, time.UTC)
	endTime := startTime.Add(time.Hour)

	instructorID := uuid.MustParse(req.InstructorID)

	existingSchedules, err := s.classScheduleRepo.GetClassSchedules()
	if err != nil {
		return err
	}

	for _, schedule := range existingSchedules {
		if schedule.InstructorID != instructorID {
			continue
		}

		for _, day := range req.DayOfWeeks {
			if int(schedule.Date.Weekday()) != day {
				continue
			}

			if schedule.ClassID == template.ClassID &&
				schedule.InstructorID == template.InstructorID &&
				schedule.StartHour == template.StartHour &&
				schedule.StartMinute == template.StartMinute {
				continue
			}

			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
				schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC)
			existEnd := existStart.Add(time.Hour)

			if startTime.Before(existEnd) && existStart.Before(endTime) {
				return fmt.Errorf("instructor already booked on day %d at %02d:%02d (from actual schedule)",
					day, req.StartHour, req.StartMinute)
			}
		}
	}

	existingTemplates, err := s.templateRepo.GetAllTemplates()
	if err != nil {
		return err
	}

	for _, tpl := range existingTemplates {
		if tpl.ID == template.ID || tpl.InstructorID != instructorID {
			continue
		}

		var tplDays []int
		if err := json.Unmarshal(tpl.DayOfWeeks, &tplDays); err != nil {
			continue
		}

		for _, day := range req.DayOfWeeks {
			for _, tplDay := range tplDays {
				if tplDay != day {
					continue
				}
				tplStart := time.Date(now.Year(), now.Month(), now.Day(), tpl.StartHour, tpl.StartMinute, 0, 0, time.UTC)
				tplEnd := tplStart.Add(time.Hour)

				if startTime.Before(tplEnd) && tplStart.Before(endTime) {
					return fmt.Errorf("instructor already booked on day %d at %02d:%02d (from another template)",
						day, req.StartHour, req.StartMinute)
				}
			}
		}
	}

	class, err := s.classScheduleRepo.GetClassByID(uuid.MustParse(req.ClassID))
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.classScheduleRepo.GetInstructorWithProfileByID(instructorID)
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
	}

	endDate, err := utils.ParseDate(req.EndDate)
	if err != nil {
		return err
	}

	template.ClassID = class.ID
	template.ClassName = class.Title
	template.InstructorID = instructor.ID
	template.InstructorName = instructor.User.Profile.Fullname
	template.DayOfWeeks = utils.IntSliceToJSON(req.DayOfWeeks)
	template.StartHour = req.StartHour
	template.StartMinute = req.StartMinute
	template.Capacity = req.Capacity
	template.EndDate = endDate

	if err := s.templateRepo.UpdateTemplate(template); err != nil {
		return err
	}

	return nil
}

func (s *scheduleTemplateService) DeleteTemplate(id string) error {
	return s.templateRepo.DeleteTemplate(id)
}

func (s *scheduleTemplateService) GenerateScheduleByTemplateID(templateID string) error {
	template, err := s.templateRepo.GetTemplateByID(templateID)
	if err != nil {
		return fmt.Errorf("failed to fetch template: %w", err)
	}

	if !template.IsActive {
		return fmt.Errorf("template is not active")
	}

	var days []int
	if err := json.Unmarshal(template.DayOfWeeks, &days); err != nil {
		return fmt.Errorf("failed to parse days of week: %w", err)
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	end := today.AddDate(0, 1, 0)

	var hasSuccess bool
	var errors []string

	for d := today; !d.After(end); d = d.AddDate(0, 0, 1) {
		if !utils.IsDayMatched(int(d.Weekday()), days) {
			continue
		}

		conflict, err := s.checkInstructorConflict(d, template.InstructorID, template.StartHour, template.StartMinute)
		if err != nil {
			return err
		}
		if conflict {
			errors = append(errors, fmt.Sprintf("conflict on %s at %02d:%02d", d.Format("2006-01-02"), template.StartHour, template.StartMinute))
			continue
		}

		schedule := models.ClassSchedule{
			ID:             uuid.New(),
			ClassID:        template.ClassID,
			ClassName:      template.ClassName,
			ClassImage:     template.ClassImage,
			InstructorID:   template.InstructorID,
			InstructorName: template.InstructorName,
			Location:       template.Location,
			Capacity:       template.Capacity,
			Color:          template.Color,
			Date:           d,
			StartHour:      template.StartHour,
			StartMinute:    template.StartMinute,
		}

		if err := s.classScheduleRepo.CreateClassSchedule(&schedule); err != nil {
			errors = append(errors, fmt.Sprintf("failed to create on %s", d.Format("2006-01-02")))
			continue
		}

		hasSuccess = true
	}

	if !hasSuccess {
		return fmt.Errorf("failed to generate any schedule: %v", errors)
	}

	now := time.Now().UTC()
	template.LastGeneratedAt = &now
	if err := s.templateRepo.UpdateTemplate(template); err != nil {
		return fmt.Errorf("schedule generated, but failed to update LastGeneratedAt: %w", err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("partial success: %v", errors)
	}

	return nil
}

func (s *scheduleTemplateService) checkInstructorConflict(date time.Time, instructorID string, hour, minute int) (bool, error) {
	schedules, err := s.classScheduleRepo.GetClassSchedules()
	if err != nil {
		return false, err
	}

	newStart := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.UTC)
	newEnd := newStart.Add(time.Hour)

	for _, s := range schedules {
		if s.InstructorID != instructorID {
			continue
		}
		if s.Date.Year() != date.Year() || s.Date.Month() != date.Month() || s.Date.Day() != date.Day() {
			continue
		}
		existStart := time.Date(s.Date.Year(), s.Date.Month(), s.Date.Day(), s.StartHour, s.StartMinute, 0, 0, time.UTC)
		existEnd := existStart.Add(time.Hour)

		if newStart.Before(existEnd) && existStart.Before(newEnd) {
			return true, nil
		}
	}

	return false, nil
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

// for cron job
func (s *scheduleTemplateService) AutoGenerateSchedules() error {
	templates, err := s.templateRepo.GetActiveTemplates()
	if err != nil {
		return fmt.Errorf("failed to fetch templates: %w", err)
	}

	if len(templates) == 0 {
		return fmt.Errorf("no active schedule templates found")
	}

	today := time.Now().Truncate(24 * time.Hour)
	var anySuccess bool
	var errs []string

	for _, template := range templates {
		if template.LastGeneratedAt != nil {
			if today.Sub(template.LastGeneratedAt.Truncate(24*time.Hour)) < 30*24*time.Hour {
				continue
			}
		}

		var days []int
		if err := json.Unmarshal(template.DayOfWeeks, &days); err != nil {
			errs = append(errs, fmt.Sprintf("template %s: failed to unmarshal days", template.ID))
			continue
		}

		end := today.AddDate(0, 0, 30)

		for d := today; !d.After(end); d = d.AddDate(0, 0, 1) {
			if !utils.IsDayMatched(int(d.Weekday()), days) {
				continue
			}

			schedule := models.ClassSchedule{
				ID:             uuid.New(),
				ClassID:        template.ClassID,
				ClassName:      template.Class.Title,
				ClassImage:     template.Class.Image,
				InstructorID:   template.InstructorID,
				InstructorName: template.InstructorName,
				Capacity:       template.Capacity,
				Date:           d,
				Color:          template.Color,
				StartHour:      template.StartHour,
				StartMinute:    template.StartMinute,
			}

			if err := s.classScheduleRepo.CreateClassSchedule(&schedule); err != nil {
				errs = append(errs, fmt.Sprintf("template %s: failed to create schedule on %s", template.ID, d.Format("2006-01-02")))
				continue
			}

			anySuccess = true
		}

		now := time.Now()
		template.LastGeneratedAt = &now
		if err := s.templateRepo.UpdateTemplate(&template); err != nil {
			errs = append(errs, fmt.Sprintf("template %s: failed to update lastGeneratedAt", template.ID))
		}
	}

	if len(errs) > 0 && !anySuccess {
		return fmt.Errorf("no schedules generated: %v", errs)
	}
	if len(errs) > 0 {
		return fmt.Errorf("some schedules generated, with errors: %v", errs)
	}

	return nil
}
