package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, handler *handlers.ProfileHandler) {
	auth := r.Group("/api/user")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/profile", handler.GetProfile)
		auth.POST("/profile", handler.SendOTP)
		auth.POST("/verify-otp", handler.VerifyOTP)
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/refresh-token", handler.RefreshToken)

		protected := auth.Group("")
		protected.Use(middleware.AuthRequired())
		protected.GET("/me", handler.AuthMe)
	}
}
