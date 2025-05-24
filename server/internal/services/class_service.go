package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"slices"

	"github.com/google/uuid"
)

type ClassService interface {
	DeleteClass(id string) error
	CreateClass(req dto.CreateClassRequest) error
	AddClassGallery(galleries []models.ClassGallery) error
	UpdateClass(id string, req dto.UpdateClassRequest) error
	GetClassByID(id string) (*dto.ClassDetailResponse, error)
	UpdateClassGallery(classID uuid.UUID, keepImages []string, newImageURLs []string) error
	GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, *dto.PaginationResponse, error)
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

	// Hapus image utama dari Cloudinary
	if class.Image != "" {
		_ = utils.DeleteFromCloudinary(class.Image)
	}

	// Hapus semua gallery dari cloudinary
	for _, gallery := range class.Galleries {
		if gallery.URL != "" {
			_ = utils.DeleteFromCloudinary(gallery.URL)
		}
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
		CreatedAt:   class.CreatedAt.Format("2006-01-02"),
	}, nil

}
func (s *classService) GetAllClasses(params dto.ClassQueryParam) ([]dto.ClassResponse, *dto.PaginationResponse, error) {
	classes, total, err := s.repo.GetAllClasses(params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.ClassResponse
	for _, c := range classes {
		galleries := make([]string, len(c.Galleries))
		for i, g := range c.Galleries {
			galleries[i] = g.URL
		}
		result = append(result, dto.ClassResponse{
			ID:            c.ID.String(),
			Title:         c.Title,
			Image:         c.Image,
			IsActive:      c.IsActive,
			Duration:      c.Duration,
			Galleries:     galleries,
			Description:   c.Description,
			Additional:    c.AdditionalList,
			TypeID:        c.TypeID.String(),
			LevelID:       c.LevelID.String(),
			LocationID:    c.LocationID.String(),
			CategoryID:    c.CategoryID.String(),
			SubcategoryID: c.SubcategoryID.String(),
			CreatedAt:     c.CreatedAt.Format("2006-01-02"),
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

func (s *classService) UpdateClassGallery(classID uuid.UUID, keepImages []string, newImageURLs []string) error {
	oldGalleries, err := s.repo.FindGalleriesByClassID(classID)
	if err != nil {
		return err
	}

	for _, gallery := range oldGalleries {
		isKept := slices.Contains(keepImages, gallery.URL)
		if !isKept {
			_ = utils.DeleteFromCloudinary(gallery.URL)
			_ = s.repo.DeleteClassGalleryByID(gallery.ID.String())
		}
	}

	if len(newImageURLs) > 0 {
		var newGalleries []models.ClassGallery
		for _, url := range newImageURLs {
			newGalleries = append(newGalleries, models.ClassGallery{
				ID:      uuid.New(),
				ClassID: classID,
				URL:     url,
			})
		}
		if err := s.repo.SaveClassGalleries(newGalleries); err != nil {
			return err
		}
	}

	return nil
}

func (s *classService) AddClassGallery(galleries []models.ClassGallery) error {
	return s.repo.SaveClassGalleries(galleries)
}
