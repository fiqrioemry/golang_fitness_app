package seeders

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"server/internal/models"
)

func SeedUsers(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	adminUser := models.User{
		ID:       uuid.New(),
		Email:    "admin@example.com",
		Password: string(password),
		Role:     "admin",
		Profile: models.Profile{
			Fullname: "Admin User",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Admin",
		},
	}

	customerUser := models.User{
		ID:       uuid.New(),
		Email:    "customer@example.com",
		Password: string(password),
		Role:     "customer",
		Profile: models.Profile{
			Fullname: "Customer User",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Customer",
		},
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Println("Failed to seed admin:", err)
	}
	if err := db.Create(&customerUser).Error; err != nil {
		log.Println("Failed to seed customer:", err)
	}

	log.Println("✅ User seeding completed")
}

func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Category{}).Count(&count)

	if count > 0 {
		log.Println("Categories already seeded, skipping...")
		return
	}

	categories := []models.Category{
		{ID: uuid.New(), Name: "Yoga"},
		{ID: uuid.New(), Name: "Strength Training"},
		{ID: uuid.New(), Name: "Cardio"},
		{ID: uuid.New(), Name: "Martial Arts"},
		{ID: uuid.New(), Name: "Pilates"},
		{ID: uuid.New(), Name: "Dance Fitness"},
		{ID: uuid.New(), Name: "Meditation"},
		{ID: uuid.New(), Name: "Crossfit"},
		{ID: uuid.New(), Name: "Bootcamp"},
		{ID: uuid.New(), Name: "Kids Classes"},
	}

	if err := db.Create(&categories).Error; err != nil {
		log.Printf("failed seeding categories: %v", err)
	} else {
		log.Println("successfully seeded categories!")
	}
}

func SeedSubcategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Subcategory{}).Count(&count)

	if count > 0 {
		log.Println("Subcategories already seeded, skipping...")
		return
	}

	// Map: categoryName -> list of subcategories
	subcategoryData := map[string][]string{
		"Pilates": {
			"Mat Pilates",
			"Reformer Pilates",
			"Clinical Pilates",
		},
		"Yoga": {
			"Hatha Yoga",
			"Vinyasa Yoga",
			"Yin Yoga",
			"Power Yoga",
		},
		"Cardio": {
			"HIIT",
			"Zumba",
			"Indoor Cycling",
		},
	}

	// Fetch all categories
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Printf("failed to fetch categories: %v", err)
		return
	}

	categoryMap := make(map[string]uuid.UUID)
	for _, c := range categories {
		categoryMap[c.Name] = c.ID
	}

	var subcategories []models.Subcategory
	for categoryName, subList := range subcategoryData {
		categoryID, exists := categoryMap[categoryName]
		if !exists {
			log.Printf("category %s not found, skipping subcategories...", categoryName)
			continue
		}
		for _, sub := range subList {
			subcategories = append(subcategories, models.Subcategory{
				ID:         uuid.New(),
				Name:       sub,
				CategoryID: categoryID,
			})
		}
	}

	if err := db.Create(&subcategories).Error; err != nil {
		log.Printf("failed seeding subcategories: %v", err)
	} else {
		log.Println("successfully seeded subcategories!")
	}
}
