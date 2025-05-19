package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TypeRoutes(r *gin.Engine, h *handlers.TypeHandler) {
	typeGroup := r.Group("/api/types")

	// Public Routes
	typeGroup.GET("", h.GetAllTypes)
	typeGroup.GET("/:id", h.GetTypeByID)

	// Admin Only Routes
	admin := typeGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateType)
	admin.PUT("/:id", h.UpdateType)
	admin.DELETE("/:id", h.DeleteType)
}
