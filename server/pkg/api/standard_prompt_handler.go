package api

import (
	"net/http"

	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStandardPrompt(c *gin.Context) {
	var prompt models.StandardPrompt
	if err := c.ShouldBindJSON(&prompt); err != nil {
		logging.Logger.Errorf("Failed to bind standard prompt: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var latestPrompt models.StandardPrompt
	if err := h.DB.Where("standard_name = ?", prompt.StandardName).Order("version DESC").First(&latestPrompt).Error; err == nil {
		prompt.Version = latestPrompt.Version + 1
	}

	if err := h.DB.Create(&prompt).Error; err != nil {
		logging.Logger.Errorf("Failed to create standard prompt: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create standard prompt"})
		return
	}

	c.JSON(http.StatusCreated, prompt)
}

func (h *Handler) UpdateStandardPrompt(c *gin.Context) {
	id := c.Param("id")

	var prompt models.StandardPrompt
	if err := c.ShouldBindJSON(&prompt); err != nil {
		logging.Logger.Errorf("Failed to bind standard prompt: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var existingPrompt models.StandardPrompt
	if err := h.DB.Where("id = ?", id).First(&existingPrompt).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch standard prompt: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Standard prompt not found"})
		return
	}

	existingPrompt.Title = prompt.Title
	existingPrompt.Prompt = prompt.Prompt
	existingPrompt.Version = existingPrompt.Version + 1

	if err := h.DB.Save(&existingPrompt).Error; err != nil {
		logging.Logger.Errorf("Failed to update standard prompt: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update standard prompt"})
		return
	}

	c.JSON(http.StatusOK, existingPrompt)
}

func (h *Handler) ListStandardPrompts(c *gin.Context) {
	var prompts []models.StandardPrompt
	if err := h.DB.Find(&prompts).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch standard prompts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch standard prompts"})
		return
	}

	c.JSON(http.StatusOK, prompts)
}
