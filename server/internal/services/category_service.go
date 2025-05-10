package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type CategoryService interface {
	DeleteCategory(id string) error
	GetAllCategories() ([]dto.CategoryResponse, error)
	CreateCategory(req dto.CreateCategoryRequest) error
	GetCategoryByID(id string) (*dto.CategoryResponse, error)
	UpdateCategory(id string, req dto.UpdateCategoryRequest) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest) error {
	category := models.Category{
		ID:   uuid.New(),
		Name: req.Name,
	}
	return s.repo.CreateCategory(&category)
}

func (s *categoryService) UpdateCategory(id string, req dto.UpdateCategoryRequest) error {
	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return err
	}

	category.Name = req.Name
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id string) error {
	_, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteCategory(id)
}

func (s *categoryService) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	var result []dto.CategoryResponse
	for _, c := range categories {
		result = append(result, dto.CategoryResponse{
			ID:   c.ID.String(),
			Name: c.Name,
		})
	}
	return result, nil
}

func (s *categoryService) GetCategoryByID(id string) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:   category.ID.String(),
		Name: category.Name,
	}, nil
}
