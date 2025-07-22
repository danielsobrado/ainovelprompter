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



func mustMarshalJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
