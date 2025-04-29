package main

import (
	"log"
	"os"
	"server/internal/config"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/seeders"
	"server/internal/services"
	"server/internal/utils"

	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	config.InitRedis()
	config.InitMailer()
	config.InitCloudinary()

	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

	db := config.DB

	seeders.ResetDatabase(db)

	// auth
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// profile
	profileRepo := repositories.NewProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)
	profileHandler := handlers.NewProfileHandler(profileService)

	// class
	classRepo := repositories.NewClassRepository(db)
	classService := services.NewClassService(classRepo)
	classHandler := handlers.NewClassHandler(classService)

	//  Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// subcategory
	subcategoryRepo := repositories.NewSubcategoryRepository(db)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)

	// type class
	typeRepo := repositories.NewTypeRepository(db)
	typeService := services.NewTypeService(typeRepo)
	typeHandler := handlers.NewTypeHandler(typeService)

	// level class
	levelRepo := repositories.NewLevelRepository(db)
	levelService := services.NewLevelService(levelRepo)
	levelHandler := handlers.NewLevelHandler(levelService)

	// location class
	locationRepo := repositories.NewLocationRepository(db)
	locationService := services.NewLocationService(locationRepo)
	locationHandler := handlers.NewLocationHandler(locationService)

	// Repository
	packageRepo := repositories.NewPackageRepository(db)
	packageService := services.NewPackageService(packageRepo)
	packageHandler := handlers.NewPackageHandler(packageService)

	// Repository
	instructorRepo := repositories.NewInstructorRepository(db)
	instructorService := services.NewInstructorService(instructorRepo, authRepo)
	instructorHandler := handlers.NewInstructorHandler(instructorService)

	// UserPackage
	userPackageRepo := repositories.NewUserPackageRepository(db)
	userPackageService := services.NewUserPackageService(userPackageRepo)
	userPackageHandler := handlers.NewUserPackageHandler(userPackageService)

	// Payment
	paymentRepo := repositories.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo, packageRepo, userPackageRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// ClassSchedule
	classScheduleRepo := repositories.NewClassScheduleRepository(db)
	classScheduleService := services.NewClassScheduleService(classScheduleRepo, classRepo)
	classScheduleHandler := handlers.NewClassScheduleHandler(classScheduleService)

	// Schedule template
	scheduleTemplateRepo := repositories.NewScheduleTemplateRepository(db)
	scheduleTemplateService := services.NewScheduleTemplateService(scheduleTemplateRepo, classRepo, classScheduleRepo)
	scheduleTemplateHandler := handlers.NewScheduleTemplateHandler(scheduleTemplateService)

	// Booking
	bookingRepo := repositories.NewBookingRepository(db)
	bookingService := services.NewBookingService(bookingRepo, classScheduleRepo, userPackageRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// Attendance
	attendanceRepo := repositories.NewAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepo, bookingRepo)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)

	// Review
	reviewRepo := repositories.NewReviewRepository(db)
	reviewService := services.NewReviewService(reviewRepo)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	routes.AuthRoutes(r, authHandler)
	routes.TypeRoutes(r, typeHandler)
	routes.ClassRoutes(r, classHandler)
	routes.LevelRoutes(r, levelHandler)
	routes.ReviewRoutes(r, reviewHandler)
	routes.PackageRoutes(r, packageHandler)
	routes.ProfileRoutes(r, profileHandler)
	routes.PaymentRoutes(r, paymentHandler)
	routes.BookingRoutes(r, bookingHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.LocationRoutes(r, locationHandler)
	routes.InstructorRoutes(r, instructorHandler)
	routes.AttendanceRoutes(r, attendanceHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)
	routes.UserPackageRoutes(r, userPackageHandler)
	routes.ClassScheduleRoutes(r, classScheduleHandler)
	routes.ScheduleTemplateRoutes(r, scheduleTemplateHandler)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
