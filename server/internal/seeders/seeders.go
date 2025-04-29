package seeders

import (
	"log"
	"time"

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

	log.Println("✅ User seeding completed!")
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
		log.Println("✅ Categories seeding completed!")
	}
}

func SeedSubcategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Subcategory{}).Count(&count)

	if count > 0 {
		log.Println("Subcategories already seeded, skipping...")
		return
	}

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
		log.Println("✅ subcategories seeding completed!")
	}
}

func SeedTypes(db *gorm.DB) {
	var count int64
	db.Model(&models.Type{}).Count(&count)

	if count > 0 {
		log.Println("Types already seeded, skipping...")
		return
	}

	types := []models.Type{
		{ID: uuid.New(), Name: "Group Class"},
		{ID: uuid.New(), Name: "Personal Training"},
		{ID: uuid.New(), Name: "Virtual Class"},
		{ID: uuid.New(), Name: "Outdoor Training"},
		{ID: uuid.New(), Name: "Workshop"},
	}

	if err := db.Create(&types).Error; err != nil {
		log.Printf("failed seeding types: %v", err)
	} else {
		log.Println("✅ Successfully seeded types!")
	}
}

func SeedLevels(db *gorm.DB) {
	var count int64
	db.Model(&models.Level{}).Count(&count)

	if count > 0 {
		log.Println("Levels already seeded, skipping...")
		return
	}

	levels := []models.Level{
		{ID: uuid.New(), Name: "Beginner"},
		{ID: uuid.New(), Name: "Intermediate"},
		{ID: uuid.New(), Name: "Advanced"},
		{ID: uuid.New(), Name: "All Levels"},
	}

	if err := db.Create(&levels).Error; err != nil {
		log.Printf("failed seeding levels: %v", err)
	} else {
		log.Println("✅ Successfully seeded levels!")
	}
}

func SeedLocations(db *gorm.DB) {
	var count int64
	db.Model(&models.Location{}).Count(&count)

	if count > 0 {
		log.Println("Locations already seeded, skipping...")
		return
	}

	locations := []models.Location{
		{
			ID:          uuid.New(),
			Name:        "Fitness Studio A",
			Address:     "123 Fitness St, New York, NY",
			GeoLocation: "40.712776,-74.005974", // Contoh lat, long
		},
		{
			ID:          uuid.New(),
			Name:        "Gym B",
			Address:     "456 Gym Ave, Los Angeles, CA",
			GeoLocation: "34.052235,-118.243683", // Contoh lat, long
		},
		{
			ID:          uuid.New(),
			Name:        "Yoga Center C",
			Address:     "789 Yoga Rd, San Francisco, CA",
			GeoLocation: "37.774929,-122.419418", // Contoh lat, long
		},
	}

	if err := db.Create(&locations).Error; err != nil {
		log.Printf("failed seeding locations: %v", err)
	} else {
		log.Println("✅ Successfully seeded locations!")
	}
}

