package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func LimitFileSize(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		c.Next()
	}
}

func RateLimiter(r rate.Limit, b int) gin.HandlerFunc {
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
