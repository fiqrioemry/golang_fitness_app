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
	pkg.GET("/class/:id", handler.GetPackagesByClassID)

	admin := pkg.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreatePackage)
	admin.PUT("/:id", handler.UpdatePackage)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeletePackage)
}
