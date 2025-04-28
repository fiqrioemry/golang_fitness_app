package handlers

import (
	"net/http"
	"server/internal/dto"
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

	user, err := h.profileService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch profile"})
		return
	}

	resp := dto.ProfileResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Fullname:  user.Profile.Fullname,
		Avatar:    user.Profile.Avatar,
		Gender:    user.Profile.Gender,
		Birthday:  "",
		Bio:       user.Profile.Bio,
		Phone:     user.Profile.Phone,
		UpdatedAt: user.UpdatedAt,
	}

	if user.Profile.Birthday != nil {
		resp.Birthday = user.Profile.Birthday.Format("2006-01-02")
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.UpdateProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := h.profileService.UpdateProfile(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update profile", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *ProfileHandler) UpdateAvatar(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Avatar file is required"})
		return
	}

	url, err := h.profileService.UpdateAvatar(userID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update avatar", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully", "avatarUrl": url})
}
