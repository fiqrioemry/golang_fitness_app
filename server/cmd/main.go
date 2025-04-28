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

	routes.AuthRoutes(r, authHandler)
	routes.TypeRoutes(r, typeHandler)
	routes.ClassRoutes(r, classHandler)
	routes.ProfileRoutes(r, profileHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)

	seeders.SeedTypes(db)
	seeders.SeedUsers(db)
	seeders.SeedCategories(db)
	seeders.SeedSubcategories(db)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
