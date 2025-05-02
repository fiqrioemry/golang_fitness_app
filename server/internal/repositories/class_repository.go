package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	CreateClass(class *models.Class) error
	UpdateClass(class *models.Class) error
	DeleteClass(id string) error
	GetClassByID(id string) (*models.Class, error)
	GetAllClasses(filter map[string]interface{}, search string, sort string, limit, offset int) ([]models.Class, int64, error)
	GetActiveClasses() ([]models.Class, error)
	SaveClassGalleries(galleries []models.ClassGallery) error
	DeleteClassGalleryByID(id string) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db}
}

func (r *classRepository) CreateClass(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *classRepository) UpdateClass(class *models.Class) error {
	return r.db.Save(class).Error
}

func (r *classRepository) DeleteClass(id string) error {
	return r.db.Delete(&models.Class{}, "id = ?", id).Error
}

func (r *classRepository) GetClassByID(id string) (*models.Class, error) {
	var class models.Class
	if err := r.db.
		Preload("Type").
		Preload("Level").
		Preload("Category").
		Preload("Subcategory").
		Preload("Location").
		Preload("Galleries").
		Preload("Reviews.User.Profile").
		First(&class, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *classRepository) GetAllClasses(filter map[string]interface{}, search string, sort string, limit, offset int) ([]models.Class, int64, error) {
	var classes []models.Class
	var count int64
	db := r.db.Model(&models.Class{})

	for key, value := range filter {
		db = db.Where(key+" = ?", value)
	}

	if search != "" {
		db = db.Where("title LIKE ?", "%"+search+"%")
	}

	switch sort {
	case "oldest":
		db = db.Order("created_at asc")
	case "title_asc":
		db = db.Order("title asc")
	case "title_desc":
		db = db.Order("title desc")
	default:
		db = db.Order("created_at desc")
	}

	db.Count(&count)
	if err := db.Preload("Galleries").Limit(limit).Offset(offset).Find(&classes).Error; err != nil {
		return nil, 0, err
	}

	return classes, count, nil
}

func (r *classRepository) GetActiveClasses() ([]models.Class, error) {
	var classes []models.Class
	if err := r.db.Where("is_active = ?", true).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) SaveClassGalleries(galleries []models.ClassGallery) error {
	return r.db.Create(&galleries).Error
}

func (r *classRepository) DeleteClassGalleryByID(id string) error {
	return r.db.Delete(&models.ClassGallery{}, "id = ?", id).Error
}
