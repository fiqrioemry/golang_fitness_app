package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine, handler *handlers.BookingHandler) {
	booking := r.Group("/api/bookings")
	// customer-protected endpoints
	booking.Use(middleware.AuthRequired(), middleware.RoleOnly("customer"))
	booking.POST("", handler.CreateBooking)
	booking.GET("", handler.GetMyBookings)
	booking.GET("/:id", handler.GetBookingDetail)
	booking.POST("/:id/check-in", handler.CheckinBookedClass)
	booking.POST("/:id/check-out", handler.CheckoutBookedClass)
}

// POST   /api/bookings               		→ Buat booking baru (user)
// GET    /api/bookings           			→ Ambil semua jadwal yang dibooking oleh user login
// GET    /api/bookings/:id          		→ Ambil detail booking berdasarkan ID
// POST   /api/bookings/:id/check-in 		→ Check-in ke kelas yang telah dibooking
// POST   /api/bookings/:id/check-out 		→ Check-out dan verifikasi kehadiran
