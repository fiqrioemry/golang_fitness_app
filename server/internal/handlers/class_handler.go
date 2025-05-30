package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClassHandler struct {
	classService services.ClassService
}

func NewClassHandler(classService services.ClassService) *ClassHandler {
	return &ClassHandler{classService}
}

func extractUploadedImages(c *gin.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		return nil, err
	}
	files := form.File["images"]
	if len(files) == 0 {
		return nil, nil
	}
	return utils.UploadMultipleImagesWithValidation(files)
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var req dto.CreateClassRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	req.IsActive, _ = utils.ParseBoolFormField(c, "isActive")

	if req.Image != nil {
		singleImageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
			return
		}
		req.ImageURL = singleImageURL
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image is required"})
		return
	}
	uploadedURLs, err := extractUploadedImages(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid image upload", "error": err.Error()})
		return
	}
	req.ImageURLs = uploadedURLs

	if err := h.classService.CreateClass(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassRequest
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

	if err := h.classService.UpdateClass(id, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id := c.Param("id")

	if err := h.classService.DeleteClass(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")

	classResponse, err := h.classService.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, classResponse)
}

func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	var params dto.ClassQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	classes, pagination, err := h.classService.GetAllClasses(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch classes", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"classes":    classes,
		"pagination": pagination,
	})
}

func (h *ClassHandler) UploadClassGallery(c *gin.Context) {
	classID := c.Param("id")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid multipart form"})
		return
	}

	keepImages := form.Value["images"]

	files := form.File["images"]
	var uploadedURLs []string

	if len(files) > 0 {
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to open file", "error": err.Error()})
				return
			}
			defer file.Close()

			imageURL, err := utils.UploadToCloudinary(file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to upload image", "error": err.Error()})
				return
			}
			uploadedURLs = append(uploadedURLs, imageURL)
		}
	}

	if err := h.classService.UpdateClassGallery(uuid.MustParse(classID), keepImages, uploadedURLs); err != nil {
		utils.CleanupImagesOnError(uploadedURLs)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update gallery", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "gallery updated successfully"})
}
