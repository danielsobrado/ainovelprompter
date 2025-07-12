package handlers

import (
	"fmt"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// MockStorage implements Storage interface for testing
type MockStorage struct {
	characters   []models.Character
	locations    []models.Location
	codexEntries []models.CodexEntry
	rules        []models.Rule
	chapters     []models.Chapter
	files        map[string][]byte
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		characters:   make([]models.Character, 0),
		locations:    make([]models.Location, 0),
		codexEntries: make([]models.CodexEntry, 0),
		rules:        make([]models.Rule, 0),
		chapters:     make([]models.Chapter, 0),
		files:        make(map[string][]byte),
	}
}

// Storage interface implementation
func (m *MockStorage) ReadFile(filename string) ([]byte, error) {
	if data, exists := m.files[filename]; exists {
		return data, nil
	}
	return []byte("[]"), nil
}

func (m *MockStorage) WriteFile(filename string, data []byte) error {
	m.files[filename] = data
	return nil
}

func (m *MockStorage) DeleteFile(filename string) error {
	delete(m.files, filename)
	return nil
}

func (m *MockStorage) ListFiles(pattern string) ([]string, error) {
	var files []string
	for filename := range m.files {
		// Simple pattern matching for testing
		if pattern == "*" || pattern == "" {
			files = append(files, filename)
		} else if len(filename) >= len(pattern) {
			// Simple contains check
			for i := 0; i <= len(filename)-len(pattern); i++ {
				if filename[i:i+len(pattern)] == pattern {
					files = append(files, filename)
					break
				}
			}
		}
	}
	return files, nil
}

// StoryContextStorage interface implementation (simplified for testing)
func (m *MockStorage) GetCharacters() ([]models.Character, error) {
	return m.characters, nil
}

func (m *MockStorage) GetCharacterByID(id string) (*models.Character, error) {
	for _, char := range m.characters {
		if char.ID == id {
			return &char, nil
		}
	}
	return nil, fmt.Errorf("character not found")
}

func (m *MockStorage) CreateCharacter(character *models.Character) error {
	m.characters = append(m.characters, *character)
	return nil
}

func (m *MockStorage) UpdateCharacter(character *models.Character) error {
	for i, char := range m.characters {
		if char.ID == character.ID {
			m.characters[i] = *character
			return nil
		}
	}
	return fmt.Errorf("character not found")
}

