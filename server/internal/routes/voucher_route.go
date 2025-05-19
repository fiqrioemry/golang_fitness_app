package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func VoucherRoutes(r *gin.Engine, h *handlers.VoucherHandler) {
	v := r.Group("/api/vouchers")
	v.POST("/apply", h.ApplyVoucher)
	v.Use(middleware.AuthRequired())
	{
		v.GET("", h.GetAllVouchers)
		v.POST("", middleware.RoleOnly("admin"), h.CreateVoucher)
		v.PUT(":id", h.UpdateVoucher)
		v.DELETE(":id", h.DeleteVoucher)
	}
}
