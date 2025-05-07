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
	templateService services.ScheduleTemplateService
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

	if req.StartHour == nil || req.StartMinute == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "startHour & startMinute required"})
		return
	}

	if !req.IsRecurring {

		if req.Date == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "date required for non-recurring schedule"})
			return
		}

		err := h.scheduleService.CreateClassSchedule(dto.CreateClassScheduleRequest{
			ClassID:      req.ClassID,
			InstructorID: req.InstructorID,
			Capacity:     req.Capacity,
			Color:        req.Color,
			Date:         *req.Date,
			StartHour:    *req.StartHour,
			StartMinute:  *req.StartMinute,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create class schedule", "error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Class schedule created successfully"})
		return
	}

	if len(req.RecurringDays) == 0 || req.EndType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "recurringDays & endType required"})
		return
	}
	if req.EndType == "until" && req.EndDate == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "endDate required for type 'until'"})
		return
	}

	for _, day := range req.RecurringDays {
		if day < 0 || day > 6 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid day in recurringDays (0 = Sunday, ..., 6 = Saturday)"})
			return
		}

		err := h.templateService.CreateScheduleTemplate(dto.CreateScheduleTemplateRequest{
			ClassID:      req.ClassID,
			InstructorID: req.InstructorID,
			DayOfWeek:    day,
			StartHour:    *req.StartHour,
			StartMinute:  *req.StartMinute,
			Capacity:     req.Capacity,
			Frequency:    "recurring",
			Color:        req.Color,
			EndType:      req.EndType,
			EndDate:      req.EndDate,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create recurring schedule template", "error": err.Error()})
			return
		}
	}

	if err := h.templateService.AutoGenerateSchedules(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to auto-generate schedules", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Recurring schedules created successfully"})
}

func (h *ClassScheduleHandler) UpdateClassSchedule(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassScheduleRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.scheduleService.UpdateClassSchedule(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
	var filter dto.ClassScheduleQueryParam
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters", "error": err.Error()})
		return
	}

	schedules, err := h.scheduleService.GetSchedulesByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch class schedules", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (h *ClassScheduleHandler) GetSchedulesWithBookingStatus(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	schedules, err := h.scheduleService.GetSchedulesWithBookingStatus(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (h *ClassScheduleHandler) GetScheduleByID(c *gin.Context) {
	scheduleID := c.Param("id")

	result, err := h.scheduleService.GetClassScheduleByID(scheduleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
