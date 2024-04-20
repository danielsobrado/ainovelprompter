package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/danielsobrado/ainovelprompter/pkg/config"
	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type GenerateResponse struct {
	Answer string `json:"answer"`
}

type ChatRequest struct {
	Model    string           `json:"model"`
	Messages []models.Message `json:"messages"`
}

type ChatResponse struct {
	Messages []models.Message `json:"messages"`
}

func (h *Handler) GenerateResponse(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.Errorf("Failed to bind generate request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		logging.Logger.Errorf("Could not marshal generate request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing request"})
		return
	}

	ollamaGenerateURL := config.GetString("ollama.generate_url")
	resp, err := http.Post(ollamaGenerateURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		logging.Logger.Errorf("Failed to send generate request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Ollama"})
		return
	}
	defer resp.Body.Close()

	var generateResp GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&generateResp); err != nil {
		logging.Logger.Errorf("Failed to decode Ollama response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode Ollama response"})
		return
	}

	logging.Logger.Infof("Decode: %v", json.NewDecoder(resp.Body).Decode(&generateResp))

	c.JSON(http.StatusOK, generateResp)
}

func (h *Handler) ChatWithModel(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.Errorf("Failed to bind chat request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		logging.Logger.Errorf("Could not marshal chat request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing request"})
		return
	}

	ollamaChatURL := config.GetString("ollama_chat_url")
	resp, err := http.Post(ollamaChatURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		logging.Logger.Errorf("Failed to send chat request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Ollama"})
		return
	}
	defer resp.Body.Close()

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		logging.Logger.Errorf("Failed to decode Ollama chat response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode chat response"})
		return
	}
}
