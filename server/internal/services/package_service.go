package services

import (
	"errors"
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PackageService interface {
	DeletePackage(id string) error
	CreatePackage(req dto.CreatePackageRequest) error
	UpdatePackage(id string, req dto.UpdatePackageRequest) error
	GetPackageByID(id string) (*dto.PackageDetailResponse, error)
	GetRelatedPackages(classID string) ([]dto.PackageListResponse, error)
	GetAllPackages(params dto.PackageQueryParam) ([]dto.PackageListResponse, *dto.PaginationResponse, error)
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

func (s *packageService) GetAllPackages(params dto.PackageQueryParam) ([]dto.PackageListResponse, *dto.PaginationResponse, error) {
	packages, total, err := s.repo.GetAllPackages(params)
	if err != nil {
		return nil, nil, err
	}

	var results []dto.PackageListResponse
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

		results = append(results, dto.PackageListResponse{
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
	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))

	pagination := &dto.PaginationResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return results, pagination, nil
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

func (s *packageService) GetRelatedPackages(classID string) ([]dto.PackageListResponse, error) {
	packages, err := s.repo.GetPackagesByClassID(classID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dto.PackageListResponse{}, nil
		}
		return nil, err
	}

	var res []dto.PackageListResponse
	for _, pkg := range packages {
		res = append(res, dto.PackageListResponse{
			ID:          pkg.ID.String(),
			Name:        pkg.Name,
			Description: pkg.Description,
			Price:       pkg.Price,
			Credit:      pkg.Credit,
			Image:       pkg.Image,
			IsActive:    pkg.IsActive,
		})
	}
	return res, nil
}

func (s *packageService) DeletePackage(id string) error {
	userPackages, err := s.repo.GetUserPackagesWithRemainingCredit(id)
	if err != nil {
		return err
	}

	if len(userPackages) > 0 {
		return fmt.Errorf("package cannot be deleted, still in use by %d user(s) with remaining credit", len(userPackages))
	}

	return s.repo.DeletePackage(id)
}
