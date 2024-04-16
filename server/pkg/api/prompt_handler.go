package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GeneratePrompt(c *gin.Context) {
	var input struct {
		TraitKeys      []string `json:"trait_keys"`
		ChapterLength  int      `json:"chapter_length"`
		ResponseFormat string   `json:"response_format"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logging.Logger.Errorf("Failed to bind input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	prompt := generatePrompt(input.TraitKeys, input.ChapterLength, input.ResponseFormat, h)

	response := struct {
		Prompt string `json:"prompt"`
	}{
		Prompt: prompt,
	}
	c.JSON(http.StatusOK, response)
}

func generatePrompt(traitKeys []string, chapterLength int, responseFormat string, h *Handler) string {
	var prompt strings.Builder

	prompt.WriteString("Generate a chapter with the following traits:\n")
	for _, traitKey := range traitKeys {
		trait := getTraitByKey(traitKey, h)
		prompt.WriteString("- ")
		prompt.WriteString(trait.TriggerText)
		prompt.WriteString("\n")
	}

	prompt.WriteString(fmt.Sprintf("\nChapter Length: %d words\n", chapterLength))
	prompt.WriteString(fmt.Sprintf("Response Format: %s\n", responseFormat))

	return prompt.String()
}

func getTraitByKey(traitKey string, h *Handler) models.TraitKey {
	var trait models.TraitKey
	if err := h.DB.Where("trait_key = ?", traitKey).First(&trait).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch trait: %v", err)
		return models.TraitKey{}
	}
	return trait
}
