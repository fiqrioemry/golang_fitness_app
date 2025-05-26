package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine, handler *handlers.BookingHandler) {
	booking := r.Group("/api/bookings")
	booking.Use(middleware.AuthRequired())
	booking.POST("", handler.CreateBooking)
	booking.GET("", handler.GetMyBookings)
	booking.POST("/:id", handler.CheckinBookedClass)
	booking.GET("/:id/qr-code", handler.RegenerateQRCode)

}
