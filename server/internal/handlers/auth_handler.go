package handlers

import (
	"net/http"
	"os"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) SendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.authService.SendOTP(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	tokens, err := h.authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email already registered", "error": err.Error()})
		return
	}
	utils.SetAccessTokenCookie(c, tokens.AccessToken)

	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.authService.VerifyOTP(req.Email, req.OTP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	tokens, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Login Successfully"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Refresh token cookie missing"})
		return
	}

	if err := h.authService.Logout(refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.ClearAccessTokenCookie(c)

	utils.ClearRefreshTokenCookie(c)

	c.JSON(http.StatusOK, gin.H{"message": "Logout Successfully"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token cookie missing"})
		return
	}

	tokens, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)

	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Refresh Successfully"})
}

func (h *AuthHandler) AuthMe(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	user, err := h.authService.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	response := dto.AuthMeResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Profile.Fullname,
		Avatar:   user.Profile.Avatar,
		Role:     user.Role,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) GoogleOAuthRedirect(c *gin.Context) {
	url := h.authService.GetGoogleOAuthURL()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleOAuthCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authorization code is missing"})
		return
	}

	tokens, err := h.authService.HandleGoogleOAuthCallback(code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)

	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_REDIRECT_URL"))
}
