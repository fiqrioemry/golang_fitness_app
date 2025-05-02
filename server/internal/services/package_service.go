package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type PackageService interface {
	CreatePackage(req dto.CreatePackageRequest) error
	UpdatePackage(id string, req dto.UpdatePackageRequest) error
	DeletePackage(id string) error
	GetAllPackages() ([]dto.PackageResponse, error)
	GetPackageByID(id string) (*dto.PackageDetailResponse, error)
}

type packageService struct {
	repo repositories.PackageRepository
}

func NewPackageService(repo repositories.PackageRepository) PackageService {
	return &packageService{repo}
}

func (s *packageService) CreatePackage(req dto.CreatePackageRequest) error {

	var classes []models.Class
	for _, classID := range req.ClassIDs {
		classUUID, err := uuid.Parse(classID)
		if err != nil {
			return fmt.Errorf("invalid class ID: %s", classID)
		}
		classes = append(classes, models.Class{ID: classUUID})
	}
	pkg := models.Package{
		ID:             uuid.New(),
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		Credit:         req.Credit,
		Discount:       req.Discount,
		Image:          req.ImageURL,
		IsActive:       req.IsActive,
		AdditionalList: req.Additional,
		Classes:        classes,
		CreatedAt:      time.Now(),
	}

	if req.Expired != 0 {
		pkg.Expired = req.Expired
	}

	return s.repo.CreatePackage(&pkg)
}

func (s *packageService) UpdatePackage(id string, req dto.UpdatePackageRequest) error {
	pkg, err := s.repo.GetPackageByID(id)
	if err != nil {
		return err
	}

	if req.Name != "" {
		pkg.Name = req.Name
	}
	if req.Description != "" {
		pkg.Description = req.Description
	}
	if req.Discount > 0 {
		pkg.Discount = req.Discount
	}
	if req.Price != 0 {
		pkg.Price = req.Price
	}
	if req.Credit != 0 {
		pkg.Credit = req.Credit
	}
	if req.Expired != 0 {
		pkg.Expired = req.Expired
	}
	if len(req.Additional) > 0 {
		pkg.AdditionalList = req.Additional
	}
	if req.ImageURL != "" {
		pkg.Image = req.ImageURL
	}
	pkg.IsActive = req.IsActive

	if len(req.ClassIDs) > 0 {
		var classes []models.Class
		for _, classID := range req.ClassIDs {
			classUUID, err := uuid.Parse(classID)
			if err != nil {
				return fmt.Errorf("invalid class ID: %s", classID)
			}
			classes = append(classes, models.Class{ID: classUUID})
		}
		pkg.Classes = classes
	}

	return s.repo.UpdatePackage(pkg)
}

func (s *packageService) DeletePackage(id string) error {
	return s.repo.DeletePackage(id)
}
func (s *packageService) GetAllPackages() ([]dto.PackageResponse, error) {
	packages, err := s.repo.GetAllPackages()
	if err != nil {
		return nil, err
	}

	var result []dto.PackageResponse
	for _, p := range packages {
		var classes []dto.ClassSummaryResponse
		for _, c := range p.Classes {
			classes = append(classes, dto.ClassSummaryResponse{
				ID:       c.ID.String(),
				Title:    c.Title,
				Image:    c.Image,
				Duration: c.Duration,
			})
		}

		result = append(result, dto.PackageResponse{
			ID:          p.ID.String(),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Credit:      p.Credit,
			Image:       p.Image,
			Discount:    p.Discount,
			Expired:     p.Expired,
			IsActive:    p.IsActive,
			Additional:  p.AdditionalList,
			Classes:     classes,
		})
	}
	return result, nil
}

func (s *packageService) GetPackageByID(id string) (*dto.PackageDetailResponse, error) {
	pkg, err := s.repo.GetPackageByID(id)
	if err != nil {
		return nil, err
	}

	var classes []dto.ClassSummaryResponse
	for _, c := range pkg.Classes {
		classes = append(classes, dto.ClassSummaryResponse{
			ID:       c.ID.String(),
			Title:    c.Title,
			Image:    c.Image,
			Duration: c.Duration,
		})
	}

	return &dto.PackageDetailResponse{
		ID:          pkg.ID.String(),
		Name:        pkg.Name,
		Description: pkg.Description,
		Price:       pkg.Price,
		Credit:      pkg.Credit,
		Discount:    pkg.Discount,
		Expired:     pkg.Expired,
		Image:       pkg.Image,
		IsActive:    pkg.IsActive,
		Additional:  pkg.AdditionalList,
		Classes:     classes,
	}, nil
}
