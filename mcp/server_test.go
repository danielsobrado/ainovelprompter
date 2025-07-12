package mcp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// TestMCPServer_NewMCPServer tests the server initialization
func TestMCPServer_NewMCPServer(t *testing.T) {
	// Setup temporary directory for testing
	tempDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	if originalHome == "" {
		originalHome = os.Getenv("USERPROFILE") // Windows
	}
	
	// Set temp directory as home for testing
	os.Setenv("HOME", tempDir)
	if originalHome != "" && os.Getenv("USERPROFILE") != "" {
		os.Setenv("USERPROFILE", tempDir)
	}
	
	defer func() {
		if originalHome != "" {
			os.Setenv("HOME", originalHome)
			os.Setenv("USERPROFILE", originalHome)
		}
	}()

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("NewMCPServer() failed: %v", err)
	}

	if server == nil {
		t.Fatal("NewMCPServer() returned nil server")
	}

	if server.storyHandler == nil {
		t.Error("storyHandler is nil")
	}

	if server.chapterHandler == nil {
		t.Error("chapterHandler is nil")
	}

	if server.proseHandler == nil {
		t.Error("proseHandler is nil")
	}

	if server.searchHandler == nil {
		t.Error("searchHandler is nil")
	}
}

// TestMCPServer_GetTools tests tool discovery
func TestMCPServer_GetTools(t *testing.T) {
	server := setupTestServer(t)

	tools := server.GetTools()

	expectedToolCount := 44
	if len(tools) != expectedToolCount {
		t.Errorf("Expected %d tools, got %d", expectedToolCount, len(tools))
	}

	// Verify required tools exist
	requiredTools := []string{
		"get_characters", "create_character", "update_character", "delete_character",
		"get_locations", "create_location",
		"get_codex_entries", "create_codex_entry",
		"get_rules", "get_rule_by_id",
		"get_chapters", "create_chapter", "update_chapter", "delete_chapter",
		"get_prose_prompts", "create_prose_prompt", "analyze_prose",
		"search_all_content", "analyze_text_traits", "get_character_mentions",
		"generate_chapter_prompt", "get_prompt_template",
	}

	toolMap := make(map[string]bool)
	for _, tool := range tools {
		toolMap[tool.Name] = true
		
		// Verify tool structure
		if tool.Name == "" {
			t.Error("Tool has empty name")
		}
		if tool.Description == "" {
			t.Errorf("Tool %s has empty description", tool.Name)
		}
	}

	for _, requiredTool := range requiredTools {
		if !toolMap[requiredTool] {
			t.Errorf("Required tool %s not found", requiredTool)
		}
	}
}

