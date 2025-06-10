package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system ENV")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func MustGetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		panic("userID not found in context")
	}
	idStr, ok := userID.(string)
	if !ok {
		panic("userID in context is not a string")
	}
	return idStr
}

func MustGetRole(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		panic("role not found in context")
	}
	userRole, ok := role.(string)
	if !ok {
		panic("role in context is not a string")
	}
	return userRole
}

func BindAndValidateJSON[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return false
	}
	return true
}
func BindAndValidateForm[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid form-data request",
			"error":   err.Error(),
		})
		return false
	}
	return true
}

func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	valStr := c.Query(key)
	val, err := strconv.Atoi(valStr)
	if err != nil || val <= 0 {
		return defaultValue
	}
	return val
}

func IsCustomer(role string) bool {
	return role == "customer"
}

func IsAdmin(role string) bool {
	return role == "admin"
}

func IsInstructor(role string) bool {
	return role == "instructor"
}

func ParseBoolFormField(c *gin.Context, field string) (bool, error) {
	val := c.PostForm(field)
	if val == "" {
		return false, nil
	}
	return strconv.ParseBool(val)
}

func GetTaxRate() float64 {
	val := os.Getenv("PAYMENT_TAX_RATE")
	if val == "" {
		return 0.05
	}
	rate, err := strconv.ParseFloat(val, 64)
	if err != nil || rate < 0 {
		return 0.05
	}
	return rate
}

func ParseDayOfWeek(day string) time.Weekday {
	switch strings.ToLower(day) {
	case "sunday":
		return time.Sunday
	case "monday":
		return time.Monday
	case "tuesday":
		return time.Tuesday
	case "wednesday":
		return time.Wednesday
	case "thursday":
		return time.Thursday
	case "friday":
		return time.Friday
	case "saturday":
		return time.Saturday
	default:
		return -1
	}
}

func IntSliceToJSON(data []int) datatypes.JSON {
	bytes, _ := json.Marshal(data)
	return datatypes.JSON(bytes)
}
func IsDayMatched(currentDay int, allowedDays []int) bool {
	return slices.Contains(allowedDays, currentDay)
}

func ParseJSONToIntSlice(jsonStr string) []int {
	var result []int
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Printf("Failed to parse JSON string to []int: %v\n", err)
		return []int{}
	}
	return result
}

func ParseDate(dateStr string) (time.Time, error) {
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t.UTC(), nil
	}
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t.UTC(), nil
	}

	return time.Time{}, fmt.Errorf("invalid date, format must be YYYY-MM-DD or ISO 8601")
}
