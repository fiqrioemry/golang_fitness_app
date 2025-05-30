package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, handler *handlers.PaymentHandler) {
	r.POST("/api/payments/stripe/notifications", handler.HandlePaymentNotifications)

	payments := r.Group("/api/payments")
	payments.Use(middleware.AuthRequired())
	payments.POST("", handler.CreatePayment)
	payments.GET("/me", handler.GetMyTransactions)

	admin := payments.Use(middleware.RoleOnly("admin"))
	admin.GET("", handler.GetAllUserPayments)

}

// POST /api/payments/stripe/notifications → Digunakan oleh Stripe untuk mengirim notifikasi ketika pembayaran berhasil
// POST /api/payments 	     → Membuat transaksi pembayaran baru untuk pembelian paket
// GET /api/payments/me		 → Mengambil daftar transaksi pembayaran milik user yang sedang login
// GET /api/payments 		 → Mengambil semua transaksi pembayaran dari seluruh user (hanya untuk admin)
