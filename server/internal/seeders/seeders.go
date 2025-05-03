package seeders

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"server/internal/models"
	"server/internal/utils"
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

	customerUsers := []models.User{
		{
			ID:       uuid.New(),
			Email:    "customer@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Customer User",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Customer",
				Gender:   "female",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "elena.morris@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Elena Morris",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=EM",
				Gender:   "female",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "brandon.tan@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Brandon Tan",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=BT",
				Gender:   "male",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "yuki.sato@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Yuki Sato",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=YS",
				Gender:   "female",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "elena.morrisaga@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Elena Morris",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=EM",
				Gender:   "female",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "elvis.presley@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Elvis Presley",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=BT",
				Gender:   "male",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "david.jovovich@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "David Jovovich",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=BT",
				Gender:   "male",
			},
		},
	}

	instructorUsers := []models.User{
		{
			ID:       uuid.New(),
			Email:    "instructor1@example.com",
			Password: string(password),
			Role:     "instructor",
			Profile: models.Profile{
				Fullname: "Nurmalasari Permata",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=AB",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "instructor2@example.com",
			Password: string(password),
			Role:     "instructor",
			Profile: models.Profile{
				Fullname: "Eisenberg Josephine",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=ZA",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "instructor3@example.com",
			Password: string(password),
			Role:     "instructor",
			Profile: models.Profile{
				Fullname: "David Lee",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=GH",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "instructor4@example.com",
			Password: string(password),
			Role:     "instructor",
			Profile: models.Profile{
				Fullname: "Ellena Morris",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=GH",
			},
		},
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Println("Failed to seed admin:", err)
	}
	if err := db.Create(&customerUsers).Error; err != nil {
		log.Println("Failed to seed customers:", err)
	}
	if err := db.Create(&instructorUsers).Error; err != nil {
		log.Println("Failed to seed instructors:", err)
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
			Image:          "https://placehold.co/400x400/blue/white",
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
			Image:          "https://placehold.co/400x400/blue/white",
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
			Image:          "https://placehold.co/400x400/purple/white",
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
			Image:          "https://placehold.co/400x400/green/white",
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

func SeedClassGalleries(db *gorm.DB) {
	var count int64
	db.Model(&models.ClassGallery{}).Count(&count)
	if count > 0 {
		log.Println("ClassGalleries already seeded, skipping...")
		return
	}

	var classes []models.Class
	if err := db.Find(&classes).Error; err != nil {
		log.Println("Failed to fetch classes:", err)
		return
	}

	if len(classes) == 0 {
		log.Println("No classes found for gallery seeding.")
		return
	}

	var galleries []models.ClassGallery
	for _, class := range classes {
		for i := 1; i <= 3; i++ {
			galleries = append(galleries, models.ClassGallery{
				ID:        uuid.New(),
				ClassID:   class.ID,
				URL:       fmt.Sprintf("https://placehold.co/400x400?text=%s+Img%d", generateGalleryText(class.Title), i),
				CreatedAt: time.Now(),
			})
		}
	}

	if err := db.Create(&galleries).Error; err != nil {
		log.Printf("failed seeding class galleries: %v", err)
	} else {
		log.Println("✅ ClassGalleries seeding completed!")
	}
}

func generateGalleryText(title string) string {
	if len(title) == 0 {
		return "CLASS"
	}
	for i, r := range title {
		if r == ' ' {
			return title[:i]
		}
	}
	return title
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
			ID:             uuid.New(),
			Name:           "Trial Session",
			Description:    "1x Class Trial for new members.",
			Price:          500000,
			Credit:         1,
			Discount:       25,
			Expired:        14, // 14 hari
			AdditionalList: []string{"Valid for 14 days after first booking."},
			Image:          "https://placehold.co/400x400/orange/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "5 Sessions Package",
			Description:    "Enjoy 5 reformer classes package.",
			Price:          2250000,
			Credit:         5,
			Discount:       15,
			Expired:        60, // 2 bulan
			AdditionalList: []string{"Valid for 2 months after first booking.", "Cannot Be Refund"},
			Image:          "https://placehold.co/400x400/blue/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "10 Sessions Package",
			Description:    "Maximize your training with 10 sessions!",
			Price:          4100000,
			Credit:         10,
			Discount:       0,
			Expired:        120, // 4 bulan
			AdditionalList: []string{"Valid for 4 months after first booking.", "Cannot Be Refund"},
			Image:          "https://placehold.co/400x400/green/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "FTM x SANE Single Visit Promo",
			Description:    "Special promo for FTM x SANE group class.",
			Price:          100000,
			Credit:         1,
			Discount:       5,
			Expired:        14,
			AdditionalList: []string{"Valid for 14 days after first booking."},
			Image:          "https://placehold.co/400x400/orange/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "FTM x SANE Bundle 2 Classes",
			Description:    "Bundle of 2 classes for group sessions.",
			Price:          275000,
			Credit:         2,
			Discount:       0,
			Expired:        20,
			AdditionalList: []string{"Valid for 20 days after first booking."},
			Image:          "https://placehold.co/400x400/blue/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
	}

	if err := db.Create(&packages).Error; err != nil {
		log.Printf("❌ Failed seeding packages: %v", err)
	} else {
		log.Println("✅ Successfully seeded packages!")
	}

	var firstPackage models.Package

	var classes []models.Class

	if err := db.First(&firstPackage).Error; err != nil {
		log.Println("❌ Failed to fetch first package:", err)
		return
	}
	if err := db.Limit(2).Find(&classes).Error; err != nil {
		log.Println("❌ Failed to fetch 2 classes:", err)
		return
	}

	if len(classes) < 2 {
		log.Println("❌ Not enough classes for seeding package_classes")
		return
	}

	// GORM many2many: cukup assign ID class ke relasi slice
	if err := db.Model(&firstPackage).Association("Classes").Replace(&classes); err != nil {
		log.Println("❌ Failed to associate package with classes:", err)
		return
	}

	log.Println("✅ Successfully associated 1 package with 2 classes!")
}

func SeedInstructors(db *gorm.DB) {
	var count int64
	db.Model(&models.Instructor{}).Count(&count)

	if count > 0 {
		log.Println("Instructors already seeded, skipping...")
		return
	}

	// Cari user dengan role instructor
	var instructorsUser []models.User
	if err := db.Where("role = ?", "instructor").Find(&instructorsUser).Error; err != nil {
		log.Println("Failed to fetch instructor users:", err)
		return
	}

	if len(instructorsUser) == 0 {
		log.Println("No instructor users found, skipping instructor seeding...")
		return
	}

	var instructors []models.Instructor
	for _, user := range instructorsUser {
		instructors = append(instructors, models.Instructor{
			ID:             uuid.New(),
			UserID:         user.ID,
			Experience:     3,
			Specialties:    "Yoga, Reformer Pilates",
			Rating:         5.0,
			TotalClass:     0,
			Certifications: "Certified Yoga Teacher, Certified Reformer Pilates Instructor",
		})
	}

	if err := db.Create(&instructors).Error; err != nil {
		log.Printf("failed seeding instructors: %v", err)
	} else {
		log.Println("✅ Successfully seeded instructors!")
	}
}

func SeedPayments(db *gorm.DB) {
	var count int64
	db.Model(&models.Payment{}).Count(&count)

	if count > 0 {
		log.Println("Payments already seeded, skipping...")
		return
	}

	var user models.User

	var pkg models.Package
	if err := db.First(&user, "role = ?", "customer").Error; err != nil {
		log.Println("Failed to find customer user:", err)
		return
	}
	if err := db.First(&pkg).Error; err != nil {
		log.Println("Failed to find package:", err)
		return
	}
	taxRate := utils.GetTaxRate()
	discounted := pkg.Price * (1 - pkg.Discount/100)
	base := discounted
	tax := base * taxRate
	total := base + tax

	payments := []models.Payment{
		{
			ID:            uuid.New(),
			UserID:        user.ID,
			PackageID:     pkg.ID,
			PaymentMethod: "bank_transfer",
			BasePrice:     base,
			Tax:           tax,
			Total:         total,
			Status:        "success",
			PaidAt:        time.Now(),
		},
	}

	if err := db.Create(&payments).Error; err != nil {
		log.Printf("failed seeding payments: %v", err)
	} else {
		log.Println("✅ Payments seeding completed!")
	}
}

func SeedUserPackages(db *gorm.DB) {
	var count int64
	db.Model(&models.UserPackage{}).Count(&count)
	if count > 0 {
		log.Println("UserPackages already seeded, skipping...")
		return
	}

	// Ambil 2 user customer pertama
	var users []models.User
	if err := db.Where("role = ?", "customer").Limit(2).Find(&users).Error; err != nil || len(users) < 2 {
		log.Println("❌ Failed to fetch at least 2 customer users")
		return
	}

	// Ambil 2 package aktif pertama
	var packages []models.Package
	if err := db.Where("is_active = ?", true).Limit(2).Find(&packages).Error; err != nil || len(packages) < 2 {
		log.Println("❌ Failed to fetch at least 2 active packages")
		return
	}

	now := time.Now()
	var userPackages []models.UserPackage

	for _, user := range users {
		for _, pkg := range packages {
			expired := now.AddDate(0, 0, getExpiredDays(pkg))
			userPackages = append(userPackages, models.UserPackage{
				ID:              uuid.New(),
				UserID:          user.ID,
				PackageID:       pkg.ID,
				RemainingCredit: pkg.Credit,
				PurchasedAt:     now,
				ExpiredAt:       &expired,
			})
		}
	}

	if err := db.Create(&userPackages).Error; err != nil {
		log.Printf("❌ Failed seeding user packages: %v", err)
	} else {
		log.Println("✅ UserPackages seeding completed!")
	}
}

func getExpiredDays(pkg models.Package) int {
	return 30
}
func SeedClassSchedules(db *gorm.DB) {
	var count int64
	db.Model(&models.ClassSchedule{}).Count(&count)
	if count > 0 {
		log.Println("ClassSchedules already seeded, skipping...")
		return
	}

	var classes []models.Class
	var instructor models.Instructor

	// Ambil minimal 2 class & 1 instructor
	if err := db.Limit(2).Find(&classes).Error; err != nil || len(classes) < 2 {
		log.Println("Failed to find classes:", err)
		return
	}
	if err := db.First(&instructor).Error; err != nil {
		log.Println("Failed to find instructor:", err)
		return
	}

	colors := []string{
		"#f87171", "#fb923c", "#60a5fa", "#a78bfa", "#4ade80",
		"#f472b6", "#34d399", "#fbbf24", "#c084fc", "#10b981",
	}

	now := time.Now()
	date := now.AddDate(0, 0, 2).Truncate(24 * time.Hour)

	schedules := []models.ClassSchedule{
		// Awal
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date, StartHour: 10, StartMinute: 0, Capacity: 10,
			IsActive: true, Color: colors[0],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 2), StartHour: 9, StartMinute: 30, Capacity: 12,
			IsActive: true, Color: colors[1],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 2), StartHour: 11, StartMinute: 30, Capacity: 8,
			IsActive: true, Color: colors[2],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 2), StartHour: 14, StartMinute: 30, Capacity: 10,
			IsActive: true, Color: colors[3],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 2), StartHour: 12, StartMinute: 30, Capacity: 6,
			IsActive: true, Color: colors[4],
		},

		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 3), StartHour: 9, StartMinute: 45, Capacity: 15,
			IsActive: true, Color: colors[5],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 3), StartHour: 11, StartMinute: 0, Capacity: 10,
			IsActive: true, Color: colors[6],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 3), StartHour: 13, StartMinute: 30, Capacity: 12,
			IsActive: true, Color: colors[7],
		},
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 3), StartHour: 15, StartMinute: 15, Capacity: 8,
			IsActive: true, Color: colors[8],
		},

		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 4), StartHour: 10, StartMinute: 0, Capacity: 20,
			IsActive: true, Color: colors[9],
		},
		{
			ID: uuid.New(), ClassID: classes[1].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 4), StartHour: 10, StartMinute: 0, Capacity: 20,
			IsActive: true, Color: colors[0],
		},
	}

	if err := db.Create(&schedules).Error; err != nil {
		log.Printf("failed seeding class schedules: %v", err)
	} else {
		log.Println("✅ ClassSchedules seeding completed!")
	}
}

