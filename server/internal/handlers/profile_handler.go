package handlers

import (
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)
}

func (h *ProfileHandler) UpdateAvatar(c *gin.Context) {
	userID := utils.MustGetUserID(c)
}
