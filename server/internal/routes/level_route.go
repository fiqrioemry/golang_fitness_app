package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func LevelRoutes(r *gin.Engine, h *handlers.LevelHandler) {
	// public endpoints
	levelGroup := r.Group("/api/levels")
	levelGroup.GET("", h.GetAllLevels)
	levelGroup.GET("/:id", h.GetLevelByID)

	// admin-protected endpoints
	admin := levelGroup.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateLevel)
	admin.PUT("/:id", h.UpdateLevel)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteLevel)
}

// GET    /api/levels             → Ambil semua level kelas (public)
// GET    /api/levels/:id         → Ambil detail level berdasarkan ID
// POST   /api/levels             → Tambah level baru (admin only)
// PUT    /api/levels/:id         → Update data level (admin only)
// DELETE /api/levels/:id         → Hapus level (khusus role owner)
