package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, handler *handlers.AuthHandler) {
	auth := r.Group("/api/auth")
	{
		// public endpoints
		auth.POST("/send-otp", handler.SendOTP)
		auth.POST("/verify-otp", handler.VerifyOTP)
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/refresh-token", handler.RefreshToken)
		auth.GET("/google", handler.GoogleOAuthRedirect)
		auth.GET("/google/callback", handler.GoogleOAuthCallback)

		// protected-endpoints
		protected := auth.Group("")
		protected.Use(middleware.AuthRequired())
		protected.GET("/me", handler.AuthMe)
	}
}

// POST   /api/auth/send-otp         → Kirim OTP ke email
// POST   /api/auth/verify-otp       → Verifikasi OTP dari email
// POST   /api/auth/register         → Registrasi akun baru dengan OTP terverifikasi
// POST   /api/auth/login            → Login user (email & password)
// POST   /api/auth/logout           → Logout user dan hapus refresh token
// POST   /api/auth/refresh-token    → Refresh JWT token (akses token baru)
// GET    /api/auth/google           → Redirect ke login Google OAuth
// GET    /api/auth/google/callback  → Callback dari Google OAuth
// GET    /api/auth/me               → Ambil data user login (hanya jika token valid)

// GET    /api/user-packages   						→ Callback dari Google OAuth
// GET    /api/user-packages/class/:id              → Ambil data user login (hanya jika token valid)