// TestMCPServer_ExecuteTool_Characters tests character operations
func TestMCPServer_ExecuteTool_Characters(t *testing.T) {
	server := setupTestServer(t)

	// Test get_characters (empty initially)
	result, err := server.ExecuteTool("get_characters", map[string]interface{}{})
	if err != nil {
		t.Fatalf("get_characters failed: %v", err)
	}

	characters, ok := result.([]models.Character)
	if !ok {
		t.Fatalf("get_characters returned wrong type: %T", result)
	}

	if len(characters) != 0 {
		t.Errorf("Expected 0 characters initially, got %d", len(characters))
	}

	// Test create_character
	createParams := map[string]interface{}{
		"name":        "Test Character",
		"description": "A test character",
		"notes":       "Created in unit test",
		"traits": map[string]interface{}{
			"brave":  true,
			"height": "tall",
		},
	}

	createResult, err := server.ExecuteTool("create_character", createParams)
	if err != nil {
		t.Fatalf("create_character failed: %v", err)
	}

	character, ok := createResult.(models.Character)
	if !ok {
		t.Fatalf("create_character returned wrong type: %T", createResult)
	}

	if character.ID == "" {
		t.Error("Created character has empty ID")
	}

	if character.Name != "Test Character" {
		t.Errorf("Expected name 'Test Character', got '%s'", character.Name)
	}

	if character.CreatedAt.IsZero() {
		t.Error("Created character has zero CreatedAt time")
	}

	// Test get_characters (should now have 1)
	result, err = server.ExecuteTool("get_characters", map[string]interface{}{})
	if err != nil {
		t.Fatalf("get_characters after create failed: %v", err)
	}

	characters, ok = result.([]models.Character)
	if !ok {
		t.Fatalf("get_characters returned wrong type: %T", result)
	}

	if len(characters) != 1 {
		t.Errorf("Expected 1 character after create, got %d", len(characters))
	}

	// Test get_character_by_id
	getResult, err := server.ExecuteTool("get_character_by_id", map[string]interface{}{
		"id": character.ID,
	})
	if err != nil {
		t.Fatalf("get_character_by_id failed: %v", err)
	}

	retrievedChar, ok := getResult.(models.Character)
	if !ok {
		t.Fatalf("get_character_by_id returned wrong type: %T", getResult)
	}

	if retrievedChar.ID != character.ID {
		t.Errorf("Retrieved character ID mismatch: expected %s, got %s", character.ID, retrievedChar.ID)
	}

	// Test update_character
	updateParams := map[string]interface{}{
		"id":          character.ID,
		"name":        "Updated Test Character",
		"description": "An updated test character",
	}

	_, err = server.ExecuteTool("update_character", updateParams)
	if err != nil {
		t.Fatalf("update_character failed: %v", err)
	}

	// Verify update
	getResult, err = server.ExecuteTool("get_character_by_id", map[string]interface{}{
		"id": character.ID,
	})
	if err != nil {
		t.Fatalf("get_character_by_id after update failed: %v", err)
	}

	updatedChar, ok := getResult.(models.Character)
	if !ok {
		t.Fatalf("get_character_by_id returned wrong type: %T", getResult)
	}

	if updatedChar.Name != "Updated Test Character" {
		t.Errorf("Character name not updated: expected 'Updated Test Character', got '%s'", updatedChar.Name)
	}

	// Test delete_character
	_, err = server.ExecuteTool("delete_character", map[string]interface{}{
		"id": character.ID,
	})
	if err != nil {
		t.Fatalf("delete_character failed: %v", err)
	}

	// Verify deletion
	result, err = server.ExecuteTool("get_characters", map[string]interface{}{})
	if err != nil {
		t.Fatalf("get_characters after delete failed: %v", err)
	}

	characters, ok = result.([]models.Character)
	if !ok {
		t.Fatalf("get_characters returned wrong type: %T", result)
	}

	if len(characters) != 0 {
		t.Errorf("Expected 0 characters after delete, got %d", len(characters))
	}
}

