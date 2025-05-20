package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req dto.CreateReviewRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	userID := utils.MustGetUserID(c)

	if err := h.service.CreateReview(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully"})
}

func (h *ReviewHandler) GetReviewsByClass(c *gin.Context) {
	classID := c.Param("classId")

	reviews, err := h.service.GetReviewsByClassID(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch reviews", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
