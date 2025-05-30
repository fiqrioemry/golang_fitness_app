package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.Engine, handler *handlers.ReviewHandler) {
	review := r.Group("/api/reviews")
	review.Use(middleware.AuthRequired())

	review.POST("/:id", handler.CreateReviewFromBookingID)
	r.GET("/api/reviews/:classId", handler.GetReviewsByClass)
}