// TestMCPServer_ExecuteTool_Chapters tests chapter operations
func TestMCPServer_ExecuteTool_Chapters(t *testing.T) {
	server := setupTestServer(t)

	// Test create_chapter
	createParams := map[string]interface{}{
		"title":   "Test Chapter",
		"content": "This is a test chapter with some content for testing purposes.",
		"summary": "A test chapter",
		"status":  "draft",
	}

	createResult, err := server.ExecuteTool("create_chapter", createParams)
	if err != nil {
		t.Fatalf("create_chapter failed: %v", err)
	}

	chapter, ok := createResult.(models.Chapter)
	if !ok {
		t.Fatalf("create_chapter returned wrong type: %T", createResult)
	}

	if chapter.ID == "" {
		t.Error("Created chapter has empty ID")
	}

	if chapter.Number == 0 {
		t.Error("Created chapter has zero number (should auto-assign)")
	}

	if chapter.WordCount == 0 {
		t.Error("Created chapter has zero word count")
	}

	expectedWordCount := 12 // "This is a test chapter with some content for testing purposes."
	if chapter.WordCount != expectedWordCount {
		t.Errorf("Expected word count %d, got %d", expectedWordCount, chapter.WordCount)
	}

	// Test get_chapters
	result, err := server.ExecuteTool("get_chapters", map[string]interface{}{})
	if err != nil {
		t.Fatalf("get_chapters failed: %v", err)
	}

	chapters, ok := result.([]models.Chapter)
	if !ok {
		t.Fatalf("get_chapters returned wrong type: %T", result)
	}

	if len(chapters) != 1 {
		t.Errorf("Expected 1 chapter, got %d", len(chapters))
	}

	// Test get_chapter_content
	contentResult, err := server.ExecuteTool("get_chapter_content", map[string]interface{}{
		"id": chapter.ID,
	})
	if err != nil {
		t.Fatalf("get_chapter_content failed: %v", err)
	}

	retrievedChapter, ok := contentResult.(models.Chapter)
	if !ok {
		t.Fatalf("get_chapter_content returned wrong type: %T", contentResult)
	}

	if retrievedChapter.Content != createParams["content"] {
		t.Errorf("Chapter content mismatch: expected '%s', got '%s'", 
			createParams["content"], retrievedChapter.Content)
	}
}

// TestMCPServer_ExecuteTool_ProseImprovement tests prose improvement operations
func TestMCPServer_ExecuteTool_ProseImprovement(t *testing.T) {
	server := setupTestServer(t)

	// Test get_prose_prompts
	result, err := server.ExecuteTool("get_prose_prompts", map[string]interface{}{})
	if err != nil {
		t.Fatalf("get_prose_prompts failed: %v", err)
	}

	prompts, ok := result.([]models.ProseImprovementPrompt)
	if !ok {
		t.Fatalf("get_prose_prompts returned wrong type: %T", result)
	}

	// Should have default prompts
	if len(prompts) == 0 {
		t.Error("Expected default prose prompts, got none")
	}

	// Verify default prompts structure
	for _, prompt := range prompts {
		if prompt.ID == "" {
			t.Error("Prose prompt has empty ID")
		}
		if prompt.Label == "" {
			t.Error("Prose prompt has empty label")
		}
		if prompt.DefaultPromptText == "" {
			t.Error("Prose prompt has empty default text")
		}
	}

	// Test create_prose_session
	sessionParams := map[string]interface{}{
		"text": "This is some sample text that needs improvement. It could be written better.",
	}

	sessionResult, err := server.ExecuteTool("create_prose_session", sessionParams)
	if err != nil {
		t.Fatalf("create_prose_session failed: %v", err)
	}

	session, ok := sessionResult.(models.ProseImprovementSession)
	if !ok {
		t.Fatalf("create_prose_session returned wrong type: %T", sessionResult)
	}

	if session.ID == "" {
		t.Error("Created session has empty ID")
	}

	if session.OriginalText != sessionParams["text"] {
		t.Errorf("Session original text mismatch: expected '%s', got '%s'",
			sessionParams["text"], session.OriginalText)
	}

	if len(session.Prompts) == 0 {
		t.Error("Session should have prompts loaded")
	}

	// Test get_prose_session
	getSessionResult, err := server.ExecuteTool("get_prose_session", map[string]interface{}{
		"sessionId": session.ID,
	})
	if err != nil {
		t.Fatalf("get_prose_session failed: %v", err)
	}

	retrievedSession, ok := getSessionResult.(models.ProseImprovementSession)
	if !ok {
		t.Fatalf("get_prose_session returned wrong type: %T", getSessionResult)
	}

	if retrievedSession.ID != session.ID {
		t.Errorf("Retrieved session ID mismatch: expected %s, got %s", 
			session.ID, retrievedSession.ID)
	}
}

