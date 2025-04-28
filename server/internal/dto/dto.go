package dto

import "time"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required,min=5"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type AuthMeResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
}

type ProfileResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Avatar    string    `json:"avatar"`
	Gender    string    `json:"gender"`
	Birthday  string    `json:"birthday"`
	Bio       string    `json:"bio"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"joinedAt"`
}

type UpdateProfileRequest struct {
	Fullname string `form:"fullname" binding:"required,min=5"`
	Birthday string `form:"birthday"`
	Gender   string `form:"gender"`
	Phone    string `form:"phone"`
	Bio      string `form:"bio"`
}
