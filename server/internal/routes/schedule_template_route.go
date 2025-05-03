package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ScheduleTemplateRoutes(r *gin.Engine, handler *handlers.ScheduleTemplateHandler) {
	template := r.Group("/api/schedule-templates")
	template.Use(middleware.AuthRequired(), middleware.AdminOnly())
	template.POST("", handler.CreateTemplate)
	template.POST("/auto-generate", handler.AutoGenerateSchedules)
	template.POST("/recurring", handler.CreateRecurringScheduleTemplate)
	template.PUT("/:id", handler.UpdateTemplate)
	template.DELETE("/:id", handler.DeleteTemplate)

}
