package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/models"
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

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var req dto.CreateClassRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// upload optional gallery images
	if len(req.Images) > 0 {
		uploadedImageURLs, err := utils.UploadMultipleImagesWithValidation(req.Images)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Gallery images upload failed", "error": err.Error()})
			return
		}
		req.ImageURLs = uploadedImageURLs
	}

	// upload cover image
	singleImageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed", "error": err.Error()})
		return
	}
	req.ImageURL = singleImageURL

	if err := h.classService.CreateClass(req); err != nil {
		utils.CleanupImageOnError(singleImageURL)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create class", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateClassRequest
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

	class, err := h.classService.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Class not found"})
		return
	}

	classResponse := dto.ClassDetailResponse{
		ID:          class.ID,
		Title:       class.Title,
		Image:       class.Image,
		IsActive:    class.IsActive,
		Duration:    class.Duration,
		Description: class.Description,
		Additional:  class.Additional,
		Type:        class.Type,
		Level:       class.Level,
		Location:    class.Location,
		Category:    class.Category,
		Subcategory: class.Subcategory,
		CreatedAt:   class.CreatedAt,
	}

	c.JSON(http.StatusOK, classResponse)
}

func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	var query dto.ClassQueryParam
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters", "error": err.Error()})
		return
	}

	classes, total, err := h.classService.GetAllClasses(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch classes", "error": err.Error()})
		return
	}

	if classes == nil {
		classes = []dto.ClassResponse{}
	}

	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
		"total":   total,
		"page":    query.Page,
		"limit":   query.Limit,
	})
}

func (h *ClassHandler) GetActiveClasses(c *gin.Context) {
	classes, err := h.classService.GetActiveClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch active classes", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (h *ClassHandler) UploadClassGallery(c *gin.Context) {
	classID := c.Param("id")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid multipart form"})
		return
	}

	files := form.File["gallery[]"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no gallery images provided"})
		return
	}

	var galleries []models.ClassGallery
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

		galleries = append(galleries, models.ClassGallery{
			ID:      uuid.New(),
			ClassID: uuid.MustParse(classID),
			URL:     imageURL,
		})
	}

	if err := h.classService.AddClassGallery(galleries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save gallery", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "gallery uploaded successfully"})
}

func (h *ClassHandler) DeleteClassGallery(c *gin.Context) {
	galleryID := c.Param("galleryId")

	if err := h.classService.DeleteClassGallery(galleryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete gallery", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "gallery deleted successfully"})
}
