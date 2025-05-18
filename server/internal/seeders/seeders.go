package seeders

import (
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

	log.Println("User seeding completed with notification settings!")
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
		{ID: uuid.New(), Name: "Pilates"},
		{ID: uuid.New(), Name: "Cardio"},
		{ID: uuid.New(), Name: "Strength Training"},
		{ID: uuid.New(), Name: "Martial Arts"},
	}

	if err := db.Create(&categories).Error; err != nil {
		log.Printf("failed seeding categories: %v", err)
	} else {
		log.Println("Categories seeding completed!")
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
		"Yoga":         {"Hatha Yoga", "Vinyasa Yoga"},
		"Pilates":      {"Mat Pilates", "Reformer Pilates"},
		"Cardio":       {"HIIT", "Zumba", "Aerobic Dance"},
		"Martial Arts": {"Boxing", "Muay Thai", "Kickboxing"},
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
		log.Println("Subcategories seeding completed!")
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
		{ID: uuid.New(), Name: "Private Class"},
		{ID: uuid.New(), Name: "Virtual Class"},
	}

	if err := db.Create(&types).Error; err != nil {
		log.Printf("failed seeding types: %v", err)
	} else {
		log.Println("Successfully seeded types!")
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
		log.Println("Successfully seeded levels!")
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
			Name:        "Sweat Up Studio A",
			Address:     "123 Fitness St, New York, NY",
			GeoLocation: "40.712776,-74.005974",
		},
		{
			ID:          uuid.New(),
			Name:        "Sweat Up Studio B",
			Address:     "456 Gym Ave, Los Angeles, CA",
			GeoLocation: "34.052235,-118.243683",
		},
	}

	if err := db.Create(&locations).Error; err != nil {
		log.Printf("failed seeding locations: %v", err)
	} else {
		log.Println("Successfully seeded locations!")
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
	var subcategories []models.Subcategory
	var types []models.Type
	var levels []models.Level
	var locations []models.Location

	db.Find(&categories)
	db.Find(&subcategories)
	db.Find(&types)
	db.Find(&levels)
	db.Find(&locations)

	categoryMap := make(map[string]uuid.UUID)
	for _, c := range categories {
		categoryMap[c.Name] = c.ID
	}

	subcategoryMap := make(map[string]uuid.UUID)
	for _, s := range subcategories {
		subcategoryMap[s.Name] = s.ID
	}

	classes := []models.Class{
		{
			ID: uuid.New(), Title: "Hatha Yoga for Beginners",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879286/fitness_booking_app/fx8mm2yumhlpen2fckgf.webp", Duration: 60,
			Description:    "This class is designed for those who are new to yoga. With a focus on breathing techniques, gentle stretches, and foundational postures, participants are introduced to the calming and strengthening principles of yoga. The goal is to cultivate awareness of body alignment and mental stillness. The instructor will guide each movement at a steady pace, ensuring comfort and safety for all levels. This class is ideal for reducing stress, improving flexibility, and setting the stage for a long-term yoga practice.",
			AdditionalList: []string{"Hatha Yoga", "Yoga"},
			TypeID:         types[0].ID, LevelID: levels[0].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Yoga"], SubcategoryID: subcategoryMap["Hatha Yoga"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Vinyasa Flow Intermediate",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879490/fitness_booking_app/o5uyu6r3qtmmb2py0c6f.webp", Duration: 75,
			Description:    "Vinyasa Flow is a dynamic yoga style that links breath with movement. This intermediate-level class focuses on building strength and flexibility through a fluid sequence of poses. Expect smooth transitions, moderate intensity, and a creative flow that challenges both body and mind. Participants should have some yoga experience, as the class moves at a faster pace. It’s perfect for deepening your practice, enhancing balance, and developing endurance in a mindful way.",
			AdditionalList: []string{"Vinyasa Yoga", "Yoga"},
			TypeID:         types[0].ID, LevelID: levels[1].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Yoga"], SubcategoryID: subcategoryMap["Vinyasa Yoga"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Mat Pilates Core Challenge",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879424/fitness_booking_app/ckauvapfizz7fbyzrd9h.webp", Duration: 60,
			Description:    "This Mat Pilates class is centered on strengthening the core muscles using bodyweight exercises. With a series of precise and controlled movements, participants develop stability, coordination, and postural alignment, which are essential for everyday mobility and injury prevention. It’s suitable for all levels and serves as a great complement to other workouts. Over time, you'll notice improvements in core strength, spinal alignment, and total-body awareness.",
			AdditionalList: []string{"Mat Pilates", "Pilates"},
			TypeID:         types[0].ID, LevelID: levels[2].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Pilates"], SubcategoryID: subcategoryMap["Mat Pilates"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Reformer Pilates Sculpt",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879388/fitness_booking_app/ew3uvgaxgecyaxz0xmgv.webp", Duration: 60,
			Description:    "Using the Reformer machine, this class adds resistance training to the Pilates experience. Participants will engage in slow, deliberate movements that sculpt lean muscles and improve flexibility without putting stress on the joints. Each session targets core, glutes, legs, and arms while enhancing balance and coordination. Reformer Pilates is ideal for anyone seeking low-impact yet highly effective full-body conditioning",
			AdditionalList: []string{"Reformer Pilates", "Pilates"},
			TypeID:         types[1].ID, LevelID: levels[1].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Pilates"], SubcategoryID: subcategoryMap["Reformer Pilates"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "HIIT Total Body Burn",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879477/fitness_booking_app/uudvfiyquve8cd4fdexh.webp", Duration: 45,
			Description:    "This High-Intensity Interval Training (HIIT) class is designed to burn maximum calories in minimum time. It combines cardio drills and bodyweight strength exercises in timed intervals, making it both efficient and effective. You’ll improve cardiovascular fitness, metabolism, and muscular endurance. Suitable for all fitness levels, this class offers modifications for beginners and challenges for advanced participants.",
			AdditionalList: []string{"HIIT", "Cardio"},
			TypeID:         types[0].ID, LevelID: levels[2].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Cardio"], SubcategoryID: subcategoryMap["HIIT"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Zumba Dance Energy",
			Image: "https://placehold.co/400x400", Duration: 60,
			Description:    "Zumba Dance Energy is a fun and high-energy workout that blends Latin rhythms with easy-to-follow dance routines. It’s a full-body cardio session that feels more like a dance party than a workout.Perfect for all levels, this class improves endurance, burns calories, and boosts your mood. No dance experience is required—just come ready to move and enjoy the beat!",
			AdditionalList: []string{"Zumba", "Cardio"},
			TypeID:         types[0].ID, LevelID: levels[3].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Cardio"], SubcategoryID: subcategoryMap["Zumba"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Boxing Fundamentals",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879364/fitness_booking_app/cfawvclumu5hjtgviini.webp", Duration: 60,
			Description:    "Learn the basics of boxing in this empowering, non-contact class. Participants will be introduced to proper stance, punches, footwork, and defensive techniques while improving coordination and reaction time. Boxing a full-body workout that enhances strength, agility, and mental focus. This class is beginner-friendly and also a great stress reliever.",
			AdditionalList: []string{"Boxing", "Martial Arts"},
			TypeID:         types[0].ID, LevelID: levels[0].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Martial Arts"], SubcategoryID: subcategoryMap["Boxing"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Muay Thai Conditioning",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879461/fitness_booking_app/lbt5otnjqiujqcokugbm.webp", Duration: 75,
			Description:    "Muay Thai is a powerful martial art that combines striking techniques with intense physical conditioning. In this class, participants will practice punches, kicks, elbows, and knees in controlled combinations, along with bodyweight drills for endurance. hether you're training for fitness or technique, Muay Thai offers an excellent way to improve power, discipline, and cardiovascular performance in one dynamic session.",
			AdditionalList: []string{"Muay Thai", "Martial Arts"},
			TypeID:         types[0].ID, LevelID: levels[1].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Martial Arts"], SubcategoryID: subcategoryMap["Muay Thai"], IsActive: true,
		},
	}

	if err := db.Create(&classes).Error; err != nil {
		log.Printf("failed seeding classes: %v", err)
	} else {
		log.Println("Successfully seeded 8 professional fitness classes!")
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

	// Mapping: class title → image URLs
	galleryMap := map[string][]string{
		"Hatha Yoga for Beginners": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880275/fitness_booking_app/zfuaa53iljyvzl9ux3wc.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880274/fitness_booking_app/uluaivkhpndaymslbkht.webp",
		},
		"Vinyasa Flow Intermediate": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879960/fitness_booking_app/ez14mm0svmysibda2rzj.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879959/fitness_booking_app/ozzdnzx5zvncibeknfqb.webp",
		},
		"HIIT Total Body Burn": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880255/fitness_booking_app/zaar3efvfdctw3qed79m.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880254/fitness_booking_app/di0zpt8a9l3rjcykpwah.webp",
		},
		"Zumba Dance Energy": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880169/fitness_booking_app/vnnrrr1nv6kt3yoq56nz.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880169/fitness_booking_app/ocmjzmyyhrkvudcqbq2q.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880168/fitness_booking_app/tolouxu0lecscg832dwi.webp",
		},
		"Mat Pilates Core Challenge": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880141/fitness_booking_app/uwlxbomqugaliffncscs.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880140/fitness_booking_app/tozp2hvl25kllrmq1amp.webp",
		},
		"Reformer Pilates Sculpt": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880139/fitness_booking_app/si4hk7k5txauk4rju517.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880141/fitness_booking_app/uwlxbomqugaliffncscs.webp",
		},
		"Boxing Fundamentals": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880115/fitness_booking_app/p7e0pl96xn8s7tnwponl.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880114/fitness_booking_app/dqdjuest25qrjqzhplup.webp",
		},
		"Muay Thai Conditioning": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880228/fitness_booking_app/ckskvayptwsjdytj1gus.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880227/fitness_booking_app/v78tjjzijecqoswyrmxf.webp",
		},
	}

	var galleries []models.ClassGallery
	for _, class := range classes {
		images := galleryMap[class.Title]
		for _, url := range images {
			galleries = append(galleries, models.ClassGallery{
				ID:        uuid.New(),
				ClassID:   class.ID,
				URL:       url,
				CreatedAt: time.Now(),
			})
		}
	}

	if err := db.Create(&galleries).Error; err != nil {
		log.Printf("failed seeding class galleries: %v", err)
	} else {
		log.Println("ClassGalleries seeding completed with Cloudinary URLs!")
	}
}

