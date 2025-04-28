package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type TypeService interface {
	CreateType(req dto.CreateTypeRequest) error
	UpdateType(id string, req dto.UpdateTypeRequest) error
	DeleteType(id string) error
	GetAllTypes() ([]dto.TypeResponse, error)
	GetTypeByID(id string) (*dto.TypeResponse, error)
}

type typeService struct {
	repo repositories.TypeRepository
}

func NewTypeService(repo repositories.TypeRepository) TypeService {
	return &typeService{repo}
}

func (s *typeService) CreateType(req dto.CreateTypeRequest) error {
	t := models.Type{
		ID:   uuid.New(),
		Name: req.Name,
	}
	return s.repo.CreateType(&t)
}

func (s *typeService) UpdateType(id string, req dto.UpdateTypeRequest) error {
	t, err := s.repo.GetTypeByID(id)
	if err != nil {
		return err
	}

	t.Name = req.Name
	return s.repo.UpdateType(t)
}

func (s *typeService) DeleteType(id string) error {
	_, err := s.repo.GetTypeByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteType(id)
}

func (s *typeService) GetAllTypes() ([]dto.TypeResponse, error) {
	types, err := s.repo.GetAllTypes()
	if err != nil {
		return nil, err
	}

	var result []dto.TypeResponse
	for _, t := range types {
		result = append(result, dto.TypeResponse{
			ID:   t.ID.String(),
			Name: t.Name,
		})
	}
	return result, nil
}

func (s *typeService) GetTypeByID(id string) (*dto.TypeResponse, error) {
	t, err := s.repo.GetTypeByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.TypeResponse{
		ID:   t.ID.String(),
		Name: t.Name,
	}, nil
}
