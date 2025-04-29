// package main

// import (
// 	"log"
// 	"os"
// 	"server/internal/config"
// 	"server/internal/handlers"
// 	"server/internal/middleware"
// 	"server/internal/repositories"
// 	"server/internal/routes"
// 	"server/internal/seeders"
// 	"server/internal/services"
// 	"server/internal/utils"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	utils.LoadEnv()
// 	config.InitDatabase()
// 	config.InitRedis()
// 	config.InitMailer()
// 	config.InitCloudinary()

// 	r := gin.Default()
// 	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

// 	db := config.DB

// 	seeders.ResetDatabase(db)

// 	authRepo := repositories.NewAuthRepository(db)
// 	typeRepo := repositories.NewTypeRepository(db)
// 	levelRepo := repositories.NewLevelRepository(db)
// 	classRepo := repositories.NewClassRepository(db)
// 	profileRepo := repositories.NewProfileRepository(db)
// 	packageRepo := repositories.NewPackageRepository(db)
// 	paymentRepo := repositories.NewPaymentRepository(db)
// 	categoryRepo := repositories.NewCategoryRepository(db)
// 	locationRepo := repositories.NewLocationRepository(db)
// 	instructorRepo := repositories.NewInstructorRepository(db)
// 	subcategoryRepo := repositories.NewSubcategoryRepository(db)
// 	userPackageRepo := repositories.NewUserPackageRepository(db)
// 	classScheduleRepo := repositories.NewClassScheduleRepository(db)
// 	scheduleTemplateRepo := repositories.NewScheduleTemplateRepository(db)

// 	authService := services.NewAuthService(authRepo)
// 	typeService := services.NewTypeService(typeRepo)
// 	classService := services.NewClassService(classRepo)
// 	levelService := services.NewLevelService(levelRepo)
// 	profileService := services.NewProfileService(profileRepo)
// 	packageService := services.NewPackageService(packageRepo)
// 	categoryService := services.NewCategoryService(categoryRepo)
// 	locationService := services.NewLocationService(locationRepo)
// 	instructorService := services.NewInstructorService(instructorRepo)
// 	userPackageService := services.NewUserPackageService(userPackageRepo)
// 	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
// 	classScheduleService := services.NewClassScheduleService(classScheduleRepo, classRepo)
// 	paymentService := services.NewPaymentService(paymentRepo, packageRepo, userPackageRepo)
// 	scheduleTemplateService := services.NewScheduleTemplateService(scheduleTemplateRepo, classRepo, classScheduleRepo)

// 	authHandler := handlers.NewAuthHandler(authService)
// 	typeHandler := handlers.NewTypeHandler(typeService)
// 	levelHandler := handlers.NewLevelHandler(levelService)
// 	classHandler := handlers.NewClassHandler(classService)
// 	profileHandler := handlers.NewProfileHandler(profileService)
// 	packageHandler := handlers.NewPackageHandler(packageService)
// 	paymentHandler := handlers.NewPaymentHandler(paymentService)
// 	categoryHandler := handlers.NewCategoryHandler(categoryService)
// 	locationHandler := handlers.NewLocationHandler(locationService)
// 	instructorHandler := handlers.NewInstructorHandler(instructorService)
// 	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)
// 	userPackageHandler := handlers.NewUserPackageHandler(userPackageService)
// 	classScheduleHandler := handlers.NewClassScheduleHandler(classScheduleService)
// 	scheduleTemplateHandler := handlers.NewScheduleTemplateHandler(scheduleTemplateService)

// 	routes.AuthRoutes(r, authHandler)
// 	routes.TypeRoutes(r, typeHandler)
// 	routes.ClassRoutes(r, classHandler)
// 	routes.LevelRoutes(r, levelHandler)
// 	routes.PackageRoutes(r, packageHandler)
// 	routes.ProfileRoutes(r, profileHandler)
// 	routes.PaymentRoutes(r, paymentHandler)
// 	routes.CategoryRoutes(r, categoryHandler)
// 	routes.LocationRoutes(r, locationHandler)
// 	routes.InstructorRoutes(r, instructorHandler)
// 	routes.SubcategoryRoutes(r, subcategoryHandler)
// 	routes.UserPackageRoutes(r, userPackageHandler)
// 	routes.ClassScheduleRoutes(r, classScheduleHandler)
// 	routes.ScheduleTemplateRoutes(r, scheduleTemplateHandler)

// 	port := os.Getenv("PORT")
// 	log.Println("server running on port:", port)
// 	log.Fatal(r.Run(":" + port))
// }