func SeedPackages(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())

	var count int64
	db.Model(&models.Package{}).Count(&count)
	if count > 0 {
		log.Println("Packages already seeded, skipping...")
		return
	}

	// Load semua class dan kelompokkan berdasarkan category
	var classes []models.Class
	if err := db.Preload("Category").Find(&classes).Error; err != nil {
		log.Fatalf("Failed fetching classes: %v", err)
	}

	categoryClasses := map[string][]models.Class{}
	for _, class := range classes {
		categoryClasses[class.Category.Name] = append(categoryClasses[class.Category.Name], class)
	}

	packages := []models.Package{
		{
			ID:             uuid.New(),
			Name:           "Yoga Wellness Trial",
			Description:    "Try 1 Yoga class to relieve stress and boost flexibility.",
			Price:          120000,
			Credit:         1,
			Discount:       20,
			Expired:        14,
			AdditionalList: []string{"Valid for 14 days after first booking."},
			Image:          "https://placehold.co/400x400/green/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "Pilates Core Pack (5x)",
			Description:    "Enjoy 5 Mat/Reformer Pilates sessions to build core strength and posture.",
			Price:          600000,
			Credit:         5,
			Discount:       10,
			Expired:        60,
			AdditionalList: []string{"Valid for 2 months", "Non-refundable"},
			Image:          "https://placehold.co/400x400/blue/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "Cardio Burnout Pass (10x)",
			Description:    "Get your heart pumping with 10 sessions of HIIT, Zumba, and Aerobic workouts.",
			Price:          1000000,
			Credit:         10,
			Discount:       15,
			Expired:        120,
			AdditionalList: []string{"Valid for 4 months", "Non-refundable"},
			Image:          "https://placehold.co/400x400/orange/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "Combat Starter (1x)",
			Description:    "Experience our Martial Arts class in one exciting session.",
			Price:          150000,
			Credit:         1,
			Discount:       0,
			Expired:        14,
			AdditionalList: []string{"Valid for 14 days after booking."},
			Image:          "https://placehold.co/400x400/red/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Name:           "Warrior Pack (5x)",
			Description:    "Boost your skills with 5 sessions of Boxing, Muay Thai, or Kickboxing.",
			Price:          650000,
			Credit:         5,
			Discount:       5,
			Expired:        60,
			AdditionalList: []string{"Valid for 2 months", "No refunds after activation."},
			Image:          "https://placehold.co/400x400/black/white",
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
	}

	if err := db.Create(&packages).Error; err != nil {
		log.Printf("Failed seeding packages: %v", err)
		return
	}
	log.Println("Successfully seeded packages!")

	packageClassMap := map[string]string{
		"Yoga Wellness Trial":       "Yoga",
		"Pilates Core Pack (5x)":    "Pilates",
		"Cardio Burnout Pass (10x)": "Cardio",
		"Combat Starter (1x)":       "Martial Arts",
		"Warrior Pack (5x)":         "Martial Arts",
	}

	for i, pkg := range packages {
		categoryName := packageClassMap[pkg.Name]
		classList := categoryClasses[categoryName]

		if len(classList) == 0 {
			log.Printf("⚠️ No class found for category: %s", categoryName)
			continue
		}

		n := min(2+rand.Intn(3), len(classList))
		rand.Shuffle(len(classList), func(i, j int) { classList[i], classList[j] = classList[j], classList[i] })
		selected := classList[:n]

		if err := db.Model(&packages[i]).Association("Classes").Replace(&selected); err != nil {
			log.Printf(" Failed associating package %s: %v", pkg.Name, err)
		} else {
			log.Printf("Associated %d classes to package %s", len(selected), pkg.Name)
		}
	}

	log.Println("Package-Class association per category completed!")
}

