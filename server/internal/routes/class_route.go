package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(r *gin.Engine, handler *handlers.ClassHandler) {
	class := r.Group("/api/classes")

	// Public
	class.GET("", handler.GetAllClasses)
	class.GET("/:id", handler.GetClassByID)

	// Admin Only
	admin := class.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreateClass)
	admin.PUT("/:id", handler.UpdateClass)
	admin.POST("/:id/gallery", handler.UploadClassGallery)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeleteClass)
}
