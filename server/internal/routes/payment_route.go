package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, handler *handlers.PaymentHandler) {
	payment := r.Group("/api/payments")

	// User - create payment
	payment.Use(middleware.AuthRequired())
	payment.POST("", handler.CreatePayment)

	// Webhook - Midtrans notification (public, tanpa auth)
	r.POST("/api/payments/notification", handler.HandlePaymentNotification)
}
