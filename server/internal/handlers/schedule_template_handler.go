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

func (h *ScheduleTemplateHandler) CreateRecurringScheduleTemplate(c *gin.Context) {
	var req dto.CreateRecurringScheduleTemplateRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateRecurringScheduleTemplate(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create recurring schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Recurring schedule template created successfully"})
}

// schedule_template_handler.go
func (h *ScheduleTemplateHandler) UpdateTemplate(c *gin.Context) {
	templateID := c.Param("id")
	var req dto.CreateScheduleTemplateRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateTemplate(templateID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update template", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template updated successfully"})
}

func (h *ScheduleTemplateHandler) DeleteTemplate(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteTemplate(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete template", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}
