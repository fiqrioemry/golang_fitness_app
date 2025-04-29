package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type InstructorService interface {
	CreateInstructor(userID string, req dto.CreateInstructorRequest) error
	UpdateInstructor(id string, req dto.UpdateInstructorRequest) error
	DeleteInstructor(id string) error
	GetInstructorByID(id string) (*dto.InstructorResponse, error)
	GetAllInstructors() ([]dto.InstructorResponse, error)
}

type instructorService struct {
	repo repositories.InstructorRepository
}

func NewInstructorService(repo repositories.InstructorRepository) InstructorService {
	return &instructorService{repo}
}

func (s *instructorService) CreateInstructor(userID string, req dto.CreateInstructorRequest) error {
	userUUID, _ := uuid.Parse(userID)
	instructor := models.Instructor{
		UserID:         userUUID,
		Experience:     req.Experience,
		Specialties:    req.Specialties,
		Certifications: req.Certifications,
	}
	return s.repo.CreateInstructor(&instructor)
}

func (s *instructorService) UpdateInstructor(id string, req dto.UpdateInstructorRequest) error {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return err
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

func (s *instructorService) DeleteInstructor(id string) error {
	return s.repo.DeleteInstructor(id)
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
