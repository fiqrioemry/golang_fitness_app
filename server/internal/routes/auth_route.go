package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, handler *handlers.AuthHandler) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/send-otp", handler.SendOTP)
		auth.POST("/verify-otp", handler.VerifyOTP)
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/refresh-token", handler.RefreshToken)

		// google
		auth.GET("/google", handler.GoogleOAuthRedirect)
		auth.GET("/google/callback", handler.GoogleOAuthCallback)

		protected := auth.Group("")
		protected.Use(middleware.AuthRequired())
		protected.GET("/me", handler.AuthMe)
	}
}
