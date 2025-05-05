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
	attendances, err := h.attendanceService.GetAllAttendances()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch attendances", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendances)
}

func (h *AttendanceHandler) ExportAttendances(c *gin.Context) {
	file, err := h.attendanceService.ExportAttendancesToExcel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to export attendances", "error": err.Error()})
		return
	}

	// Set header untuk download
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=attendances.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	// Kirim file excel ke client
	if err := file.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to write excel file", "error": err.Error()})
	}
}

func (h *AttendanceHandler) CheckinAttendance(c *gin.Context) {
	bookingID := c.Param("bookingId")
	userID := utils.MustGetUserID(c)

	qr, err := h.attendanceService.CheckinAttendance(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"qr": qr})
}

// khusus buat cron job kelas
func (h *AttendanceHandler) MarkAbsentAttendances(c *gin.Context) {
	if err := h.attendanceService.MarkAbsentAttendances(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "absent marked"})
}

func (h *AttendanceHandler) RegenerateQRCode(c *gin.Context) {
	bookingID := c.Param("bookingId")
	userID := utils.MustGetUserID(c)

	qr, err := h.attendanceService.GetQRCode(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"qr": qr})
}