func SeedInstructors(db *gorm.DB) {
	var count int64
	db.Model(&models.Instructor{}).Count(&count)

	if count > 0 {
		log.Println("Instructors already seeded, skipping...")
		return
	}
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
		log.Println("Successfully seeded instructors!")
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

	now := time.Now()
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
			PaidAt:        now,
		},
		{
			ID:            uuid.New(),
			UserID:        user.ID,
			PackageID:     pkg.ID,
			PaymentMethod: "bank_transfer",
			BasePrice:     base,
			Tax:           tax,
			Total:         total,
			Status:        "success",
			PaidAt:        now,
		},
		{
			ID:            uuid.New(),
			UserID:        user.ID,
			PackageID:     pkg.ID,
			PaymentMethod: "bank_transfer",
			BasePrice:     base,
			Tax:           tax,
			Total:         total,
			Status:        "success",
			PaidAt:        now.AddDate(0, 0, -1),
		},
		{
			ID:            uuid.New(),
			UserID:        user.ID,
			PackageID:     pkg.ID,
			PaymentMethod: "bank_transfer",
			BasePrice:     base,
			Tax:           tax,
			Total:         total,
			Status:        "success",
			PaidAt:        now.AddDate(0, 0, -2),
		},
		{
			ID:            uuid.New(),
			UserID:        user.ID,
			PackageID:     pkg.ID,
			PaymentMethod: "bank_transfer",
			BasePrice:     base,
			Tax:           tax,
			Total:         total,
			Status:        "failed",
			PaidAt:        now.AddDate(0, 0, -2),
		},
	}

	if err := db.Create(&payments).Error; err != nil {
		log.Printf("failed seeding payments: %v", err)
	} else {
		log.Println("Payments seeding completed with multiple dates!")
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
	for i := range 2 {
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
		log.Println("UserPackages seeding completed!")
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
			ID: uuid.New(), ClassID: classes[0].ID, InstructorID: instructor.ID, ClassName: classes[0].Title, ClassImage: classes[0].Image,
			Date: date, StartHour: 9, StartMinute: 0, Capacity: 10, Duration: classes[0].Duration,
			IsActive: true, Color: "#60a5fa",
		},
		{
			ID: uuid.New(), ClassID: classes[1].ID, ClassName: classes[1].Title, ClassImage: classes[1].Image, InstructorID: instructor.ID, InstructorName: instructor.User.Profile.Fullname,
			Date: date.AddDate(0, 0, 1), StartHour: 10, StartMinute: 30, Duration: classes[1].Duration, Capacity: 12,
			IsActive: true, Color: "#a78bfa",
		},
	}

	if err := db.Create(&schedules).Error; err != nil {
		log.Printf("failed seeding class schedules: %v", err)
	} else {
		log.Println("ClassSchedules seeding completed!")
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
		log.Println("Bookings seeding completed!")
	}
}

