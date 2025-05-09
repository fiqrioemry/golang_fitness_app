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

func (h *NotificationHandler) GetAllNotifications(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	notifications, err := h.service.GetAllNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
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

func (h *NotificationHandler) MarkAllNotificationsAsRead(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	if err := h.service.MarkAllAsRead(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to mark all as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "all notifications marked as read"})
}
func (h *NotificationHandler) SendPromoNotification(c *gin.Context) {
	var req dto.SendPromoNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.service.SendPromoNotification(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send promo notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "promo notification sent"})
}
