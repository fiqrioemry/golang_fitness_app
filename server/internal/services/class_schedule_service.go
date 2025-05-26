package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ClassScheduleService interface {
	DeleteClassSchedule(id string) error
	CreateClassSchedule(req dto.CreateClassScheduleRequest) error
	UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error
	GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error)
	GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error)
	GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleResponse, error)

	// instructor only
	CloseClassSchedule(id string) (string, error)
	OpenClassSchedule(id string, req dto.OpenClassScheduleRequest) error
	GetClassScheduleAttendances(scheduleID string) ([]dto.ClassAttendanceStruct, error)
	GetSchedulesByInstructor(userID string, filter dto.InstructorScheduleQueryParam) ([]dto.ClassScheduleResponse, error)
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

	class, err := s.repo.GetClassByID(uuid.MustParse(req.ClassID))
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.repo.GetInstructorWithProfileByID(uuid.MustParse(req.InstructorID))
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
	}

	schedule := models.ClassSchedule{
		ID:             uuid.New(),
		ClassID:        class.ID,
		ClassName:      class.Title,
		ClassImage:     class.Image,
		InstructorID:   instructor.ID,
		InstructorName: instructor.User.Profile.Fullname,
		Location:       class.Location.Name,
		Duration:       class.Duration,
		Capacity:       req.Capacity,
		IsActive:       true,
		Color:          req.Color,
		Date:           localDate,
		StartHour:      req.StartHour,
		StartMinute:    req.StartMinute,
	}

	return s.repo.CreateClassSchedule(&schedule)
}

func (s *classScheduleService) UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}

	instructorID := uuid.MustParse(req.InstructorID)
	classID := uuid.MustParse(req.ClassID)
	localDate := req.Date.In(time.Local)

	newStart := time.Date(
		localDate.Year(), localDate.Month(), localDate.Day(),
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
		if other.InstructorID == instructorID &&
			other.Date.Year() == localDate.Year() &&
			other.Date.Month() == localDate.Month() &&
			other.Date.Day() == localDate.Day() {

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

	class, err := s.repo.GetClassByID(classID)
	if err != nil {
		return fmt.Errorf("class not found: %w", err)
	}

	instructor, err := s.repo.GetInstructorWithProfileByID(instructorID)
	if err != nil {
		return fmt.Errorf("instructor not found: %w", err)
	}

	schedule.Date = localDate
	schedule.StartHour = req.StartHour
	schedule.StartMinute = req.StartMinute
	schedule.ClassID = class.ID
	schedule.ClassName = class.Title
	schedule.ClassImage = class.Image
	schedule.InstructorID = instructor.ID
	schedule.InstructorName = instructor.User.Profile.Fullname
	schedule.Capacity = req.Capacity
	schedule.Color = req.Color
	schedule.Duration = class.Duration

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

func (s *classScheduleService) GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error) {
	schedule, err := s.repo.GetClassScheduleByID(scheduleID)
	if err != nil {
		return nil, err
	}

	packages, err := s.packageRepo.GetPackagesByClassID(schedule.ClassID.String())
	if err != nil {
		return nil, err
	}

	var pkgResponses []dto.PackageListResponse
	for _, p := range packages {
		pkgResponses = append(pkgResponses, dto.PackageListResponse{
			ID:    p.ID.String(),
			Name:  p.Name,
			Price: p.Price,
			Image: p.Image,
		})
	}

	isBooked, _ := s.bookingRepo.IsUserBookedSchedule(userID, scheduleID)

	return &dto.ClassScheduleDetailResponse{
		ClassScheduleResponse: dto.ClassScheduleResponse{
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Location:       schedule.Location,
			Date:           schedule.Date,
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Capacity:       schedule.Capacity,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
			Duration:       schedule.Duration,
			IsBooked:       isBooked,
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
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Location:       schedule.Location,
			Date:           schedule.Date,
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Capacity:       schedule.Capacity,
			Duration:       schedule.Duration,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
			IsBooked:       false,
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
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Date:           schedule.Date,
			StartHour:      schedule.StartHour,
			Location:       schedule.Location,
			StartMinute:    schedule.StartMinute,
			Duration:       schedule.Duration,
			Capacity:       schedule.Capacity,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
			IsBooked:       isBooked,
		})
	}

	return result, nil
}

func (s *classScheduleService) GetSchedulesByInstructor(userID string, filter dto.InstructorScheduleQueryParam) ([]dto.ClassScheduleResponse, error) {
	instructorID := uuid.MustParse(userID)

	schedules, err := s.repo.GetSchedulesByInstructorID(instructorID, filter)
	if err != nil {
		return nil, err
	}

	var result []dto.ClassScheduleResponse
	for _, schedule := range schedules {
		result = append(result, dto.ClassScheduleResponse{
			ID:             schedule.ID.String(),
			ClassID:        schedule.ClassID.String(),
			ClassName:      schedule.ClassName,
			ClassImage:     schedule.ClassImage,
			InstructorID:   schedule.InstructorID.String(),
			InstructorName: schedule.InstructorName,
			Location:       schedule.Location,
			Date:           schedule.Date,
			StartHour:      schedule.StartHour,
			StartMinute:    schedule.StartMinute,
			Capacity:       schedule.Capacity,
			Duration:       schedule.Duration,
			BookedCount:    schedule.Booked,
			Color:          schedule.Color,
		})
	}
	return result, nil
}

func (s *classScheduleService) OpenClassSchedule(id string, req dto.OpenClassScheduleRequest) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}
	if schedule.IsOpened {
		return fmt.Errorf("schedule already opened")
	}
	if req.IsOnline && (req.ZoomLink == "" || req.ZoomLink == "null") {
		return fmt.Errorf("zoom link is required for online class")
	}
	var zoomPtr *string
	if req.IsOnline {
		zoomPtr = &req.ZoomLink
	}
	return s.repo.OpenSchedule(schedule.ID, req.IsOnline, zoomPtr)
}

func (s *classScheduleService) CloseClassSchedule(id string) (string, error) {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return "", fmt.Errorf("schedule not found")
	}
	if schedule.IsClosed {
		return "", fmt.Errorf("schedule already closed")
	}

	code := utils.GenerateVerificationCode(6) // misal 6 karakter acak

	err = s.repo.CloseScheduleWithCode(schedule.ID, code)
	if err != nil {
		return "", fmt.Errorf("failed to close schedule: %w", err)
	}

	return code, nil
}

func (s *classScheduleService) GetClassScheduleAttendances(scheduleID string) ([]dto.ClassAttendanceStruct, error) {
	id := uuid.MustParse(scheduleID)

	attendances, err := s.repo.GetAttendancesByScheduleID(id)
	if err != nil {
		return nil, err
	}

	var results []dto.ClassAttendanceStruct
	for _, a := range attendances {
		results = append(results, dto.ClassAttendanceStruct{
			ID:         a.ID.String(),
			UserID:     a.Booking.UserID.String(),
			Status:     a.Status,
			Verified:   a.Verified,
			CheckinAt:  a.CheckedAt.Format("2006-01-02 15:04:05"),
			CheckoutAt: a.VerifiedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return results, nil
}
