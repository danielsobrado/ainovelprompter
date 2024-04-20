package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/danielsobrado/ainovelprompter/pkg/config"
	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/gin-gonic/gin"
)

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type GenerateResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

type ChatRequest struct {
	Model    string           `json:"model"`
	Messages []models.Message `json:"messages"`
	Stream   bool             `json:"stream"`
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

	// Set stream to false if not provided in the request
	if !req.Stream {
		req.Stream = false
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

	logging.Logger.Infof("Ollama API response status: %s", resp.Status)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Logger.Errorf("Failed to read Ollama response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Ollama response"})
		return
	}
	logging.Logger.Infof("Ollama API response body: %s", string(respBody))

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	var generateResp GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&generateResp); err != nil {
		logging.Logger.Errorf("Failed to decode Ollama response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode Ollama response"})
		return
	}
	logging.Logger.Infof("Decoded Ollama response: %+v", generateResp)

	c.JSON(http.StatusOK, generateResp)
}

func (h *Handler) ChatWithModel(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.Errorf("Failed to bind chat request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Set stream to false if not provided in the request
	if !req.Stream {
		req.Stream = false
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

	logging.Logger.Infof("Ollama API response status: %s", resp.Status)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Logger.Errorf("Failed to read Ollama response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Ollama response"})
		return
	}
	logging.Logger.Infof("Ollama API response body: %s", string(respBody))

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		logging.Logger.Errorf("Failed to decode Ollama chat response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode chat response"})
		return
	}
	logging.Logger.Infof("Decoded Ollama chat response: %+v", chatResp)

	c.JSON(http.StatusOK, chatResp)
}
