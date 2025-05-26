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
	booking.POST("/:id/check-in", handler.CheckinBookedClass)
	booking.POST("/:id/check-out", handler.CheckinBookedClass)

}

// TODO : Add booking routes for customer to cancel out their booking
// booking.DELETE("/:id", handler.CancelBooking
