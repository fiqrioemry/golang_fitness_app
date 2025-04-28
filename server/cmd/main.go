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
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	profileRepo := repositories.NewProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)
	profileHandler := handlers.NewProfileHandler(profileService)

	classRepo := repositories.NewClassRepository(db)
	classService := services.NewClassService(classRepo)
	classHandler := handlers.NewClassHandler(classService)

	routes.AuthRoutes(r, authHandler)
	routes.ClassRoutes(r, classHandler)
	routes.ProfileRoutes(r, profileHandler)

	seeders.SeedUsers(db)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
