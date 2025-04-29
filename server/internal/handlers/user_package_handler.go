package handlers

import (
	"net/http"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserPackageHandler struct {
	userPackageService services.UserPackageService
}

func NewUserPackageHandler(userPackageService services.UserPackageService) *UserPackageHandler {
	return &UserPackageHandler{userPackageService}
}

func (h *UserPackageHandler) GetUserPackages(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	userPackages, err := h.userPackageService.GetUserPackages(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user packages", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userPackages)
}
