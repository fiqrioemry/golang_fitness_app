package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

// AUTHENTICATION  ==============
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
	Role     string `json:"role"`
}

// AUTHENTICATION  =============

// CLASS  ======================
type CreateClassRequest struct {
	Title         string                  `form:"title" binding:"required"`
	Duration      int                     `form:"duration" binding:"required,min=15"`
	Description   string                  `form:"description" binding:"required"`
	Additional    []string                `form:"additional[]"`
	IsActive      bool                    `form:"isActive"`
	TypeID        string                  `form:"typeId" binding:"required"`
	LevelID       string                  `form:"levelId" binding:"required"`
	LocationID    string                  `form:"locationId" binding:"required"`
	CategoryID    string                  `form:"categoryId" binding:"required"`
	SubcategoryID string                  `form:"subcategoryId" binding:"required"`
	Image         *multipart.FileHeader   `form:"image" binding:"required"`
	ImageURL      string                  `form:"-"`
	Images        []*multipart.FileHeader `form:"images" binding:"omitempty"`
	ImageURLs     []string                `form:"-"`
}

type UpdateClassRequest struct {
	Title         string                `form:"title"`
	Duration      int                   `form:"duration"`
	IsActive      bool                  `form:"isActive"`
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
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Image       string   `json:"image"`
	IsActive    bool     `json:"isActive"`
	Duration    int      `json:"duration"`
	Description string   `json:"description"`
	Additional  []string `json:"additional"`
	Type        string   `json:"type"`
	Level       string   `json:"level"`
	Location    string   `json:"location"`
	Category    string   `json:"category"`
	Subcategory string   `json:"subcategory"`
	Galleries   []string `json:"galleries"`
	CreatedAt   string   `json:"createdAt"`
}

type GalleryResponse struct {
	ID        string `json:"id"`
	ImageURL  string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
}

type ClassResponse struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Image         string   `json:"image"`
	IsActive      bool     `json:"isActive"`
	Duration      int      `json:"duration"`
	Description   string   `json:"description"`
	Additional    []string `json:"additional"`
	TypeID        string   `json:"typeId"`
	LevelID       string   `json:"levelId"`
	LocationID    string   `json:"locationId"`
	Galleries     []string `json:"galleries"`
	CategoryID    string   `json:"categoryId"`
	SubcategoryID string   `json:"subcategoryId"`
	CreatedAt     string   `json:"createdAt"`
}

// CLASS  ========================

// CLASS CATEGORY ================
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

// CLASS CATEGORY ================

// CLASS SUBCATEGORY =============
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

// CLASS SUBCATEGORY =============

// CLASS TYPE ====================
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

// CLASS TYPE ===================

// CLASS LEVEL ==================
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

// CLASS LEVEL ==================

// LOCATION =====================
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

// LOCATION =========================

// PACKAGE ==========================
type CreatePackageRequest struct {
	Name        string                `form:"name" binding:"required,min=6"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"required,gt=0"`
	Credit      int                   `form:"credit" binding:"required,gt=0"`
	Expired     int                   `form:"expired" binding:"required,gt=0"`
	Discount    float64               `form:"discount"`
	Additional  []string              `form:"additional[]"`
	IsActive    bool                  `form:"isActive"`
	ClassIDs    []string              `form:"classIds[]"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL    string                `form:"-"`
}

type UpdatePackageRequest struct {
	Name        string                `form:"name" binding:"required,min=6"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"required,gt=0"`
	Credit      int                   `form:"credit" binding:"required,gt=0"`
	Expired     int                   `form:"expired" binding:"required,gt=0"`
	Discount    float64               `form:"discount"`
	Additional  []string              `form:"additional[]"`
	IsActive    bool                  `form:"isActive"`
	ClassIDs    []string              `form:"classIds[]"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL    string                `form:"-"`
}

type PackageResponse struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Credit      int                    `json:"credit"`
	Expired     int                    `json:"expired"`
	Image       string                 `json:"image"`
	Discount    float64                `json:"discount"`
	IsActive    bool                   `json:"isActive"`
	Additional  []string               `json:"additional"`
	Classes     []ClassSummaryResponse `json:"classes"`
}

