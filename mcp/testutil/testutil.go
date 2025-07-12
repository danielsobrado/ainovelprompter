// Package testutil provides testing utilities for the MCP server
package testutil

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp"
	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// TestServer wraps MCPServer with test utilities
type TestServer struct {
	*mcp.MCPServer
	tempDir string
	t       *testing.T
}

// NewTestServer creates a test server with temporary storage
func NewTestServer(t *testing.T) *TestServer {
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

	server, err := mcp.NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create test server: %v", err)
	}

	return &TestServer{
		MCPServer: server,
		tempDir:   tempDir,
		t:         t,
	}
}

// CreateTestCharacter creates a test character and returns it
func (ts *TestServer) CreateTestCharacter(name string) models.Character {
	ts.t.Helper()

	result, err := ts.ExecuteTool("create_character", map[string]interface{}{
		"name":        name,
		"description": "A test character created by " + name,
		"notes":       "Test character for unit testing",
		"traits": map[string]interface{}{
			"test": true,
		},
	})
	if err != nil {
		ts.t.Fatalf("Failed to create test character %s: %v", name, err)
	}

	character, ok := result.(models.Character)
	if !ok {
		ts.t.Fatalf("create_character returned wrong type: %T", result)
	}

	return character
}

// CreateTestChapter creates a test chapter and returns it
func (ts *TestServer) CreateTestChapter(title string) models.Chapter {
	ts.t.Helper()

	result, err := ts.ExecuteTool("create_chapter", map[string]interface{}{
		"title":   title,
		"content": "This is test content for chapter: " + title,
		"summary": "Test chapter summary",
		"status":  "draft",
	})
	if err != nil {
		ts.t.Fatalf("Failed to create test chapter %s: %v", title, err)
	}

	chapter, ok := result.(models.Chapter)
	if !ok {
		ts.t.Fatalf("create_chapter returned wrong type: %T", result)
	}

	return chapter
}

// AssertToolExists verifies a tool exists in the server
func (ts *TestServer) AssertToolExists(toolName string) {
	ts.t.Helper()

	tools := ts.GetTools()
	for _, tool := range tools {
		if tool.Name == toolName {
			return
		}
	}
	ts.t.Errorf("Tool %s not found in server", toolName)
}

// AssertToolCount verifies the expected number of tools
func (ts *TestServer) AssertToolCount(expected int) {
	ts.t.Helper()

	tools := ts.GetTools()
	if len(tools) != expected {
		ts.t.Errorf("Expected %d tools, got %d", expected, len(tools))
	}
}

// ExecuteToolAssertNoError executes a tool and asserts no error occurs
func (ts *TestServer) ExecuteToolAssertNoError(toolName string, params map[string]interface{}) interface{} {
	ts.t.Helper()

	result, err := ts.ExecuteTool(toolName, params)
	if err != nil {
		ts.t.Fatalf("Tool %s failed: %v", toolName, err)
	}
	return result
}

// ExecuteToolAssertError executes a tool and asserts an error occurs
func (ts *TestServer) ExecuteToolAssertError(toolName string, params map[string]interface{}) {
	ts.t.Helper()

	_, err := ts.ExecuteTool(toolName, params)
	if err == nil {
		ts.t.Errorf("Expected error for tool %s, but got none", toolName)
	}
}

// WaitForCondition waits for a condition to be true with timeout
func (ts *TestServer) WaitForCondition(condition func() bool, timeout time.Duration, message string) {
	ts.t.Helper()

	start := time.Now()
	for time.Since(start) < timeout {
		if condition() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	ts.t.Fatalf("Condition not met within %v: %s", timeout, message)
}

// CleanupTestData removes all test data from the server
func (ts *TestServer) CleanupTestData() {
	ts.t.Helper()

	// Get and delete all characters
	if result, err := ts.ExecuteTool("get_characters", map[string]interface{}{}); err == nil {
		if characters, ok := result.([]models.Character); ok {
			for _, char := range characters {
				ts.ExecuteTool("delete_character", map[string]interface{}{"id": char.ID})
			}
		}
	}

	// Get and delete all chapters
	if result, err := ts.ExecuteTool("get_chapters", map[string]interface{}{}); err == nil {
		if chapters, ok := result.([]models.Chapter); ok {
			for _, chapter := range chapters {
				ts.ExecuteTool("delete_chapter", map[string]interface{}{"id": chapter.ID})
			}
		}
	}
}

// TestDataBuilder helps build test data sets
type TestDataBuilder struct {
	server *TestServer
}

// NewTestDataBuilder creates a new test data builder
func NewTestDataBuilder(server *TestServer) *TestDataBuilder {
	return &TestDataBuilder{server: server}
}

// WithCharacters creates multiple test characters
func (tdb *TestDataBuilder) WithCharacters(count int) *TestDataBuilder {
	for i := 0; i < count; i++ {
		tdb.server.CreateTestCharacter(fmt.Sprintf("Character_%d", i))
	}
	return tdb
}

// WithChapters creates multiple test chapters
func (tdb *TestDataBuilder) WithChapters(count int) *TestDataBuilder {
	for i := 0; i < count; i++ {
		tdb.server.CreateTestChapter(fmt.Sprintf("Chapter_%d", i))
	}
	return tdb
}

// WithSpecificCharacter creates a character with specific properties
func (tdb *TestDataBuilder) WithSpecificCharacter(name, description string, traits map[string]interface{}) models.Character {
	result, err := tdb.server.ExecuteTool("create_character", map[string]interface{}{
		"name":        name,
		"description": description,
		"traits":      traits,
	})
	if err != nil {
		tdb.server.t.Fatalf("Failed to create specific character: %v", err)
	}

	character, ok := result.(models.Character)
	if !ok {
		tdb.server.t.Fatalf("create_character returned wrong type: %T", result)
	}

	return character
}

// Build completes the test data setup
func (tdb *TestDataBuilder) Build() {
	// Test data is already created, this is just for API completeness
}

// Performance test utilities

// BenchmarkHelper provides utilities for benchmark tests
type BenchmarkHelper struct {
	server *TestServer
}

// NewBenchmarkHelper creates a benchmark helper
func NewBenchmarkHelper(server *TestServer) *BenchmarkHelper {
	return &BenchmarkHelper{server: server}
}

// CreateBenchmarkData creates a large dataset for performance testing
func (bh *BenchmarkHelper) CreateBenchmarkData(characters, chapters int) {
	// Create characters
	for i := 0; i < characters; i++ {
		bh.server.CreateTestCharacter(fmt.Sprintf("BenchChar_%d", i))
	}

	// Create chapters
	for i := 0; i < chapters; i++ {
		bh.server.CreateTestChapter(fmt.Sprintf("BenchChapter_%d", i))
	}
}

// MeasureOperation measures the execution time of an operation
func (bh *BenchmarkHelper) MeasureOperation(operation func()) time.Duration {
	start := time.Now()
	operation()
	return time.Since(start)
}

// AssertPerformance asserts that an operation completes within expected time
func (bh *BenchmarkHelper) AssertPerformance(operation func(), maxDuration time.Duration, operationName string) {
	duration := bh.MeasureOperation(operation)
	if duration > maxDuration {
		bh.server.t.Errorf("Operation %s took %v, expected less than %v", 
			operationName, duration, maxDuration)
	}
}
