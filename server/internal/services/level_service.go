package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type LevelService interface {
	CreateLevel(req dto.CreateLevelRequest) error
	UpdateLevel(id string, req dto.UpdateLevelRequest) error
	DeleteLevel(id string) error
	GetAllLevels() ([]dto.LevelResponse, error)
	GetLevelByID(id string) (*dto.LevelResponse, error)
}

type levelService struct {
	repo repositories.LevelRepository
}

func NewLevelService(repo repositories.LevelRepository) LevelService {
	return &levelService{repo}
}

func (s *levelService) CreateLevel(req dto.CreateLevelRequest) error {
	level := models.Level{
		ID:   uuid.New(),
		Name: req.Name,
	}
	return s.repo.CreateLevel(&level)
}

func (s *levelService) UpdateLevel(id string, req dto.UpdateLevelRequest) error {
	level, err := s.repo.GetLevelByID(id)
	if err != nil {
		return err
	}

	level.Name = req.Name
	return s.repo.UpdateLevel(level)
}

func (s *levelService) DeleteLevel(id string) error {
	return s.repo.DeleteLevel(id)
}

func (s *levelService) GetAllLevels() ([]dto.LevelResponse, error) {
	levels, err := s.repo.GetAllLevels()
	if err != nil {
		return nil, err
	}

	var result []dto.LevelResponse
	for _, l := range levels {
		result = append(result, dto.LevelResponse{
			ID:   l.ID.String(),
			Name: l.Name,
		})
	}
	return result, nil
}

func (s *levelService) GetLevelByID(id string) (*dto.LevelResponse, error) {
	level, err := s.repo.GetLevelByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.LevelResponse{
		ID:   level.ID.String(),
		Name: level.Name,
	}, nil
}
