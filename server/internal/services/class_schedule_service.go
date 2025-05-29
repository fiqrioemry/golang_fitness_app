package services

import (
	"fmt"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ClassScheduleService interface {
	// admin
	DeleteClassSchedule(id string) error
	CreateClassSchedule(req dto.CreateClassScheduleRequest) error
	UpdateClassSchedule(id string, req dto.UpdateClassScheduleRequest) error

	// customer
	GetSchedulesWithBookingStatus(userID string) ([]dto.ClassScheduleResponse, error)
	GetClassScheduleByID(scheduleID, userID string) (*dto.ClassScheduleDetailResponse, error)
	GetSchedulesByFilter(filter dto.ClassScheduleQueryParam) ([]dto.ClassScheduleResponse, error)

	// instructor only
	OpenClassSchedule(id string, req dto.OpenClassScheduleRequest) error
	GetAttendancesForSchedule(scheduleID string) ([]dto.AttendanceWithUserResponse, error)
	GetSchedulesByInstructor(userID string, params dto.InstructorScheduleQueryParam) ([]dto.InstructorScheduleResponse, *dto.PaginationResponse, error)
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

func (s *classScheduleService) GetSchedulesByInstructor(userID string, params dto.InstructorScheduleQueryParam) ([]dto.InstructorScheduleResponse, *dto.PaginationResponse, error) {
	ID := uuid.MustParse(userID)

	instructor, err := s.repo.GetInstructorByUserID(ID)
	log.Printf("instructor id %s", instructor.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("instructor not found: %w with %s", err, instructor.ID)
	}

	schedules, total, err := s.repo.GetSchedulesByInstructorID(instructor.ID, params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.InstructorScheduleResponse
	for _, schedule := range schedules {
		result = append(result, dto.InstructorScheduleResponse{
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
			IsOpened:       schedule.IsOpened,
		})
	}
	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))
	pagination := &dto.PaginationResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil
}

func (s *classScheduleService) OpenClassSchedule(id string, req dto.OpenClassScheduleRequest) error {
	schedule, err := s.repo.GetClassScheduleByID(id)
	if err != nil {
		return fmt.Errorf("schedule not found")
	}
	if schedule.IsOpened {
		return fmt.Errorf("schedule already opened")
	}
	return s.repo.OpenSchedule(schedule.ID, req.ZoomLink)
}

func (s *classScheduleService) GetAttendancesForSchedule(scheduleID string) ([]dto.AttendanceWithUserResponse, error) {
	bookings, err := s.repo.GetAttendancesByScheduleID(scheduleID)
	if err != nil {
		return nil, err
	}

	var result []dto.AttendanceWithUserResponse
	for _, b := range bookings {
		attendance := b.Attendance
		user := b.User

		resp := dto.AttendanceWithUserResponse{
			UserName:   user.Profile.Fullname,
			Email:      user.Email,
			Status:     attendance.Status,
			CheckedIn:  attendance.CheckedIn,
			CheckedOut: attendance.CheckedOut,
		}

		if !attendance.CheckedAt.IsZero() {
			resp.CheckedAt = attendance.CheckedAt.Format("2006-01-02 15:04:05")
		}
		if !attendance.VerifiedAt.IsZero() {
			resp.VerifiedAt = attendance.VerifiedAt.Format("2006-01-02 15:04:05")
		}

		result = append(result, resp)
	}

	return result, nil
}
