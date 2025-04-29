package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ScheduleTemplateService interface {
	CreateTemplate(req dto.CreateScheduleTemplateRequest) error
	AutoGenerateSchedules() error
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
		DayOfWeek:    req.DayOfWeek,
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

	for _, template := range templates {
		startDate := nextWeekday(time.Now(), time.Weekday(template.DayOfWeek))
		startTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(),
			template.StartHour, template.StartMinute, 0, 0, time.UTC)

		class, err := s.classRepo.GetClassByID(template.ClassID.String())
		if err != nil {
			continue
		}

		endTime := startTime.Add(time.Minute * time.Duration(class.Duration))

		schedule := models.ClassSchedule{
			ID:           uuid.New(),
			ClassID:      template.ClassID,
			InstructorID: template.InstructorID,
			StartTime:    startTime,
			EndTime:      endTime,
			Capacity:     template.Capacity,
			IsActive:     true,
		}

		s.classScheduleRepo.CreateClassSchedule(&schedule)
	}

	return nil
}

// Helper untuk cari tanggal weekday berikutnya
func nextWeekday(t time.Time, weekday time.Weekday) time.Time {
	daysUntil := (int(weekday) - int(t.Weekday()) + 7) % 7
	if daysUntil == 0 {
		daysUntil = 7
	}
	return t.AddDate(0, 0, daysUntil)
}
