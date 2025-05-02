package utils

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
