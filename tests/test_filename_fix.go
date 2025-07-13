package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFilenameParsingFix specifically tests the filename parsing fix
func TestFilenameParsingFix(t *testing.T) {
	tempDir := t.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	// Create a character to generate a file
	character := &models.Character{
		Name:        "Filename Test Character",
		Description: "Testing filename parsing",
	}

	version, err := fs.Create(storage.EntityCharacters, character)
	require.NoError(t, err)
	t.Logf("Created version: %+v", version)

	// Check the actual filename that was created
	charDir := filepath.Join(tempDir, storage.EntityCharacters)
	entries, err := os.ReadDir(charDir)
	require.NoError(t, err)
	require.Len(t, entries, 1)

	filename := entries[0].Name()
	t.Logf("Generated filename: %s", filename)

	// Create a new storage instance to test cache rebuild
	fs2 := storage.NewFolderStorage(tempDir)
	
	// Test that GetAll works
	entities, err := fs2.GetAll(storage.EntityCharacters)
	require.NoError(t, err)
	t.Logf("GetAll returned %d entities", len(entities))

	if len(entities) > 0 {
		char := entities[0].(*models.Character)
		t.Logf("Retrieved character: ID=%s, Name=%s", char.ID, char.Name)
		assert.Equal(t, character.ID, char.ID)
		assert.Equal(t, character.Name, char.Name)
	}

	// Test that GetCharacters works
	characters, err := fs2.GetCharacters()
	require.NoError(t, err)
	t.Logf("GetCharacters returned %d characters", len(characters))
	
	require.Len(t, characters, 1)
	assert.Equal(t, character.ID, characters[0].ID)
	assert.Equal(t, character.Name, characters[0].Name)
}

// TestMigrationDebug tests migration with detailed debugging
func TestMigrationDebug(t *testing.T) {
	oldDir := filepath.Join(t.TempDir(), "old")
	newDir := filepath.Join(t.TempDir(), "new")

	// Create old directory
	err := os.MkdirAll(oldDir, 0755)
	require.NoError(t, err)

	// Create test character data
	character := models.Character{
		ID:          "test_char_1",
		Name:        "Debug Test Character", 
		Description: "Character for debugging migration",
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now().Add(-12 * time.Hour),
	}

	// Write JSON file
	jsonData := fmt.Sprintf(`[%s]`, mustMarshalJSON(character))
	err = os.WriteFile(filepath.Join(oldDir, "characters.json"), []byte(jsonData), 0644)
	require.NoError(t, err)
	
	t.Logf("Created test data: %s", jsonData)

	// Perform migration
	fs := storage.NewFolderStorage(newDir)
	err = fs.MigrateFromJSON(oldDir)
	require.NoError(t, err)

	// Check files created
	charDir := filepath.Join(newDir, storage.EntityCharacters)
	entries, err := os.ReadDir(charDir)
	require.NoError(t, err)
	t.Logf("Migration created %d files", len(entries))

	// Create fresh storage instance
	fs2 := storage.NewFolderStorage(newDir)

	// Test retrieval
	characters, err := fs2.GetCharacters()
	require.NoError(t, err)
	t.Logf("Retrieved %d characters", len(characters))

	if len(characters) > 0 {
		t.Logf("First character: ID=%s, Name=%s", characters[0].ID, characters[0].Name)
		assert.Equal(t, "test_char_1", characters[0].ID)
		assert.Equal(t, "Debug Test Character", characters[0].Name)
	}
}

func mustMarshalJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
