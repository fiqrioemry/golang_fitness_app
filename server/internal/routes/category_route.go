package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, h *handlers.CategoryHandler) {
	category := r.Group("/api/categories")

	// Public Routes
	category.GET("", h.GetAllCategories)
	category.GET("/:id", h.GetCategoryByID)

	// Admin Only Routes
	admin := category.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateCategory)
	admin.PUT("/:id", h.UpdateCategory)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteCategory)
}
