package seeders

import (
	"log"
	"server/internal/models"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) {
	log.Println("⚠️ Dropping all tables...")

	err := db.Migrator().DropTable(
		&models.Token{},
		&models.Profile{},
		&models.User{},
		&models.Package{},
		&models.PackageClass{},
		&models.UserPackage{},
		&models.Category{},
		&models.Subcategory{},
		&models.Type{},
		&models.Level{},
		&models.Class{},
		&models.ClassGallery{},
		&models.ClassSchedule{},
		&models.ScheduleTemplate{},
		&models.Booking{},
		&models.Payment{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Review{},
		&models.Attendance{},
		&models.Instructor{},
		&models.Location{},
	)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	log.Println("All tables dropped successfully.")

	log.Println("Migrating tables...")

	err = db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Profile{},
		&models.Package{},
		&models.PackageClass{},
		&models.UserPackage{},
		&models.Category{},
		&models.Subcategory{},
		&models.Type{},
		&models.Level{},
		&models.Class{},
		&models.ClassGallery{},
		&models.ClassSchedule{},
		&models.ScheduleTemplate{},
		&models.Booking{},
		&models.Payment{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Review{},
		&models.Attendance{},
		&models.Instructor{},
		&models.Location{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	log.Println("Migration completed successfully.")

	log.Println("Seeding dummy data...")

	SeedNotificationTypes(db)
	SeedUsers(db)
	SeedCategories(db)
	SeedSubcategories(db)
	SeedTypes(db)
	SeedLevels(db)
	SeedLocations(db)
	SeedClasses(db)
	SeedClassGalleries(db)
	SeedPackages(db)
	SeedInstructors(db)
	SeedPayments(db)
	SeedUserPackages(db)
	SeedClassSchedules(db)
	SeedReviews(db)
	SeedDummyNotifications(db)
	SeedVouchers(db)
	SeedFutureBookingForCustomer1(db)

	log.Println("Seeding completed successfully.")
}
