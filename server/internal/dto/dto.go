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
	ImageURL      string                `form:"-"`
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
	ImageURL      string                `form:"-"`
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

type ClassDetailResponse struct {
	ID             string            `json:"id"`
	Title          string            `json:"title"`
	Image          string            `json:"image"`
	IsActive       bool              `json:"isActive"`
	Duration       int               `json:"duration"`
	Description    string            `json:"description"`
	AdditionalList []string          `json:"additional"`
	Type           string            `json:"type"`
	Level          string            `json:"level"`
	Location       string            `json:"location"`
	Category       string            `json:"category"`
	Subcategory    string            `json:"subcategory"`
	Galleries      []GalleryResponse `json:"galleries"`
	Reviews        []ReviewResponse  `json:"reviews"`
	CreatedAt      string            `json:"createdAt"`
}

type GalleryResponse struct {
	ID        string `json:"id"`
	ImageURL  string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
}

type ClassResponse struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Image          string   `json:"image"`
	IsActive       bool     `json:"isActive"`
	Duration       int      `json:"duration"`
	Description    string   `json:"description"`
	AdditionalList []string `json:"additional"`
	TypeID         string   `json:"typeId"`
	LevelID        string   `json:"levelId"`
	LocationID     string   `json:"locationId"`
	CategoryID     string   `json:"categoryId"`
	SubcategoryID  string   `json:"subcategoryId"`
	CreatedAt      string   `json:"createdAt"`
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

// package
// Create
type CreatePackageRequest struct {
	Name        string                `form:"name" binding:"required,min=2"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"required,gt=0"`
	Credit      int                   `form:"credit" binding:"required,gt=0"`
	Expired     int                   `form:"expired"`
	Information string                `form:"information"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL    string                `form:"-"`
}

// Update
type UpdatePackageRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Price       float64               `form:"price"`
	Credit      int                   `form:"credit"`
	Expired     int                   `form:"expired"`
	Information string                `form:"information"`
	Image       *multipart.FileHeader `form:"image"`
	ImageURL    string                `form:"-"`
}

// Response
type PackageResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Credit      int     `json:"credit"`
	Image       string  `json:"image"`
	IsActive    bool    `json:"isActive"`
}

type PackageDetailResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Credit      int     `json:"credit"`
	Expired     int     `json:"expired"`
	Image       string  `json:"image"`
	IsActive    bool    `json:"isActive"`
	Information string  `json:"information"`
}

// instructor

type CreateInstructorRequest struct {
	Experience     int    `json:"experience" binding:"required,min=0"`
	Specialties    string `json:"specialties" binding:"required"`
	Certifications string `json:"certifications"`
}

type UpdateInstructorRequest struct {
	Experience     int    `json:"experience"`
	Specialties    string `json:"specialties"`
	Certifications string `json:"certifications"`
}

type InstructorResponse struct {
	ID             string  `json:"id"`
	UserID         string  `json:"userId"`
	Fullname       string  `json:"fullname"`
	Avatar         string  `json:"avatar"`
	Experience     int     `json:"experience"`
	Specialties    string  `json:"specialties"`
	Certifications string  `json:"certifications"`
	Rating         float64 `json:"rating"`
	TotalClass     int     `json:"totalClass"`
}

// Create Payment Request (user beli package)
type CreatePaymentRequest struct {
	PackageID string `json:"packageId" binding:"required"`
}

// Response Setelah Create Payment
type CreatePaymentResponse struct {
	PaymentID string `json:"paymentId"`
	SnapToken string `json:"snapToken"`
}

// Request Midtrans Webhook (notif dari Midtrans ke backend)
type MidtransNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

// Payment Detail Response (optional buat lihat status payment)
type PaymentResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"userId"`
	PackageID     string  `json:"packageId"`
	PackageName   string  `json:"packageName"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
	PaidAt        string  `json:"paidAt"`
}

// user package
type UserPackageResponse struct {
	ID              string `json:"id"`
	PackageName     string `json:"packageName"`
	RemainingCredit int    `json:"remainingCredit"`
	ExpiredAt       string `json:"expiredAt"`
	PurchasedAt     string `json:"purchasedAt"`
}

type CreateClassScheduleRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	StartTime    string `json:"startTime" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Capacity     int    `json:"capacity" binding:"required,gt=0"`
}

type UpdateClassScheduleRequest struct {
	StartTime string `json:"startTime" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	EndTime   string `json:"endTime" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	Capacity  int    `json:"capacity" binding:"omitempty,gt=0"`
}

type ClassScheduleResponse struct {
	ID             string `json:"id"`
	ClassID        string `json:"classId"`
	ClassTitle     string `json:"classTitle"`
	InstructorID   string `json:"instructorId"`
	InstructorName string `json:"instructorName"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	Capacity       int    `json:"capacity"`
	BookedCount    int    `json:"bookedCount"`
}

type CreateScheduleTemplateRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	DayOfWeek    int    `json:"dayOfWeek" binding:"required,min=0,max=6"`
	StartHour    int    `json:"startHour" binding:"required,min=0,max=23"`
	StartMinute  int    `json:"startMinute" binding:"required,min=0,max=59"`
	Capacity     int    `json:"capacity" binding:"required,gt=0"`
}

// booking

type CreateBookingRequest struct {
	ClassScheduleID string `json:"classScheduleId" binding:"required"`
}

type BookingResponse struct {
	ID         string `json:"id"`
	ClassTitle string `json:"classTitle"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Status     string `json:"status"`
}

// Attendance

type MarkAttendanceRequest struct {
	BookingID string `json:"bookingId" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=attended absent cancelled"`
}

type AttendanceResponse struct {
	ID           string `json:"id"`
	ClassTitle   string `json:"classTitle"`
	UserFullname string `json:"userFullname"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Status       string `json:"status"`
	CheckedAt    string `json:"checkedAt"`
}

// Review
type CreateReviewRequest struct {
	ClassID string `json:"classId" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment" binding:"omitempty"`
}

type ReviewResponse struct {
	ID         string `json:"id"`
	UserName   string `json:"userName"`
	ClassTitle string `json:"classTitle"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"createdAt"`
}
