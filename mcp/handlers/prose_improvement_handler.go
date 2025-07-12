package handlers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/google/uuid"
)

// ProseImprovementHandler handles prose improvement operations
type ProseImprovementHandler struct {
	storage storage.Storage
}

// NewProseImprovementHandler creates a new prose improvement handler
func NewProseImprovementHandler(storage storage.Storage) *ProseImprovementHandler {
	return &ProseImprovementHandler{
		storage: storage,
	}
}

// GetProsePrompts retrieves prose improvement prompts
func (h *ProseImprovementHandler) GetProsePrompts(params map[string]interface{}) (interface{}, error) {
	category, _ := params["category"].(string)
	activeOnly, _ := params["activeOnly"].(bool)

	filename := "prose_improvement_prompts.json"
	data, err := h.storage.ReadFile(filename)
	if err != nil {
		// Return default prompts if file doesn't exist
		return h.getDefaultProsePrompts(category), nil
	}

	var prompts []models.ProseImprovementPrompt
	if err := json.Unmarshal(data, &prompts); err != nil {
		return nil, fmt.Errorf("failed to parse prose prompts: %v", err)
	}

	// Filter by category if specified
	if category != "" {
		var filtered []models.ProseImprovementPrompt
		for _, prompt := range prompts {
			if prompt.Category == category {
				filtered = append(filtered, prompt)
			}
		}
		prompts = filtered
	}

	// TODO: Implement activeOnly filtering if needed
	_ = activeOnly

	return prompts, nil
}

// GetProsePromptByID retrieves a specific prose improvement prompt
func (h *ProseImprovementHandler) GetProsePromptByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id parameter is required")
	}

	filename := "prose_improvement_prompts.json"
	data, err := h.storage.ReadFile(filename)
	if err != nil {
		// Check default prompts
		defaults := h.getDefaultProsePrompts("")
		for _, prompt := range defaults {
			if prompt.ID == id {
				return prompt, nil
			}
		}
		return nil, fmt.Errorf("prose prompt not found: %s", id)
	}

	var prompts []models.ProseImprovementPrompt
	if err := json.Unmarshal(data, &prompts); err != nil {
		return nil, fmt.Errorf("failed to parse prose prompts: %v", err)
	}

	for _, prompt := range prompts {
		if prompt.ID == id {
			return prompt, nil
		}
	}

	return nil, fmt.Errorf("prose prompt not found: %s", id)
}

// CreateProsePrompt creates a new prose improvement prompt
func (h *ProseImprovementHandler) CreateProsePrompt(params map[string]interface{}) (interface{}, error) {
	label, ok := params["label"].(string)
	if !ok {
		return nil, fmt.Errorf("label parameter is required")
	}

	defaultPromptText, ok := params["defaultPromptText"].(string)
	if !ok {
		return nil, fmt.Errorf("defaultPromptText parameter is required")
	}

	category, _ := params["category"].(string)
	if category == "" {
		category = "custom"
	}

	description, _ := params["description"].(string)

	prompt := models.ProseImprovementPrompt{
		ID:                generateID(),
		Label:             label,
		Category:          category,
		Description:       description,
		DefaultPromptText: defaultPromptText,
		Variants:          []models.ProsePromptVariant{},
		Order:             0, // Will be set when loading all prompts
	}

	// Load existing prompts
	filename := "prose_improvement_prompts.json"
	var prompts []models.ProseImprovementPrompt
	
	data, err := h.storage.ReadFile(filename)
	if err == nil {
		json.Unmarshal(data, &prompts)
	}

	// Set order
	prompt.Order = len(prompts)
	prompts = append(prompts, prompt)

	// Save back to file
	data, err = json.Marshal(prompts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal prose prompts: %v", err)
	}

	if err := h.storage.WriteFile(filename, data); err != nil {
		return nil, fmt.Errorf("failed to save prose prompt: %v", err)
	}

	return prompt, nil
}

// AnalyzeProse applies a specific prose improvement prompt to text
func (h *ProseImprovementHandler) AnalyzeProse(params map[string]interface{}) (interface{}, error) {
	text, ok := params["text"].(string)
	if !ok {
		return nil, fmt.Errorf("text parameter is required")
	}

	promptID, ok := params["promptId"].(string)
	if !ok {
		return nil, fmt.Errorf("promptId parameter is required")
	}

	provider, _ := params["provider"].(string)
	if provider == "" {
		provider = "manual"
	}

	// Get the prompt
	promptResult, err := h.GetProsePromptByID(map[string]interface{}{"id": promptID})
	if err != nil {
		return nil, err
	}

	prompt, ok := promptResult.(models.ProseImprovementPrompt)
	if !ok {
		return nil, fmt.Errorf("invalid prompt type")
	}

	// For now, return a structure that the frontend can use
	// The actual LLM integration would be done on the frontend side
	result := models.ProseAnalysisResult{
		OriginalText: text,
		Changes:      []models.ProseChange{}, // Would be populated by LLM analysis
		PromptUsed:   prompt.DefaultPromptText,
		Provider:     provider,
		ProcessedAt:  time.Now(),
	}

	return result, nil
}