func (m *MockStorage) DeleteCharacter(id string) error {
	for i, char := range m.characters {
		if char.ID == id {
			m.characters = append(m.characters[:i], m.characters[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("character not found")
}

func (m *MockStorage) SearchCharacters(query string) ([]models.Character, error) {
	var results []models.Character
	for _, char := range m.characters {
		if contains(char.Name, query) || contains(char.Description, query) {
			results = append(results, char)
		}
	}
	return results, nil
}

// Implement other required interfaces (simplified)
func (m *MockStorage) GetLocations() ([]models.Location, error)                           { return m.locations, nil }
func (m *MockStorage) GetLocationByID(id string) (*models.Location, error)               { return nil, fmt.Errorf("not found") }
func (m *MockStorage) CreateLocation(location *models.Location) error                    { return nil }
func (m *MockStorage) UpdateLocation(location *models.Location) error                    { return nil }
func (m *MockStorage) DeleteLocation(id string) error                                    { return nil }
func (m *MockStorage) SearchLocations(query string) ([]models.Location, error)          { return nil, nil }
func (m *MockStorage) GetCodexEntries() ([]models.CodexEntry, error)                     { return m.codexEntries, nil }
func (m *MockStorage) GetCodexEntryByID(id string) (*models.CodexEntry, error)           { return nil, fmt.Errorf("not found") }
func (m *MockStorage) CreateCodexEntry(entry *models.CodexEntry) error                   { return nil }
func (m *MockStorage) UpdateCodexEntry(entry *models.CodexEntry) error                   { return nil }
func (m *MockStorage) DeleteCodexEntry(id string) error                                  { return nil }
func (m *MockStorage) SearchCodex(query string) ([]models.CodexEntry, error)            { return nil, nil }
func (m *MockStorage) GetRules() ([]models.Rule, error)                                  { return m.rules, nil }
func (m *MockStorage) GetRuleByID(id string) (*models.Rule, error)                       { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetActiveRules() ([]models.Rule, error)                            { return nil, nil }
func (m *MockStorage) GetChapters() ([]models.Chapter, error)                            { return m.chapters, nil }
func (m *MockStorage) GetChapterByID(id string) (*models.Chapter, error)                 { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetChapterByNumber(number int) (*models.Chapter, error)            { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetPreviousChapter(currentNumber int) (*models.Chapter, error)     { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetChapterRange(start, end int) ([]models.Chapter, error)          { return nil, nil }
func (m *MockStorage) CreateChapter(chapter *models.Chapter) error                       { return nil }
func (m *MockStorage) UpdateChapter(chapter *models.Chapter) error                       { return nil }
func (m *MockStorage) DeleteChapter(id string) error                                     { return nil }
func (m *MockStorage) SearchChapters(query string) ([]models.Chapter, error)            { return nil, nil }
func (m *MockStorage) GetStoryBeats(chapterNumber int) (*models.StoryBeats, error)       { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetAllStoryBeats() ([]models.StoryBeats, error)                    { return nil, nil }
func (m *MockStorage) SaveStoryBeats(beats *models.StoryBeats) error                     { return nil }
func (m *MockStorage) GetFutureNotes() ([]models.FutureNotes, error)                     { return nil, nil }
func (m *MockStorage) GetFutureNotesByChapter(chapterNum int) ([]models.FutureNotes, error) { return nil, nil }
func (m *MockStorage) CreateFutureNote(note *models.FutureNotes) error                   { return nil }
func (m *MockStorage) UpdateFutureNote(note *models.FutureNotes) error                   { return nil }
func (m *MockStorage) DeleteFutureNote(id string) error                                  { return nil }
func (m *MockStorage) GetSampleChapters() ([]models.SampleChapter, error)                { return nil, nil }
func (m *MockStorage) GetSampleChapterByID(id string) (*models.SampleChapter, error)     { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetSampleChaptersByPurpose(purpose string) ([]models.SampleChapter, error) { return nil, nil }
func (m *MockStorage) CreateSampleChapter(sample *models.SampleChapter) error            { return nil }
func (m *MockStorage) GetTaskTypes() ([]models.TaskType, error)                          { return nil, nil }
func (m *MockStorage) GetTaskTypeByID(id string) (*models.TaskType, error)               { return nil, fmt.Errorf("not found") }
func (m *MockStorage) GetActiveTaskTypes() ([]models.TaskType, error)                    { return nil, nil }

// TestStoryContextHandler_Characters tests character operations
func TestStoryContextHandler_Characters(t *testing.T) {
	mockStorage := NewMockStorage()
	handler := NewStoryContextHandler(mockStorage)

	// Test GetCharacters (empty initially)
	result, err := handler.GetCharacters(map[string]interface{}{})
	if err != nil {
		t.Fatalf("GetCharacters failed: %v", err)
	}

	characters, ok := result.([]models.Character)
	if !ok {
		t.Fatalf("GetCharacters returned wrong type: %T", result)
	}

	if len(characters) != 0 {
		t.Errorf("Expected 0 characters initially, got %d", len(characters))
	}

	// Test CreateCharacter
	createParams := map[string]interface{}{
		"name":        "Test Character",
		"description": "A test character for unit testing",
		"notes":       "Created in handler test",
		"traits": map[string]interface{}{
			"brave":  true,
			"height": "tall",
		},
	}

	createResult, err := handler.CreateCharacter(createParams)
	if err != nil {
		t.Fatalf("CreateCharacter failed: %v", err)
	}

	character, ok := createResult.(models.Character)
	if !ok {
		t.Fatalf("CreateCharacter returned wrong type: %T", createResult)
	}

	if character.ID == "" {
		t.Error("Created character should have an ID")
	}

	if character.Name != "Test Character" {
		t.Errorf("Expected name 'Test Character', got '%s'", character.Name)
	}

	if character.CreatedAt.IsZero() {
		t.Error("Created character should have CreatedAt timestamp")
	}

	// Test GetCharacterByID
	getResult, err := handler.GetCharacterByID(map[string]interface{}{
		"id": character.ID,
	})
	if err != nil {
		t.Fatalf("GetCharacterByID failed: %v", err)
	}

	retrievedChar, ok := getResult.(models.Character)
	if !ok {
		t.Fatalf("GetCharacterByID returned wrong type: %T", getResult)
	}

	if retrievedChar.ID != character.ID {
		t.Errorf("Retrieved character ID mismatch: expected %s, got %s", 
			character.ID, retrievedChar.ID)
	}

	// Test UpdateCharacter
	updateParams := map[string]interface{}{
		"id":          character.ID,
		"name":        "Updated Test Character",
		"description": "An updated test character",
	}

	_, err = handler.UpdateCharacter(updateParams)
	if err != nil {
		t.Fatalf("UpdateCharacter failed: %v", err)
	}

	// Test DeleteCharacter
	_, err = handler.DeleteCharacter(map[string]interface{}{
		"id": character.ID,
	})
	if err != nil {
		t.Fatalf("DeleteCharacter failed: %v", err)
	}

	// Verify deletion
	_, err = handler.GetCharacterByID(map[string]interface{}{
		"id": character.ID,
	})
	if err == nil {
		t.Error("Expected error when getting deleted character")
	}
}

// TestStoryContextHandler_ErrorHandling tests error conditions
func TestStoryContextHandler_ErrorHandling(t *testing.T) {
	mockStorage := NewMockStorage()
	handler := NewStoryContextHandler(mockStorage)

	// Test CreateCharacter with missing name
	_, err := handler.CreateCharacter(map[string]interface{}{
		"description": "A character without a name",
	})
	if err == nil {
		t.Error("Expected error when creating character without name")
	}

	// Test GetCharacterByID with missing ID
	_, err = handler.GetCharacterByID(map[string]interface{}{})
	if err == nil {
		t.Error("Expected error when getting character without ID")
	}

	// Test GetCharacterByID with non-string ID
	_, err = handler.GetCharacterByID(map[string]interface{}{
		"id": 123,
	})
	if err == nil {
		t.Error("Expected error when getting character with non-string ID")
	}

	// Test UpdateCharacter with missing ID
	_, err = handler.UpdateCharacter(map[string]interface{}{
		"name": "Test",
	})
	if err == nil {
		t.Error("Expected error when updating character without ID")
	}

	// Test DeleteCharacter with missing ID
	_, err = handler.DeleteCharacter(map[string]interface{}{})
	if err == nil {
		t.Error("Expected error when deleting character without ID")
	}
}

// TestStoryContextHandler_BuildWritingContext tests context building
func TestStoryContextHandler_BuildWritingContext(t *testing.T) {
	mockStorage := NewMockStorage()
	handler := NewStoryContextHandler(mockStorage)

	// Add some test data
	testChar := models.Character{
		ID:          "char1",
		Name:        "Hero",
		Description: "The brave protagonist",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockStorage.CreateCharacter(&testChar)

	// Test BuildWritingContext
	contextParams := map[string]interface{}{
		"characterIds": []interface{}{"char1"},
		"includeRules": true,
	}

	result, err := handler.BuildWritingContext(contextParams)
	if err != nil {
		t.Fatalf("BuildWritingContext failed: %v", err)
	}

	context, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("BuildWritingContext returned wrong type: %T", result)
	}

	// Verify context structure
	if _, exists := context["characters"]; !exists {
		t.Error("Context should include characters section")
	}

	if _, exists := context["rules"]; !exists {
		t.Error("Context should include rules section when includeRules is true")
	}

	if _, exists := context["generatedAt"]; !exists {
		t.Error("Context should include generatedAt timestamp")
	}
}

// Helper functions
func contains(str, substr string) bool {
	return len(str) >= len(substr) && substr != "" &&
		func() bool {
			for i := 0; i <= len(str)-len(substr); i++ {
				if str[i:i+len(substr)] == substr {
					return true
				}
			}
			return false
		}()
}

// Benchmark tests for handlers
func BenchmarkStoryContextHandler_GetCharacters(b *testing.B) {
	mockStorage := NewMockStorage()
	handler := NewStoryContextHandler(mockStorage)

	// Add test data
	for i := 0; i < 100; i++ {
		testChar := models.Character{
			ID:          fmt.Sprintf("char%d", i),
			Name:        fmt.Sprintf("Character %d", i),
			Description: fmt.Sprintf("Description for character %d", i),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		mockStorage.CreateCharacter(&testChar)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := handler.GetCharacters(map[string]interface{}{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStoryContextHandler_CreateCharacter(b *testing.B) {
	mockStorage := NewMockStorage()
	handler := NewStoryContextHandler(mockStorage)

	params := map[string]interface{}{
		"name":        "Benchmark Character",
		"description": "A character created during benchmarking",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		params["name"] = fmt.Sprintf("Benchmark Character %d", i)
		_, err := handler.CreateCharacter(params)
		if err != nil {
			b.Fatal(err)
		}
	}
}
