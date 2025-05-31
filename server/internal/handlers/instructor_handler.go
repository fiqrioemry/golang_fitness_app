package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type InstructorHandler struct {
	service services.InstructorService
}

func NewInstructorHandler(service services.InstructorService) *InstructorHandler {
	return &InstructorHandler{service}
}

func (h *InstructorHandler) CreateInstructor(c *gin.Context) {
	var req dto.CreateInstructorRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateInstructor(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create instructor", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Instructor created successfully"})
}

func (h *InstructorHandler) UpdateInstructor(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateInstructorRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateInstructor(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update instructor", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Instructor updated successfully"})
}

func (h *InstructorHandler) DeleteInstructor(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteInstructor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete instructor", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instructor deleted successfully"})
}

func (h *InstructorHandler) GetInstructorByID(c *gin.Context) {
	id := c.Param("id")

	instructor, err := h.service.GetInstructorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Instructor not found"})
		return
	}

	c.JSON(http.StatusOK, instructor)
}

func (h *InstructorHandler) GetAllInstructors(c *gin.Context) {
	instructors, err := h.service.GetAllInstructors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch instructors", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, instructors)
}
