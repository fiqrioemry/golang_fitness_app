// internal/routes/user_route.go
package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, handler *handlers.UserHandler) {
	admin := r.Group("/api/admin/users")
	admin.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))

	admin.GET("", handler.GetAllUsers)
	admin.GET(":id", handler.GetUserDetail)
	admin.GET("/stats", handler.GetUserStats)
}