type PackageDetailResponse struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Credit      int                    `json:"credit"`
	Expired     int                    `json:"expired"`
	Discount    float64                `form:"discount"`
	Image       string                 `json:"image"`
	IsActive    bool                   `json:"isActive"`
	Additional  []string               `json:"additional"`
	Classes     []ClassSummaryResponse `json:"classes"`
}

type ClassSummaryResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Duration int    `json:"duration"`
}

// PACKAGE ==========================

// INSTRUCTOR ==========================
type CreateInstructorRequest struct {
	UserID         string `json:"userId" binding:"required,uuid"`
	Experience     int    `json:"experience" binding:"required,min=0"`
	Specialties    string `json:"specialties" binding:"required"`
	Certifications string `json:"certifications"`
}

type UpdateInstructorRequest struct {
	UserID         string `json:"userId" binding:"required,uuid"`
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

// INSTRUCTOR ==========================

// PAYMENT ================================
type CreatePaymentRequest struct {
	PackageID   string  `json:"packageId" binding:"required"`
	VoucherCode *string `json:"voucherCode"` // Optional
}

type CreatePaymentResponse struct {
	PaymentID string `json:"paymentId"`
	SnapToken string `json:"snapToken"`
	SnapURL   string `json:"snapUrl"`
}

type MidtransNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

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

type PaymentQueryParam struct {
	Q     string `form:"q"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}

type AdminPaymentResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"userId"`
	UserEmail     string  `json:"userEmail"`
	Fullname      string  `json:"fullname"`
	PackageID     string  `json:"packageId"`
	PackageName   string  `json:"packageName"`
	Price         float64 `json:"price"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	PaidAt        string  `json:"paidAt"`
}

type AdminPaymentListResponse struct {
	Payments []AdminPaymentResponse `json:"payments"`
	Total    int64                  `json:"total"`
	Page     int                    `json:"page"`
	Limit    int                    `json:"limit"`
}

// CLASS-SCHEDULE ====================
type CreateClassScheduleRequest struct {
	ClassID      string    `json:"classId" binding:"required"`
	InstructorID string    `json:"instructorId" binding:"required"`
	Date         time.Time `json:"date" binding:"required"`
	StartHour    int       `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int       `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	Capacity     int       `json:"capacity" validate:"required,min=1"`
	Color        string    `json:"color"`
}

type UpdateClassScheduleRequest struct {
	ClassID      string    `json:"classId" binding:"required"`
	InstructorID string    `json:"instructorId" binding:"required"`
	Date         time.Time `json:"date" binding:"required"`
	StartHour    int       `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int       `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	Capacity     int       `json:"capacity" validate:"required,min=1"`
	Color        string    `json:"color"`
}

type CreateScheduleRequest struct {
	IsRecurring  bool       `json:"isRecurring"`
	ClassID      string     `json:"classId" binding:"required"`
	InstructorID string     `json:"instructorId" binding:"required"`
	Capacity     int        `json:"capacity" binding:"required"`
	Color        string     `json:"color"`
	Date         *time.Time `json:"date,omitempty"`
	StartHour    int        `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int        `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	DayOfWeeks   []int      `json:"dayOfWeeks" binding:"required,dive,min=0,max=6"`
	EndDate      *time.Time `json:"endDate,omitempty"`
}

type ScheduleTemplateToggleRequest struct {
	IsActive bool `json:"isActive" binding:"required"`
}

type ClassScheduleResponse struct {
	ID          string          `json:"id"`
	Class       ClassBrief      `json:"class"`
	Instructor  InstructorBrief `json:"instructor"`
	Category    string          `json:"category"`
	Date        time.Time       `json:"date"`
	StartHour   int             `json:"startHour"`
	StartMinute int             `json:"startMinute"`
	Capacity    int             `json:"capacity"`
	BookedCount int             `json:"bookedCount"`
	Color       string          `json:"color"`
	IsBooked    bool            `json:"isBooked"`
}

type ClassScheduleDetailResponse struct {
	ClassScheduleResponse
	Packages []PackageResponse `json:"packages"`
}

type ClassScheduleQueryParam struct {
	StartDate  string `form:"startDate"`
	EndDate    string `form:"endDate"`
	CategoryID string `form:"categoryId"`
}

type ScheduleTemplateResponse struct {
	ID           uuid.UUID `json:"id"`
	ClassID      uuid.UUID `json:"classId"`
	ClassName    string    `json:"className"`
	InstructorID uuid.UUID `json:"instructorId"`
	Instructor   string    `json:"instructor"`
	DayOfWeeks   []int     `json:"dayOfWeeks"`
	StartHour    int       `json:"startHour"`
	StartMinute  int       `json:"startMinute"`
	Capacity     int       `json:"capacity"`
	IsActive     bool      `json:"isActive"`
	Frequency    string    `json:"frequency"`
	EndDate      string    `json:"endDate"`
	CreatedAt    string    `json:"createdAt"`
}

type CreateScheduleTemplateRequest struct {
	ClassID      string     `json:"classId" binding:"required"`
	InstructorID string     `json:"instructorId" binding:"required"`
	DayOfWeeks   []int      `json:"dayOfWeeks" binding:"required,dive,min=0,max=6"`
	StartHour    int        `json:"startHour" binding:"required,min=0,max=23"`
	StartMinute  int        `json:"startMinute" binding:"required,min=0,max=59"`
	Capacity     int        `json:"capacity" binding:"required,gt=0"`
	Color        string     `json:"color" binding:"required"`
	EndDate      *time.Time `json:"endDate" binding:"required"`
}

type UpdateScheduleTemplateRequest struct {
	ClassID      string     `json:"classId" binding:"required"`
	InstructorID string     `json:"instructorId" binding:"required"`
	DayOfWeeks   []int      `json:"dayOfWeeks" binding:"required,dive,min=0,max=6"`
	StartHour    int        `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int        `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	Capacity     int        `json:"capacity" binding:"required,gt=0"`
	EndDate      *time.Time `json:"endDate,omitempty"`
}

// CLASS-SCHEDULE =====================

// BOOKINGS ===========================
type CreateBookingRequest struct {
	PackageID       string `json:"packageId" binding:"required,uuid"`
	ClassScheduleID string `json:"scheduleId" binding:"required,uuid"`
}

type BookingResponse struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	BookedAt string `json:"bookedAt"`

	ClassID    string `json:"classId"`
	ClassTitle string `json:"classTitle"`
	ClassImage string `json:"classImage"`
	Duration   int    `json:"duration"`

	Date        string `json:"date"`
	StartHour   int    `json:"startHour"`
	StartMinute int    `json:"startMinute"`
	Location    string `json:"location"`

	Instructor  string `json:"instructor"`
	Participant int    `json:"participant"`
}
type QRCodeAttendanceResponse struct {
	QR         string `json:"qr"`
	ClassTitle string `json:"classTitle"`
	Date       string `json:"date"`
	Instructor string `json:"instructor"`
	StartTime  string `json:"startTime"`
}

// ATTENDANCE ==========================

type ValidateQRRequest struct {
	QRCode string `json:"qrCode"`
}

type MarkAttendanceRequest struct {
	BookingID string `json:"bookingId" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=attended absent cancelled"`
}

type AttendanceResponse struct {
	ID          string          `json:"id"`
	Class       ClassBrief      `json:"class"`
	Instructor  InstructorBrief `json:"instructor"`
	Fullname    string          `json:"fullname"`
	Date        string          `json:"date"`
	StartHour   int             `json:"startHour"`
	StartMinute int             `json:"startMinute"`
	Status      string          `json:"status"`
	CheckedAt   string          `json:"checkedAt"`
	Reviewed    bool            `json:"reviewed"`
	Verified    bool            `json:"verified"`
}

type ClassBrief struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Duration int    `json:"duration"`
}

