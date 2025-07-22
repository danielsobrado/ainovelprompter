package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// TestFileStorage_NewFileStorage tests storage initialization
func TestFileStorage_NewFileStorage(t *testing.T) {
	tempDir := t.TempDir()
	
	storage := NewFileStorage(tempDir)
	
	if storage == nil {
		t.Fatal("NewFileStorage returned nil")
	}
	
	if storage.basePath != tempDir {
		t.Errorf("Expected basePath %s, got %s", tempDir, storage.basePath)
	}
}

// TestFileStorage_Characters tests character CRUD operations
func TestFileStorage_Characters(t *testing.T) {
	tempDir := t.TempDir()
	storage := NewFileStorage(tempDir)
	
	// Test GetCharacters (empty initially)
	characters, err := storage.GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters failed: %v", err)
	}
	
	if len(characters) != 0 {
		t.Errorf("Expected 0 characters initially, got %d", len(characters))
	}
	
	// Test CreateCharacter
	testChar := &models.Character{
		Name:        "Test Character",
		Description: "A test character for storage testing",
		Notes:       "Created in storage test",
		Traits: map[string]interface{}{
			"brave":  "true",
			"height": "tall",
		},
	}
	
	err = storage.CreateCharacter(testChar)
	if err != nil {
		t.Fatalf("CreateCharacter failed: %v", err)
	}
	
	if testChar.ID == "" {
		t.Error("CreateCharacter should assign an ID")
	}
	
	if testChar.CreatedAt.IsZero() {
		t.Error("CreateCharacter should set CreatedAt")
	}
	
	if testChar.UpdatedAt.IsZero() {
		t.Error("CreateCharacter should set UpdatedAt")
	}
	
	// Test GetCharacters (should now have 1)
	characters, err = storage.GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters after create failed: %v", err)
	}
	
	if len(characters) != 1 {
		t.Errorf("Expected 1 character after create, got %d", len(characters))
	}
	
	if characters[0].Name != "Test Character" {
		t.Errorf("Expected character name 'Test Character', got '%s'", characters[0].Name)
	}
	
	// Test GetCharacterByID
	retrievedChar, err := storage.GetCharacterByID(testChar.ID)
	if err != nil {
		t.Fatalf("GetCharacterByID failed: %v", err)
	}
	
	if retrievedChar.ID != testChar.ID {
		t.Errorf("Retrieved character ID mismatch: expected %s, got %s", 
			testChar.ID, retrievedChar.ID)
	}
	
	// Test UpdateCharacter
	testChar.Name = "Updated Test Character"
	testChar.Description = "An updated test character"
	originalUpdatedAt := testChar.UpdatedAt
	
	// Wait a moment to ensure UpdatedAt changes
	time.Sleep(time.Millisecond)
	
	err = storage.UpdateCharacter(testChar)
	if err != nil {
		t.Fatalf("UpdateCharacter failed: %v", err)
	}
	
	if !testChar.UpdatedAt.After(originalUpdatedAt) {
		t.Error("UpdateCharacter should update UpdatedAt timestamp")
	}
	
	// Verify update
	retrievedChar, err = storage.GetCharacterByID(testChar.ID)
	if err != nil {
		t.Fatalf("GetCharacterByID after update failed: %v", err)
	}
	
	if retrievedChar.Name != "Updated Test Character" {
		t.Errorf("Character name not updated: expected 'Updated Test Character', got '%s'", 
			retrievedChar.Name)
	}
	
	// Test SearchCharacters
	results, err := storage.SearchCharacters("updated")
	if err != nil {
		t.Fatalf("SearchCharacters failed: %v", err)
	}
	
	if len(results) != 1 {
		t.Errorf("Expected 1 search result, got %d", len(results))
	}
	
	if results[0].Name != "Updated Test Character" {
		t.Errorf("Search result name mismatch: expected 'Updated Test Character', got '%s'", 
			results[0].Name)
	}
	
	// Test DeleteCharacter
	err = storage.DeleteCharacter(testChar.ID)
	if err != nil {
		t.Fatalf("DeleteCharacter failed: %v", err)
	}
	
	// Verify deletion
	_, err = storage.GetCharacterByID(testChar.ID)
	if err == nil {
		t.Error("Expected error when getting deleted character")
	}
	
	characters, err = storage.GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters after delete failed: %v", err)
	}
	
	if len(characters) != 0 {
		t.Errorf("Expected 0 characters after delete, got %d", len(characters))
	}
}

