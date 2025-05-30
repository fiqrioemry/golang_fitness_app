package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClassScheduleHandler struct {
	service services.ClassScheduleService
}

func NewClassScheduleHandler(
	service services.ClassScheduleService,
) *ClassScheduleHandler {
	return &ClassScheduleHandler{
		service,
	}
}

// Handler
func (h *ClassScheduleHandler) CreateClassSchedule(c *gin.Context) {
	var req dto.CreateScheduleRequest
	if ok := utils.BindAndValidateJSON(c, &req); !ok {
		return
	}

	if err := h.service.CreateClassSchedule(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class schedule created successfully"})
}

func (h *ClassScheduleHandler) CreateRecurringSchedule(c *gin.Context) {
	var req dto.CreateRecurringScheduleRequest
	if ok := utils.BindAndValidateJSON(c, &req); !ok {
		return
	}

	err := h.service.CreateRecurringSchedule(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Recurring schedule created successfully"})
}

func (h *ClassScheduleHandler) UpdateClassSchedule(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassScheduleRequest
	if !utils.BindAndValidateJSON(c, &req) {
		utils.LogServiceWarn("ClassScheduleHandler", "Update", "invalid input", zap.String("scheduleID", id))
		return
	}

	if err := h.service.UpdateClassSchedule(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule updated successfully"})
}

func (h *ClassScheduleHandler) DeleteClassSchedule(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteClassSchedule(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to delete schedule", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule deleted successfully"})
}

// public
func (h *ClassScheduleHandler) GetAllClassSchedules(c *gin.Context) {
	var param dto.ClassScheduleQueryParam
	if !utils.BindAndValidateForm(c, &param) {
		return
	}

	schedules, err := h.service.GetSchedulesByFilter(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch class schedules", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// authenticated user
func (h *ClassScheduleHandler) GetSchedulesWithStatus(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	schedules, err := h.service.GetSchedulesWithBookingStatus(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (h *ClassScheduleHandler) GetScheduleByID(c *gin.Context) {
	scheduleID := c.Param("id")
	userID := utils.MustGetUserID(c)

	result, err := h.service.GetClassScheduleByID(scheduleID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// for instructor only

func (h *ClassScheduleHandler) GetInstructorSchedules(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var param dto.InstructorScheduleQueryParam
	if !utils.BindAndValidateForm(c, &param) {
		return
	}

	if param.Page == 0 {
		param.Page = 1
	}
	if param.Limit == 0 {
		param.Limit = 10
	}

	data, pagination, err := h.service.GetSchedulesByInstructor(userID, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch instructor schedules", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"pagination": pagination,
	})
}

func (h *ClassScheduleHandler) GetClassAttendances(c *gin.Context) {
	scheduleID := c.Param("id")

	result, err := h.service.GetAttendancesForSchedule(scheduleID)
	if err != nil {
		utils.LogServiceError("classScheduleHandler", "GetClassAttendances", err,
			zap.String("scheduleID", scheduleID),
			zap.String("ip", c.ClientIP()),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	utils.LogServiceInfo("classScheduleHandler", "GetClassAttendances", "returned attendance list",
		zap.String("scheduleID", scheduleID),
		zap.Int("resultCount", len(result)),
		zap.String("ip", c.ClientIP()),
	)

	c.JSON(http.StatusOK, result)
}

func (h *ClassScheduleHandler) OpenClassSchedule(c *gin.Context) {
	id := c.Param("id")

	var req dto.OpenClassScheduleRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.OpenClassSchedule(id, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class schedule opened successfully"})
}
