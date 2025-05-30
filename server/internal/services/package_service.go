package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"

	"github.com/google/uuid"
)

type PackageService interface {
	DeletePackage(id string) error
	CreatePackage(req dto.CreatePackageRequest) error
	UpdatePackage(id string, req dto.UpdatePackageRequest) error
	GetPackageByID(id string) (*dto.PackageDetailResponse, error)
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
		Price:          req.Price,
		Credit:         req.Credit,
		Discount:       req.Discount,
		Image:          req.ImageURL,
		IsActive:       req.IsActive,
		AdditionalList: req.Additional,
		Description:    req.Description,
		Classes:        classes,
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

	pkg.Name = req.Name
	pkg.Price = req.Price
	pkg.Credit = req.Credit
	pkg.Expired = req.Expired
	pkg.Description = req.Description

	if req.ImageURL != "" {
		pkg.Image = req.ImageURL
	}

	if req.Discount > 0 {
		pkg.Discount = req.Discount
	}

	if len(req.Additional) > 0 {
		pkg.AdditionalList = req.Additional
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

	pagination := utils.Paginate(total, params.Page, params.Limit)
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