// TestFileStorage_Chapters tests chapter CRUD operations
func TestFileStorage_Chapters(t *testing.T) {
	tempDir := t.TempDir()
	storage := NewFileStorage(tempDir)
	
	// Test CreateChapter
	testChapter := &models.Chapter{
		Title:   "Test Chapter",
		Content: "This is a test chapter with some content for testing purposes.",
		Summary: "A test chapter for storage testing",
		Status:  "draft",
	}
	
	err := storage.CreateChapter(testChapter)
	if err != nil {
		t.Fatalf("CreateChapter failed: %v", err)
	}
	
	if testChapter.ID == "" {
		t.Error("CreateChapter should assign an ID")
	}
	
	if testChapter.Number == 0 {
		t.Error("CreateChapter should auto-assign chapter number")
	}
	
	expectedWordCount := 12 // "This is a test chapter with some content for testing purposes."
	if testChapter.WordCount != expectedWordCount {
		t.Errorf("Expected word count %d, got %d", expectedWordCount, testChapter.WordCount)
	}
	
	// Test GetChapters
	chapters, err := storage.GetChapters()
	if err != nil {
		t.Fatalf("GetChapters failed: %v", err)
	}
	
	if len(chapters) != 1 {
		t.Errorf("Expected 1 chapter, got %d", len(chapters))
	}
	
	// Test GetChapterByID
	retrievedChapter, err := storage.GetChapterByID(testChapter.ID)
	if err != nil {
		t.Fatalf("GetChapterByID failed: %v", err)
	}
	
	if retrievedChapter.Content != testChapter.Content {
		t.Errorf("Chapter content mismatch: expected '%s', got '%s'", 
			testChapter.Content, retrievedChapter.Content)
	}
	
	// Test GetChapterByNumber
	retrievedByNumber, err := storage.GetChapterByNumber(testChapter.Number)
	if err != nil {
		t.Fatalf("GetChapterByNumber failed: %v", err)
	}
	
	if retrievedByNumber.ID != testChapter.ID {
		t.Errorf("GetChapterByNumber ID mismatch: expected %s, got %s", 
			testChapter.ID, retrievedByNumber.ID)
	}
	
	// Test SearchChapters
	searchResults, err := storage.SearchChapters("test chapter")
	if err != nil {
		t.Fatalf("SearchChapters failed: %v", err)
	}
	
	if len(searchResults) != 1 {
		t.Errorf("Expected 1 search result, got %d", len(searchResults))
	}
}

// TestFileStorage_BasicFileOperations tests the Storage interface
func TestFileStorage_BasicFileOperations(t *testing.T) {
	tempDir := t.TempDir()
	storage := NewFileStorage(tempDir)
	
	testData := []byte(`{"test": "data"}`)
	filename := "test_file.json"
	
	// Test WriteFile
	err := storage.WriteFile(filename, testData)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	
	// Verify file exists
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("File should exist after WriteFile")
	}
	
	// Test ReadFile
	readData, err := storage.ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	
	if string(readData) != string(testData) {
		t.Errorf("Read data mismatch: expected '%s', got '%s'", 
			string(testData), string(readData))
	}
	
	// Test ListFiles
	files, err := storage.ListFiles("*.json")
	if err != nil {
		t.Fatalf("ListFiles failed: %v", err)
	}
	
	found := false
	for _, file := range files {
		if file == filename {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("ListFiles should include %s", filename)
	}
	
	// Test DeleteFile
	err = storage.DeleteFile(filename)
	if err != nil {
		t.Fatalf("DeleteFile failed: %v", err)
	}
	
	// Verify file is deleted
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Error("File should not exist after DeleteFile")
	}
}

