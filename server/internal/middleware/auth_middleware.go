package middleware

import (
	"net/http"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("accessToken")
		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized!! Token missing"})
			return
		}

		claims, err := utils.DecodeAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			return
		}

		c.Set("role", claims.Role)
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.MustGetRole(c)
		if !utils.IsAdmin(role) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden: Admin only"})
			return
		}
		c.Next()
	}
}
