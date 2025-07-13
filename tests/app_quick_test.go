package main

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBasicStorageOperations tests the basic storage operations work correctly
func TestBasicStorageOperations(t *testing.T) {
	tempDir := t.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	// Test creating a character
	character := &models.Character{
		Name:        "Test Character",
		Description: "A character for testing storage operations",
		Traits:      map[string]string{"role": "protagonist"},
		Notes:       "Created during testing",
	}

	version, err := fs.Create(storage.EntityCharacters, character)
	require.NoError(t, err)
	assert.Equal(t, storage.OperationCreate, version.Operation)
	assert.True(t, version.Active)
	assert.NotEmpty(t, character.ID)

	// Test retrieving the character
	retrieved, err := fs.GetLatest(storage.EntityCharacters, character.ID)
	require.NoError(t, err)
	retrievedChar := retrieved.(*models.Character)
	assert.Equal(t, character.Name, retrievedChar.Name)
	assert.Equal(t, character.Description, retrievedChar.Description)

	// Test updating the character
	character.Description = "Updated description for testing"
	updateVersion, err := fs.Update(storage.EntityCharacters, character.ID, character)
	require.NoError(t, err)
	assert.Equal(t, storage.OperationUpdate, updateVersion.Operation)
	assert.True(t, updateVersion.Active)

	// Test version history
	versions, err := fs.GetVersions(storage.EntityCharacters, character.ID)
	require.NoError(t, err)
	assert.Len(t, versions, 2) // Create + Update

	// Test storage statistics
	stats, err := fs.GetStorageStats()
	require.NoError(t, err)
	assert.Equal(t, 2, stats.TotalFiles) // 2 versions
	assert.Equal(t, 1, stats.EntitiesByType[storage.EntityCharacters])
	assert.Equal(t, 2, stats.VersionsByType[storage.EntityCharacters])
	assert.Greater(t, stats.TotalSize, int64(0))
}

// TestLegacyInterfaceCompatibility tests that the legacy interface still works
func TestLegacyInterfaceCompatibility(t *testing.T) {
	tempDir := t.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	// Test creating character through legacy interface
	character := &models.Character{
		Name:        "Legacy Test Character",
		Description: "Testing legacy interface compatibility",
	}

	err := fs.CreateCharacter(character)
	require.NoError(t, err)
	assert.NotEmpty(t, character.ID)

	// Test retrieving through legacy interface
	characters, err := fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 1)
	assert.Equal(t, character.Name, characters[0].Name)

	// Test getting by ID through legacy interface
	retrieved, err := fs.GetCharacterByID(character.ID)
	require.NoError(t, err)
	assert.Equal(t, character.Name, retrieved.Name)

	// Test updating through legacy interface
	character.Description = "Updated through legacy interface"
	err = fs.UpdateCharacter(character)
	require.NoError(t, err)

	// Verify update worked
	updated, err := fs.GetCharacterByID(character.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated through legacy interface", updated.Description)
}

// TestFileNamingConventions tests the file naming system
func TestFileNamingConventions(t *testing.T) {
	tempDir := t.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	character := &models.Character{
		Name:        "Test Character With Spaces & Special!@# Characters",
		Description: "Testing file naming conventions",
	}

	_, err := fs.Create(storage.EntityCharacters, character)
	require.NoError(t, err)

	// Check that files follow naming convention
	files, err := os.ReadDir(tempDir + "/" + storage.EntityCharacters)
	require.NoError(t, err)
	require.Len(t, files, 1)

	filename := files[0].Name()

	// Should follow pattern: {entity_name}_{YYYYMMDD_HHMMSS}_{operation}.json
	assert.Contains(t, filename, "test_character_with_spaces")
	assert.Contains(t, filename, "_create.json")
	assert.Regexp(t, `\d{8}_\d{6}`, filename) // Date and time pattern

	// Should not contain special characters
	assert.NotContains(t, filename, " ")
	assert.NotContains(t, filename, "&")
	assert.NotContains(t, filename, "!")
	assert.NotContains(t, filename, "@")
	assert.NotContains(t, filename, "#")

	// Test that we can parse the timestamp from filename
	parts := strings.Split(strings.TrimSuffix(filename, ".json"), "_")
	assert.GreaterOrEqual(t, len(parts), 4) // name parts + date + time + operation

	// Find date and time parts
	dateTimeFound := false
	for i := 0; i < len(parts)-1; i++ {
		if len(parts[i]) == 8 && len(parts[i+1]) == 6 {
			// Try to parse as timestamp
			timestampStr := parts[i] + "_" + parts[i+1]
			_, err := time.Parse("20060102_150405", timestampStr)
			if err == nil {
				dateTimeFound = true
				break
			}
		}
	}
	assert.True(t, dateTimeFound, "Should contain valid timestamp in filename")
}

// TestAppIntegration tests the integration with the App struct
func TestAppIntegration(t *testing.T) {
	tempDir := t.TempDir()
	app := NewApp(tempDir)

	// Test GetDataDirectory
	assert.Equal(t, tempDir, app.GetDataDirectory())

	// Test SetDataDirectory
	newDir := t.TempDir()
	err := app.SetDataDirectory(newDir)
	require.NoError(t, err)
	assert.Equal(t, newDir, app.GetDataDirectory())

	// Test ValidateDataDirectory
	err = app.ValidateDataDirectory()
	require.NoError(t, err)

	// Create some test data using the app's storage
	fs := storage.NewFolderStorage(app.GetDataDirectory())
	character := &models.Character{
		Name:        "Stats Test Character",
		Description: "Character for testing storage stats",
	}
	_, err = fs.Create(storage.EntityCharacters, character)
	require.NoError(t, err)

	// Test GetStorageStats
	stats, err := app.GetStorageStats()
	require.NoError(t, err)
	assert.Greater(t, stats.TotalFiles, 0)
	assert.Greater(t, stats.TotalSize, int64(0))
	assert.Equal(t, 1, stats.EntitiesByType[storage.EntityCharacters])
}
