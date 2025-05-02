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

type ClassService interface {
	CreateClass(req dto.CreateClassRequest) error
	UpdateClass(id string, req dto.UpdateClassRequest) error
	DeleteClass(id string) error
	GetClassByID(id string) (*dto.ClassDetailResponse, error)
	GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, int64, error)
	GetActiveClasses() ([]dto.ClassResponse, error)
	AddClassGallery(galleries []models.ClassGallery) error
	DeleteClassGallery(galleryID string) error
}

type classService struct {
	repo repositories.ClassRepository
}

func NewClassService(repo repositories.ClassRepository) ClassService {
	return &classService{repo}
}

func (s *classService) CreateClass(req dto.CreateClassRequest) error {
	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return err
	}
	levelID, err := uuid.Parse(req.LevelID)
	if err != nil {
		return err
	}
	locationID, err := uuid.Parse(req.LocationID)
	if err != nil {
		return err
	}
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return err
	}
	subcategoryID, err := uuid.Parse(req.SubcategoryID)
	if err != nil {
		return err
	}

	class := models.Class{
		Title:          req.Title,
		Image:          req.ImageURL,
		Duration:       req.Duration,
		Description:    req.Description,
		AdditionalList: req.Additional,
		TypeID:         typeID,
		LevelID:        levelID,
		LocationID:     locationID,
		CategoryID:     categoryID,
		SubcategoryID:  subcategoryID,
		IsActive:       req.IsActive,
		CreatedAt:      time.Now(),
	}

	err = s.repo.CreateClass(&class)
	if err != nil {
		return err
	}
	fmt.Println("imageURLs", req.ImageURLs)

	if len(req.ImageURLs) > 0 {
		var galleries []models.ClassGallery
		for _, url := range req.ImageURLs {
			galleries = append(galleries, models.ClassGallery{
				ID:        uuid.New(),
				ClassID:   class.ID,
				URL:       url,
				CreatedAt: time.Now(),
			})
		}

		if err := s.repo.SaveClassGalleries(galleries); err != nil {
			return err
		}

	}

	return nil
}

func (s *classService) UpdateClass(id string, req dto.UpdateClassRequest) error {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return err
	}

	class.IsActive = req.IsActive

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

	if req.SubcategoryID != "" {
		subcategoryID, _ := uuid.Parse(req.SubcategoryID)
		class.SubcategoryID = subcategoryID
	}

	if req.ImageURL != "" {
		class.Image = req.ImageURL
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

func (s *classService) GetClassByID(id string) (*dto.ClassDetailResponse, error) {
	class, err := s.repo.GetClassByID(id)
	if err != nil {
		return nil, err
	}

	var galleries []string
	for _, g := range class.Galleries {
		galleries = append(galleries, g.URL)
	}

	reviews := make([]dto.ReviewResponse, len(class.Reviews))
	for i, r := range class.Reviews {
		reviews[i] = dto.ReviewResponse{
			ID:         r.ID.String(),
			UserName:   r.User.Profile.Fullname,
			ClassTitle: class.Title,
			Rating:     r.Rating,
			Comment:    r.Comment,
			CreatedAt:  r.CreatedAt.Format(time.RFC3339),
		}
	}

	return &dto.ClassDetailResponse{
		ID:          class.ID.String(),
		Title:       class.Title,
		Image:       class.Image,
		IsActive:    class.IsActive,
		Duration:    class.Duration,
		Description: class.Description,
		Additional:  class.AdditionalList,
		Type:        class.Type.Name,
		Level:       class.Level.Name,
		Location:    class.Location.Name,
		Category:    class.Category.Name,
		Subcategory: class.Subcategory.Name,
		Galleries:   galleries,
		Reviews:     reviews,
		CreatedAt:   class.CreatedAt.Format(time.RFC3339),
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
		var galleries []string
		for _, g := range c.Galleries {
			galleries = append(galleries, g.URL)
		}

		result = append(result, dto.ClassResponse{
			ID:            c.ID.String(),
			Title:         c.Title,
			Image:         c.Image,
			IsActive:      c.IsActive,
			Duration:      c.Duration,
			Description:   c.Description,
			Additional:    c.AdditionalList,
			TypeID:        c.TypeID.String(),
			LevelID:       c.LevelID.String(),
			LocationID:    c.LocationID.String(),
			CategoryID:    c.CategoryID.String(),
			SubcategoryID: c.SubcategoryID.String(),
			CreatedAt:     c.CreatedAt.Format(time.RFC3339),
			Galleries:     galleries,
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
			ID:            c.ID.String(),
			Title:         c.Title,
			Image:         c.Image,
			IsActive:      c.IsActive,
			Duration:      c.Duration,
			Description:   c.Description,
			Additional:    c.AdditionalList,
			TypeID:        c.TypeID.String(),
			LevelID:       c.LevelID.String(),
			LocationID:    c.LocationID.String(),
			CategoryID:    c.CategoryID.String(),
			SubcategoryID: c.SubcategoryID.String(),
			CreatedAt:     c.CreatedAt.String(),
		})
	}

	return result, nil
}

func (s *classService) AddClassGallery(galleries []models.ClassGallery) error {
	return s.repo.SaveClassGalleries(galleries)
}

func (s *classService) DeleteClassGallery(galleryID string) error {
	return s.repo.DeleteClassGalleryByID(galleryID)
}
