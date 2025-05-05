package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClassScheduleRoutes(r *gin.Engine, handler *handlers.ClassScheduleHandler) {
	schedule := r.Group("/api/schedules")

	// Public
	schedule.GET("", handler.GetAllClassSchedules)
	customer := schedule.Use(middleware.AuthRequired())
	customer.GET("/status", handler.GetSchedulesWithBookingStatus)
	customer.GET("/:id", handler.GetScheduleByID)

	// Admin Only
	admin := schedule.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", handler.CreateClassSchedule)
	admin.PUT("/:id", handler.UpdateClassSchedule)
	admin.DELETE("/:id", handler.DeleteClassSchedule)
}
