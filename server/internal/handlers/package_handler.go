package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type PackageHandler struct {
	packageService services.PackageService
}

func NewPackageHandler(packageService services.PackageService) *PackageHandler {
	return &PackageHandler{packageService}
}

func (h *PackageHandler) CreatePackage(c *gin.Context) {
	var req dto.CreatePackageRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
		return
	}
	req.ImageURL = imageURL

	if err := h.packageService.CreatePackage(req); err != nil {
		utils.CleanupImageOnError(imageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create package", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Package created successfully"})
}

func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdatePackageRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if req.Image != nil {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
			return
		}
		req.ImageURL = imageURL
	}

	if err := h.packageService.UpdatePackage(id, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update package", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package updated successfully"})
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	id := c.Param("id")

	if err := h.packageService.DeletePackage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete package", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}

func (h *PackageHandler) GetAllPackages(c *gin.Context) {
	packages, err := h.packageService.GetAllPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch packages", "error": err.Error()})
		return
	}

	if packages == nil {
		packages = []dto.PackageResponse{}
	}

	c.JSON(http.StatusOK, gin.H{"packages": packages})
}

func (h *PackageHandler) GetPackageByID(c *gin.Context) {
	id := c.Param("id")

	pkg, err := h.packageService.GetPackageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Package not found"})
		return
	}

	c.JSON(http.StatusOK, pkg)
}
