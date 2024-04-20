package api

import (
	"net/http"

	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListTraitTypes(c *gin.Context) {
	var traitTypes []models.TraitType
	if err := h.DB.Find(&traitTypes).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch trait types: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trait types"})
		return
	}

	c.JSON(http.StatusOK, traitTypes)
}