func SeedScheduleTemplates(db *gorm.DB) {
	var count int64
	db.Model(&models.ScheduleTemplate{}).Count(&count)
	if count > 0 {
		log.Println("ScheduleTemplates already seeded, skipping...")
		return
	}

	var classes []models.Class
	var instructors []models.Instructor
	if err := db.Find(&classes).Error; err != nil || len(classes) == 0 {
		log.Println("Failed to fetch classes:", err)
		return
	}
	if err := db.Find(&instructors).Error; err != nil || len(instructors) == 0 {
		log.Println("Failed to fetch instructors:", err)
		return
	}

	jsonDays := func(days []int) datatypes.JSON {
		j, _ := json.Marshal(days)
		return j
	}

	templates := []models.ScheduleTemplate{
		{
			ID:           uuid.New(),
			ClassID:      classes[0].ID,
			InstructorID: instructors[0].ID,
			DayOfWeeks:   jsonDays([]int{1, 3, 5}),
			StartHour:    10,
			StartMinute:  0,
			Capacity:     10,
			IsActive:     true,
		},
		{
			ID:           uuid.New(),
			ClassID:      classes[1%len(classes)].ID,
			InstructorID: instructors[1%len(instructors)].ID,
			DayOfWeeks:   jsonDays([]int{2}), // Selasa
			StartHour:    14,
			StartMinute:  30,
			Capacity:     12,
			IsActive:     true,
		},
		{
			ID:           uuid.New(),
			ClassID:      classes[2%len(classes)].ID,
			InstructorID: instructors[2%len(instructors)].ID,
			DayOfWeeks:   jsonDays([]int{0, 6}), // Minggu, Sabtu
			StartHour:    18,
			StartMinute:  0,
			Capacity:     8,
			IsActive:     true,
		},
		// Tambahan 4 data baru
		{
			ID:           uuid.New(),
			ClassID:      classes[0].ID,
			InstructorID: instructors[0].ID,
			DayOfWeeks:   jsonDays([]int{1}),
			StartHour:    9,
			StartMinute:  0,
			Capacity:     15,
			IsActive:     true,
		},
		{
			ID:           uuid.New(),
			ClassID:      classes[1].ID,
			InstructorID: instructors[1].ID,
			DayOfWeeks:   jsonDays([]int{3}),
			StartHour:    11,
			StartMinute:  30,
			Capacity:     20,
			IsActive:     true,
		},
		{
			ID:           uuid.New(),
			ClassID:      classes[2].ID,
			InstructorID: instructors[2].ID,
			DayOfWeeks:   jsonDays([]int{5}),
			StartHour:    15,
			StartMinute:  0,
			Capacity:     18,
			IsActive:     true,
		},
		// Bentrok di hari dan jam yang sama
		{
			ID:           uuid.New(),
			ClassID:      classes[0].ID,
			InstructorID: instructors[0].ID,
			DayOfWeeks:   jsonDays([]int{4}),
			StartHour:    13,
			StartMinute:  0,
			Capacity:     10,
			IsActive:     true,
		},
		{
			ID:           uuid.New(),
			ClassID:      classes[1].ID,
			InstructorID: instructors[1].ID,
			DayOfWeeks:   jsonDays([]int{4}),
			StartHour:    13,
			StartMinute:  0,
			Capacity:     10,
			IsActive:     true,
		},
	}

	if err := db.Create(&templates).Error; err != nil {
		log.Printf("failed seeding schedule templates: %v", err)
	} else {
		log.Println("✅ ScheduleTemplates seeding completed!")
	}
}

