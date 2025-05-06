package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.Engine, h *handlers.AttendanceHandler) {
	attendance := r.Group("/api/attendances")
	attendance.Use(middleware.AuthRequired())

	attendance.GET("", h.GetAllAttendances)
	attendance.GET("/:bookingId/qr", middleware.AuthRequired(), h.RegenerateQRCode)
	attendance.POST("/:bookingId/checkin", middleware.AuthRequired(), h.CheckinAttendance)

	adminGroup := attendance.Use(middleware.AuthRequired(), middleware.AdminOnly())
	adminGroup.POST("/validate", h.ValidateQRCodeScan)

}
