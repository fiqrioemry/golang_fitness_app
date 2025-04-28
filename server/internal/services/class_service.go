package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ClassService interface {
	CreateClass(req dto.CreateClassRequest) error
	UpdateClass(id string, req dto.UpdateClassRequest) error
	DeleteClass(id string) error
	GetClassByID(id string) (*dto.ClassResponse, error)
	GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, int64, error)
	GetActiveClasses() ([]dto.ClassResponse, error)
}

type classService struct {
	repo repositories.ClassRepository
}

func NewClassService(repo repositories.ClassRepository) ClassService {
	return &classService{repo}
}

func (s *classService) CreateClass(req dto.CreateClassRequest) error {
	file, err := req.Image.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	imageUrl, err := utils.UploadToCloudinary(file)
	if err != nil {
		return err
	}

	typeID, _ := uuid.Parse(req.TypeID)
	levelID, _ := uuid.Parse(req.LevelID)
	locationID, _ := uuid.Parse(req.LocationID)
	categoryID, _ := uuid.Parse(req.CategoryID)

	class := models.Class{
		Title:          req.Title,
		Image:          imageUrl,
		Duration:       req.Duration,
		Description:    req.Description,
		AdditionalList: req.Additional,
		TypeID:         typeID,
		LevelID:        levelID,
		LocationID:     locationID,
		CategoryID:     categoryID,
		IsActive:       true,
		CreatedAt:      time.Now(),
	}

	return s.repo.CreateClass(&class)
}

func (s *classService) UpdateClass(id string, req dto.UpdateClassRequest) error {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return err
	}

	if req.Title != "" {
		class.Title = req.Title
	}
	if req.Description != "" {
		class.Description = req.Description
	}
	if req.Duration != 0 {
		class.Duration = req.Duration
	}
	if len(req.Additional) > 0 {
		class.AdditionalList = req.Additional
	}
	if req.TypeID != "" {
		typeID, _ := uuid.Parse(req.TypeID)
		class.TypeID = typeID
	}
	if req.LevelID != "" {
		levelID, _ := uuid.Parse(req.LevelID)
		class.LevelID = levelID
	}
	if req.LocationID != "" {
		locationID, _ := uuid.Parse(req.LocationID)
		class.LocationID = locationID
	}
	if req.CategoryID != "" {
		categoryID, _ := uuid.Parse(req.CategoryID)
		class.CategoryID = categoryID
	}

	if req.Image != nil {
		file, err := req.Image.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		newImageUrl, err := utils.UploadToCloudinary(file)
		if err != nil {
			return err
		}

		if class.Image != "" {
			_ = utils.DeleteFromCloudinary(class.Image)
		}
		class.Image = newImageUrl
	}

	return s.repo.UpdateClass(class)
}

func (s *classService) DeleteClass(id string) error {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return err
	}

	if class.Image != "" {
		_ = utils.DeleteFromCloudinary(class.Image)
	}

	return s.repo.DeleteClass(id)
}

func (s *classService) GetClassByID(id string) (*dto.ClassResponse, error) {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ClassResponse{
		ID:          class.ID.String(),
		Title:       class.Title,
		Image:       class.Image,
		IsActive:    class.IsActive,
		Duration:    class.Duration,
		Description: class.Description,
		Additional:  class.AdditionalList,
		TypeID:      class.TypeID.String(),
		LevelID:     class.LevelID.String(),
		LocationID:  class.LocationID.String(),
		CategoryID:  class.CategoryID.String(),
		CreatedAt:   class.CreatedAt,
	}, nil
}

func (s *classService) GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, int64, error) {
	filter := map[string]interface{}{}
	if params.TypeID != "" {
		filter["type_id"] = params.TypeID
	}
	if params.LevelID != "" {
		filter["level_id"] = params.LevelID
	}
	if params.LocationID != "" {
		filter["location_id"] = params.LocationID
	}
	if params.CategoryID != "" {
		filter["category_id"] = params.CategoryID
	}
	if params.IsActive != "" {
		if params.IsActive == "true" {
			filter["is_active"] = true
		} else if params.IsActive == "false" {
			filter["is_active"] = false
		}
	}

	offset := (params.Page - 1) * params.Limit

	classes, total, err := s.repo.GetAllClasses(filter, params.Q, params.Sort, params.Limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var result []dto.ClassResponse
	for _, c := range classes {
		result = append(result, dto.ClassResponse{
			ID:          c.ID.String(),
			Title:       c.Title,
			Image:       c.Image,
			IsActive:    c.IsActive,
			Duration:    c.Duration,
			Description: c.Description,
			Additional:  c.AdditionalList,
			TypeID:      c.TypeID.String(),
			LevelID:     c.LevelID.String(),
			LocationID:  c.LocationID.String(),
			CategoryID:  c.CategoryID.String(),
			CreatedAt:   c.CreatedAt,
		})
	}

	return result, total, nil
}

func (s *classService) GetActiveClasses() ([]dto.ClassResponse, error) {
	classes, err := s.repo.GetActiveClasses()
	if err != nil {
		return nil, err
	}

	var result []dto.ClassResponse
	for _, c := range classes {
		result = append(result, dto.ClassResponse{
			ID:          c.ID.String(),
			Title:       c.Title,
			Image:       c.Image,
			IsActive:    c.IsActive,
			Duration:    c.Duration,
			Description: c.Description,
			Additional:  c.AdditionalList,
			TypeID:      c.TypeID.String(),
			LevelID:     c.LevelID.String(),
			LocationID:  c.LocationID.String(),
			CategoryID:  c.CategoryID.String(),
			CreatedAt:   c.CreatedAt,
		})
	}

	return result, nil
}
