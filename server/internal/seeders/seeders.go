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
