package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.CreatePaymentRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	response, err := h.paymentService.CreatePayment(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create payment", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *PaymentHandler) HandlePaymentNotification(c *gin.Context) {
	var notif dto.MidtransNotificationRequest
	if !utils.BindAndValidateJSON(c, &notif) {
		return
	}

	if err := h.paymentService.HandlePaymentNotification(notif); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to process payment notification", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment succesfully"})
}

func (h *PaymentHandler) GetAllUserPayments(c *gin.Context) {
	var params dto.PaymentQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	payments, pagination, err := h.paymentService.GetAllUserPayments(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch payments", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       payments,
		"pagination": pagination,
	})

}
