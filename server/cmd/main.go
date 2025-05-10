package main

import (
	"log"
	"os"
	"server/internal/config"
	"server/internal/cron"
	"server/internal/handlers"
	"server/internal/middleware"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/seeders"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	config.InitRedis()
	config.InitMailer()
	config.InitCloudinary()
	config.InitMidtrans()

	db := config.DB
	r := gin.Default()
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(5, 10),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/payments"}),
	)

	// ========== Seeder ==========
	seeders.ResetDatabase(db)

	// ========== Repository Layer ==========
	authRepo := repositories.NewAuthRepository(db)
	userRepo := repositories.NewUserRepository(db)
	typeRepo := repositories.NewTypeRepository(db)
	voucherRepo := repositories.NewVoucherRepository(db)

	classRepo := repositories.NewClassRepository(db)
	levelRepo := repositories.NewLevelRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)
	profileRepo := repositories.NewProfileRepository(db)
	packageRepo := repositories.NewPackageRepository(db)
	paymentRepo := repositories.NewPaymentRepository(db)
	bookingRepo := repositories.NewBookingRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	locationRepo := repositories.NewLocationRepository(db)
	instructorRepo := repositories.NewInstructorRepository(db)
	attendanceRepo := repositories.NewAttendanceRepository(db)
	subcategoryRepo := repositories.NewSubcategoryRepository(db)
	userPackageRepo := repositories.NewUserPackageRepository(db)
	notificationRepo := repositories.NewNotificationRepository(db)
	classScheduleRepo := repositories.NewClassScheduleRepository(db)
	scheduleTemplateRepo := repositories.NewScheduleTemplateRepository(db)

	// ========== Service Layer ==========

	authService := services.NewAuthService(authRepo, notificationRepo)
	typeService := services.NewTypeService(typeRepo)
	userService := services.NewUserService(userRepo)
	classService := services.NewClassService(classRepo)
	levelService := services.NewLevelService(levelRepo)
	reviewService := services.NewReviewService(reviewRepo)
	profileService := services.NewProfileService(profileRepo)
	packageService := services.NewPackageService(packageRepo)
	locationService := services.NewLocationService(locationRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	voucherService := services.NewVoucherService(voucherRepo)

	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	notificationService := services.NewNotificationService(notificationRepo)
	instructorService := services.NewInstructorService(instructorRepo, authRepo)
	attendanceService := services.NewAttendanceService(attendanceRepo, bookingRepo, reviewRepo)
	paymentService := services.NewPaymentService(paymentRepo, packageRepo, userPackageRepo, authRepo, voucherService)
	bookingService := services.NewBookingService(bookingRepo, classScheduleRepo, userPackageRepo, packageRepo)
	scheduleTemplateService := services.NewScheduleTemplateService(scheduleTemplateRepo, classRepo, classScheduleRepo)
	classScheduleService := services.NewClassScheduleService(classScheduleRepo, classRepo, packageRepo, userPackageRepo, bookingRepo)

	// ========== Handler Layer ==========
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	voucherHandler := handlers.NewVoucherHandler(voucherService)
	typeHandler := handlers.NewTypeHandler(typeService)
	levelHandler := handlers.NewLevelHandler(levelService)
	classHandler := handlers.NewClassHandler(classService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	profileHandler := handlers.NewProfileHandler(profileService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	packageHandler := handlers.NewPackageHandler(packageService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	locationHandler := handlers.NewLocationHandler(locationService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	instructorHandler := handlers.NewInstructorHandler(instructorService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)
	scheduleTemplateHandler := handlers.NewScheduleTemplateHandler(scheduleTemplateService)
	classScheduleHandler := handlers.NewClassScheduleHandler(classScheduleService, scheduleTemplateService)

	// ========== Cron Job ==========
	cronManager := cron.NewCronManager(paymentService, scheduleTemplateService, notificationService)
	cronManager.RegisterJobs()
	cronManager.Start()

	// StartBackupJob()

	// ========== Route Binding ==========
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler)
	routes.TypeRoutes(r, typeHandler)
	routes.ClassRoutes(r, classHandler)
	routes.LevelRoutes(r, levelHandler)
	routes.ReviewRoutes(r, reviewHandler)
	routes.VoucherRoutes(r, voucherHandler)
	routes.PackageRoutes(r, packageHandler)
	routes.ProfileRoutes(r, profileHandler)
	routes.PaymentRoutes(r, paymentHandler)
	routes.BookingRoutes(r, bookingHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.LocationRoutes(r, locationHandler)
	routes.InstructorRoutes(r, instructorHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)
	routes.AttendanceRoutes(r, attendanceHandler)
	routes.NotificationRoutes(r, notificationHandler)
	routes.ClassScheduleRoutes(r, classScheduleHandler)
	routes.ScheduleTemplateRoutes(r, scheduleTemplateHandler)

	// ========== Start Server ==========
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
