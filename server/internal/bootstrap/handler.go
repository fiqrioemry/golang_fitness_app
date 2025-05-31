package bootstrap

import (
	"server/internal/handlers"
)

type HandlerContainer struct {
	DashboardHandler        *handlers.DashboardHandler
	AuthHandler             *handlers.AuthHandler
	UserHandler             *handlers.UserHandler
	TypeHandler             *handlers.TypeHandler
	ClassHandler            *handlers.ClassHandler
	LevelHandler            *handlers.LevelHandler
	ReviewHandler           *handlers.ReviewHandler
	ProfileHandler          *handlers.ProfileHandler
	PackageHandler          *handlers.PackageHandler
	UserPackageHandler      *handlers.UserPackageHandler
	BookingHandler          *handlers.BookingHandler
	VoucherHandler          *handlers.VoucherHandler
	PaymentHandler          *handlers.PaymentHandler
	CategoryHandler         *handlers.CategoryHandler
	LocationHandler         *handlers.LocationHandler
	InstructorHandler       *handlers.InstructorHandler
	SubcategoryHandler      *handlers.SubcategoryHandler
	NotificationHandler     *handlers.NotificationHandler
	ClassScheduleHandler    *handlers.ClassScheduleHandler
	ScheduleTemplateHandler *handlers.ScheduleTemplateHandler
}

func InitHandlers(svc *ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		AuthHandler:             handlers.NewAuthHandler(svc.AuthService),
		UserHandler:             handlers.NewUserHandler(svc.UserService),
		TypeHandler:             handlers.NewTypeHandler(svc.TypeService),
		ClassHandler:            handlers.NewClassHandler(svc.ClassService),
		LevelHandler:            handlers.NewLevelHandler(svc.LevelService),
		ReviewHandler:           handlers.NewReviewHandler(svc.ReviewService),
		ProfileHandler:          handlers.NewProfileHandler(svc.ProfileService),
		PaymentHandler:          handlers.NewPaymentHandler(svc.PaymentService),
		PackageHandler:          handlers.NewPackageHandler(svc.PackageService),
		VoucherHandler:          handlers.NewVoucherHandler(svc.VoucherService),
		BookingHandler:          handlers.NewBookingHandler(svc.BookingService),
		CategoryHandler:         handlers.NewCategoryHandler(svc.CategoryService),
		LocationHandler:         handlers.NewLocationHandler(svc.LocationService),
		DashboardHandler:        handlers.NewDashboardHandler(svc.DashboardService),
		InstructorHandler:       handlers.NewInstructorHandler(svc.InstructorService),
		UserPackageHandler:      handlers.NewUserPackageHandler(svc.UserPackageService),
		SubcategoryHandler:      handlers.NewSubcategoryHandler(svc.SubcategoryService),
		NotificationHandler:     handlers.NewNotificationHandler(svc.NotificationService),
		ClassScheduleHandler:    handlers.NewClassScheduleHandler(svc.ClassScheduleService),
		ScheduleTemplateHandler: handlers.NewScheduleTemplateHandler(svc.ScheduleTemplateService),
	}
}
