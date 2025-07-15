package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/danielsobrado/ainovelprompter/server/pkg/config"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ReadSettingsFile reads the settings from the settings.json file
func (a *App) ReadSettingsFile() (string, error) {
	settingsPath := filepath.Join(a.getAppDataDir(), "settings.json")
	content, err := os.ReadFile(settingsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "{}", nil
		}
		return "", fmt.Errorf("error reading settings file: %v", err)
	}
	return string(content), nil
}

// WriteSettingsFile writes the settings to the settings.json file
func (a *App) WriteSettingsFile(content string) error {
	settingsPath := filepath.Join(a.getAppDataDir(), "settings.json")
	err := os.MkdirAll(filepath.Dir(settingsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating settings directory: %v", err)
	}
	err = os.WriteFile(settingsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing settings file: %v", err)
	}
	return nil
}

// getAppDataDir returns the path to the application data directory
func (a *App) getAppDataDir() string {
	// Debug logging to file
	debugLog := fmt.Sprintf("DEBUG: getAppDataDir called, a.dataDir='%s'\n", a.dataDir)
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}

	// Use the configured data directory if available
	if a.dataDir != "" {
		debugLog2 := fmt.Sprintf("DEBUG: Using configured data directory: '%s'\n", a.dataDir)
		if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
			file.WriteString(debugLog2)
			file.Close()
		}
		return a.dataDir
	}

	// Fallback to default directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("Error getting user home directory: %v", err))
		return ""
	}
	defaultDir := filepath.Join(homeDir, ".ai-novel-prompter")
	debugLog3 := fmt.Sprintf("DEBUG: Using fallback default directory: '%s'\n", defaultDir)
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog3)
		file.Close()
	}
	return defaultDir
}

func (a *App) ReadTaskTypesFile() (string, error) {
	return a.readMCPVersionedEntities("task-types", mcpTaskTypeToServer)
}

func (a *App) WriteTaskTypesFile(content string) error {
	return a.writeMCPVersionedEntities("task-types", content, serverTaskTypeToMCP)
}

func (a *App) ReadRulessFile() (string, error) {
	return a.readMCPVersionedEntities("rules", mcpRuleToServer)
}

func (a *App) WriteRulessFile(content string) error {
	return a.writeMCPVersionedEntities("rules", content, serverRuleToMCP)
}

// =============================================================================
// MCP VERSIONED STORAGE - ALL ENTITY TYPES
// =============================================================================

// Universal helper function for reading MCP versioned entities
func (a *App) readMCPVersionedEntities(entityType string, convertToServer func(interface{}) interface{}) (string, error) {
	debugLog := fmt.Sprintf("DEBUG: Read%sFile called (MCP versioned storage)\n", entityType)
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}

	entitiesDir := filepath.Join(a.getAppDataDir(), entityType)

	// Check if entities directory exists
	if _, err := os.Stat(entitiesDir); os.IsNotExist(err) {
		return "[]", nil
	}

	serverEntities := make([]interface{}, 0)

	// Read all entity directories
	entries, err := os.ReadDir(entitiesDir)
	if err != nil {
		return "[]", nil
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		entityDir := filepath.Join(entitiesDir, entry.Name())

		// Find the latest version file
		latestFile, err := a.findLatestVersionFile(entityDir)
		if err != nil {
			continue
		}

		// Read and convert the entity
		content, err := os.ReadFile(latestFile)
		if err != nil {
			continue
		}

		var mcpEntity map[string]interface{}
		if err := json.Unmarshal(content, &mcpEntity); err != nil {
			continue
		}

		// Convert MCP format to Server format
		serverEntity := convertToServer(mcpEntity)
		if serverEntity != nil {
			serverEntities = append(serverEntities, serverEntity)
		}
	}

	// Marshal to JSON
	if serverEntities == nil || len(serverEntities) == 0 {
		return "[]", nil
	}

	result, err := json.MarshalIndent(serverEntities, "", "  ")
	if err != nil {
		return "[]", err
	}

	return string(result), nil
}

