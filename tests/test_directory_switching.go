package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSetDataDirectoryDebug specifically tests directory switching
func TestSetDataDirectoryDebug(t *testing.T) {
	dir1 := filepath.Join(t.TempDir(), "project1")
	dir2 := filepath.Join(t.TempDir(), "project2")

	t.Logf("Directory 1: %s", dir1)
	t.Logf("Directory 2: %s", dir2)

	// Create storage instance for dir1
	fs := storage.NewFolderStorage(dir1)
	t.Logf("Created storage for dir1")

	// Create character in dir1
	char1 := &models.Character{
		Name:        "Project 1 Character",
		Description: "Character in first project",
	}
	version1, err := fs.Create(storage.EntityCharacters, char1)
	require.NoError(t, err)
	t.Logf("Created character in dir1: ID=%s, Name=%s", char1.ID, char1.Name)
	t.Logf("Version: %+v", version1)

	// Verify character exists in dir1
	characters1, err := fs.GetCharacters()
	require.NoError(t, err)
	t.Logf("Characters in dir1 after creation: %d", len(characters1))
	require.Len(t, characters1, 1)

	// Check files created in dir1
	charDir1 := filepath.Join(dir1, storage.EntityCharacters)
	if entries, err := os.ReadDir(charDir1); err == nil {
		t.Logf("Files in dir1/characters: %d", len(entries))
		for _, entry := range entries {
			t.Logf("  File: %s", entry.Name())
		}
	}

	// Switch to dir2
	t.Logf("Switching to dir2...")
	err = fs.SetDataDirectory(dir2)
	require.NoError(t, err)
	t.Logf("Successfully switched to dir2")

	// Verify dir2 is empty
	characters2, err := fs.GetCharacters()
	require.NoError(t, err)
	t.Logf("Characters in dir2 after switch: %d", len(characters2))
	assert.Len(t, characters2, 0)

	// Create character in dir2
	char2 := &models.Character{
		Name:        "Project 2 Character",
		Description: "Character in second project",
	}
	version2, err := fs.Create(storage.EntityCharacters, char2)
	require.NoError(t, err)
	t.Logf("Created character in dir2: ID=%s, Name=%s", char2.ID, char2.Name)

	// Verify character exists in dir2
	characters2b, err := fs.GetCharacters()
	require.NoError(t, err)
	t.Logf("Characters in dir2 after creation: %d", len(characters2b))
	require.Len(t, characters2b, 1)

	// Check files created in dir2
	charDir2 := filepath.Join(dir2, storage.EntityCharacters)
	if entries, err := os.ReadDir(charDir2); err == nil {
		t.Logf("Files in dir2/characters: %d", len(entries))
		for _, entry := range entries {
			t.Logf("  File: %s", entry.Name())
		}
	}

	// Switch back to dir1
	t.Logf("Switching back to dir1...")
	err = fs.SetDataDirectory(dir1)
	require.NoError(t, err)
	t.Logf("Successfully switched back to dir1")

	// Check if files still exist in dir1
	if entries, err := os.ReadDir(charDir1); err == nil {
		t.Logf("Files in dir1/characters after switch back: %d", len(entries))
		for _, entry := range entries {
			t.Logf("  File: %s", entry.Name())
		}
	} else {
		t.Logf("Error reading dir1/characters: %v", err)
	}

	// Try to get characters from dir1
	characters1b, err := fs.GetCharacters()
	require.NoError(t, err)
	t.Logf("Characters retrieved from dir1 after switch back: %d", len(characters1b))

	if len(characters1b) == 0 {
		t.Logf("ERROR: No characters found in dir1 after switch back!")

		// Debug the cache state
		if entities, err := fs.GetAll(storage.EntityCharacters); err == nil {
			t.Logf("GetAll returned %d entities", len(entities))
		} else {
			t.Logf("GetAll failed: %v", err)
		}
	} else {
		t.Logf("SUCCESS: Found %d characters in dir1", len(characters1b))
		for i, char := range characters1b {
			t.Logf("  Character %d: ID=%s, Name=%s", i, char.ID, char.Name)
		}
	}

	// This should pass
	require.Len(t, characters1b, 1)
	assert.Equal(t, "Project 1 Character", characters1b[0].Name)
}
