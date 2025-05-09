package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine, h *handlers.NotificationHandler) {
	route := r.Group("/api/notifications")
	route.Use(middleware.AuthRequired())

	route.GET("/settings", h.GetNotificationSettings)
	route.PUT("/settings", h.UpdateNotificationSetting)

	route.GET("/unread", h.GetUnreadNotifications)
	route.PATCH("/:id/read", h.MarkNotificationRead)
}
