package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClassScheduleRoutes(r *gin.Engine, h *handlers.ClassScheduleHandler) {
	schedule := r.Group("/api/schedules")

	// Public
	schedule.GET("", h.GetAllClassSchedules)

	// Authenticated
	auth := schedule.Use(middleware.AuthRequired())

	// Instructor
	instructor := auth.Use(middleware.RoleOnly("instructor"))
	instructor.GET("/instructor", h.GetInstructorSchedules)
	instructor.PATCH("/:id/open", h.OpenClassSchedule)
	instructor.GET("/:id/attendance", h.GetClassAttendances)

	// Customer
	customer := auth.Use(middleware.RoleOnly("customer"))
	customer.GET("/status", h.GetSchedulesWithStatus)
	customer.GET("/:id", h.GetScheduleByID)

	// Admin
	admin := auth.Use(middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClassSchedule)
	admin.PUT("/:id", h.UpdateClassSchedule)
	admin.DELETE("/:id", h.DeleteClassSchedule)
}
