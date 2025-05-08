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

func (h *ScheduleTemplateHandler) GetAllTemplates(c *gin.Context) {
	templates, err := h.service.GetAllTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, templates)
}

func (h *ScheduleTemplateHandler) UpdateScheduleTemplate(c *gin.Context) {

	templateID := c.Param("id")
	var req dto.UpdateScheduleTemplateRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateScheduleTemplate(templateID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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

func (h *ScheduleTemplateHandler) RunScheduleTemplate(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.RunTemplate(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to run template", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template  Activated successfully"})
}

func (h *ScheduleTemplateHandler) StopScheduleTemplate(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.StopTemplate(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to stop template", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template deactivated successfully"})
}

// func (h *ScheduleTemplateHandler) AutoGenerateSchedules(c *gin.Context) {
// 	if err := h.service.AutoGenerateSchedules(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to auto generate schedules", "error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Schedules generated successfully"})
// }
// func (h *ScheduleTemplateHandler) CreateScheduleTemplate(c *gin.Context) {
// 	var req dto.CreateScheduleTemplateRequest
// 	if !utils.BindAndValidateJSON(c, &req) {
// 		return
// 	}

// 	if err := h.service.CreateScheduleTemplate(req); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create recurring schedule", "error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Recurring schedule template created successfully"})
// }
