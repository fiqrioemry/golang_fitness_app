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
	LocationHandler         *handlers.LocationHandler
	CategoryHandler         *handlers.CategoryHandler
	VoucherHandler          *handlers.VoucherHandler
	SubcategoryHandler      *handlers.SubcategoryHandler
	NotificationHandler     *handlers.NotificationHandler
	InstructorHandler       *handlers.InstructorHandler
	PaymentHandler          *handlers.PaymentHandler
	BookingHandler          *handlers.BookingHandler
	ScheduleTemplateHandler *handlers.ScheduleTemplateHandler
	ClassScheduleHandler    *handlers.ClassScheduleHandler
}

func InitHandlers(svc *ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		DashboardHandler:        handlers.NewDashboardHandler(svc.DashboardService),
		AuthHandler:             handlers.NewAuthHandler(svc.AuthService),
		UserHandler:             handlers.NewUserHandler(svc.UserService),
		TypeHandler:             handlers.NewTypeHandler(svc.TypeService),
		ClassHandler:            handlers.NewClassHandler(svc.ClassService),
		LevelHandler:            handlers.NewLevelHandler(svc.LevelService),
		ReviewHandler:           handlers.NewReviewHandler(svc.ReviewService),
		ProfileHandler:          handlers.NewProfileHandler(svc.ProfileService),
		PackageHandler:          handlers.NewPackageHandler(svc.PackageService),
		LocationHandler:         handlers.NewLocationHandler(svc.LocationService),
		CategoryHandler:         handlers.NewCategoryHandler(svc.CategoryService),
		VoucherHandler:          handlers.NewVoucherHandler(svc.VoucherService),
		SubcategoryHandler:      handlers.NewSubcategoryHandler(svc.SubcategoryService),
		NotificationHandler:     handlers.NewNotificationHandler(svc.NotificationService),
		InstructorHandler:       handlers.NewInstructorHandler(svc.InstructorService),
		PaymentHandler:          handlers.NewPaymentHandler(svc.PaymentService),
		BookingHandler:          handlers.NewBookingHandler(svc.BookingService),
		ScheduleTemplateHandler: handlers.NewScheduleTemplateHandler(svc.ScheduleTemplateService),
		ClassScheduleHandler:    handlers.NewClassScheduleHandler(svc.ClassScheduleService),
	}
}
