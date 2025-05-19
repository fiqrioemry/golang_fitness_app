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

	// authenticate customer
	customer := schedule.Use(middleware.AuthRequired())
	customer.GET("/:id", h.GetScheduleByID)
	customer.GET("/status", h.GetSchedulesWithBookingStatus)

	// Admin Only
	admin := schedule.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateClassSchedule)
	admin.PUT("/:id", h.UpdateClassSchedule)
	admin.DELETE("/:id", h.DeleteClassSchedule)

}