// TestMCPServer_ExecuteTool_Search tests search and analysis operations
func TestMCPServer_ExecuteTool_Search(t *testing.T) {
	server := setupTestServer(t)

	// Create some test data first
	_, err := server.ExecuteTool("create_character", map[string]interface{}{
		"name":        "Search Test Character",
		"description": "A character for search testing",
	})
	if err != nil {
		t.Fatalf("Failed to create test character: %v", err)
	}

	// Test search_all_content
	searchResult, err := server.ExecuteTool("search_all_content", map[string]interface{}{
		"query": "search",
		"limit": 10,
	})
	if err != nil {
		t.Fatalf("search_all_content failed: %v", err)
	}

	searchResults, ok := searchResult.([]models.SearchResult)
	if !ok {
		t.Fatalf("search_all_content returned wrong type: %T", searchResult)
	}

	// Should find the character we created
	found := false
	for _, result := range searchResults {
		if result.Type == "character" && result.Title == "Search Test Character" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Search should have found the test character")
	}

	// Test analyze_text_traits
	traitsResult, err := server.ExecuteTool("analyze_text_traits", map[string]interface{}{
		"text": "This is a sample text for analysis. It has multiple sentences! Does it work? Yes, it should work well.",
	})
	if err != nil {
		t.Fatalf("analyze_text_traits failed: %v", err)
	}

	traits, ok := traitsResult.(map[string]interface{})
	if !ok {
		t.Fatalf("analyze_text_traits returned wrong type: %T", traitsResult)
	}

	// Verify expected analysis fields
	expectedFields := []string{"wordCount", "characterCount", "sentenceCount", "tone", "writingStyle"}
	for _, field := range expectedFields {
		if _, exists := traits[field]; !exists {
			t.Errorf("analyze_text_traits missing field: %s", field)
		}
	}

	// Verify basic counts
	if wordCount, ok := traits["wordCount"].(int); ok {
		if wordCount == 0 {
			t.Error("Word count should not be zero")
		}
	} else {
		t.Error("wordCount should be an integer")
	}
}

// TestMCPServer_ExecuteTool_PromptGeneration tests prompt generation
func TestMCPServer_ExecuteTool_PromptGeneration(t *testing.T) {
	server := setupTestServer(t)

	// Test get_prompt_template
	templateResult, err := server.ExecuteTool("get_prompt_template", map[string]interface{}{
		"format": "ChatGPT",
	})
	if err != nil {
		t.Fatalf("get_prompt_template failed: %v", err)
	}

	template, ok := templateResult.(map[string]interface{})
	if !ok {
		t.Fatalf("get_prompt_template returned wrong type: %T", templateResult)
	}

	if template["template"] == "" {
		t.Error("Template should not be empty")
	}

	if template["format"] != "ChatGPT" {
		t.Errorf("Expected format 'ChatGPT', got '%v'", template["format"])
	}

	// Test generate_chapter_prompt
	promptResult, err := server.ExecuteTool("generate_chapter_prompt", map[string]interface{}{
		"promptType":       "ChatGPT",
		"taskType":         "Write the next chapter",
		"nextChapterBeats": "The protagonist discovers something important",
	})
	if err != nil {
		t.Fatalf("generate_chapter_prompt failed: %v", err)
	}

	prompt, ok := promptResult.(map[string]interface{})
	if !ok {
		t.Fatalf("generate_chapter_prompt returned wrong type: %T", promptResult)
	}

	if prompt["prompt"] == "" {
		t.Error("Generated prompt should not be empty")
	}

	if prompt["promptType"] != "ChatGPT" {
		t.Errorf("Expected promptType 'ChatGPT', got '%v'", prompt["promptType"])
	}

	// Verify prompt contains our task
	promptText, ok := prompt["prompt"].(string)
	if !ok {
		t.Fatal("Prompt text should be a string")
	}

	if !contains(promptText, "Write the next chapter") {
		t.Error("Generated prompt should contain task type")
	}

	if !contains(promptText, "discovers something important") {
		t.Error("Generated prompt should contain story beats")
	}
}

