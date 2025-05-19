package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func LevelRoutes(r *gin.Engine, h *handlers.LevelHandler) {
	levelGroup := r.Group("/api/levels")

	// Public Routes
	levelGroup.GET("", h.GetAllLevels)
	levelGroup.GET("/:id", h.GetLevelByID)

	// Admin Only Routes
	admin := levelGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLevel)
	admin.PUT("/:id", h.UpdateLevel)
	admin.DELETE("/:id", h.DeleteLevel)
}
