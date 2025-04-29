package services

import (
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
	startTime, _ := time.Parse(time.RFC3339, req.StartTime)

	class, err := s.classRepo.GetClassByID(req.ClassID)
	if err != nil {
		return err
	}

	endTime := startTime.Add(time.Minute * time.Duration(class.Duration))

	schedule := models.ClassSchedule{
		ClassID:      uuid.MustParse(req.ClassID),
		InstructorID: uuid.MustParse(req.InstructorID),
		StartTime:    startTime,
		EndTime:      endTime,
		Capacity:     req.Capacity,
	}

	return s.repo.CreateClassSchedule(&schedule)
}

func (s *classScheduleService) UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return err
	}

	if req.StartTime != "" {
		startTime, _ := time.Parse(time.RFC3339, req.StartTime)
		schedule.StartTime = startTime
	}
	if req.EndTime != "" {
		endTime, _ := time.Parse(time.RFC3339, req.EndTime)
		schedule.EndTime = endTime
	}
	if req.Capacity != 0 {
		schedule.Capacity = req.Capacity
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
			ID:             s.ID.String(),
			ClassID:        s.ClassID.String(),
			ClassTitle:     s.Class.Title,
			InstructorID:   s.InstructorID.String(),
			InstructorName: s.Instructor.User.Profile.Fullname,
			StartTime:      s.StartTime.Format(time.RFC3339),
			EndTime:        s.EndTime.Format(time.RFC3339),
			Capacity:       s.Capacity,
			BookedCount:    s.BookedCount,
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
		ID:             schedule.ID.String(),
		ClassID:        schedule.ClassID.String(),
		ClassTitle:     schedule.Class.Title,
		InstructorID:   schedule.InstructorID.String(),
		InstructorName: schedule.Instructor.User.Profile.Fullname,
		StartTime:      schedule.StartTime.Format(time.RFC3339),
		EndTime:        schedule.EndTime.Format(time.RFC3339),
		Capacity:       schedule.Capacity,
		BookedCount:    schedule.BookedCount,
	}, nil
}