func SeedBookings(db *gorm.DB) {
	var count int64
	db.Model(&models.Booking{}).Count(&count)

	if count > 0 {
		log.Println("Bookings already seeded, skipping...")
		return
	}

	// Fetch customer user
	var user models.User
	if err := db.First(&user, "role = ?", "customer").Error; err != nil {
		log.Println("Failed to find customer user:", err)
		return
	}

	// Fetch class schedules
	var schedules []models.ClassSchedule
	if err := db.Limit(3).Find(&schedules).Error; err != nil {
		log.Println("Failed to fetch class schedules:", err)
		return
	}

	if len(schedules) == 0 {
		log.Println("No class schedules available for booking seeding.")
		return
	}

	// Create dummy bookings
	var bookings []models.Booking
	for _, schedule := range schedules {
		booking := models.Booking{
			ID:              uuid.New(),
			UserID:          user.ID,
			ClassScheduleID: schedule.ID,
			Status:          "booked",
		}
		bookings = append(bookings, booking)
	}

	if err := db.Create(&bookings).Error; err != nil {
		log.Printf("failed seeding bookings: %v", err)
	} else {
		log.Println("✅ Bookings seeding completed!")
	}
}

func SeedAttendances(db *gorm.DB) {
	var count int64
	db.Model(&models.Attendance{}).Count(&count)

	if count > 0 {
		log.Println("Attendances already seeded, skipping...")
		return
	}

	// Fetch bookings
	var bookings []models.Booking
	if err := db.Limit(3).Find(&bookings).Error; err != nil {
		log.Println("Failed to fetch bookings:", err)
		return
	}

	if len(bookings) == 0 {
		log.Println("No bookings available for attendance seeding.")
		return
	}

	var attendances []models.Attendance
	now := time.Now()

	for _, booking := range bookings {
		attendance := models.Attendance{
			ID:              uuid.New(),
			UserID:          booking.UserID,
			ClassScheduleID: booking.ClassScheduleID,
			Status:          "attended",
			CheckedAt:       &now,
		}
		attendances = append(attendances, attendance)
	}

	if err := db.Create(&attendances).Error; err != nil {
		log.Printf("failed seeding attendances: %v", err)
	} else {
		log.Println("✅ Attendances seeding completed!")
	}
}

func SeedReviews(db *gorm.DB) {
	var count int64
	db.Model(&models.Review{}).Count(&count)

	if count > 0 {
		log.Println("Reviews already seeded, skipping...")
		return
	}

	// Fetch user (customer)
	var user models.User
	if err := db.First(&user, "role = ?", "customer").Error; err != nil {
		log.Println("Failed to find customer user:", err)
		return
	}

	// Fetch classes
	var classes []models.Class
	if err := db.Limit(3).Find(&classes).Error; err != nil {
		log.Println("Failed to fetch classes:", err)
		return
	}

	if len(classes) == 0 {
		log.Println("No classes found for review seeding.")
		return
	}

	// Create dummy reviews
	var reviews []models.Review
	for i, class := range classes {
		review := models.Review{
			ID:      uuid.New(),
			UserID:  user.ID,
			ClassID: class.ID,
			Rating:  4 + (i % 2), // 4 atau 5 bervariasi
			Comment: "Great class experience!",
		}
		reviews = append(reviews, review)
	}

	if err := db.Create(&reviews).Error; err != nil {
		log.Printf("failed seeding reviews: %v", err)
	} else {
		log.Println("✅ Reviews seeding completed!")
	}
}
