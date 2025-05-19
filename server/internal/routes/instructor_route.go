package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InstructorRoutes(r *gin.Engine, handler *handlers.InstructorHandler) {
	instructor := r.Group("/api/instructors")

	// Public
	instructor.GET("", handler.GetAllInstructors)
	instructor.GET("/:id", handler.GetInstructorByID)

	// Admin Only
	admin := instructor.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", handler.CreateInstructor)
	admin.PUT("/:id", handler.UpdateInstructor)
	admin.DELETE("/:id", handler.DeleteInstructor)
}
