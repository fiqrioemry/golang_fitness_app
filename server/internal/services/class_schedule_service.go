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
	GetClassScheduleByID(id string) (*dto.ClassScheduleDetailResponse, error)
	GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error)
	GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleWithBookingStatusResponse, error)
}

type classScheduleService struct {
	repo            repositories.ClassScheduleRepository
	classRepo       repositories.ClassRepository
	packageRepo     repositories.PackageRepository
	userPackageRepo repositories.UserPackageRepository
	bookingRepo     repositories.BookingRepository
}

func NewClassScheduleService(repo repositories.ClassScheduleRepository, classRepo repositories.ClassRepository, packageRepo repositories.PackageRepository,
	userPackageRepo repositories.UserPackageRepository, bookingRepo repositories.BookingRepository) ClassScheduleService {
	return &classScheduleService{
		repo:            repo,
		classRepo:       classRepo,
		packageRepo:     packageRepo,
		userPackageRepo: userPackageRepo,
		bookingRepo:     bookingRepo,
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

	newStart := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), req.StartHour, req.StartMinute, 0, 0, time.Local)
	newEnd := newStart.Add(time.Hour)

	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}

	for _, schedule := range existingSchedules {
		if schedule.InstructorID == uuid.MustParse(req.InstructorID) && schedule.Date.Equal(parsedDate) {
			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(), schedule.StartHour, schedule.StartMinute, 0, 0, time.Local)
			existEnd := existStart.Add(time.Hour)

			if newStart.Before(existEnd) && existStart.Before(newEnd) {
				return fmt.Errorf("instructor already booked at this time")
			}
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
			BookedCount:  s.Booked,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetClassScheduleByID(id string) (*dto.ClassScheduleDetailResponse, error) {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return nil, err
	}

	packages, err := s.packageRepo.GetPackagesByClassID(schedule.ClassID.String())
	if err != nil {
		return nil, err
	}

	var pkgResponses []dto.PackageResponse
	for _, p := range packages {
		pkgResponses = append(pkgResponses, dto.PackageResponse{
			ID:    p.ID.String(),
			Name:  p.Name,
			Price: p.Price,
		})
	}

	return &dto.ClassScheduleDetailResponse{
		ClassScheduleResponse: dto.ClassScheduleResponse{
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
			BookedCount:  schedule.Booked,
		},
		Packages: pkgResponses,
	}, nil
}

func (s *classScheduleService) GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error) {
	schedules, err := s.repo.GetClassSchedulesWithFilter(filter)
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
			Instructor:   s.Instructor.User.Profile.Fullname,
			Date:         s.Date,
			Color:        s.Color,
			StartHour:    s.StartHour,
			StartMinute:  s.StartMinute,
			Capacity:     s.Capacity,
			BookedCount:  s.Booked,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleWithBookingStatusResponse, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	schedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleWithBookingStatusResponse
	for _, schedule := range schedules {
		isBooked, _ := s.bookingRepo.IsUserBookedSchedule(parsedUserID, schedule.ID)

		result = append(result, dto.ClassScheduleWithBookingStatusResponse{
			ClassScheduleResponse: dto.ClassScheduleResponse{
				ID:           schedule.ID.String(),
				ClassID:      schedule.ClassID.String(),
				ClassTitle:   schedule.Class.Title,
				Category:     schedule.Class.Category.Name,
				InstructorID: schedule.InstructorID.String(),
				Color:        schedule.Color,
				Instructor:   schedule.Instructor.User.Profile.Fullname,
				Date:         schedule.Date,
				StartHour:    schedule.StartHour,
				StartMinute:  schedule.StartMinute,
				Capacity:     schedule.Capacity,
				BookedCount:  schedule.Booked,
			},
			IsBooked: isBooked,
		})
	}

	return result, nil
}
