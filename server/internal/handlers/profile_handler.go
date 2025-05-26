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
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Profile.Fullname,
		Avatar:   user.Profile.Avatar,
		Gender:   user.Profile.Gender,
		Birthday: "",
		Bio:      user.Profile.Bio,
		Phone:    user.Profile.Phone,
		JoinedAt: user.CreatedAt.Format("2006-01-02"),
	}

	if user.Profile.Birthday != nil {
		resp.Birthday = user.Profile.Birthday.Format("2006-01-02")
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.UpdateProfileRequest
	if !utils.BindAndValidateJSON(c, &req) {
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

	if err := h.profileService.UpdateAvatar(userID, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update avatar", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully"})
}

func (h *ProfileHandler) GetUserTransactions(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	page := utils.GetQueryInt(c, "page", 1)
	limit := utils.GetQueryInt(c, "limit", 10)

	resp, err := h.profileService.GetUserTransactions(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch transactions", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileHandler) GetUserPackages(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	page := utils.GetQueryInt(c, "page", 1)
	limit := utils.GetQueryInt(c, "limit", 10)

	resp, err := h.profileService.GetUserPackages(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch packages", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileHandler) GetUserBookings(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	page := utils.GetQueryInt(c, "page", 1)
	limit := utils.GetQueryInt(c, "limit", 10)

	resp, err := h.profileService.GetUserBookings(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bookings", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileHandler) GetUserPackagesByClassID(c *gin.Context) {
	classID := c.Param("id")
	userID := utils.MustGetUserID(c)

	userPackages, err := h.profileService.GetUserPackagesByClassID(userID, classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user packages", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userPackages)
}
