package handlers

import (
	"fmt"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/danielsobrado/ainovelprompter/mcp/validation"
)

type StoryContextHandler struct {
	storage storage.StoryContextStorage
}

func NewStoryContextHandler(storage storage.StoryContextStorage) *StoryContextHandler {
	return &StoryContextHandler{
		storage: storage,
	}
}

// Character handlers
func (h *StoryContextHandler) GetCharacters(params map[string]interface{}) (interface{}, error) {
	// Check for search parameter
	if searchQuery, ok := params["search"].(string); ok && searchQuery != "" {
		return h.storage.SearchCharacters(searchQuery)
	}

	// Check for specific IDs
	if ids, ok := params["ids"].([]string); ok {
		var characters []models.Character
		for _, id := range ids {
			char, err := h.storage.GetCharacterByID(id)
			if err == nil {
				characters = append(characters, *char)
			}
		}
		return characters, nil
	}

	// Return all characters
	return h.storage.GetCharacters()
}

func (h *StoryContextHandler) GetCharacterByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("character ID is required")
	}

	return h.storage.GetCharacterByID(id)
}

func (h *StoryContextHandler) CreateCharacter(params map[string]interface{}) (interface{}, error) {
	var character models.Character
	validator := validation.NewEntityValidator()

	// Validate and sanitize name
	if name, ok := params["name"].(string); ok {
		sanitizedName, err := validator.ValidateAndSanitize(name, "name", "Character", true)
		if err != nil {
			return nil, err
		}
		character.Name = sanitizedName
	} else {
		return nil, fmt.Errorf("character name is required")
	}

	// Validate and sanitize description
	if desc, ok := params["description"].(string); ok {
		sanitizedDesc, err := validator.ValidateAndSanitize(desc, "description", "Character", true)
		if err != nil {
			return nil, err
		}
		character.Description = sanitizedDesc
	} else {
		return nil, fmt.Errorf("character description is required")
	}

	if traits, ok := params["traits"].(map[string]interface{}); ok {
		character.Traits = make(map[string]interface{})
		for k, v := range traits {
			character.Traits[k] = v
		}
	}

	if notes, ok := params["notes"].(string); ok {
		character.Notes = notes
	}

	err := h.storage.CreateCharacter(&character)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func (h *StoryContextHandler) UpdateCharacter(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("character ID is required")
	}

	character, err := h.storage.GetCharacterByID(id)
	if err != nil {
		return nil, err
	}

	validator := validation.NewEntityValidator()

	if name, ok := params["name"].(string); ok {
		sanitizedName, err := validator.ValidateAndSanitize(name, "name", "Character", true)
		if err != nil {
			return nil, err
		}
		character.Name = sanitizedName
	}

	if desc, ok := params["description"].(string); ok {
		sanitizedDesc, err := validator.ValidateAndSanitize(desc, "description", "Character", true)
		if err != nil {
			return nil, err
		}
		character.Description = sanitizedDesc
	}

	if traits, ok := params["traits"].(map[string]interface{}); ok {
		character.Traits = make(map[string]interface{})
		for k, v := range traits {
			character.Traits[k] = v
		}
	}

	if notes, ok := params["notes"].(string); ok {
		character.Notes = notes
	}

	err = h.storage.UpdateCharacter(character)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func (h *StoryContextHandler) DeleteCharacter(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("character ID is required")
	}

	err := h.storage.DeleteCharacter(id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"status": "deleted", "id": id}, nil
}

// Location handlers
func (h *StoryContextHandler) GetLocations(params map[string]interface{}) (interface{}, error) {
	if searchQuery, ok := params["search"].(string); ok && searchQuery != "" {
		return h.storage.SearchLocations(searchQuery)
	}

	return h.storage.GetLocations()
}

func (h *StoryContextHandler) GetLocationByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("location ID is required")
	}

	return h.storage.GetLocationByID(id)
}

