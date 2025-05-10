package bootstrap

import (
	"server/internal/repositories"

	"gorm.io/gorm"
)

type RepositoryContainer struct {
	DashboardRepo        repositories.DashboardRepository
	AuthRepo             repositories.AuthRepository
	UserRepo             repositories.UserRepository
	TypeRepo             repositories.TypeRepository
	VoucherRepo          repositories.VoucherRepository
	ClassRepo            repositories.ClassRepository
	LevelRepo            repositories.LevelRepository
	ReviewRepo           repositories.ReviewRepository
	ProfileRepo          repositories.ProfileRepository
	PackageRepo          repositories.PackageRepository
	PaymentRepo          repositories.PaymentRepository
	BookingRepo          repositories.BookingRepository
	CategoryRepo         repositories.CategoryRepository
	LocationRepo         repositories.LocationRepository
	InstructorRepo       repositories.InstructorRepository
	AttendanceRepo       repositories.AttendanceRepository
	SubcategoryRepo      repositories.SubcategoryRepository
	UserPackageRepo      repositories.UserPackageRepository
	NotificationRepo     repositories.NotificationRepository
	ClassScheduleRepo    repositories.ClassScheduleRepository
	ScheduleTemplateRepo repositories.ScheduleTemplateRepository
}

func InitRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		DashboardRepo:        repositories.NewDashboardRepository(db),
		AuthRepo:             repositories.NewAuthRepository(db),
		UserRepo:             repositories.NewUserRepository(db),
		TypeRepo:             repositories.NewTypeRepository(db),
		VoucherRepo:          repositories.NewVoucherRepository(db),
		ClassRepo:            repositories.NewClassRepository(db),
		LevelRepo:            repositories.NewLevelRepository(db),
		ReviewRepo:           repositories.NewReviewRepository(db),
		ProfileRepo:          repositories.NewProfileRepository(db),
		PackageRepo:          repositories.NewPackageRepository(db),
		PaymentRepo:          repositories.NewPaymentRepository(db),
		BookingRepo:          repositories.NewBookingRepository(db),
		CategoryRepo:         repositories.NewCategoryRepository(db),
		LocationRepo:         repositories.NewLocationRepository(db),
		InstructorRepo:       repositories.NewInstructorRepository(db),
		AttendanceRepo:       repositories.NewAttendanceRepository(db),
		SubcategoryRepo:      repositories.NewSubcategoryRepository(db),
		UserPackageRepo:      repositories.NewUserPackageRepository(db),
		NotificationRepo:     repositories.NewNotificationRepository(db),
		ClassScheduleRepo:    repositories.NewClassScheduleRepository(db),
		ScheduleTemplateRepo: repositories.NewScheduleTemplateRepository(db),
	}
}
