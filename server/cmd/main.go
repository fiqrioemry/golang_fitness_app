package main

import (
	"log"
	"os"
	"server/internal/config"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/routes"
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

	routes.AuthRoutes(r, authHandler)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