// TestMCPServer_ExecuteTool_ErrorHandling tests error conditions
func TestMCPServer_ExecuteTool_ErrorHandling(t *testing.T) {
	server := setupTestServer(t)

	// Test invalid tool name
	_, err := server.ExecuteTool("invalid_tool_name", map[string]interface{}{})
	if err == nil {
		t.Error("Expected error for invalid tool name")
	}

	expectedError := "unknown tool: invalid_tool_name"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	// Test missing required parameters
	_, err = server.ExecuteTool("create_character", map[string]interface{}{})
	if err == nil {
		t.Error("Expected error for missing required parameters")
	}

	// Test invalid parameter types
	_, err = server.ExecuteTool("get_character_by_id", map[string]interface{}{
		"id": 123, // Should be string
	})
	if err == nil {
		t.Error("Expected error for invalid parameter type")
	}

	// Test non-existent resource
	_, err = server.ExecuteTool("get_character_by_id", map[string]interface{}{
		"id": "non-existent-id",
	})
	if err == nil {
		t.Error("Expected error for non-existent character")
	}
}

// Helper functions

func setupTestServer(t *testing.T) *MCPServer {
	t.Helper()

	// Setup temporary directory for testing
	tempDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	if originalHome == "" {
		originalHome = os.Getenv("USERPROFILE") // Windows
	}
	
	// Set temp directory as home for testing
	os.Setenv("HOME", tempDir)
	if os.Getenv("USERPROFILE") != "" {
		os.Setenv("USERPROFILE", tempDir)
	}
	
	t.Cleanup(func() {
		if originalHome != "" {
			os.Setenv("HOME", originalHome)
			if os.Getenv("USERPROFILE") != "" {
				os.Setenv("USERPROFILE", originalHome)
			}
		}
	})

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create test server: %v", err)
	}

	return server
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && 
		   (len(substr) == 0 || str != str[:len(str)-len(substr)]+str[len(str)-len(substr)+1:])
}

// Benchmark tests

func BenchmarkMCPServer_GetTools(b *testing.B) {
	server := setupBenchmarkServer(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = server.GetTools()
	}
}

func BenchmarkMCPServer_ExecuteTool_GetCharacters(b *testing.B) {
	server := setupBenchmarkServer(b)

	// Create some test data
	for i := 0; i < 10; i++ {
		server.ExecuteTool("create_character", map[string]interface{}{
			"name":        fmt.Sprintf("Character %d", i),
			"description": fmt.Sprintf("Description for character %d", i),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := server.ExecuteTool("get_characters", map[string]interface{}{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMCPServer_ExecuteTool_Search(b *testing.B) {
	server := setupBenchmarkServer(b)

	// Create test data
	for i := 0; i < 50; i++ {
		server.ExecuteTool("create_character", map[string]interface{}{
			"name":        fmt.Sprintf("Character %d", i),
			"description": fmt.Sprintf("Description for character %d with search term", i),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := server.ExecuteTool("search_all_content", map[string]interface{}{
			"query": "search term",
			"limit": 10,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func setupBenchmarkServer(b *testing.B) *MCPServer {
	b.Helper()

	tempDir := b.TempDir()
	originalHome := os.Getenv("HOME")
	if originalHome == "" {
		originalHome = os.Getenv("USERPROFILE")
	}
	
	os.Setenv("HOME", tempDir)
	if os.Getenv("USERPROFILE") != "" {
		os.Setenv("USERPROFILE", tempDir)
	}
	
	b.Cleanup(func() {
		if originalHome != "" {
			os.Setenv("HOME", originalHome)
			if os.Getenv("USERPROFILE") != "" {
				os.Setenv("USERPROFILE", originalHome)
			}
		}
	})

	server, err := NewMCPServer()
	if err != nil {
		b.Fatalf("Failed to create benchmark server: %v", err)
	}

	return server
}
