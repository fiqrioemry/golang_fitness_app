package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SubcategoryRoutes(r *gin.Engine, h *handlers.SubcategoryHandler) {
	subcategory := r.Group("/api/subcategories")

	// Public Routes
	subcategory.GET("", h.GetAllSubcategories)
	subcategory.GET("/:id", h.GetSubcategoryByID)
	subcategory.GET("/category/:categoryId", h.GetSubcategoriesByCategoryID)

	// Admin Only Routes
	admin := subcategory.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.CreateSubcategory)
	admin.PUT("/:id", h.UpdateSubcategory)
	admin.DELETE("/:id", h.DeleteSubcategory)
}
