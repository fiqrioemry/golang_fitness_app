package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InstructorRoutes(r *gin.Engine, handler *handlers.InstructorHandler) {
	// public endpoints
	instructor := r.Group("/api/instructors")
	instructor.GET("", handler.GetAllInstructors)
	instructor.GET("/:id", handler.GetInstructorByID)
	// admin-protected endpoints
	admin := instructor.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreateInstructor)
	admin.PUT("/:id", handler.UpdateInstructor)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeleteInstructor)
}

// GET    /api/instructors             → Ambil semua instruktur (public)
// GET    /api/instructors/:id         → Ambil detail instruktur berdasarkan ID
// POST   /api/instructors             → Tambah instruktur baru (admin only)
// PUT    /api/instructors/:id         → Update data instruktur (admin only)
// DELETE /api/instructors/:id         → Hapus instruktur (khusus role owner)
