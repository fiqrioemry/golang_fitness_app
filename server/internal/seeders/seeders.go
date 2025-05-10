package seeders

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
			Email:    "customer01@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Customer User 01",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Customer",
				Gender:   "female",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "customer02@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Customer User 02",
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

	allUsers := []models.User{adminUser}
	allUsers = append(allUsers, customerUsers...)
	allUsers = append(allUsers, instructorUsers...)

	for _, user := range allUsers {
		generateNotificationSettingsForUser(db, user)
	}

	log.Println("✅ User seeding completed with notification settings!")
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
	rand.Seed(time.Now().UnixNano())

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
		log.Printf("Failed seeding packages: %v", err)
	} else {
		log.Println("✅ Successfully seeded packages!")
	}

	var classes []models.Class
	if err := db.Find(&classes).Error; err != nil {
		log.Println("Failed to fetch classes:", err)
		return
	}

	var allPackages []models.Package
	if err := db.Find(&allPackages).Error; err != nil {
		log.Println("Failed to fetch packages:", err)
		return
	}

	// Pastikan setiap class punya minimal 1 relasi package
	classHasPackage := map[uuid.UUID]bool{}

	for i, pkg := range allPackages {
		n := 1 + rand.Intn(3)

		selected := map[uuid.UUID]bool{}
		var selectedClasses []models.Class
		for len(selectedClasses) < n {
			idx := rand.Intn(len(classes))
			class := classes[idx]
			if !selected[class.ID] {
				selected[class.ID] = true
				selectedClasses = append(selectedClasses, class)
				classHasPackage[class.ID] = true
			}
		}

		if err := db.Model(&allPackages[i]).Association("Classes").Replace(&selectedClasses); err != nil {
			log.Printf("Failed associating package %s: %v", pkg.Name, err)
		} else {
			log.Printf("✅ Associated %d classes to package %s", len(selectedClasses), pkg.Name)
		}
	}

	for _, class := range classes {
		if !classHasPackage[class.ID] {
			randomPkg := allPackages[rand.Intn(len(allPackages))]
			if err := db.Model(&randomPkg).Association("Classes").Append(&class); err != nil {
				log.Printf("Failed assigning fallback package to class %s: %v", class.Title, err)
			} else {
				log.Printf("Fallback assigned 1 package to class %s", class.Title)
			}
		}
	}

	log.Println("✅ Package-Class association completed!")
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

	// Ambil user dengan email customer01 dan customer02
	var users []models.User
	if err := db.Where("email IN ?", []string{"customer01@example.com", "customer02@example.com"}).Find(&users).Error; err != nil || len(users) < 2 {
		log.Println("Failed to fetch users with specified emails")
		return
	}

	// Ambil 2 package aktif pertama
	var packages []models.Package
	if err := db.Where("is_active = ?", true).Limit(2).Find(&packages).Error; err != nil || len(packages) < 2 {
		log.Println("Failed to fetch at least 2 active packages")
		return
	}

	now := time.Now()
	var userPackages []models.UserPackage

	// Assign 1 package untuk setiap user
	for i := 0; i < 2; i++ {
		pkg := packages[i]
		user := users[i]
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

	if err := db.Create(&userPackages).Error; err != nil {
		log.Printf("Failed seeding user packages: %v", err)
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

	now := time.Now()
	date := now.AddDate(0, 0, 2).Truncate(24 * time.Hour)

	schedules := []models.ClassSchedule{
		{
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID,
			Date: date, StartHour: 9, StartMinute: 0, Capacity: 10,
			IsActive: true, Color: "#60a5fa",
		},
		{
			ID: uuid.New(), ClassID: classes[1].ID, InstructorID: instructor.ID,
			Date: date.AddDate(0, 0, 1), StartHour: 10, StartMinute: 30, Capacity: 12,
			IsActive: true, Color: "#a78bfa",
		},
	}

	if err := db.Create(&schedules).Error; err != nil {
		log.Printf("failed seeding class schedules: %v", err)
	} else {
		log.Println("✅ ClassSchedules seeding completed!")
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

func SeedScheduleTemplate(db *gorm.DB) {
	if err := db.Exec("DELETE FROM recurrence_rules").Error; err != nil {
		log.Println("Failed to clear recurrence_rules:", err)
	} else {
		log.Println("🧹 RecurrenceRules cleared")
	}

	if err := db.Exec("DELETE FROM schedule_templates").Error; err != nil {
		log.Println("Failed to clear schedule_templates:", err)
	} else {
		log.Println("🧹 ScheduleTemplates cleared")
	}

	log.Println("✅ ScheduleTemplate reset completed!")
}

func SeedNotificationTypes(db *gorm.DB) {
	defaultTypes := []models.NotificationType{
		{ID: uuid.New(), Code: "system_message", Title: "System Announcement", Category: "announcement", DefaultEnabled: true}, // for login info, payment notification, booking success, and other
		{ID: uuid.New(), Code: "class_reminder", Title: "Class Reminder", Category: "reminder", DefaultEnabled: false},
		{ID: uuid.New(), Code: "daily_reminder", Title: "Daily Class Reminder", Category: "reminder", DefaultEnabled: false},
		{ID: uuid.New(), Code: "promo_offer", Title: "New Promotion Available", Category: "promotion", DefaultEnabled: false},
	}
	for _, t := range defaultTypes {
		db.FirstOrCreate(&t, "code = ?", t.Code)
	}
}

func generateNotificationSettingsForUser(db *gorm.DB, user models.User) {
	var notifTypes []models.NotificationType
	if err := db.Find(&notifTypes).Error; err != nil {
		log.Printf("Failed to get notification types for user %s: %v", user.Email, err)
		return
	}

	for _, nt := range notifTypes {
		for _, channel := range []string{"email", "browser"} {
			setting := models.NotificationSetting{
				ID:                 uuid.New(),
				UserID:             user.ID,
				NotificationTypeID: nt.ID,
				Channel:            channel,
				Enabled:            nt.DefaultEnabled,
			}
			if err := db.Create(&setting).Error; err != nil {
				log.Printf("Failed to create notification setting for user %s: %v", user.Email, err)
			}
		}
	}
}

func SeedDummyNotifications(db *gorm.DB) {
	var user models.User
	if err := db.Where("email = ?", "customer01@example.com").First(&user).Error; err != nil {
		log.Println("customer01@example.com not found")
		return
	}

	notifications := []models.Notification{
		{
			ID:       uuid.New(),
			UserID:   user.ID,
			TypeCode: "class_reminder",
			Title:    "Upcoming Class Reminder",
			Message:  "Don't forget your class starts in 1 hour!",
			Channel:  "browser",
			IsRead:   false,
		},
		{
			ID:       uuid.New(),
			UserID:   user.ID,
			TypeCode: "promo_offer",
			Title:    "Special Promo Just for You",
			Message:  "Get 20% off your next class using code: FIT20",
			Channel:  "browser",
			IsRead:   false,
		},
	}

	if err := db.Create(&notifications).Error; err != nil {
		log.Printf("Failed to seed dummy notifications: %v", err)
	} else {
		log.Println("✅ Dummy notifications for customer01@example.com seeded!")
	}
}
func SeedVouchers(db *gorm.DB) {
	var count int64
	db.Model(&models.Voucher{}).Count(&count)
	if count > 0 {
		log.Println("Vouchers already seeded, skipping...")
		return
	}

	now := time.Now()
	expired := now.AddDate(0, 1, 0)

	max1 := 30000.0
	max2 := 50000.0

	voucher1 := models.Voucher{
		ID:           uuid.New(),
		Code:         "FIT50",
		Description:  "Dapatkan diskon 50% hingga 30.000",
		DiscountType: "percentage",
		Discount:     50,
		MaxDiscount:  &max1,
		Quota:        10,
		IsReusable:   false,
		ExpiredAt:    expired,
		CreatedAt:    now,
	}

	voucher2 := models.Voucher{
		ID:           uuid.New(),
		Code:         "HEALTHY100K",
		Description:  "Diskon langsung 100.000",
		DiscountType: "fixed",
		Discount:     100000,
		MaxDiscount:  &max2,
		Quota:        10,
		IsReusable:   true,
		ExpiredAt:    expired,
		CreatedAt:    now,
	}

	if err := db.Create([]models.Voucher{voucher1, voucher2}).Error; err != nil {
		log.Printf("Failed to seed vouchers: %v", err)
		return
	}

	log.Println("✅ Vouchers seeding completed!")

}
