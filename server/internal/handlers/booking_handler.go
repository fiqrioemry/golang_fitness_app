package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *BookingHandler {
	return &BookingHandler{bookingService}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req dto.CreateBookingRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	userID := utils.MustGetUserID(c)

	err := h.bookingService.CreateBooking(userID, req.PackageID, req.ClassScheduleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking successful"})
}

func (h *BookingHandler) GetMyBookings(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var params dto.BookingQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	data, pagination, err := h.bookingService.GetBookingByUser(userID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bookings", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"pagination": pagination,
	})
}

func (h *BookingHandler) GetBookingDetail(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	response, err := h.bookingService.GetBookingByID(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bookings", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{response})
}

func (h *BookingHandler) CheckinBookedClass(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	response, err := h.bookingService.CheckedInClassSchedule(userID, bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BookingHandler) CheckoutBookedClass(c *gin.Context) {
	bookingID := c.Param("id")
	userID := utils.MustGetUserID(c)

	var req dto.ValidateCheckoutRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.bookingService.CheckoutClassSchedule(userID, bookingID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance verified successfully"})
}
