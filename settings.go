package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	taskTypesPath := filepath.Join(a.getAppDataDir(), "task_types.json")
	content, err := os.ReadFile(taskTypesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading task types file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteTaskTypesFile(content string) error {
	taskTypesPath := filepath.Join(a.getAppDataDir(), "task_types.json")
	err := os.MkdirAll(filepath.Dir(taskTypesPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating task types directory: %v", err)
	}
	err = os.WriteFile(taskTypesPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing task types file: %v", err)
	}
	return nil
}

func (a *App) ReadRulessFile() (string, error) {
	RulessPath := filepath.Join(a.getAppDataDir(), "custom_instructions.json")
	content, err := os.ReadFile(RulessPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading custom instructions file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteRulessFile(content string) error {
	RulessPath := filepath.Join(a.getAppDataDir(), "custom_instructions.json")
	err := os.MkdirAll(filepath.Dir(RulessPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating custom instructions directory: %v", err)
	}
	err = os.WriteFile(RulessPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing custom instructions file: %v", err)
	}
	return nil
}

// Characters
func (a *App) ReadCharactersFile() (string, error) {
	// Debug logging for character file reading
	debugLog := fmt.Sprintf("DEBUG: ReadCharactersFile called\n")
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}
	
	charactersPath := filepath.Join(a.getAppDataDir(), "characters.json")
	content, err := os.ReadFile(charactersPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading characters file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteCharactersFile(content string) error {
	// Debug logging for character file writing
	debugLog := fmt.Sprintf("DEBUG: WriteCharactersFile called with content length: %d\n", len(content))
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}
	
	charactersPath := filepath.Join(a.getAppDataDir(), "characters.json")
	err := os.MkdirAll(filepath.Dir(charactersPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating characters directory: %v", err)
	}
	err = os.WriteFile(charactersPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing characters file: %v", err)
	}
	return nil
}

// Locations
func (a *App) ReadLocationsFile() (string, error) {
	locationsPath := filepath.Join(a.getAppDataDir(), "locations.json")
	content, err := os.ReadFile(locationsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading locations file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteLocationsFile(content string) error {
	locationsPath := filepath.Join(a.getAppDataDir(), "locations.json")
	err := os.MkdirAll(filepath.Dir(locationsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating locations directory: %v", err)
	}
	err = os.WriteFile(locationsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing locations file: %v", err)
	}
	return nil
}

// Codex
func (a *App) ReadCodexFile() (string, error) {
	codexPath := filepath.Join(a.getAppDataDir(), "codex.json")
	content, err := os.ReadFile(codexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading codex file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteCodexFile(content string) error {
	codexPath := filepath.Join(a.getAppDataDir(), "codex.json")
	err := os.MkdirAll(filepath.Dir(codexPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating codex directory: %v", err)
	}
	err = os.WriteFile(codexPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing codex file: %v", err)
	}
	return nil
}

// Sample Chapters
func (a *App) ReadSampleChaptersFile() (string, error) {
	chaptersPath := filepath.Join(a.getAppDataDir(), "sample_chapters.json")
	content, err := os.ReadFile(chaptersPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading sample chapters file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteSampleChaptersFile(content string) error {
	chaptersPath := filepath.Join(a.getAppDataDir(), "sample_chapters.json")
	err := os.MkdirAll(filepath.Dir(chaptersPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating sample chapters directory: %v", err)
	}
	err = os.WriteFile(chaptersPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing sample chapters file: %v", err)
	}
	return nil
}

// Fix the typo in ReadRulessFile
func (a *App) ReadRulesFile() (string, error) {
	rulesPath := filepath.Join(a.getAppDataDir(), "rules.json")
	content, err := os.ReadFile(rulesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading rules file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteRulesFile(content string) error {
	rulesPath := filepath.Join(a.getAppDataDir(), "rules.json")
	err := os.MkdirAll(filepath.Dir(rulesPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating rules directory: %v", err)
	}
	err = os.WriteFile(rulesPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing rules file: %v", err)
	}
	return nil
}

// Prose Improvement Prompts
func (a *App) ReadProsePromptsFile() (string, error) {
	prosePromptsPath := filepath.Join(a.getAppDataDir(), "prose_prompts.json")
	content, err := os.ReadFile(prosePromptsPath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return an empty JSON array string
			// The frontend will handle populating with defaults if necessary
			return "[]", nil
		}
		return "", fmt.Errorf("error reading prose prompts file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteProsePromptsFile(content string) error {
	prosePromptsPath := filepath.Join(a.getAppDataDir(), "prose_prompts.json")
	err := os.MkdirAll(filepath.Dir(prosePromptsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating prose prompts directory: %v", err)
	}
	err = os.WriteFile(prosePromptsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing prose prompts file: %v", err)
	}
	return nil
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
