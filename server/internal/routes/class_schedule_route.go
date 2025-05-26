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
	customer := schedule.Use(middleware.AuthRequired())
	customer.GET("/status", h.GetSchedulesWithStatus)
	customer.GET("/:id", h.GetScheduleByID)

	// Admin Only
	admin := schedule.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClassSchedule)
	admin.PUT("/:id", h.UpdateClassSchedule)
	admin.DELETE("/:id", middleware.RoleOnly("owner"), h.DeleteClassSchedule)

	// instructor only
	instructor := schedule.Use(middleware.AuthRequired(), middleware.RoleOnly("instructor"))
	instructor.GET("/instructor", h.GetInstructorSchedules)
	instructor.GET("/:id/detail", h.GetInstructorSchedules)
	instructor.GET("/:id/attendance", h.GetClassAttendances)

	instructor.PATCH("/:id/open", h.OpenClassSchedule)
	instructor.PATCH("/:id/close", h.CloseClassSchedule)

}
