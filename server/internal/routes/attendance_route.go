package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.Engine, handler *handlers.AttendanceHandler) {
	attendance := r.Group("/api/attendances")
	attendance.Use(middleware.AuthRequired())

	attendance.POST("", handler.MarkAttendance)
	attendance.GET("", handler.GetAllAttendances)
	attendance.GET("/export", handler.ExportAttendances)
}
