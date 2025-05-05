package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type PackageRepository interface {
	CreatePackage(pkg *models.Package) error
	UpdatePackage(pkg *models.Package) error
	DeletePackage(id string) error
	GetAllPackages() ([]models.Package, error)
	GetPackageByID(id string) (*models.Package, error)
	GetPackagesByClassID(classID string) ([]models.Package, error)
}

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) PackageRepository {
	return &packageRepository{db}
}

func (r *packageRepository) CreatePackage(pkg *models.Package) error {
	return r.db.Create(pkg).Error
}

func (r *packageRepository) UpdatePackage(pkg *models.Package) error {
	return r.db.Save(pkg).Error
}

func (r *packageRepository) DeletePackage(id string) error {
	return r.db.Delete(&models.Package{}, "id = ?", id).Error
}

func (r *packageRepository) GetAllPackages() ([]models.Package, error) {
	var packages []models.Package
	if err := r.db.Preload("Classes").Order("created_at desc").Find(&packages).Error; err != nil {
		return nil, err
	}
	return packages, nil
}

func (r *packageRepository) GetPackageByID(id string) (*models.Package, error) {
	var pkg models.Package
	if err := r.db.Preload("Classes").First(&pkg, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *packageRepository) GetPackagesByClassID(classID string) ([]models.Package, error) {
	var packages []models.Package
	err := r.db.Joins("JOIN package_classes ON packages.id = package_classes.package_id").
		Where("package_classes.class_id = ?", classID).
		Preload("Classes").
		Find(&packages).Error
	return packages, err
}
