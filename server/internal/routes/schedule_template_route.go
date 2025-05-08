package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ScheduleTemplateRoutes(r *gin.Engine, handler *handlers.ScheduleTemplateHandler) {
	template := r.Group("/api/schedule-templates")
	template.Use(middleware.AuthRequired(), middleware.AdminOnly())

	template.GET("", handler.GetAllTemplates)
	template.DELETE("/:id", handler.DeleteTemplate)
	template.POST("/:id/run", handler.RunScheduleTemplate)
	template.POST("/:id/stop", handler.StopScheduleTemplate)
	template.PUT("/:id", handler.UpdateScheduleTemplate)

	//
	// template.PUT("", handler.CreateScheduleTemplate)
	// template.POST("/generate", handler.AutoGenerateSchedules)
}
