package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	classService services.ClassService
}

func NewClassHandler(classService services.ClassService) *ClassHandler {
	return &ClassHandler{classService}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var req dto.CreateClassRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.classService.CreateClass(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.classService.UpdateClass(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id := c.Param("id")

	if err := h.classService.DeleteClass(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")

	class, err := h.classService.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Class not found"})
		return
	}

	c.JSON(http.StatusOK, class)
}

func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	var query dto.ClassQueryParam
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters", "error": err.Error()})
		return
	}

	classes, total, err := h.classService.GetAllClasses(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch classes", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
		"total":   total,
		"page":    query.Page,
		"limit":   query.Limit,
	})
}

func (h *ClassHandler) GetActiveClasses(c *gin.Context) {
	classes, err := h.classService.GetActiveClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch active classes", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}
