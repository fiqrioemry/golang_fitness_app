package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserPackageRoutes(r *gin.Engine, handler *handlers.UserPackageHandler) {
	userPackage := r.Group("/api/user/packages")
	userPackage.Use(middleware.AuthRequired())
	userPackage.GET("", handler.GetUserPackages)
}
