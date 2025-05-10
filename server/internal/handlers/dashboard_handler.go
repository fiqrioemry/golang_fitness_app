package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService services.DashboardService
}

func NewDashboardHandler(dashboardService services.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	data, err := h.dashboardService.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch dashboard summary"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *DashboardHandler) GetRevenueStats(c *gin.Context) {
	var req dto.RevenueStatRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid query parameter"})
		return
	}

	stats, err := h.dashboardService.GetRevenueStats(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get revenue stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
