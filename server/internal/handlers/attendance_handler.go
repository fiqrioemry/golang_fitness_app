package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	service services.AttendanceService
}

func NewAttendanceHandler(service services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service}
}

func (h *AttendanceHandler) MarkAttendance(c *gin.Context) {
	var req dto.MarkAttendanceRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	userID := utils.MustGetUserID(c)

	if err := h.service.MarkAttendance(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance marked successfully"})
}

func (h *AttendanceHandler) GetAllAttendances(c *gin.Context) {
	attendances, err := h.service.GetAllAttendances()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch attendances", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attendances": attendances})
}

func (h *AttendanceHandler) ExportAttendances(c *gin.Context) {
	file, err := h.service.ExportAttendancesToExcel()
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