// GetProsePromptByCategory retrieves prompts filtered by category
func (h *ProseImprovementHandler) GetProsePromptByCategory(params map[string]interface{}) (interface{}, error) {
	category, ok := params["category"].(string)
	if !ok {
		return nil, fmt.Errorf("category parameter is required")
	}

	return h.GetProsePrompts(map[string]interface{}{
		"category": category,
	})
}

// CreateProseSession creates a new prose improvement session
func (h *ProseImprovementHandler) CreateProseSession(params map[string]interface{}) (interface{}, error) {
	text, ok := params["text"].(string)
	if !ok {
		return nil, fmt.Errorf("text parameter is required")
	}

	// Get available prompts
	promptsResult, err := h.GetProsePrompts(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	prompts, ok := promptsResult.([]models.ProseImprovementPrompt)
	if !ok {
		return nil, fmt.Errorf("failed to get prose prompts")
	}

	session := models.NewProseSession(text, prompts)

	// Save session to temporary storage
	filename := fmt.Sprintf("prose_session_%s.json", session.ID)
	data, err := json.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session: %v", err)
	}

	if err := h.storage.WriteFile(filename, data); err != nil {
		return nil, fmt.Errorf("failed to save session: %v", err)
	}

	return session, nil
}

// GetProseSession retrieves a prose improvement session
func (h *ProseImprovementHandler) GetProseSession(params map[string]interface{}) (interface{}, error) {
	sessionID, ok := params["sessionId"].(string)
	if !ok {
		return nil, fmt.Errorf("sessionId parameter is required")
	}

	filename := fmt.Sprintf("prose_session_%s.json", sessionID)
	data, err := h.storage.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("session not found: %s", sessionID)
	}

	var session models.ProseImprovementSession
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to parse session: %v", err)
	}

	return session, nil
}

// UpdateProseSession updates a prose improvement session
func (h *ProseImprovementHandler) UpdateProseSession(params map[string]interface{}) (interface{}, error) {
	sessionID, ok := params["sessionId"].(string)
	if !ok {
		return nil, fmt.Errorf("sessionId parameter is required")
	}

	// Get existing session
	sessionResult, err := h.GetProseSession(map[string]interface{}{"sessionId": sessionID})
	if err != nil {
		return nil, err
	}

	session, ok := sessionResult.(models.ProseImprovementSession)
	if !ok {
		return nil, fmt.Errorf("invalid session type")
	}

	// Update fields if provided
	if currentText, ok := params["currentText"].(string); ok {
		session.CurrentText = currentText
	}

	if currentPromptIndex, ok := params["currentPromptIndex"].(float64); ok {
		session.CurrentPromptIndex = int(currentPromptIndex)
	}

	// Update changes if provided
	if changesData, ok := params["changes"].([]interface{}); ok {
		var changes []models.ProseChange
		for _, changeData := range changesData {
			changeBytes, _ := json.Marshal(changeData)
			var change models.ProseChange
			if json.Unmarshal(changeBytes, &change) == nil {
				changes = append(changes, change)
			}
		}
		session.Changes = changes
	}

	session.UpdatedAt = time.Now()

	// Save updated session
	filename := fmt.Sprintf("prose_session_%s.json", sessionID)
	data, err := json.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session: %v", err)
	}

	if err := h.storage.WriteFile(filename, data); err != nil {
		return nil, fmt.Errorf("failed to save session: %v", err)
	}

	return session, nil
}

// getDefaultProsePrompts returns default prose improvement prompts
func (h *ProseImprovementHandler) getDefaultProsePrompts(category string) []models.ProseImprovementPrompt {
	defaults := []models.ProseImprovementPrompt{
		{
			ID:                "enhance-imagery",
			Label:             "Enhance Imagery",
			Category:          "tropes",
			Order:             0,
			Description:       "Enhance sensory details and vivid imagery",
			DefaultPromptText: "Analyze the following text and enhance its imagery by adding more vivid sensory details, metaphors, and descriptive language. Return your suggestions as a JSON array of changes.",
			Variants:          []models.ProsePromptVariant{},
		},
		{
			ID:                "strengthen-verbs",
			Label:             "Strengthen Verbs",
			Category:          "style",
			Order:             1,
			Description:       "Replace weak verbs with stronger alternatives",
			DefaultPromptText: "Identify weak verbs in the following text and suggest stronger, more specific alternatives. Return your suggestions as a JSON array of changes.",
			Variants:          []models.ProsePromptVariant{},
		},
		{
			ID:                "check-grammar",
			Label:             "Check Grammar",
			Category:          "grammar",
			Order:             2,
			Description:       "Grammar and punctuation fixes",
			DefaultPromptText: "Check the following text for grammar, punctuation, and syntax errors. Provide corrections as a JSON array of changes.",
			Variants:          []models.ProsePromptVariant{},
		},
		{
			ID:                "improve-dialogue",
			Label:             "Improve Dialogue",
			Category:          "style",
			Order:             3,
			Description:       "Enhance dialogue authenticity and flow",
			DefaultPromptText: "Analyze the dialogue in the following text and suggest improvements for authenticity, character voice, and natural flow. Return your suggestions as a JSON array of changes.",
			Variants:          []models.ProsePromptVariant{},
		},
	}

	if category != "" {
		var filtered []models.ProseImprovementPrompt
		for _, prompt := range defaults {
			if prompt.Category == category {
				filtered = append(filtered, prompt)
			}
		}
		return filtered
	}

	return defaults
}

