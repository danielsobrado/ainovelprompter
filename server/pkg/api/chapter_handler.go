package api

import (
	"net/http"
	"strconv"

	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListChapters(c *gin.Context) {
	search := c.Query("search")

	var chapters []models.Chapter
	query := h.DB.Order("created_at DESC")
	if search != "" {
		query = query.Where("chapter_title ILIKE ?", "%"+search+"%")
	}

	if err := query.Find(&chapters).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch chapters: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chapters"})
		return
	}

	c.JSON(http.StatusOK, chapters)
}

func (h *Handler) CreateChapter(c *gin.Context) {
	var chapter models.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		logging.Logger.Errorf("Failed to bind chapter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Create(&chapter).Error; err != nil {
		logging.Logger.Errorf("Failed to create chapter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chapter"})
		return
	}

	c.JSON(http.StatusCreated, chapter)
}

func (h *Handler) GetChapter(c *gin.Context) {
	chapterID, err := strconv.ParseUint(c.Param("chapterId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid chapter ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	var chapter models.Chapter
	if err := h.DB.First(&chapter, chapterID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch chapter: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

func (h *Handler) UpdateChapter(c *gin.Context) {
	chapterID, err := strconv.ParseUint(c.Param("chapterId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid chapter ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	var chapter models.Chapter
	if err := h.DB.First(&chapter, chapterID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch chapter: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	if err := c.ShouldBindJSON(&chapter); err != nil {
		logging.Logger.Errorf("Failed to bind chapter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Save(&chapter).Error; err != nil {
		logging.Logger.Errorf("Failed to update chapter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update chapter"})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

func (h *Handler) DeleteChapter(c *gin.Context) {
	chapterID, err := strconv.ParseUint(c.Param("chapterId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid chapter ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	var chapter models.Chapter
	if err := h.DB.First(&chapter, chapterID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch chapter: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}

	if err := h.DB.Delete(&chapter).Error; err != nil {
		logging.Logger.Errorf("Failed to delete chapter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete chapter"})
		return
	}

	c.Status(http.StatusNoContent)
}
