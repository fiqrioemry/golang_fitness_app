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
	var req dto.CreatePaymentRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	userID := utils.MustGetUserID(c)

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

	c.JSON(http.StatusOK, gin.H{"message": "Notification processed successfully"})
}

func (h *PaymentHandler) GetAllUserPayments(c *gin.Context) {
	var q dto.PaymentQueryParam
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query params", "error": err.Error()})
		return
	}

	resp, err := h.paymentService.GetAllUserPayments(q.Q, q.Page, q.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch payments", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