// Universal helper function for writing MCP versioned entities
func (a *App) writeMCPVersionedEntities(entityType string, content string, convertToMCP func(interface{}) map[string]interface{}) error {
	debugLog := fmt.Sprintf("DEBUG: Write%sFile called with content length: %d (MCP versioned storage)\n", entityType, len(content))
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}

	// Parse server format entities
	serverEntities := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(content), &serverEntities); err != nil {
		return fmt.Errorf("error parsing %s JSON: %v", entityType, err)
	}

	entitiesDir := filepath.Join(a.getAppDataDir(), entityType)

	// Ensure entities directory exists
	if err := os.MkdirAll(entitiesDir, 0755); err != nil {
		return fmt.Errorf("error creating %s directory: %v", entityType, err)
	}

	// Write each entity in MCP versioned format
	for _, serverEntity := range serverEntities {
		mcpEntity := convertToMCP(serverEntity)
		if mcpEntity == nil {
			continue
		}

		// Set timestamps
		now := time.Now().Format(time.RFC3339)
		mcpEntity["createdAt"] = now
		mcpEntity["updatedAt"] = now

		// Get entity ID
		entityID, exists := mcpEntity["id"].(string)
		if !exists || entityID == "" {
			runtime.LogWarning(a.ctx, fmt.Sprintf("%s missing ID, skipping", entityType))
			continue
		}

		// Create entity subdirectory
		entityDir := filepath.Join(entitiesDir, entityID)
		if err := os.MkdirAll(entityDir, 0755); err != nil {
			continue
		}

		// Write entity file with timestamp
		timestamp := time.Now().Format("2006-01-02T15-04-05.000Z07-00")
		filename := fmt.Sprintf("%s.json", timestamp)
		filePath := filepath.Join(entityDir, filename)

		entityData, err := json.MarshalIndent(mcpEntity, "", "  ")
		if err != nil {
			continue
		}

		if err := os.WriteFile(filePath, entityData, 0644); err != nil {
			continue
		}
	}

	return nil
}

// Utility functions
func getString(m map[string]interface{}, key string) string {
	if val, exists := m[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getStringSlice(m map[string]interface{}, key string) []string {
	if val, exists := m[key]; exists {
		if slice, ok := val.([]interface{}); ok {
			result := make([]string, 0)
			for _, item := range slice {
				if str, ok := item.(string); ok {
					result = append(result, str)
				}
			}
			return result
		}
	}
	return []string{}
}

// =============================================================================
// DATA STRUCTURE DEFINITIONS
// =============================================================================

// Characters - MCP Versioned Storage
// ServerCharacter represents the format expected by frontend (simple)
type ServerCharacter struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// MCPCharacter represents the enhanced MCP format with versioning
type MCPCharacter struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Traits      map[string]interface{} `json:"traits,omitempty"`
	Notes       string                 `json:"notes,omitempty"`
	CreatedAt   string                 `json:"createdAt,omitempty"`
	UpdatedAt   string                 `json:"updatedAt,omitempty"`
}

// Task Types
type ServerTaskType struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Template    string `json:"template"`
}

type MCPTaskType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Template    string `json:"template"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

// Rules
type ServerRule struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type MCPRule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

// Locations
type ServerLocation struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

type MCPLocation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Details     string `json:"details"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

// Codex
type ServerCodex struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Category string   `json:"category"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

type MCPCodex struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Category  string   `json:"category"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

// Sample Chapters
type ServerSampleChapter struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Purpose string   `json:"purpose"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Source  string   `json:"source"`
	Tags    []string `json:"tags"`
}

type MCPSampleChapter struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Purpose   string   `json:"purpose"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	Source    string   `json:"source"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

// Prose Prompts
type ServerProsePrompt struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Category    string `json:"category"`
	Description string `json:"description"`
	PromptText  string `json:"defaultPromptText"`
}

type MCPProsePrompt struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	PromptText  string `json:"defaultPromptText"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

// =============================================================================
// CONVERSION FUNCTIONS
// =============================================================================

// Character conversions
func mcpCharacterToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerCharacter{
		ID:          getString(mcpMap, "id"),
		Label:       getString(mcpMap, "name"),
		Description: getString(mcpMap, "description"),
	}
}

func serverCharacterToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":          getString(serverMap, "id"),
		"name":        getString(serverMap, "label"),
		"description": getString(serverMap, "description"),
		"traits":      make(map[string]interface{}),
		"notes":       "",
	}
}

// Task Type conversions
func mcpTaskTypeToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerTaskType{
		ID:          getString(mcpMap, "id"),
		Label:       getString(mcpMap, "name"),
		Category:    getString(mcpMap, "category"),
		Description: getString(mcpMap, "description"),
		Template:    getString(mcpMap, "template"),
	}
}

func serverTaskTypeToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":          getString(serverMap, "id"),
		"name":        getString(serverMap, "label"),
		"category":    getString(serverMap, "category"),
		"description": getString(serverMap, "description"),
		"template":    getString(serverMap, "template"),
	}
}

// Rule conversions
func mcpRuleToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerRule{
		ID:          getString(mcpMap, "id"),
		Label:       getString(mcpMap, "name"),
		Category:    getString(mcpMap, "category"),
		Description: getString(mcpMap, "description"),
		Content:     getString(mcpMap, "content"),
	}
}

func serverRuleToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":          getString(serverMap, "id"),
		"name":        getString(serverMap, "label"),
		"category":    getString(serverMap, "category"),
		"description": getString(serverMap, "description"),
		"content":     getString(serverMap, "content"),
	}
}

// Location conversions
func mcpLocationToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerLocation{
		ID:          getString(mcpMap, "id"),
		Label:       getString(mcpMap, "name"),
		Description: getString(mcpMap, "description"),
		Details:     getString(mcpMap, "details"),
	}
}

func serverLocationToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":          getString(serverMap, "id"),
		"name":        getString(serverMap, "label"),
		"description": getString(serverMap, "description"),
		"details":     getString(serverMap, "details"),
	}
}

// Codex conversions
func mcpCodexToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerCodex{
		ID:       getString(mcpMap, "id"),
		Title:    getString(mcpMap, "title"),
		Category: getString(mcpMap, "category"),
		Content:  getString(mcpMap, "content"),
		Tags:     getStringSlice(mcpMap, "tags"),
	}
}

func serverCodexToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":       getString(serverMap, "id"),
		"title":    getString(serverMap, "title"),
		"category": getString(serverMap, "category"),
		"content":  getString(serverMap, "content"),
		"tags":     getStringSlice(serverMap, "tags"),
	}
}

// Sample Chapter conversions
func mcpSampleChapterToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerSampleChapter{
		ID:      getString(mcpMap, "id"),
		Title:   getString(mcpMap, "title"),
		Purpose: getString(mcpMap, "purpose"),
		Content: getString(mcpMap, "content"),
		Author:  getString(mcpMap, "author"),
		Source:  getString(mcpMap, "source"),
		Tags:    getStringSlice(mcpMap, "tags"),
	}
}

func serverSampleChapterToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":      getString(serverMap, "id"),
		"title":   getString(serverMap, "title"),
		"purpose": getString(serverMap, "purpose"),
		"content": getString(serverMap, "content"),
		"author":  getString(serverMap, "author"),
		"source":  getString(serverMap, "source"),
		"tags":    getStringSlice(serverMap, "tags"),
	}
}

// Prose Prompt conversions
func mcpProsePromptToServer(mcpData interface{}) interface{} {
	mcpMap, ok := mcpData.(map[string]interface{})
	if !ok {
		return nil
	}

	return ServerProsePrompt{
		ID:          getString(mcpMap, "id"),
		Label:       getString(mcpMap, "name"),
		Category:    getString(mcpMap, "category"),
		Description: getString(mcpMap, "description"),
		PromptText:  getString(mcpMap, "defaultPromptText"),
	}
}

func serverProsePromptToMCP(serverData interface{}) map[string]interface{} {
	serverMap, ok := serverData.(map[string]interface{})
	if !ok {
		return nil
	}

	return map[string]interface{}{
		"id":                getString(serverMap, "id"),
		"name":              getString(serverMap, "label"),
		"category":          getString(serverMap, "category"),
		"description":       getString(serverMap, "description"),
		"defaultPromptText": getString(serverMap, "defaultPromptText"),
	}
}

func (a *App) ReadCharactersFile() (string, error) {
	return a.readMCPVersionedEntities("characters", mcpCharacterToServer)
}

func (a *App) WriteCharactersFile(content string) error {
	return a.writeMCPVersionedEntities("characters", content, serverCharacterToMCP)
}

// findLatestVersionFile finds the most recent version file in an MCP entity directory
func (a *App) findLatestVersionFile(entityDir string) (string, error) {
	entries, err := os.ReadDir(entityDir)
	if err != nil {
		return "", err
	}

	var latestFile string
	var latestTime time.Time

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(entityDir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.ModTime().After(latestTime) {
			latestTime = info.ModTime()
			latestFile = filePath
		}
	}

	if latestFile == "" {
		return "", fmt.Errorf("no version files found")
	}

	return latestFile, nil
}

