package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine, h *handlers.NotificationHandler) {
	notif := r.Group("/api/notifications")
	notif.Use(middleware.AuthRequired())
	notif.GET("/settings", h.GetNotificationSettings)
	notif.PUT("/settings", h.UpdateNotificationSetting)
	notif.GET("", h.GetAllNotifications)
	notif.PATCH("/read", h.MarkAllNotificationsAsRead)
	notif.POST("/broadcast", middleware.RoleOnly("admin"), h.SendNewNotificatioon)
}

// GET    /api/notifications/settings     → Ambil pengaturan notifikasi milik user
// PUT    /api/notifications/settings     → Update pengaturan notifikasi user
// GET    /api/notifications              → Ambil semua notifikasi milik user
// PATCH  /api/notifications/read         → Tandai semua notifikasi sebagai sudah dibaca
// POST   /api/notifications/broadcast    → Kirim notifikasi broadcast ke semua user (admin on
