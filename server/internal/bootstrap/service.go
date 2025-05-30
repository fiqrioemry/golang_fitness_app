package bootstrap

import (
	"server/internal/services"

	"gorm.io/gorm"
)

type ServiceContainer struct {
	DashboardService        services.DashboardService
	AuthService             services.AuthService
	UserService             services.UserService
	TypeService             services.TypeService
	ClassService            services.ClassService
	LevelService            services.LevelService
	ReviewService           services.ReviewService
	ProfileService          services.ProfileService
	PackageService          services.PackageService
	LocationService         services.LocationService
	CategoryService         services.CategoryService
	VoucherService          services.VoucherService
	PaymentService          services.PaymentService
	BookingService          services.BookingService
	InstructorService       services.InstructorService
	SubcategoryService      services.SubcategoryService
	NotificationService     services.NotificationService
	ClassScheduleService    services.ClassScheduleService
	ScheduleTemplateService services.ScheduleTemplateService
}

func InitServices(repo *RepositoryContainer, db *gorm.DB) *ServiceContainer {
	voucherService := services.NewVoucherService(repo.VoucherRepo)
	notificationService := services.NewNotificationService(repo.NotificationRepo)
	scheduleTemplateService :=
		services.NewScheduleTemplateService(repo.ScheduleTemplateRepo, repo.ClassRepo, repo.ClassScheduleRepo)

	return &ServiceContainer{
		VoucherService:          voucherService,
		NotificationService:     notificationService,
		ScheduleTemplateService: scheduleTemplateService,
		UserService:             services.NewUserService(repo.UserRepo),
		TypeService:             services.NewTypeService(repo.TypeRepo),
		ClassService:            services.NewClassService(repo.ClassRepo),
		LevelService:            services.NewLevelService(repo.LevelRepo),
		ProfileService:          services.NewProfileService(repo.ProfileRepo),
		PackageService:          services.NewPackageService(repo.PackageRepo),
		LocationService:         services.NewLocationService(repo.LocationRepo),
		CategoryService:         services.NewCategoryService(repo.CategoryRepo),
		DashboardService:        services.NewDashboardService(repo.DashboardRepo),
		SubcategoryService:      services.NewSubcategoryService(repo.SubcategoryRepo),
		AuthService:             services.NewAuthService(repo.AuthRepo, repo.NotificationRepo),
		InstructorService:       services.NewInstructorService(repo.InstructorRepo, repo.AuthRepo),
		ReviewService:           services.NewReviewService(repo.ReviewRepo, repo.BookingRepo, repo.InstructorRepo),
		PaymentService:          services.NewPaymentService(repo.PaymentRepo, repo.PackageRepo, repo.UserPackageRepo, repo.AuthRepo, voucherService, notificationService),
		BookingService:          services.NewBookingService(db, repo.BookingRepo, repo.ClassScheduleRepo, repo.UserPackageRepo, repo.PackageRepo, notificationService),
		ClassScheduleService:    services.NewClassScheduleService(repo.ClassScheduleRepo, repo.ClassRepo, repo.PackageRepo, repo.BookingRepo, repo.ScheduleTemplateRepo, scheduleTemplateService),
	}
}