// TestFileStorage_ConcurrentAccess tests thread safety
func TestFileStorage_ConcurrentAccess(t *testing.T) {
	tempDir := t.TempDir()
	storage := NewFileStorage(tempDir)
	
	const numGoroutines = 10
	const charactersPerGoroutine = 5
	
	done := make(chan bool, numGoroutines)
	
	// Create characters concurrently
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer func() { done <- true }()
			
			for j := 0; j < charactersPerGoroutine; j++ {
				testChar := &models.Character{
					Name:        fmt.Sprintf("Character_%d_%d", goroutineID, j),
					Description: fmt.Sprintf("Character created by goroutine %d", goroutineID),
				}
				
				err := storage.CreateCharacter(testChar)
				if err != nil {
					t.Errorf("Concurrent CreateCharacter failed: %v", err)
					return
				}
			}
		}(i)
	}
	
	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
	
	// Verify all characters were created
	characters, err := storage.GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters after concurrent creation failed: %v", err)
	}
	
	expectedCount := numGoroutines * charactersPerGoroutine
	if len(characters) != expectedCount {
		t.Errorf("Expected %d characters after concurrent creation, got %d", 
			expectedCount, len(characters))
	}
}

// TestFileStorage_ErrorHandling tests error conditions
func TestFileStorage_ErrorHandling(t *testing.T) {
	tempDir := t.TempDir()
	storage := NewFileStorage(tempDir)
	
	// Test GetCharacterByID with non-existent ID
	_, err := storage.GetCharacterByID("non-existent-id")
	if err == nil {
		t.Error("Expected error when getting non-existent character")
	}
	
	// Test UpdateCharacter with non-existent character
	nonExistentChar := &models.Character{
		ID:   "non-existent-id",
		Name: "Non-existent",
	}
	
	err = storage.UpdateCharacter(nonExistentChar)
	if err == nil {
		t.Error("Expected error when updating non-existent character")
	}
	
	// Test DeleteCharacter with non-existent ID
	err = storage.DeleteCharacter("non-existent-id")
	if err == nil {
		t.Error("Expected error when deleting non-existent character")
	}
}

// TestFileStorage_DataPersistence tests data persistence across instances
func TestFileStorage_DataPersistence(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create first storage instance and add data
	storage1 := NewFileStorage(tempDir)
	
	testChar := &models.Character{
		Name:        "Persistent Character",
		Description: "A character to test persistence",
	}
	
	err := storage1.CreateCharacter(testChar)
	if err != nil {
		t.Fatalf("CreateCharacter failed: %v", err)
	}
	
	// Create second storage instance using same directory
	storage2 := NewFileStorage(tempDir)
	
	// Verify data persists
	characters, err := storage2.GetCharacters()
	if err != nil {
		t.Fatalf("GetCharacters from second instance failed: %v", err)
	}
	
	if len(characters) != 1 {
		t.Errorf("Expected 1 character in second instance, got %d", len(characters))
	}
	
	if characters[0].Name != "Persistent Character" {
		t.Errorf("Character name not persisted: expected 'Persistent Character', got '%s'", 
			characters[0].Name)
	}
}

// Benchmark tests for storage operations
func BenchmarkFileStorage_CreateCharacter(b *testing.B) {
	tempDir := b.TempDir()
	storage := NewFileStorage(tempDir)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testChar := &models.Character{
			Name:        fmt.Sprintf("Benchmark Character %d", i),
			Description: "A character created during benchmarking",
		}
		
		err := storage.CreateCharacter(testChar)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFileStorage_GetCharacters(b *testing.B) {
	tempDir := b.TempDir()
	storage := NewFileStorage(tempDir)
	
	// Create test data
	for i := 0; i < 100; i++ {
		testChar := &models.Character{
			Name:        fmt.Sprintf("Character %d", i),
			Description: fmt.Sprintf("Description for character %d", i),
		}
		storage.CreateCharacter(testChar)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := storage.GetCharacters()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFileStorage_SearchCharacters(b *testing.B) {
	tempDir := b.TempDir()
	storage := NewFileStorage(tempDir)
	
	// Create test data
	for i := 0; i < 100; i++ {
		testChar := &models.Character{
			Name:        fmt.Sprintf("Character %d", i),
			Description: fmt.Sprintf("Description for character %d with searchable term", i),
		}
		storage.CreateCharacter(testChar)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := storage.SearchCharacters("searchable")
		if err != nil {
			b.Fatal(err)
		}
	}
}
