package handlers

import (
	"fmt"
	"net/http"
	"server/internal/dto"
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

func (h *AttendanceHandler) GetAttendanceDetail(c *gin.Context) {
	scheduleID := c.Param("id")
	result, err := h.attendanceService.GetAttendanceDetail(scheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch attendances", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *AttendanceHandler) CheckinAttendance(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	qr, err := h.attendanceService.CheckinAttendance(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking, _ := h.attendanceService.GetBookingInfo(bookingID)

	response := dto.QRCodeAttendanceResponse{
		QR:         qr,
		ClassTitle: booking.ClassSchedule.Class.Title,
		Date:       booking.ClassSchedule.Date.Format("2006-01-02"),
		Instructor: booking.ClassSchedule.Instructor.User.Profile.Fullname,
		StartTime:  fmt.Sprintf("%02d:%02d", booking.ClassSchedule.StartHour, booking.ClassSchedule.StartMinute),
	}

	c.JSON(http.StatusOK, response)
}

func (h *AttendanceHandler) RegenerateQRCode(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	qr, info, err := h.attendanceService.GetQRCode(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dto.QRCodeAttendanceResponse{
		QR:         qr,
		ClassTitle: info.ClassSchedule.Class.Title,
		Date:       info.ClassSchedule.Date.Format("2006-01-02"),
		Instructor: info.ClassSchedule.Instructor.User.Profile.Fullname,
		StartTime:  fmt.Sprintf("%02d:%02d", info.ClassSchedule.StartHour, info.ClassSchedule.StartMinute),
	}

	c.JSON(http.StatusOK, response)
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
