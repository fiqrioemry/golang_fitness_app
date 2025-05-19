package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.Engine, h *handlers.AttendanceHandler) {
	attendance := r.Group("/api/attendances")
	attendance.Use(middleware.AuthRequired())

	// customer attendance
	attendance.GET("", h.GetAllAttendances)
	attendance.GET("/:id", h.GetAttendanceDetail)
	attendance.POST("/:id", h.CheckinAttendance)
	attendance.GET("/:id/qr-code", h.RegenerateQRCode)

	admin := attendance.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("/validate", h.ValidateQRCodeScan)
}