func SeedClasses(db *gorm.DB) {
	var count int64
	db.Model(&models.Class{}).Count(&count)

	if count > 0 {
		log.Println("Classes already seeded, skipping...")
		return
	}

	// Fetch necessary related data
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Printf("failed to fetch categories: %v", err)
		return
	}

	var subcategories []models.Subcategory
	if err := db.Find(&subcategories).Error; err != nil {
		log.Printf("failed to fetch subcategories: %v", err)
		return
	}

	var types []models.Type
	if err := db.Find(&types).Error; err != nil {
		log.Printf("failed to fetch types: %v", err)
		return
	}

	var levels []models.Level
	if err := db.Find(&levels).Error; err != nil {
		log.Printf("failed to fetch levels: %v", err)
		return
	}

	var locations []models.Location
	if err := db.Find(&locations).Error; err != nil {
		log.Printf("failed to fetch locations: %v", err)
		return
	}

	// Create sample class data
	classes := []models.Class{
		{
			ID:             uuid.New(),
			Title:          "Yoga Beginners",
			Image:          "https://example.com/yoga-beginner.jpg",
			Duration:       60,
			Description:    "A gentle introduction to yoga.",
			AdditionalList: []string{"Beginner", "Stretching", "Breathing"},
			TypeID:         types[0].ID,  // Group Class
			LevelID:        levels[0].ID, // Beginner
			LocationID:     locations[0].ID,
			CategoryID:     categories[0].ID,    // Yoga
			SubcategoryID:  subcategories[0].ID, // Hatha Yoga
			IsActive:       true,
		},
		{
			ID:             uuid.New(),
			Title:          "Strength Training - Intermediate",
			Image:          "https://example.com/strength-training.jpg",
			Duration:       90,
			Description:    "A strength-building session for intermediate athletes.",
			AdditionalList: []string{"Intermediate", "Strength", "Weightlifting"},
			TypeID:         types[0].ID,  // Group Class
			LevelID:        levels[1].ID, // Intermediate
			LocationID:     locations[1].ID,
			CategoryID:     categories[1].ID,    // Strength Training
			SubcategoryID:  subcategories[1].ID, // Bodyweight Training
			IsActive:       true,
		},
		{
			ID:             uuid.New(),
			Title:          "Zumba Dance Party",
			Image:          "https://example.com/zumba-dance.jpg",
			Duration:       60,
			Description:    "A high-energy, fun dance workout.",
			AdditionalList: []string{"Dance", "Cardio", "Party"},
			TypeID:         types[0].ID,  // Group Class
			LevelID:        levels[2].ID, // Advanced
			LocationID:     locations[2].ID,
			CategoryID:     categories[2].ID,    // Cardio
			SubcategoryID:  subcategories[2].ID, // Zumba
			IsActive:       true,
		},
		{
			ID:             uuid.New(),
			Title:          "Private Yoga Session",
			Image:          "https://example.com/private-yoga.jpg",
			Duration:       45,
			Description:    "A one-on-one session with a yoga instructor.",
			AdditionalList: []string{"Private", "Yoga", "Therapy"},
			TypeID:         types[1].ID,
			LevelID:        levels[0].ID,
			LocationID:     locations[0].ID,
			CategoryID:     categories[0].ID,
			SubcategoryID:  subcategories[0].ID,
			IsActive:       true,
		},
	}

	if err := db.Create(&classes).Error; err != nil {
		log.Printf("failed seeding classes: %v", err)
	} else {
		log.Println("✅ Successfully seeded classes!")
	}
}

func SeedPackages(db *gorm.DB) {
	var count int64
	db.Model(&models.Package{}).Count(&count)

	if count > 0 {
		log.Println("Packages already seeded, skipping...")
		return
	}

	packages := []models.Package{
		{
			ID:          uuid.New(),
			Name:        "Trial Session",
			Description: "1x Class Trial for new members.",
			Price:       500000,
			Credit:      1,
			Expired:     intPtr(14), // 14 hari
			Information: "Valid for 14 days after first booking.",
			Image:       "https://example.com/package-trial.jpg",
			IsActive:    true,
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "5 Sessions Package",
			Description: "Enjoy 5 reformer classes package.",
			Price:       2250000,
			Credit:      5,
			Expired:     intPtr(60), // 2 bulan
			Information: "Valid for 2 months after first booking.",
			Image:       "https://example.com/package-5sessions.jpg",
			IsActive:    true,
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "10 Sessions Package",
			Description: "Maximize your training with 10 sessions!",
			Price:       4100000,
			Credit:      10,
			Expired:     intPtr(120), // 4 bulan
			Information: "Valid for 4 months after first booking.",
			Image:       "https://example.com/package-10sessions.jpg",
			IsActive:    true,
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "FTM x SANE Single Visit Promo",
			Description: "Special promo for FTM x SANE group class.",
			Price:       100000,
			Credit:      1,
			Expired:     intPtr(14),
			Information: "Valid for 14 days after first booking.",
			Image:       "https://example.com/package-ftm-promo.jpg",
			IsActive:    true,
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "FTM x SANE Bundle 2 Classes",
			Description: "Bundle of 2 classes for group sessions.",
			Price:       275000,
			Credit:      2,
			Expired:     intPtr(20),
			Information: "Valid for 20 days after first booking.",
			Image:       "https://example.com/package-ftm-bundle.jpg",
			IsActive:    true,
			CreatedAt:   time.Now(),
		},
	}

	if err := db.Create(&packages).Error; err != nil {
		log.Printf("failed seeding packages: %v", err)
	} else {
		log.Println("✅ Successfully seeded packages!")
	}
}

func intPtr(i int) *int {
	return &i
}