// Locations
func (a *App) ReadLocationsFile() (string, error) {
	return a.readMCPVersionedEntities("locations", mcpLocationToServer)
}

func (a *App) WriteLocationsFile(content string) error {
	return a.writeMCPVersionedEntities("locations", content, serverLocationToMCP)
}

// Codex
func (a *App) ReadCodexFile() (string, error) {
	return a.readMCPVersionedEntities("codex", mcpCodexToServer)
}

func (a *App) WriteCodexFile(content string) error {
	return a.writeMCPVersionedEntities("codex", content, serverCodexToMCP)
}

// Sample Chapters
func (a *App) ReadSampleChaptersFile() (string, error) {
	return a.readMCPVersionedEntities("sample-chapters", mcpSampleChapterToServer)
}

func (a *App) WriteSampleChaptersFile(content string) error {
	return a.writeMCPVersionedEntities("sample-chapters", content, serverSampleChapterToMCP)
}

// Standardized ReadRulesFile (same as ReadRulessFile)
func (a *App) ReadRulesFile() (string, error) {
	return a.readMCPVersionedEntities("rules", mcpRuleToServer)
}

func (a *App) WriteRulesFile(content string) error {
	return a.writeMCPVersionedEntities("rules", content, serverRuleToMCP)
}

// Prose Improvement Prompts
func (a *App) ReadProsePromptsFile() (string, error) {
	return a.readMCPVersionedEntities("prose-prompts", mcpProsePromptToServer)
}

func (a *App) WriteProsePromptsFile(content string) error {
	return a.writeMCPVersionedEntities("prose-prompts", content, serverProsePromptToMCP)
}

// GetInitialLLMSettings retrieves LLM provider settings from config (env/.env/yaml)
func (a *App) GetInitialLLMSettings() (map[string]string, error) {
	settings := make(map[string]string)
	// Viper automatically handles precedence: actual env var > .env file > config.yaml > defaults set in LoadConfig
	settings["openrouter_api_key"] = config.GetString("openrouter.api_key")
	settings["openrouter_default_model"] = config.GetString("openrouter.default_model")
	settings["lmstudio_api_url"] = config.GetString("lmstudio.api_url")
	settings["lmstudio_default_model"] = config.GetString("lmstudio.default_model")

	// Log what's being loaded by Go to help debug
	runtime.LogInfof(a.ctx, "Go backend GetInitialLLMSettings: OR Key Set: %t, OR Model: %s, LMStudio URL: %s, LMStudio Model: %s",
		settings["openrouter_api_key"] != "",
		settings["openrouter_default_model"],
		settings["lmstudio_api_url"],
		settings["lmstudio_default_model"],
	)

	return settings, nil
}

