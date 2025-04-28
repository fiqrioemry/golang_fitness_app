package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type SubcategoryService interface {
	CreateSubcategory(req dto.CreateSubcategoryRequest) error
	UpdateSubcategory(id string, req dto.UpdateSubcategoryRequest) error
	DeleteSubcategory(id string) error
	GetAllSubcategories() ([]dto.SubcategoryResponse, error)
	GetSubcategoryByID(id string) (*dto.SubcategoryResponse, error)
	GetSubcategoriesByCategoryID(categoryID string) ([]dto.SubcategoryResponse, error)
}

type subcategoryService struct {
	repo repositories.SubcategoryRepository
}

func NewSubcategoryService(repo repositories.SubcategoryRepository) SubcategoryService {
	return &subcategoryService{repo}
}

func (s *subcategoryService) CreateSubcategory(req dto.CreateSubcategoryRequest) error {
	categoryID, _ := uuid.Parse(req.CategoryID)

	subcategory := models.Subcategory{
		ID:         uuid.New(),
		Name:       req.Name,
		CategoryID: categoryID,
	}

	return s.repo.CreateSubcategory(&subcategory)
}

func (s *subcategoryService) UpdateSubcategory(id string, req dto.UpdateSubcategoryRequest) error {
	subcategory, err := s.repo.GetSubcategoryByID(id)
	if err != nil {
		return err
	}

	categoryID, _ := uuid.Parse(req.CategoryID)

	subcategory.Name = req.Name
	subcategory.CategoryID = categoryID

	return s.repo.UpdateSubcategory(subcategory)
}

func (s *subcategoryService) DeleteSubcategory(id string) error {
	_, err := s.repo.GetSubcategoryByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteSubcategory(id)
}

func (s *subcategoryService) GetAllSubcategories() ([]dto.SubcategoryResponse, error) {
	subcategories, err := s.repo.GetAllSubcategories()
	if err != nil {
		return nil, err
	}

	var result []dto.SubcategoryResponse
	for _, s := range subcategories {
		result = append(result, dto.SubcategoryResponse{
			ID:         s.ID.String(),
			Name:       s.Name,
			CategoryID: s.CategoryID.String(),
		})
	}
	return result, nil
}

func (s *subcategoryService) GetSubcategoryByID(id string) (*dto.SubcategoryResponse, error) {
	subcategory, err := s.repo.GetSubcategoryByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.SubcategoryResponse{
		ID:         subcategory.ID.String(),
		Name:       subcategory.Name,
		CategoryID: subcategory.CategoryID.String(),
	}, nil
}

func (s *subcategoryService) GetSubcategoriesByCategoryID(categoryID string) ([]dto.SubcategoryResponse, error) {
	subcategories, err := s.repo.GetSubcategoriesByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	var result []dto.SubcategoryResponse
	for _, s := range subcategories {
		result = append(result, dto.SubcategoryResponse{
			ID:         s.ID.String(),
			Name:       s.Name,
			CategoryID: s.CategoryID.String(),
		})
	}
	return result, nil
}
