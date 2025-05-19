package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, h *handlers.LocationHandler) {
	locationGroup := r.Group("/api/locations")

	// Public Routes
	locationGroup.GET("", h.GetAllLocations)
	locationGroup.GET("/:id", h.GetLocationByID)

	// Admin Only Routes
	admin := locationGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLocation)
	admin.PUT("/:id", h.UpdateLocation)
	admin.DELETE("/:id", h.DeleteLocation)
}