// LLM Provider Settings File I/O
func (a *App) ReadLLMSettingsFile() (string, error) {
	settingsPath := filepath.Join(a.getAppDataDir(), "llm_provider_settings.json")
	content, err := os.ReadFile(settingsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "{}", nil // Return empty object string if not found
		}
		return "", fmt.Errorf("error reading llm_provider_settings.json: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteLLMSettingsFile(content string) error {
	settingsPath := filepath.Join(a.getAppDataDir(), "llm_provider_settings.json")
	err := os.MkdirAll(filepath.Dir(settingsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating directory for llm_provider_settings.json: %v", err)
	}
	err = os.WriteFile(settingsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing llm_provider_settings.json: %v", err)
	}
	return nil
}

// --- New Prompt System Structs ---
type PromptVariant struct {
	VariantLabel        string   `json:"variantLabel,omitempty"`
	TargetModelFamilies []string `json:"targetModelFamilies,omitempty"`
	TargetModels        []string `json:"targetModels,omitempty"`
	PromptText          string   `json:"promptText"`
}

type PromptDefinition struct {
	ID                string          `json:"id"`
	Label             string          `json:"label"`
	Category          string          `json:"category"`
	Order             int             `json:"order"`
	Description       string          `json:"description,omitempty"`
	DefaultPromptText string          `json:"defaultPromptText"`
	Variants          []PromptVariant `json:"variants,omitempty"`
}

type LLMProviderConfig struct {
	APIKey string `json:"apiKey,omitempty"`
	APIURL string `json:"apiUrl,omitempty"`
	Model  string `json:"model,omitempty"`
}

type LLMProviderFrontend struct {
	Type   string            `json:"type"` // "manual", "lmstudio", "openrouter"
	Config LLMProviderConfig `json:"config,omitempty"`
}

// --- Prompt Management Functions ---
func (a *App) ReadPromptsFile() (string, error) {
	promptsPath := filepath.Join(a.getAppDataDir(), "prompts.json")
	content, err := os.ReadFile(promptsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading prompts file: %v", err)
	}
	return string(content), nil
}

func (a *App) WritePromptsFile(content string) error {
	promptsPath := filepath.Join(a.getAppDataDir(), "prompts.json")
	err := os.MkdirAll(filepath.Dir(promptsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating prompts directory: %v", err)
	}
	err = os.WriteFile(promptsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing prompts file: %v", err)
	}
	return nil
}

// LoadPromptDefinitions loads prompt definitions from the prompts.json file
func (a *App) LoadPromptDefinitions() ([]PromptDefinition, error) {
	content, err := a.ReadPromptsFile()
	if err != nil {
		return nil, err
	}

	var prompts []PromptDefinition
	err = json.Unmarshal([]byte(content), &prompts)
	if err != nil {
		return nil, fmt.Errorf("error parsing prompts.json: %v", err)
	}

	return prompts, nil
}

// SavePromptDefinitions saves prompt definitions to the prompts.json file
func (a *App) SavePromptDefinitions(prompts []PromptDefinition) error {
	content, err := json.MarshalIndent(prompts, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling prompts to JSON: %v", err)
	}

	return a.WritePromptsFile(string(content))
}

// FindPromptDefinitionByID finds a prompt definition by its ID
func (a *App) FindPromptDefinitionByID(id string) (*PromptDefinition, error) {
	prompts, err := a.LoadPromptDefinitions()
	if err != nil {
		return nil, err
	}

	for _, prompt := range prompts {
		if prompt.ID == id {
			return &prompt, nil
		}
	}

	return nil, fmt.Errorf("prompt definition with ID %s not found", id)
}

// UpdatePromptDefinition updates an existing prompt definition
func (a *App) UpdatePromptDefinition(updatedPrompt PromptDefinition) error {
	prompts, err := a.LoadPromptDefinitions()
	if err != nil {
		return err
	}

	for i, prompt := range prompts {
		if prompt.ID == updatedPrompt.ID {
			// Update the prompt definition
			prompts[i] = updatedPrompt

			// Save the updated prompts back to the file
			return a.SavePromptDefinitions(prompts)
		}
	}

	return fmt.Errorf("prompt definition with ID %s not found", updatedPrompt.ID)
}

// DeletePromptDefinition deletes a prompt definition by its ID
func (a *App) DeletePromptDefinition(id string) error {
	prompts, err := a.LoadPromptDefinitions()
	if err != nil {
		return err
	}

	var updatedPrompts []PromptDefinition
	for _, prompt := range prompts {
		if prompt.ID != id {
			updatedPrompts = append(updatedPrompts, prompt)
		}
	}

	return a.SavePromptDefinitions(updatedPrompts)
}

// GetPromptVariantsForModel retrieves prompt variants for a specific model
func (a *App) GetPromptVariantsForModel(modelName string) ([]PromptVariant, error) {
	prompts, err := a.LoadPromptDefinitions()
	if err != nil {
		return nil, err
	}

	var variants []PromptVariant
	for _, prompt := range prompts {
		for _, variant := range prompt.Variants {
			// Check if the variant is applicable to the given model
			if contains(variant.TargetModels, modelName) || contains(variant.TargetModelFamilies, getModelFamily(modelName)) {
				variants = append(variants, variant)
			}
		}
	}

	return variants, nil
}

// contains checks if a string slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// getModelFamily derives the model family from the model name
func getModelFamily(modelName string) string {
	// Example: "gpt-3.5-turbo" -> "gpt-3.5"
	parts := strings.Split(modelName, "-")
	if len(parts) >= 3 {
		return fmt.Sprintf("%s-%s", parts[0], parts[1])
	}
	return modelName
}

// --- Context Management ---
func (a *App) getContext() context.Context {
	return a.ctx
}

func (a *App) setContext(ctx context.Context) {
	a.ctx = ctx
}

// --- Prose Prompt Definition File I/O ---
func (a *App) ReadPromptDefinitionsFile() (string, error) {
	filePath := filepath.Join(a.getAppDataDir(), "prose_prompt_definitions.json")
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			runtime.LogInfo(a.ctx, "Go: prose_prompt_definitions.json not found, returning empty array for initialization.")
			// Return an empty array string so frontend can populate with defaults if needed
			return "[]", nil
		}
		runtime.LogErrorf(a.ctx, "Go: Error reading prose_prompt_definitions.json: %v", err)
		return "", fmt.Errorf("error reading prose_prompt_definitions.json: %v", err)
	}
	runtime.LogInfof(a.ctx, "Go: ReadPromptDefinitionsFile successful.")
	return string(content), nil
}

func (a *App) WritePromptDefinitionsFile(jsonData string) error {
	filePath := filepath.Join(a.getAppDataDir(), "prose_prompt_definitions.json")
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		runtime.LogErrorf(a.ctx, "Go: Error creating directory for prose_prompt_definitions.json: %v", err)
		return fmt.Errorf("error creating directory: %v", err)
	}
	err = os.WriteFile(filePath, []byte(jsonData), 0644)
	if err != nil {
		runtime.LogErrorf(a.ctx, "Go: Error writing prose_prompt_definitions.json: %v", err)
		return fmt.Errorf("error writing file: %v", err)
	}
	runtime.LogInfo(a.ctx, "Go: WritePromptDefinitionsFile successful.")
	return nil
}

