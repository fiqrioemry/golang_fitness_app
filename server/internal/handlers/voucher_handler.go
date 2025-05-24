package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	service services.VoucherService
}

func NewVoucherHandler(service services.VoucherService) *VoucherHandler {
	return &VoucherHandler{service}
}

func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
	var req dto.CreateVoucherRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateVoucher(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Voucher created"})
}

func (h *VoucherHandler) UpdateVoucher(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateVoucherRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateVoucher(id, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voucher updated"})
}

func (h *VoucherHandler) GetAllVouchers(c *gin.Context) {
	vouchers, err := h.service.GetAllVouchers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get vouchers", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vouchers)
}
func (h *VoucherHandler) ApplyVoucher(c *gin.Context) {
	var req dto.ApplyVoucherRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	res, err := h.service.ApplyVoucher(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *VoucherHandler) DeleteVoucher(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteVoucher(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voucher deleted"})
}
