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
	fmt.Println("DEBUG IsActive in service:", req.IsActive)

	pkg := models.Package{
		ID:             uuid.New(),
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		Credit:         req.Credit,
		Image:          req.ImageURL,
		IsActive:       req.IsActive,
		AdditionalList: req.Additional,
		CreatedAt:      time.Now(),
	}
	if req.Expired != 0 {
		pkg.Expired = req.Expired
	}
	fmt.Println("DEBUG IsActive in service:", pkg.IsActive)
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
	if req.Price != 0 {
		pkg.Price = req.Price
	}
	if req.Credit != 0 {
		pkg.Credit = req.Credit
	}
	if len(req.Additional) > 0 {
		pkg.AdditionalList = req.Additional
	}

	pkg.IsActive = req.IsActive

	if req.Expired != 0 {
		pkg.Expired = req.Expired
	}
	if req.ImageURL != "" {
		pkg.Image = req.ImageURL
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
		result = append(result, dto.PackageResponse{
			ID:          p.ID.String(),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Credit:      p.Credit,
			Image:       p.Image,
			Expired:     p.Expired,
			IsActive:    p.IsActive,
			Additional:  p.AdditionalList,
		})
	}
	return result, nil
}

func (s *packageService) GetPackageByID(id string) (*dto.PackageDetailResponse, error) {
	pkg, err := s.repo.GetPackageByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.PackageDetailResponse{
		ID:          pkg.ID.String(),
		Name:        pkg.Name,
		Description: pkg.Description,
		Price:       pkg.Price,
		Credit:      pkg.Credit,
		Expired:     pkg.Expired,
		Image:       pkg.Image,
		IsActive:    pkg.IsActive,
		Additional:  pkg.AdditionalList,
	}, nil
}