// SearchProseContent searches across prose improvement content
func (h *ProseImprovementHandler) SearchProseContent(params map[string]interface{}) (interface{}, error) {
	query, ok := params["query"].(string)
	if !ok {
		return nil, fmt.Errorf("query parameter is required")
	}

	var results []models.SearchResult

	// Search prose prompts
	prosePromptsResult, err := h.GetProsePrompts(map[string]interface{}{})
	if err == nil {
		if prompts, ok := prosePromptsResult.([]models.ProseImprovementPrompt); ok {
			for _, prompt := range prompts {
				if h.matchesQuery(strings.ToLower(query), strings.ToLower(prompt.Label), strings.ToLower(prompt.Description), strings.ToLower(prompt.DefaultPromptText)) {
					results = append(results, models.SearchResult{
						ID:       prompt.ID,
						Type:     "prose_prompt",
						Title:    prompt.Label,
						Content:  prompt.DefaultPromptText,
						Snippet:  h.createSnippet(prompt.DefaultPromptText, query),
						Score:    h.calculateScore(query, prompt.Label, prompt.Description),
						Metadata: map[string]interface{}{
							"category": prompt.Category,
							"order":    prompt.Order,
						},
						MatchedAt: time.Now(),
					})
				}
			}
		}
	}

	// Search prose sessions
	sessionFiles, _ := h.storage.ListFiles("prose_session_*.json")
	for _, sessionFile := range sessionFiles {
		data, err := h.storage.ReadFile(sessionFile)
		if err != nil {
			continue
		}

		var session models.ProseImprovementSession
		if json.Unmarshal(data, &session) == nil {
			if h.matchesQuery(strings.ToLower(query), strings.ToLower(session.OriginalText), strings.ToLower(session.CurrentText)) {
				results = append(results, models.SearchResult{
					ID:       session.ID,
					Type:     "prose_session",
					Title:    fmt.Sprintf("Prose Session - %s", session.CreatedAt.Format("2006-01-02")),
					Content:  session.OriginalText,
					Snippet:  h.createSnippet(session.OriginalText, query),
					Score:    h.calculateScore(query, session.OriginalText),
					Metadata: map[string]interface{}{
						"changesCount":      len(session.Changes),
						"currentPromptIndex": session.CurrentPromptIndex,
						"createdAt":         session.CreatedAt,
					},
					MatchedAt: time.Now(),
				})
			}
		}
	}

	return results, nil
}

// Helper methods

func (h *ProseImprovementHandler) matchesQuery(query string, texts ...string) bool {
	for _, text := range texts {
		if strings.Contains(text, query) {
			return true
		}
	}
	return false
}

func (h *ProseImprovementHandler) createSnippet(text, query string) string {
	text = strings.ToLower(text)
	query = strings.ToLower(query)
	
	index := strings.Index(text, query)
	if index == -1 {
		// Return first 100 characters if no match
		if len(text) > 100 {
			return text[:100] + "..."
		}
		return text
	}

	start := index - 50
	if start < 0 {
		start = 0
	}

	end := index + len(query) + 50
	if end > len(text) {
		end = len(text)
	}

	snippet := text[start:end]
	if start > 0 {
		snippet = "..." + snippet
	}
	if end < len(text) {
		snippet = snippet + "..."
	}

	return snippet
}

func (h *ProseImprovementHandler) calculateScore(query string, texts ...string) float64 {
	score := 0.0
	query = strings.ToLower(query)

	for _, text := range texts {
		text = strings.ToLower(text)
		
		// Exact match gets highest score
		if strings.Contains(text, query) {
			score += 1.0
		}
		
		// Word matches get partial score
		queryWords := strings.Fields(query)
		textWords := strings.Fields(text)
		matches := 0
		
		for _, qWord := range queryWords {
			for _, tWord := range textWords {
				if strings.Contains(tWord, qWord) {
					matches++
					break
				}
			}
		}
		
		if len(queryWords) > 0 {
			score += float64(matches) / float64(len(queryWords)) * 0.5
		}
	}

	return score
}

// generateID creates a new UUID
func generateID() string {
	return uuid.New().String()
}
