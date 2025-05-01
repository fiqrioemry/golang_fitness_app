package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, handler *handlers.ProfileHandler) {
	user := r.Group("/api/user")
	user.Use(middleware.AuthRequired())
	user.GET("/profile", handler.GetProfile)
	user.PUT("/profile", handler.UpdateProfile)
	user.PUT("/profile/avatar", handler.UpdateAvatar)
	user.GET("/transactions", handler.GetUserTransactions)
	user.GET("/packages", handler.GetUserPackages)
	user.GET("/bookings", handler.GetUserBookings)
}
