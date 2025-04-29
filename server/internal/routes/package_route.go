package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(r *gin.Engine, handler *handlers.PackageHandler) {
	pkg := r.Group("/api/packages")

	pkg.GET("", handler.GetAllPackages)
	pkg.GET("/:id", handler.GetPackageByID)

	// Admin Only
	admin := pkg.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", handler.CreatePackage)
	admin.PUT("/:id", handler.UpdatePackage)
	admin.DELETE("/:id", handler.DeletePackage)
}
