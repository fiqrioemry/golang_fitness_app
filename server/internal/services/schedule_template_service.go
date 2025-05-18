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
			EndDate:      t.EndDate.Format(time.RFC3339),
			CreatedAt:    t.CreatedAt.Format(time.RFC3339),
		}
		result = append(result, resp)
	}
	return result, nil
}

func (s *scheduleTemplateService) CreateScheduleTemplate(req dto.CreateScheduleTemplateRequest) error {
	now := time.Now().In(time.Local)
	startTime := time.Date(now.Year(), now.Month(), now.Day(), req.StartHour, req.StartMinute, 0, 0, time.Local)
	endTime := startTime.Add(time.Hour)

	instructorID := uuid.MustParse(req.InstructorID)

	// Cek konflik dengan jadwal aktual (class schedules)
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
			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
				schedule.StartHour, schedule.StartMinute, 0, 0, time.Local)
			existEnd := existStart.Add(time.Hour)

			// Cek overlap waktu
			if startTime.Before(existEnd) && existStart.Before(endTime) {
				return fmt.Errorf("instructor already booked on %s at %02d:%02d (from actual schedule)",
					schedule.Date.Weekday().String(), req.StartHour, req.StartMinute)
			}
		}
	}

	// Cek konflik dengan template lain
	existingTemplates, err := s.templateRepo.GetAllTemplates()
	if err != nil {
		return err
	}

	for _, tpl := range existingTemplates {
		if tpl.InstructorID != instructorID {
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
				tplStart := time.Date(now.Year(), now.Month(), now.Day(), tpl.StartHour, tpl.StartMinute, 0, 0, time.Local)
				tplEnd := tplStart.Add(time.Hour)

				if startTime.Before(tplEnd) && tplStart.Before(endTime) {
					return fmt.Errorf("instructor already booked on day %d at %02d:%02d (from recurring template)",
						day, req.StartHour, req.StartMinute)
				}
			}
		}
	}

	// Simpan Template Baru
	template := models.ScheduleTemplate{
		ID:           uuid.New(),
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: instructorID,
		DayOfWeeks:   utils.IntSliceToJSON(req.DayOfWeeks),
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
		Capacity:     req.Capacity,
		IsActive:     false,
		Color:        req.Color,
		EndDate:      req.EndDate,
	}

	return s.templateRepo.CreateTemplate(&template)
}

func (s *scheduleTemplateService) UpdateScheduleTemplate(id string, req dto.UpdateScheduleTemplateRequest) error {
	template, err := s.templateRepo.GetTemplateByID(id)
	if err != nil {
		return err
	}

	if template.IsActive {
		return fmt.Errorf("cannot update an active template, please stop it first")
	}

	now := time.Now().In(time.Local)
	startTime := time.Date(now.Year(), now.Month(), now.Day(), req.StartHour, req.StartMinute, 0, 0, time.Local)
	endTime := startTime.Add(time.Hour)

	instructorID := uuid.MustParse(req.InstructorID)

	// 🔍 Cek konflik dengan jadwal aktual
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

			// skip jika itu adalah jadwal yang sama (sudah berasal dari template ini)
			if schedule.ClassID == template.ClassID &&
				schedule.InstructorID == template.InstructorID &&
				schedule.StartHour == template.StartHour &&
				schedule.StartMinute == template.StartMinute {
				continue
			}

			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
				schedule.StartHour, schedule.StartMinute, 0, 0, time.Local)
			existEnd := existStart.Add(time.Hour)

			if startTime.Before(existEnd) && existStart.Before(endTime) {
				return fmt.Errorf("instructor already booked on day %d at %02d:%02d (from actual schedule)",
					day, req.StartHour, req.StartMinute)
			}
		}
	}

	// 🔍 Cek konflik dengan template lain
	existingTemplates, err := s.templateRepo.GetAllTemplates()
	if err != nil {
		return err
	}

	for _, tpl := range existingTemplates {
		if tpl.ID == template.ID || tpl.InstructorID != instructorID {
			continue // skip diri sendiri atau instruktur berbeda
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
				tplStart := time.Date(now.Year(), now.Month(), now.Day(), tpl.StartHour, tpl.StartMinute, 0, 0, time.Local)
				tplEnd := tplStart.Add(time.Hour)

				if startTime.Before(tplEnd) && tplStart.Before(endTime) {
					return fmt.Errorf("instructor already booked on day %d at %02d:%02d (from another template)",
						day, req.StartHour, req.StartMinute)
				}
			}
		}
	}

	// ✅ Update data template
	template.ClassID = uuid.MustParse(req.ClassID)
	template.InstructorID = instructorID
	template.DayOfWeeks = utils.IntSliceToJSON(req.DayOfWeeks)
	template.StartHour = req.StartHour
	template.StartMinute = req.StartMinute
	template.Capacity = req.Capacity
	template.EndDate = req.EndDate

	if err := s.templateRepo.UpdateTemplate(template); err != nil {
		return err
	}

	return nil
}

func (s *scheduleTemplateService) DeleteTemplate(id string) error {
	return s.templateRepo.DeleteTemplate(id)
}

func (s *scheduleTemplateService) AutoGenerateSchedules() error {
	templates, err := s.templateRepo.GetActiveTemplates()
	if err != nil {
		return fmt.Errorf("failed to fetch templates: %w", err)
	}

	if len(templates) == 0 {
		return fmt.Errorf("no active schedule templates found")
	}

	today := time.Now().Truncate(24 * time.Hour)
	oneMonthAhead := today.AddDate(0, 1, 0)

	var anySuccess bool
	var errs []string

	for _, template := range templates {
		var days []int
		if err := json.Unmarshal(template.DayOfWeeks, &days); err != nil {
			errs = append(errs, fmt.Sprintf("template %s: failed to unmarshal days", template.ID))
			continue
		}

		for d := today; !d.After(oneMonthAhead); d = d.AddDate(0, 0, 1) {
			if !utils.IsDayMatched(int(d.Weekday()), days) {
				continue
			}

			schedule := models.ClassSchedule{
				ID:           uuid.New(),
				ClassID:      template.ClassID,
				InstructorID: template.InstructorID,
				Capacity:     template.Capacity,
				IsActive:     true,
				Date:         d,
				Color:        template.Color,
				StartHour:    template.StartHour,
				StartMinute:  template.StartMinute,
			}

			if err := s.classScheduleRepo.CreateClassSchedule(&schedule); err != nil {
				errs = append(errs, fmt.Sprintf("template %s: failed to create schedule on %s", template.ID, d.Format("2006-01-02")))
				continue
			}

			anySuccess = true
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
