package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type SubcategoryHandler struct {
	subcategoryService services.SubcategoryService
}

func NewSubcategoryHandler(subcategoryService services.SubcategoryService) *SubcategoryHandler {
	return &SubcategoryHandler{subcategoryService}
}

func (h *SubcategoryHandler) CreateSubcategory(c *gin.Context) {
	var req dto.CreateSubcategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.subcategoryService.CreateSubcategory(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create subcategory", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subcategory created successfully"})
}

func (h *SubcategoryHandler) UpdateSubcategory(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateSubcategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.subcategoryService.UpdateSubcategory(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update subcategory", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcategory updated successfully"})
}

func (h *SubcategoryHandler) DeleteSubcategory(c *gin.Context) {
	id := c.Param("id")

	if err := h.subcategoryService.DeleteSubcategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete subcategory", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcategory deleted successfully"})
}

func (h *SubcategoryHandler) GetAllSubcategories(c *gin.Context) {
	subcategories, err := h.subcategoryService.GetAllSubcategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch subcategories", "error": err.Error()})
		return
	}

	if subcategories == nil {
		subcategories = []dto.SubcategoryResponse{}
	}

	c.JSON(http.StatusOK, gin.H{"subcategories": subcategories})
}

func (h *SubcategoryHandler) GetSubcategoryByID(c *gin.Context) {
	id := c.Param("id")

	subcategory, err := h.subcategoryService.GetSubcategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Subcategory not found"})
		return
	}

	c.JSON(http.StatusOK, subcategory)
}

func (h *SubcategoryHandler) GetSubcategoriesByCategoryID(c *gin.Context) {
	categoryID := c.Param("categoryId")

	subcategories, err := h.subcategoryService.GetSubcategoriesByCategoryID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch subcategories by category", "error": err.Error()})
		return
	}

	if subcategories == nil {
		subcategories = []dto.SubcategoryResponse{}
	}

	c.JSON(http.StatusOK, gin.H{"subcategories": subcategories})
}
