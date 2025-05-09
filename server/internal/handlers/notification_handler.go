package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service}
}

func (h *NotificationHandler) GetNotificationSettings(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	settings, err := h.service.GetSettingsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch settings"})
		return
	}

	c.JSON(http.StatusOK, settings)
}

func (h *NotificationHandler) UpdateNotificationSetting(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.UpdateNotificationSettingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.service.UpdateSetting(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update setting"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *NotificationHandler) GetUnreadNotifications(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	notifs, err := h.service.GetUnreadNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, notifs)
}

func (h *NotificationHandler) MarkNotificationRead(c *gin.Context) {
	notifID := c.Param("id")
	userID := utils.MustGetUserID(c)

	if err := h.service.MarkAsRead(userID, notifID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "marked as read"})
}
