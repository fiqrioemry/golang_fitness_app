package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ClassScheduleService interface {
	CreateClassSchedule(req dto.CreateClassScheduleRequest) error
	UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error
	DeleteClassSchedule(id string) error
	GetAllClassSchedules() ([]dto.ClassScheduleResponse, error)
	GetClassScheduleByID(id string) (*dto.ClassScheduleResponse, error)
}

type classScheduleService struct {
	repo      repositories.ClassScheduleRepository
	classRepo repositories.ClassRepository
}

func NewClassScheduleService(repo repositories.ClassScheduleRepository, classRepo repositories.ClassRepository) ClassScheduleService {
	return &classScheduleService{
		repo:      repo,
		classRepo: classRepo,
	}
}
func (s *classScheduleService) CreateClassSchedule(req dto.CreateClassScheduleRequest) error {
	_, err := s.classRepo.GetClassByID(req.ClassID)
	if err != nil {
		return err
	}

	dateOnly := req.Date.Format("2006-01-02")
	parsedDate, err := time.Parse("2006-01-02", dateOnly)
	if err != nil {
		return err
	}

	// Validasi: Instruktur tidak boleh punya kelas lain di waktu sama
	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}

	for _, schedule := range existingSchedules {
		if schedule.InstructorID == uuid.MustParse(req.InstructorID) &&
			schedule.Date.Equal(parsedDate) &&
			schedule.StartHour == req.StartHour &&
			schedule.StartMinute == req.StartMinute {
			return fmt.Errorf("instructor is already assigned to another class at the same time")
		}
	}

	schedule := models.ClassSchedule{
		ID:           uuid.New(),
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		Capacity:     req.Capacity,
		IsActive:     true,
		Color:        req.Color,
		Date:         parsedDate,
		StartHour:    req.StartHour,
		StartMinute:  req.StartMinute,
	}

	return s.repo.CreateClassSchedule(&schedule)
}

func (s *classScheduleService) UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return err
	}

	dateOnly := req.Date.Format("2006-01-02")
	parsedDate, err := time.Parse("2006-01-02", dateOnly)
	if err != nil {
		return err
	}

	schedule.Date = parsedDate
	schedule.StartHour = req.StartHour
	schedule.StartMinute = req.StartMinute
	if req.Capacity > 0 {
		schedule.Capacity = req.Capacity
	}

	if req.Color != "" {
		schedule.Color = req.Color
	}

	return s.repo.UpdateClassSchedule(schedule)
}

func (s *classScheduleService) DeleteClassSchedule(id string) error {
	return s.repo.DeleteClassSchedule(id)
}

func (s *classScheduleService) GetAllClassSchedules() ([]dto.ClassScheduleResponse, error) {
	schedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleResponse
	for _, s := range schedules {
		result = append(result, dto.ClassScheduleResponse{
			ID:           s.ID.String(),
			ClassID:      s.ClassID.String(),
			ClassTitle:   s.Class.Title,
			Category:     s.Class.Category.Name,
			InstructorID: s.InstructorID.String(),
			Color:        s.Color,
			Instructor:   s.Instructor.User.Profile.Fullname,
			Date:         s.Date,
			StartHour:    s.StartHour,
			StartMinute:  s.StartMinute,
			Capacity:     s.Capacity,
			BookedCount:  s.BookedCount,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetClassScheduleByID(id string) (*dto.ClassScheduleResponse, error) {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ClassScheduleResponse{
		ID:           schedule.ID.String(),
		ClassID:      schedule.ClassID.String(),
		ClassTitle:   schedule.Class.Title,
		Category:     schedule.Class.Category.Name,
		InstructorID: schedule.InstructorID.String(),
		Instructor:   schedule.Instructor.User.Profile.Fullname,
		Date:         schedule.Date,
		Color:        schedule.Color,
		StartHour:    schedule.StartHour,
		StartMinute:  schedule.StartMinute,
		Capacity:     schedule.Capacity,
		BookedCount:  schedule.BookedCount,
	}, nil
}
