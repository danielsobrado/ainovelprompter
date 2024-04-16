package api

import (
	"net/http"
	"strconv"

	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListFeedback(c *gin.Context) {
	search := c.Query("search")

	var feedbacks []models.Feedback
	query := h.DB.Order("created_at DESC")
	if search != "" {
		query = query.Where("content ILIKE ?", "%"+search+"%")
	}

	if err := query.Find(&feedbacks).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch feedbacks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feedbacks"})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}

func (h *Handler) CreateFeedback(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		logging.Logger.Errorf("Failed to bind feedback: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Create(&feedback).Error; err != nil {
		logging.Logger.Errorf("Failed to create feedback: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feedback"})
		return
	}

	c.JSON(http.StatusCreated, feedback)
}

func (h *Handler) GetFeedback(c *gin.Context) {
	feedbackID, err := strconv.ParseUint(c.Param("feedbackId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid feedback ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feedback ID"})
		return
	}

	var feedback models.Feedback
	if err := h.DB.First(&feedback, feedbackID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch feedback: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

func (h *Handler) UpdateFeedback(c *gin.Context) {
	feedbackID, err := strconv.ParseUint(c.Param("feedbackId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid feedback ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feedback ID"})
		return
	}

	var feedback models.Feedback
	if err := h.DB.First(&feedback, feedbackID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch feedback: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	if err := c.ShouldBindJSON(&feedback); err != nil {
		logging.Logger.Errorf("Failed to bind feedback: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Save(&feedback).Error; err != nil {
		logging.Logger.Errorf("Failed to update feedback: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feedback"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

func (h *Handler) DeleteFeedback(c *gin.Context) {
	feedbackID, err := strconv.ParseUint(c.Param("feedbackId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid feedback ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feedback ID"})
		return
	}

	var feedback models.Feedback
	if err := h.DB.First(&feedback, feedbackID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch feedback: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	if err := h.DB.Delete(&feedback).Error; err != nil {
		logging.Logger.Errorf("Failed to delete feedback: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feedback"})
		return
	}

	c.Status(http.StatusNoContent)
}
