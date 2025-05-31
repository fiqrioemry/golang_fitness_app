package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserPackageRoutes(r *gin.Engine, handler *handlers.UserPackageHandler) {
	user := r.Group("/api/user-packages")
	user.Use(middleware.AuthRequired())
	user.GET("", handler.GetUserPackages)
	user.GET("/class/:id", handler.GetUserPackagesByClassID)
}
