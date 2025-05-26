package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, handler *handlers.PaymentHandler) {
	p := r.Group("/api/payments")
	p.POST("", middleware.AuthRequired(), handler.CreatePayment)
	p.POST("/stripe/notifications", handler.HandlePaymentNotifications)

	admin := p.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.GET("", handler.GetAllUserPayments)
}
