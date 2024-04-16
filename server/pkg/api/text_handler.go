package api

import (
	"net/http"
	"strconv"

	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListTexts(c *gin.Context) {
	search := c.Query("search")

	var texts []models.Text
	query := h.DB.Order("created_at DESC")
	if search != "" {
		query = query.Where("content ILIKE ?", "%"+search+"%")
	}

	if err := query.Find(&texts).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch texts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch texts"})
		return
	}

	c.JSON(http.StatusOK, texts)
}

func (h *Handler) CreateText(c *gin.Context) {
	var text models.Text
	if err := c.ShouldBindJSON(&text); err != nil {
		logging.Logger.Errorf("Failed to bind text: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Create(&text).Error; err != nil {
		logging.Logger.Errorf("Failed to create text: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create text"})
		return
	}

	c.JSON(http.StatusCreated, text)
}

func (h *Handler) GetText(c *gin.Context) {
	textID, err := strconv.ParseUint(c.Param("textId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid text ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text ID"})
		return
	}

	var text models.Text
	if err := h.DB.First(&text, textID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch text: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Text not found"})
		return
	}

	c.JSON(http.StatusOK, text)
}

func (h *Handler) UpdateText(c *gin.Context) {
	textID, err := strconv.ParseUint(c.Param("textId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid text ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text ID"})
		return
	}

	var text models.Text
	if err := h.DB.First(&text, textID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch text: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Text not found"})
		return
	}

	if err := c.ShouldBindJSON(&text); err != nil {
		logging.Logger.Errorf("Failed to bind text: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Save(&text).Error; err != nil {
		logging.Logger.Errorf("Failed to update text: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update text"})
		return
	}

	c.JSON(http.StatusOK, text)
}

func (h *Handler) DeleteText(c *gin.Context) {
	textID, err := strconv.ParseUint(c.Param("textId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid text ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text ID"})
		return
	}

	var text models.Text
	if err := h.DB.First(&text, textID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch text: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Text not found"})
		return
	}

	if err := h.DB.Delete(&text).Error; err != nil {
		logging.Logger.Errorf("Failed to delete text: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete text"})
		return
	}

	c.Status(http.StatusNoContent)
}
