package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserPackageHandler struct {
	service services.UserPackageService
}

func NewUserPackageHandler(service services.UserPackageService) *UserPackageHandler {
	return &UserPackageHandler{service}
}

func (h *UserPackageHandler) GetUserPackages(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var params dto.PackageQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	response, pagination, err := h.service.GetUserPackages(userID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch packages", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       response,
		"pagination": pagination,
	})
}

func (h *UserPackageHandler) GetUserPackagesByClassID(c *gin.Context) {
	classID := c.Param("id")
	userID := utils.MustGetUserID(c)

	userPackages, err := h.service.GetUserPackagesByClassID(userID, classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user packages", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userPackages)
}
