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
	if !utils.BindAndValidateForm(c, &req) {
		return
	}
	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")

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
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")

	if req.Image != nil && req.Image.Filename != "" {
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

func (h *PackageHandler) GetAllPackages(c *gin.Context) {
	var params dto.PackageQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	packages, pagination, err := h.packageService.GetAllPackages(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch payments", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       packages,
		"pagination": pagination,
	})
}

func (h *PackageHandler) GetPackageByID(c *gin.Context) {
	id := c.Param("id")

	classPackage, err := h.packageService.GetPackageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Package not found"})
		return
	}

	c.JSON(http.StatusOK, classPackage)
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	id := c.Param("id")

	if err := h.packageService.DeletePackage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}
