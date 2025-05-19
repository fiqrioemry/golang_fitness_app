package bootstrap

import (
	"server/internal/services"
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
	SubcategoryService      services.SubcategoryService
	NotificationService     services.NotificationService
	InstructorService       services.InstructorService
	AttendanceService       services.AttendanceService
	PaymentService          services.PaymentService
	BookingService          services.BookingService
	ScheduleTemplateService services.ScheduleTemplateService
	ClassScheduleService    services.ClassScheduleService
}

func InitServices(repo *RepositoryContainer) *ServiceContainer {
	voucherSvc := services.NewVoucherService(repo.VoucherRepo)

	return &ServiceContainer{
		DashboardService:        services.NewDashboardService(repo.DashboardRepo),
		AuthService:             services.NewAuthService(repo.AuthRepo, repo.NotificationRepo),
		UserService:             services.NewUserService(repo.UserRepo),
		TypeService:             services.NewTypeService(repo.TypeRepo),
		ClassService:            services.NewClassService(repo.ClassRepo),
		LevelService:            services.NewLevelService(repo.LevelRepo),
		ReviewService:           services.NewReviewService(repo.ReviewRepo, repo.ClassScheduleRepo, repo.InstructorRepo),
		ProfileService:          services.NewProfileService(repo.ProfileRepo),
		PackageService:          services.NewPackageService(repo.PackageRepo),
		LocationService:         services.NewLocationService(repo.LocationRepo),
		CategoryService:         services.NewCategoryService(repo.CategoryRepo),
		VoucherService:          voucherSvc,
		SubcategoryService:      services.NewSubcategoryService(repo.SubcategoryRepo),
		NotificationService:     services.NewNotificationService(repo.NotificationRepo),
		InstructorService:       services.NewInstructorService(repo.InstructorRepo, repo.AuthRepo),
		AttendanceService:       services.NewAttendanceService(repo.AttendanceRepo, repo.BookingRepo, repo.ReviewRepo),
		PaymentService:          services.NewPaymentService(repo.PaymentRepo, repo.PackageRepo, repo.UserPackageRepo, repo.AuthRepo, voucherSvc),
		BookingService:          services.NewBookingService(repo.BookingRepo, repo.ClassScheduleRepo, repo.UserPackageRepo, repo.PackageRepo),
		ScheduleTemplateService: services.NewScheduleTemplateService(repo.ScheduleTemplateRepo, repo.ClassRepo, repo.ClassScheduleRepo),
		ClassScheduleService:    services.NewClassScheduleService(repo.ClassScheduleRepo, repo.ClassRepo, repo.PackageRepo, repo.UserPackageRepo, repo.BookingRepo),
	}
}
