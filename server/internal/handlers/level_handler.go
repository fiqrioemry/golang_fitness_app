package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	levelService services.LevelService
}

func NewLevelHandler(levelService services.LevelService) *LevelHandler {
	return &LevelHandler{levelService}
}

func (h *LevelHandler) CreateLevel(c *gin.Context) {
	var req dto.CreateLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.levelService.CreateLevel(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create level", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Level created successfully"})
}

func (h *LevelHandler) UpdateLevel(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.levelService.UpdateLevel(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update level", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Level updated successfully"})
}

func (h *LevelHandler) DeleteLevel(c *gin.Context) {
	id := c.Param("id")

	if err := h.levelService.DeleteLevel(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete level", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Level deleted successfully"})
}

func (h *LevelHandler) GetAllLevels(c *gin.Context) {
	levels, err := h.levelService.GetAllLevels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch levels", "error": err.Error()})
		return
	}

	if levels == nil {
		levels = []dto.LevelResponse{}
	}

	c.JSON(http.StatusOK, levels)
}

func (h *LevelHandler) GetLevelByID(c *gin.Context) {
	id := c.Param("id")

	level, err := h.levelService.GetLevelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Level not found"})
		return
	}

	c.JSON(http.StatusOK, level)
}