type InstructorBrief struct {
	ID       string  `json:"id"`
	Fullname string  `json:"fullname"`
	Rating   float64 `json:"rating"`
}

type AttendanceDetailResponse struct {
	ID        string `json:"id"`
	Fullname  string `json:"fullname"`
	Avatar    string `json:"avatar"`
	CheckedAt string `json:"checkedAt"`
}

// ATTENDANCE ==========================

// REVIEWS ==========================
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

// ATTENDANCE ==========================

// USER-LIST ==========================
type UserQueryParam struct {
	Q     string `form:"q"`
	Role  string `form:"role"`
	Sort  string `form:"sort"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}

type UserListResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

type UserDetailResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday,omitempty"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserStatsResponse struct {
	Total        int64 `json:"total"`
	Customers    int64 `json:"customers"`
	Instructors  int64 `json:"instructors"`
	Admins       int64 `json:"admins"`
	NewThisMonth int64 `json:"newThisMonth"`
}

// USER-LIST ==========================

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

type TransactionResponse struct {
	ID            string  `json:"id"`
	PackageID     string  `json:"packageId"`
	PackageName   string  `json:"packageName"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	BasePrice     float64 `json:"basePrice"`
	Tax           float64 `json:"taxRate"`
	Price         float64 `json:"price"`
	PaidAt        string  `json:"paidAt"`
}

type TransactionListResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
	Total        int64                 `json:"total"`
	Page         int                   `json:"page"`
	Limit        int                   `json:"limit"`
}

type UserPackageListResponse struct {
	Packages []UserPackageResponse `json:"packages"`
	Total    int64                 `json:"total"`
	Page     int                   `json:"page"`
	Limit    int                   `json:"limit"`
}

type UserPackageResponse struct {
	ID              string `json:"id"`
	PackageID       string `json:"packageId"`
	PackageName     string `json:"packageName"`
	RemainingCredit int    `json:"remainingCredit"`
	ExpiredAt       string `json:"expiredAt,omitempty"`
	ExpiredInDays   int    `json:"expiredInDays,omitempty"`
	PurchasedAt     string `json:"purchasedAt"`
}

type BookingListResponse struct {
	Bookings []BookingResponse `json:"bookings"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}

// NOTIFICATIONS
type NotificationSettingResponse struct {
	TypeID  string `json:"typeId"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Channel string `json:"channel"`
	Enabled bool   `json:"enabled"`
}

type UpdateNotificationSettingRequest struct {
	TypeID  string `json:"typeId" binding:"required"`
	Channel string `json:"channel" binding:"required,oneof=email browser"`
	Enabled bool   `json:"enabled"`
}

type CreateNotificationRequest struct {
	UserID   string `json:"userId"`
	TypeCode string `json:"typeCode"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Channel  string `json:"channel"` // "email" / "browser"
}

type NotificationResponse struct {
	ID        string `json:"id"`
	TypeCode  string `json:"typeCode"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}

type SendPromoNotificationRequest struct {
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// VOUCHER
type CreateVoucherRequest struct {
	Code         string   `json:"code" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	DiscountType string   `json:"discountType" binding:"required,oneof=fixed percentage"`
	Discount     float64  `json:"discount" binding:"required,gt=0"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota" binding:"required,gt=0"`
	ExpiredAt    string   `json:"expiredAt" binding:"required,datetime=2006-01-02"`
}

type VoucherResponse struct {
	ID           string   `json:"id"`
	Code         string   `json:"code"`
	Description  string   `json:"description"`
	DiscountType string   `json:"discountType"`
	Discount     float64  `json:"discount"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota"`
	ExpiredAt    string   `json:"expiredAt"`
	CreatedAt    string   `json:"createdAt"`
}

type ApplyVoucherRequest struct {
	UserID *string `json:"userId"`
	Code   string  `json:"code" binding:"required"`
	Total  float64 `json:"total" binding:"required"`
}

type ApplyVoucherResponse struct {
	Code          string   `json:"code"`
	DiscountType  string   `json:"discountType"`
	Discount      float64  `json:"discount"`
	MaxDiscount   *float64 `json:"maxDiscount,omitempty"`
	DiscountValue float64  `json:"discountValue"`
	FinalTotal    float64  `json:"finalTotal"`
}