func (h *StoryContextHandler) CreateLocation(params map[string]interface{}) (interface{}, error) {
	var location models.Location
	validator := validation.NewEntityValidator()

	// Validate and sanitize name
	if name, ok := params["name"].(string); ok {
		sanitizedName, err := validator.ValidateAndSanitize(name, "name", "Location", true)
		if err != nil {
			return nil, err
		}
		location.Name = sanitizedName
	} else {
		return nil, fmt.Errorf("location name is required")
	}

	// Validate and sanitize description
	if desc, ok := params["description"].(string); ok {
		sanitizedDesc, err := validator.ValidateAndSanitize(desc, "description", "Location", true)
		if err != nil {
			return nil, err
		}
		location.Description = sanitizedDesc
	} else {
		return nil, fmt.Errorf("location description is required")
	}

	if details, ok := params["details"].(string); ok {
		location.Details = details
	}

	if notes, ok := params["notes"].(string); ok {
		location.Notes = notes
	}

	err := h.storage.CreateLocation(&location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// Codex handlers
func (h *StoryContextHandler) GetCodexEntries(params map[string]interface{}) (interface{}, error) {
	if searchQuery, ok := params["search"].(string); ok && searchQuery != "" {
		return h.storage.SearchCodex(searchQuery)
	}

	if category, ok := params["category"].(string); ok && category != "" {
		entries, err := h.storage.GetCodexEntries()
		if err != nil {
			return nil, err
		}

		var filtered []models.CodexEntry
		for _, entry := range entries {
			if entry.Category == category {
				filtered = append(filtered, entry)
			}
		}
		return filtered, nil
	}

	return h.storage.GetCodexEntries()
}

func (h *StoryContextHandler) GetCodexEntryByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("codex entry ID is required")
	}

	return h.storage.GetCodexEntryByID(id)
}

func (h *StoryContextHandler) CreateCodexEntry(params map[string]interface{}) (interface{}, error) {
	var entry models.CodexEntry
	validator := validation.NewEntityValidator()

	// Validate and sanitize title
	if title, ok := params["title"].(string); ok {
		sanitizedTitle, err := validator.ValidateAndSanitize(title, "title", "Codex entry", true)
		if err != nil {
			return nil, err
		}
		entry.Title = sanitizedTitle
	} else {
		return nil, fmt.Errorf("codex entry title is required")
	}

	// Validate and sanitize content
	if content, ok := params["content"].(string); ok {
		sanitizedContent, err := validator.ValidateAndSanitize(content, "content", "Codex entry", true)
		if err != nil {
			return nil, err
		}
		entry.Content = sanitizedContent
	} else {
		return nil, fmt.Errorf("codex entry content is required")
	}

	if category, ok := params["category"].(string); ok {
		entry.Category = category
	}

	if tags, ok := params["tags"].([]interface{}); ok {
		for _, tag := range tags {
			if tagStr, ok := tag.(string); ok {
				entry.Tags = append(entry.Tags, tagStr)
			}
		}
	}

	err := h.storage.CreateCodexEntry(&entry)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

// Rules handlers
func (h *StoryContextHandler) GetRules(params map[string]interface{}) (interface{}, error) {
	if activeOnly, ok := params["activeOnly"].(bool); ok && activeOnly {
		return h.storage.GetActiveRules()
	}

	if category, ok := params["category"].(string); ok && category != "" {
		rules, err := h.storage.GetRules()
		if err != nil {
			return nil, err
		}

		var filtered []models.Rule
		for _, rule := range rules {
			if rule.Category == category {
				filtered = append(filtered, rule)
			}
		}
		return filtered, nil
	}

	return h.storage.GetRules()
}

func (h *StoryContextHandler) GetRuleByID(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok {
		return nil, fmt.Errorf("rule ID is required")
	}

	return h.storage.GetRuleByID(id)
}

// Context building handler
func (h *StoryContextHandler) BuildWritingContext(params map[string]interface{}) (interface{}, error) {
	context := make(map[string]interface{})

	// Get relevant characters
	if charIDs, ok := params["characterIds"].([]string); ok {
		var characters []models.Character
		for _, id := range charIDs {
			if char, err := h.storage.GetCharacterByID(id); err == nil {
				characters = append(characters, *char)
			}
		}
		context["characters"] = characters
	}

	// Get relevant locations
	if locIDs, ok := params["locationIds"].([]string); ok {
		var locations []models.Location
		for _, id := range locIDs {
			if loc, err := h.storage.GetLocationByID(id); err == nil {
				locations = append(locations, *loc)
			}
		}
		context["locations"] = locations
	}

	// Get active rules
	if includeRules, ok := params["includeRules"].(bool); ok && includeRules {
		rules, _ := h.storage.GetActiveRules()
		context["rules"] = rules
	}

	// Get relevant codex entries
	if codexCategories, ok := params["codexCategories"].([]string); ok {
		entries, _ := h.storage.GetCodexEntries()
		var relevantEntries []models.CodexEntry
		for _, entry := range entries {
			for _, cat := range codexCategories {
				if entry.Category == cat {
					relevantEntries = append(relevantEntries, entry)
					break
				}
			}
		}
		context["codex"] = relevantEntries
	}

	return context, nil
}
