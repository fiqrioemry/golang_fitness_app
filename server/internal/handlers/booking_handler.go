package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service services.BookingService
}

func NewBookingHandler(service services.BookingService) *BookingHandler {
	return &BookingHandler{service}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req dto.CreateBookingRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	userID := utils.MustGetUserID(c)

	if err := h.service.CreateBooking(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking successful"})
}

func (h *BookingHandler) GetUserBookings(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	bookings, err := h.service.GetUserBookings(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bookings", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
