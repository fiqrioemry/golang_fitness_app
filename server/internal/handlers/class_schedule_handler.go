package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ClassScheduleHandler struct {
	scheduleService services.ClassScheduleService
	templateService services.ScheduleTemplateService // ✅ tambahkan ini
}

func NewClassScheduleHandler(
	scheduleService services.ClassScheduleService,
	templateService services.ScheduleTemplateService,
) *ClassScheduleHandler {
	return &ClassScheduleHandler{
		scheduleService: scheduleService,
		templateService: templateService,
	}
}

func (h *ClassScheduleHandler) CreateClassSchedule(c *gin.Context) {
	var req dto.CreateScheduleRequest
	if ok := utils.BindAndValidateJSON(c, &req); !ok {
		return
	}

	if !req.IsRecurring {
		// Non-recurring schedule
		if req.Date == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "date is required for non-recurring schedule"})
			return
		}

		err := h.scheduleService.CreateClassSchedule(dto.CreateClassScheduleRequest{
			ClassID:      req.ClassID,
			InstructorID: req.InstructorID,
			Capacity:     req.Capacity,
			Color:        req.Color,
			Date:         *req.Date,
			StartHour:    req.StartHour,
			StartMinute:  req.StartMinute,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create schedule", "error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Class schedule created"})
		return
	}

	// Recurring schedule
	if len(req.RecurringDays) == 0 || req.EndType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "recurringDays and endType are required for recurring schedule"})
		return
	}

	if req.EndType == "until" && req.EndDate == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "endDate is required for 'until' endType"})
		return
	}

	for _, day := range req.RecurringDays {
		dayEnum := utils.ParseDayOfWeek(day)
		if dayEnum == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid day in recurringDays: " + day})
			return
		}

		err := h.templateService.CreateRecurringScheduleTemplate(dto.CreateRecurringScheduleTemplateRequest{
			ClassID:      req.ClassID,
			InstructorID: req.InstructorID,
			DayOfWeek:    int(dayEnum), // ✅ FIX here
			StartHour:    req.StartHour,
			StartMinute:  req.StartMinute,
			Capacity:     req.Capacity,
			Frequency:    "recurring",
			EndType:      req.EndType,
			EndDate:      req.EndDate,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create recurring template", "error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Recurring schedule templates created"})
}

func (h *ClassScheduleHandler) UpdateClassSchedule(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassScheduleRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.scheduleService.UpdateClassSchedule(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule updated successfully"})
}

func (h *ClassScheduleHandler) DeleteClassSchedule(c *gin.Context) {
	id := c.Param("id")

	if err := h.scheduleService.DeleteClassSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule deleted successfully"})
}

func (h *ClassScheduleHandler) GetAllClassSchedules(c *gin.Context) {
	schedules, err := h.scheduleService.GetAllClassSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get schedules", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}
