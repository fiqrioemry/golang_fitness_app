package handlers

import (
	"net/http"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceService services.AttendanceService
}

func NewAttendanceHandler(attendanceService services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{attendanceService}
}

func (h *AttendanceHandler) GetAllAttendances(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	attendances, err := h.attendanceService.GetAllAttendances(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch attendances", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attendances)
}

func (h *AttendanceHandler) CheckinAttendance(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	qr, err := h.attendanceService.CheckinAttendance(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"qr": qr})
}

func (h *AttendanceHandler) RegenerateQRCode(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	qr, err := h.attendanceService.GetQRCode(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"qr": qr})
}

func (h *AttendanceHandler) ValidateQRCodeScan(c *gin.Context) {
	var req struct {
		QR string `json:"qr" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR payload"})
		return
	}

	info, err := h.attendanceService.ValidateQRCodeData(req.QR)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": info})
}

// buat cron job
func (h *AttendanceHandler) MarkAbsentAttendances(c *gin.Context) {
	if err := h.attendanceService.MarkAbsentAttendances(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "absent marked"})
}
