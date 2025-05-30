package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type InstructorService interface {
	DeleteInstructor(id string) error
	GetAllInstructors() ([]dto.InstructorResponse, error)
	CreateInstructor(req dto.CreateInstructorRequest) error
	GetInstructorByID(id string) (*dto.InstructorResponse, error)
	UpdateInstructor(id string, req dto.UpdateInstructorRequest) error
}

type instructorService struct {
	repo     repositories.InstructorRepository
	userRepo repositories.AuthRepository
}

func NewInstructorService(repo repositories.InstructorRepository, userRepo repositories.AuthRepository) InstructorService {
	return &instructorService{repo, userRepo}
}

func (s *instructorService) DeleteInstructor(id string) error {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetUserByID(instructor.UserID.String())
	if err == nil {
		user.Role = "customer"
		_ = s.userRepo.UpdateUser(user)
	}

	return s.repo.DeleteInstructor(id)
}

func (s *instructorService) GetAllInstructors() ([]dto.InstructorResponse, error) {
	instructors, err := s.repo.GetAllInstructors()
	if err != nil {
		return nil, err
	}

	var result []dto.InstructorResponse
	for _, i := range instructors {
		result = append(result, dto.InstructorResponse{
			ID:             i.ID.String(),
			UserID:         i.UserID.String(),
			Fullname:       i.User.Profile.Fullname,
			Avatar:         i.User.Profile.Avatar,
			Experience:     i.Experience,
			Specialties:    i.Specialties,
			Certifications: i.Certifications,
			Rating:         i.Rating,
			TotalClass:     i.TotalClass,
		})
	}
	return result, nil
}

func (s *instructorService) CreateInstructor(req dto.CreateInstructorRequest) error {
	user, err := s.userRepo.GetUserByID(req.UserID)
	if err != nil {
		return errors.New("user not found")
	}
	user.Role = "instructor"
	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	instructor := models.Instructor{
		UserID:         uuid.MustParse(req.UserID),
		Experience:     req.Experience,
		Specialties:    req.Specialties,
		Certifications: req.Certifications,
	}
	return s.repo.CreateInstructor(&instructor)
}

func (s *instructorService) GetInstructorByID(id string) (*dto.InstructorResponse, error) {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.InstructorResponse{
		ID:             instructor.ID.String(),
		UserID:         instructor.UserID.String(),
		Fullname:       instructor.User.Profile.Fullname,
		Avatar:         instructor.User.Profile.Avatar,
		Experience:     instructor.Experience,
		Specialties:    instructor.Specialties,
		Certifications: instructor.Certifications,
		Rating:         instructor.Rating,
		TotalClass:     instructor.TotalClass,
	}, nil
}

func (s *instructorService) UpdateInstructor(id string, req dto.UpdateInstructorRequest) error {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return err
	}

	if instructor.UserID.String() != req.UserID {
		return errors.New("userId mismatch")
	}

	if req.Experience != 0 {
		instructor.Experience = req.Experience
	}
	if req.Specialties != "" {
		instructor.Specialties = req.Specialties
	}
	if req.Certifications != "" {
		instructor.Certifications = req.Certifications
	}

	return s.repo.UpdateInstructor(instructor)
}
