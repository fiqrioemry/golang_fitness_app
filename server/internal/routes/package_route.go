package routes

import (
	"server/internal/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(r *gin.Engine, handler *handlers.PackageHandler) {
	pkg := r.Group("/api/packages")
	{
		// public endpoints
		pkg.GET("", handler.GetAllPackages)
		pkg.GET("/:id", handler.GetPackageByID)

		// admin-protected endpoints
		admin := pkg.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
		admin.POST("", handler.CreatePackage)
		admin.PUT("/:id", handler.UpdatePackage)
		admin.DELETE("/:id", middleware.RoleOnly("owner"), handler.DeletePackage)
	}
}

// Endpoint List:
// GET    /api/packages              → Ambil semua paket yang tersedia
// POST   /api/packages              → Buat paket baru (admin only)
// GET    /api/packages/:id          → Ambil detail paket berdasarkan ID
// PUT    /api/packages/:id          → Update data paket (admin only)
// DELETE /api/packages/:id          → Hapus paket (khusus role owner)
