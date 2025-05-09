package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine, h *handlers.NotificationHandler) {
	route := r.Group("/api/notifications")
	route.Use(middleware.AuthRequired())
	route.GET("", h.GetAllNotifications)

	route.GET("/settings", h.GetNotificationSettings)
	route.PUT("/settings", h.UpdateNotificationSetting)
	route.PATCH("/:id/read", h.MarkNotificationRead)
}
