package middleware

import (
	"time"

	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		status := c.Writer.Status()
		log.Printf("| %3d | %13v | %15s | %-7s  %s\n",
			status,
			latency,
			c.ClientIP(),
			c.Request.Method,
			path,
		)
	}
}
