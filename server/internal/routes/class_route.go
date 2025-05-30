package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(r *gin.Engine, handler *handlers.ClassHandler) {
	class := r.Group("/api/classes")

	// public endpoints
	class.GET("", handler.GetAllClasses)
	class.GET("/:id", handler.GetClassByID)

	// admin-protected endpoints
	admin := class.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreateClass)
	admin.PUT("/:id", handler.UpdateClass)
	admin.POST("/:id/gallery", handler.UploadClassGallery)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeleteClass)
}

// GET    /api/classes                  → Ambil semua kelas (public)
// GET    /api/classes/:id              → Ambil detail kelas berdasarkan ID
// POST   /api/classes                  → Buat kelas baru (admin only)
// PUT    /api/classes/:id              → Update data kelas (admin only)
// DELETE /api/classes/:id              → Hapus kelas (khusus role owner)
// POST   /api/classes/:id/gallery      → Upload gambar galeri kelas (admin only)