// GetResolvedProsePrompt finds the best prompt text for a given task and provider
func (a *App) GetResolvedProsePrompt(taskID string, providerJSON string) (string, error) {
	runtime.LogInfof(a.ctx, "Go: GetResolvedProsePrompt called for taskID: %s", taskID)

	// Parse the provider data from frontend
	var provider LLMProviderFrontend
	if err := json.Unmarshal([]byte(providerJSON), &provider); err != nil {
		runtime.LogErrorf(a.ctx, "Go: Error unmarshalling provider data: %v", err)
		return "", fmt.Errorf("could not parse provider data: %w", err)
	}

	runtime.LogInfof(a.ctx, "Go: Provider type: %s, model: %s", provider.Type, provider.Config.Model)

	jsonData, err := a.ReadPromptDefinitionsFile()
	if err != nil {
		return "", fmt.Errorf("could not read prompt definitions: %w", err)
	}

	var definitions []PromptDefinition
	if err := json.Unmarshal([]byte(jsonData), &definitions); err != nil {
		runtime.LogErrorf(a.ctx, "Go: Error unmarshalling prompt definitions: %v. JSON data: %s", err, jsonData)
		return "", fmt.Errorf("could not parse prompt definitions: %w", err)
	}

	for _, def := range definitions {
		if def.ID == taskID {
			// Found the task definition, now check variants
			selectedModel := provider.Config.Model
			selectedFamily := provider.Type

			// Priority 1: Match specific model ID
			for _, variant := range def.Variants {
				for _, targetModel := range variant.TargetModels {
					if targetModel == selectedModel {
						runtime.LogInfof(a.ctx, "Go: Matched variant by specific model ID '%s' for task '%s'. VariantLabel: '%s'", selectedModel, taskID, variant.VariantLabel)
						return variant.PromptText, nil
					}
				}
			}

			// Priority 2: Match model family
			for _, variant := range def.Variants {
				for _, targetFamily := range variant.TargetModelFamilies {
					// Normalize provider.Type to common families if needed
					// For OpenRouter, try to infer family from model string (e.g., "openai/gpt-4o" -> "openai")
					normalizedProviderFamily := selectedFamily
					if selectedFamily == "openrouter" && selectedModel != "" {
						parts := strings.Split(selectedModel, "/")
						if len(parts) > 0 {
							normalizedProviderFamily = parts[0]
						}
					}

					if targetFamily == normalizedProviderFamily || targetFamily == selectedFamily {
						runtime.LogInfof(a.ctx, "Go: Matched variant by model family '%s' for task '%s'. VariantLabel: '%s'", targetFamily, taskID, variant.VariantLabel)
						return variant.PromptText, nil
					}
				}
			}

			// Priority 3: Use default prompt text for the definition
			runtime.LogInfof(a.ctx, "Go: No specific variant matched for task '%s'. Using default prompt.", taskID)
			return def.DefaultPromptText, nil
		}
	}

	runtime.LogWarning(a.ctx, fmt.Sprintf("Go: No prompt definition found for taskID: %s", taskID))
	return "", fmt.Errorf("prompt definition not found for taskID: %s", taskID)
}
