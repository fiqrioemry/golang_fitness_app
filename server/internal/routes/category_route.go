package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	category := r.Group("/api/categories")

	// public endpoints
	category.GET("", h.GetAllCategories)
	category.GET("/:id", h.GetCategoryByID)

	// admin-protected endpoints
	admin := category.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateCategory)
	admin.PUT("/:id", h.UpdateCategory)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteCategory)
}

// GET    /api/categories             → Ambil semua kategori (public)
// GET    /api/categories/:id         → Ambil detail kategori berdasarkan ID
// POST   /api/categories             → Tambah kategori baru (admin only)
// PUT    /api/categories/:id         → Update data kategori (admin only)
// DELETE /api/categories/:id         → Hapus kategori (khusus role owner)
