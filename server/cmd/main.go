package main

import (
	"log"
	"os"
	"server/internal/bootstrap"
	"server/internal/config"
	"server/internal/cron"
	"server/internal/middleware"
	"server/internal/routes"
	"server/internal/seeders"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitRedis()
	config.InitMailer()
	config.InitDatabase()
	config.InitCloudinary()
	config.InitMidtrans()

	db := config.DB
	// ========== Seeder ==========
	seeders.ResetDatabase(db)

	// middleware config
	r := gin.Default()
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(5, 10),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/payments"}),
	)

	// ========== layer ==========
	repo := bootstrap.InitRepositories(db)
	s := bootstrap.InitServices(repo)
	h := bootstrap.InitHandlers(s)

	// ========== Cron Job ==========
	cronManager := cron.NewCronManager(s.PaymentService, s.ScheduleTemplateService, s.NotificationService)
	cronManager.RegisterJobs()
	cronManager.Start()

	// ========== Route Binding ==========
	routes.AuthRoutes(r, h.AuthHandler)
	routes.UserRoutes(r, h.UserHandler)
	routes.TypeRoutes(r, h.TypeHandler)
	routes.ClassRoutes(r, h.ClassHandler)
	routes.LevelRoutes(r, h.LevelHandler)
	routes.ReviewRoutes(r, h.ReviewHandler)
	routes.VoucherRoutes(r, h.VoucherHandler)
	routes.PackageRoutes(r, h.PackageHandler)
	routes.ProfileRoutes(r, h.ProfileHandler)
	routes.PaymentRoutes(r, h.PaymentHandler)
	routes.BookingRoutes(r, h.BookingHandler)
	routes.CategoryRoutes(r, h.CategoryHandler)
	routes.LocationRoutes(r, h.LocationHandler)
	routes.InstructorRoutes(r, h.InstructorHandler)
	routes.AttendanceRoutes(r, h.AttendanceHandler)
	routes.SubcategoryRoutes(r, h.SubcategoryHandler)
	routes.NotificationRoutes(r, h.NotificationHandler)
	routes.ClassScheduleRoutes(r, h.ClassScheduleHandler)
	routes.ScheduleTemplateRoutes(r, h.ScheduleTemplateHandler)

	// ========== Start Server ==========
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
