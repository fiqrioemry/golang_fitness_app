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

	// Authenticated base
	auth := schedule.Group("")
	auth.Use(middleware.AuthRequired())

	// Instructor
	instructor := auth.Group("/instructor")
	instructor.Use(middleware.RoleOnly("instructor"))
	instructor.GET("", h.GetInstructorSchedules)

	// PATCH /:id/open → role instructor
	instructor.PATCH("/:id/open", h.OpenClassSchedule)
	instructor.GET("/:id/attendance", h.GetClassAttendances)

	// Customer
	customer := auth.Group("")
	customer.Use(middleware.RoleOnly("customer"))
	customer.GET("/status", h.GetSchedulesWithStatus)
	customer.GET("/:id", h.GetScheduleByID)

	// Admin
	admin := auth.Group("")
	admin.Use(middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClassSchedule)
	admin.POST("/recurring", h.CreateRecurringSchedule)
	admin.PUT("/:id", h.UpdateClassSchedule)
	admin.DELETE("/:id", h.DeleteClassSchedule)
}
