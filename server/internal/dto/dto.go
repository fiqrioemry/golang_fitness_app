package dto

import (
	"mime/multipart"
	"time"
)

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

// class request
type CreateClassRequest struct {
	Title         string                `form:"title" binding:"required"`
	Duration      int                   `form:"duration" binding:"required,min=15"`
	Description   string                `form:"description" binding:"required"`
	Additional    []string              `form:"additional[]"`
	TypeID        string                `form:"typeId" binding:"required"`
	LevelID       string                `form:"levelId" binding:"required"`
	LocationID    string                `form:"locationId" binding:"required"`
	CategoryID    string                `form:"categoryId" binding:"required"`
	SubcategoryID string                `form:"subcategoryId" binding:"required"`
	Image         *multipart.FileHeader `form:"image" binding:"required"`
}

type UpdateClassRequest struct {
	Title         string                `form:"title"`
	Duration      int                   `form:"duration"`
	Description   string                `form:"description"`
	Additional    []string              `form:"additional[]"`
	TypeID        string                `form:"typeId"`
	LevelID       string                `form:"levelId"`
	LocationID    string                `form:"locationId"`
	CategoryID    string                `form:"categoryId"`
	SubcategoryID string                `form:"subcategoryId"`
	Image         *multipart.FileHeader `form:"image"`
}

// class response
type ClassQueryParam struct {
	Q          string `form:"q"`
	TypeID     string `form:"typeId"`
	LevelID    string `form:"levelId"`
	LocationID string `form:"locationId"`
	CategoryID string `form:"categoryId"`
	IsActive   string `form:"isActive"`
	Page       int    `form:"page,default=1"`
	Limit      int    `form:"limit,default=10"`
	Sort       string `form:"sort,default=latest"`
}

type ClassResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Image         string    `json:"image"`
	IsActive      bool      `json:"isActive"`
	Duration      int       `json:"duration"`
	Description   string    `json:"description"`
	Additional    []string  `json:"additional"`
	TypeID        string    `json:"typeId"`
	LevelID       string    `json:"levelId"`
	LocationID    string    `json:"locationId"`
	CategoryID    string    `json:"categoryId"`
	SubcategoryID string    `json:"subcategoryId"`
	CreatedAt     time.Time `json:"createdAt"`
}

// Category
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Subcategory

type CreateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required,min=2"`
	CategoryID string `json:"categoryId" binding:"required"`
}

type UpdateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required,min=2"`
	CategoryID string `json:"categoryId" binding:"required"`
}

type SubcategoryResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"categoryId"`
}

// Type class
type CreateTypeRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type UpdateTypeRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type TypeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Level class
type CreateLevelRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type UpdateLevelRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type LevelResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// location
type CreateLocationRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Address     string `json:"address" binding:"required"`
	GeoLocation string `json:"geoLocation" binding:"required"`
}

type UpdateLocationRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Address     string `json:"address" binding:"required"`
	GeoLocation string `json:"geoLocation" binding:"required"`
}

type LocationResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	GeoLocation string `json:"geoLocation"`
}
