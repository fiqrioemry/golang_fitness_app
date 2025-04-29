package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ClassScheduleHandler struct {
	service services.ClassScheduleService
}

func NewClassScheduleHandler(service services.ClassScheduleService) *ClassScheduleHandler {
	return &ClassScheduleHandler{service}
}

func (h *ClassScheduleHandler) CreateClassSchedule(c *gin.Context) {
	var req dto.CreateClassScheduleRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateClassSchedule(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class schedule created successfully"})
}

func (h *ClassScheduleHandler) UpdateClassSchedule(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassScheduleRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateClassSchedule(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule updated successfully"})
}

func (h *ClassScheduleHandler) DeleteClassSchedule(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteClassSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule deleted successfully"})
}

func (h *ClassScheduleHandler) GetAllClassSchedules(c *gin.Context) {
	schedules, err := h.service.GetAllClassSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get schedules", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}