func SeedAttendances(db *gorm.DB) {
	var count int64
	db.Model(&models.Attendance{}).Count(&count)

	if count > 0 {
		log.Println("Attendances already seeded, skipping...")
		return
	}

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
		log.Println("Attendances seeding completed!")
	}
}

func SeedReviews(db *gorm.DB) {
	var count int64
	db.Model(&models.Review{}).Count(&count)

	if count > 0 {
		log.Println("Reviews already seeded, skipping...")
		return
	}

	var user models.User
	if err := db.First(&user, "role = ?", "customer").Error; err != nil {
		log.Println("Failed to find customer user:", err)
		return
	}

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
			Rating:  4 + (i % 2),
			Comment: "Great class experience!",
		}
		reviews = append(reviews, review)
	}

	if err := db.Create(&reviews).Error; err != nil {
		log.Printf("failed seeding reviews: %v", err)
	} else {
		log.Println("Reviews seeding completed!")
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

	log.Println("ScheduleTemplate reset completed!")
}

func SeedNotificationTypes(db *gorm.DB) {
	defaultTypes := []models.NotificationType{
		{ID: uuid.New(), Code: "system_message", Title: "System Announcement", Category: "announcement", DefaultEnabled: false},
		{ID: uuid.New(), Code: "class_reminder", Title: "Class Reminder", Category: "reminder", DefaultEnabled: false},
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
		log.Println("Dummy notifications for customer01@example.com seeded!")
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

	log.Println("Vouchers seeding completed!")

}
