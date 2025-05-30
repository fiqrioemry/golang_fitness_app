package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(r *gin.Engine, h *handlers.LocationHandler) {
	locationGroup := r.Group("/api/locations")

	// public endpoints
	locationGroup.GET("", h.GetAllLocations)
	locationGroup.GET("/:id", h.GetLocationByID)

	// admin-protected endpoints
	admin := locationGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLocation)
	admin.PUT("/:id", h.UpdateLocation)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteLocation)
}

// GET    /api/locations             → Ambil semua lokasi (public)
// GET    /api/locations/:id         → Ambil detail lokasi berdasarkan ID
// POST   /api/locations             → Tambah lokasi baru (admin only)
// PUT    /api/locations/:id         → Update data lokasi (admin only)
// DELETE /api/locations/:id         → Hapus lokasi (khusus role owner)
