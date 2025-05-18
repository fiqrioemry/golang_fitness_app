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
	DeleteClassSchedule(id string) error
	GetAllClassSchedules() ([]dto.ClassScheduleResponse, error)
	CreateClassSchedule(req dto.CreateClassScheduleRequest) error
	UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error
	GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error)
	GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error)
	GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleResponse, error)
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
	localDate := req.Date.In(time.Local)

	newStart := time.Date(
		localDate.Year(), localDate.Month(), localDate.Day(),
		req.StartHour, req.StartMinute, 0, 0, time.Local,
	)

	now := time.Now().In(time.Local)

	if newStart.Before(now) {
		return fmt.Errorf("cannot create schedule in the past")
	}

	newEnd := newStart.Add(time.Hour)

	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}
	for _, schedule := range existingSchedules {
		if schedule.InstructorID == uuid.MustParse(req.InstructorID) && schedule.Date.Equal(localDate) {
			existStart := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
				schedule.StartHour, schedule.StartMinute, 0, 0, time.Local)
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
		Date:         localDate,
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

	instructorID := uuid.MustParse(req.InstructorID)
	parsedDate := req.Date.In(time.Local)

	newStart := time.Date(
		parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		req.StartHour, req.StartMinute, 0, 0, time.Local,
	)

	now := time.Now().In(time.Local)

	if newStart.Before(now) {
		return fmt.Errorf("cannot update schedule to a past time")
	}

	newEnd := newStart.Add(time.Hour)

	existingSchedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return err
	}

	for _, other := range existingSchedules {
		if other.ID == schedule.ID {
			continue
		}
		if other.InstructorID == instructorID && other.Date.Equal(parsedDate) {
			existStart := time.Date(other.Date.Year(), other.Date.Month(), other.Date.Day(),
				other.StartHour, other.StartMinute, 0, 0, time.Local)
			existEnd := existStart.Add(time.Hour)

			if newStart.Before(existEnd) && existStart.Before(newEnd) {
				return fmt.Errorf("instructor already booked at this time")
			}
		}
	}

	if req.Capacity < schedule.Booked {
		return fmt.Errorf("capacity cannot be less than booked participant (%d)", schedule.Booked)
	}

	schedule.Date = parsedDate
	schedule.StartHour = req.StartHour
	schedule.StartMinute = req.StartMinute
	schedule.ClassID = uuid.MustParse(req.ClassID)
	schedule.InstructorID = instructorID
	schedule.Capacity = req.Capacity
	schedule.Color = req.Color

	return s.repo.UpdateClassSchedule(schedule)
}

func (s *classScheduleService) DeleteClassSchedule(id string) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}

	// Hitung waktu mulai kelas
	startTime := time.Date(
		schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
		schedule.StartHour, schedule.StartMinute, 0, 0, time.Local,
	)

	if startTime.Before(time.Now().In(time.Local)) {
		return fmt.Errorf("cannot delete past or ongoing class schedule")
	}

	// Cek apakah sudah dibooking
	isBooked, err := s.repo.HasActiveBooking(schedule.ID)
	if err != nil {
		return fmt.Errorf("failed to check booking: %w", err)
	}
	if isBooked {
		return fmt.Errorf("cannot delete schedule with active bookings")
	}

	return s.repo.DeleteClassSchedule(id)
}

func (s *classScheduleService) GetAllClassSchedules() ([]dto.ClassScheduleResponse, error) {
	schedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleResponse
	for _, schedule := range schedules {
		result = append(result, dto.ClassScheduleResponse{
			ID: schedule.ID.String(),
			Class: dto.ClassBrief{
				ID:       schedule.Class.ID.String(),
				Title:    schedule.Class.Title,
				Image:    schedule.Class.Image,
				Duration: schedule.Class.Duration,
			},
			Instructor: dto.InstructorBrief{
				ID:       schedule.Instructor.ID.String(),
				Fullname: schedule.Instructor.User.Profile.Fullname,
				Rating:   schedule.Instructor.Rating,
			},
			Category:    schedule.Class.Category.Name,
			Date:        schedule.Date,
			StartHour:   schedule.StartHour,
			StartMinute: schedule.StartMinute,
			Capacity:    schedule.Capacity,
			BookedCount: schedule.Booked,
			Color:       schedule.Color,
			IsBooked:    false,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error) {
	schedule, err := s.repo.GetClassScheduleByID(scheduleID)
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
			Image: p.Image,
		})
	}

	isBooked, _ := s.bookingRepo.IsUserBookedSchedule(userID, scheduleID)

	return &dto.ClassScheduleDetailResponse{
		ClassScheduleResponse: dto.ClassScheduleResponse{
			ID: schedule.ID.String(),
			Class: dto.ClassBrief{
				ID:       schedule.Class.ID.String(),
				Title:    schedule.Class.Title,
				Image:    schedule.Class.Image,
				Duration: schedule.Class.Duration,
			},
			Instructor: dto.InstructorBrief{
				ID:       schedule.Instructor.ID.String(),
				Fullname: schedule.Instructor.User.Profile.Fullname,
				Rating:   schedule.Instructor.Rating,
			},
			Category:    schedule.Class.Category.Name,
			Date:        schedule.Date,
			StartHour:   schedule.StartHour,
			StartMinute: schedule.StartMinute,
			Capacity:    schedule.Capacity,
			BookedCount: schedule.Booked,
			Color:       schedule.Color,
			IsBooked:    isBooked,
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
	for _, schedule := range schedules {
		result = append(result, dto.ClassScheduleResponse{
			ID: schedule.ID.String(),
			Class: dto.ClassBrief{
				ID:       schedule.Class.ID.String(),
				Title:    schedule.Class.Title,
				Image:    schedule.Class.Image,
				Duration: schedule.Class.Duration,
			},
			Instructor: dto.InstructorBrief{
				ID:       schedule.Instructor.ID.String(),
				Fullname: schedule.Instructor.User.Profile.Fullname,
				Rating:   schedule.Instructor.Rating,
			},
			Category:    schedule.Class.Category.Name,
			Date:        schedule.Date,
			StartHour:   schedule.StartHour,
			StartMinute: schedule.StartMinute,
			Capacity:    schedule.Capacity,
			BookedCount: schedule.Booked,
			Color:       schedule.Color,
			IsBooked:    false,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleResponse, error) {
	schedules, err := s.repo.GetClassSchedules()
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleResponse
	for _, schedule := range schedules {
		isBooked, _ := s.bookingRepo.IsUserBookedSchedule(userID, schedule.ID.String())

		result = append(result, dto.ClassScheduleResponse{
			ID: schedule.ID.String(),
			Class: dto.ClassBrief{
				ID:       schedule.Class.ID.String(),
				Title:    schedule.Class.Title,
				Image:    schedule.Class.Image,
				Duration: schedule.Class.Duration,
			},
			Instructor: dto.InstructorBrief{
				ID:       schedule.Instructor.ID.String(),
				Fullname: schedule.Instructor.User.Profile.Fullname,
				Rating:   schedule.Instructor.Rating,
			},
			Category:    schedule.Class.Category.Name,
			Date:        schedule.Date,
			StartHour:   schedule.StartHour,
			StartMinute: schedule.StartMinute,
			Capacity:    schedule.Capacity,
			BookedCount: schedule.Booked,
			Color:       schedule.Color,
			IsBooked:    isBooked,
		})
	}

	return result, nil
}
