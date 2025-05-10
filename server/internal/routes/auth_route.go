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
		auth.POST("/google-signin", handler.GoogleSignIn)
		auth.POST("/refresh-token", handler.RefreshToken)

		protected := auth.Group("")
		protected.Use(middleware.AuthRequired())
		protected.GET("/me", handler.AuthMe)
	}
}
