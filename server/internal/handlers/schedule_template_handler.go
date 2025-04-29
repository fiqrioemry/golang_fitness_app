package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ScheduleTemplateHandler struct {
	service services.ScheduleTemplateService
}

func NewScheduleTemplateHandler(service services.ScheduleTemplateService) *ScheduleTemplateHandler {
	return &ScheduleTemplateHandler{service}
}

func (h *ScheduleTemplateHandler) CreateTemplate(c *gin.Context) {
	var req dto.CreateScheduleTemplateRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateTemplate(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create schedule template", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule template created successfully"})
}

func (h *ScheduleTemplateHandler) AutoGenerateSchedules(c *gin.Context) {
	if err := h.service.AutoGenerateSchedules(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to auto generate schedules", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedules generated successfully"})
}
