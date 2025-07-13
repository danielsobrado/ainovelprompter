package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigration_FromJSON(t *testing.T) {
	// Setup old JSON format directory
	oldDir := filepath.Join(os.TempDir(), "old_data_"+time.Now().Format("20060102_150405"))
	newDir := filepath.Join(os.TempDir(), "new_data_"+time.Now().Format("20060102_150405"))
	
	defer func() {
		os.RemoveAll(oldDir)
		os.RemoveAll(newDir)
	}()

	err := os.MkdirAll(oldDir, 0755)
	require.NoError(t, err)

	// Create test data in old JSON format
	testCharacters := []models.Character{
		{
			ID:          "char1",
			Name:        "Test Character 1",
			Description: "First test character",
			Traits:      map[string]string{"brave": "true"},
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now().Add(-12 * time.Hour),
		},
		{
			ID:          "char2",
			Name:        "Test Character 2",
			Description: "Second test character",
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   time.Now().Add(-6 * time.Hour),
		},
	}

	testLocations := []models.Location{
		{
			ID:          "loc1",
			Name:        "Test Location",
			Description: "A test location",
			Details:     "Detailed information",
			CreatedAt:   time.Now().Add(-72 * time.Hour),
			UpdatedAt:   time.Now().Add(-1 * time.Hour),
		},
	}

	testCodex := []models.CodexEntry{
		{
			ID:        "codex1",
			Title:     "Test Codex Entry",
			Category:  "Lore",
			Content:   "Test content for codex",
			Tags:      []string{"test", "migration"},
			CreatedAt: time.Now().Add(-96 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
	}

	// Write test data to JSON files
	writeTestJSON(t, oldDir, "characters.json", testCharacters)
	writeTestJSON(t, oldDir, "locations.json", testLocations)
	writeTestJSON(t, oldDir, "codex.json", testCodex)

	// Create new folder storage
	fs := NewFolderStorage(newDir)

	// Perform migration
	err = fs.MigrateFromJSON(oldDir)
	require.NoError(t, err)

	// Verify characters migrated correctly
	t.Run("Characters Migration", func(t *testing.T) {
		characters, err := fs.GetCharacters()
		require.NoError(t, err)
		assert.Len(t, characters, 2)

		// Find specific character
		var char1 *models.Character
		for _, c := range characters {
			if c.ID == "char1" {
				char1 = &c
				break
			}
		}
		require.NotNil(t, char1)
		assert.Equal(t, "Test Character 1", char1.Name)
		assert.Equal(t, "First test character", char1.Description)
		assert.Equal(t, "true", char1.Traits["brave"])
		assert.False(t, char1.CreatedAt.IsZero())
		assert.False(t, char1.UpdatedAt.IsZero())
	})

	// Verify locations migrated correctly
	t.Run("Locations Migration", func(t *testing.T) {
		locations, err := fs.GetLocations()
		require.NoError(t, err)
		assert.Len(t, locations, 1)

		loc := locations[0]
		assert.Equal(t, "loc1", loc.ID)
		assert.Equal(t, "Test Location", loc.Name)
		assert.Equal(t, "A test location", loc.Description)
		assert.Equal(t, "Detailed information", loc.Details)
	})

	// Verify codex migrated correctly
	t.Run("Codex Migration", func(t *testing.T) {
		entries, err := fs.GetCodexEntries()
		require.NoError(t, err)
		assert.Len(t, entries, 1)

		entry := entries[0]
		assert.Equal(t, "codex1", entry.ID)
		assert.Equal(t, "Test Codex Entry", entry.Title)
		assert.Equal(t, "Lore", entry.Category)
		assert.Equal(t, "Test content for codex", entry.Content)
		assert.Contains(t, entry.Tags, "test")
		assert.Contains(t, entry.Tags, "migration")
	})

	// Verify version history created
	t.Run("Version History", func(t *testing.T) {
		versions, err := fs.GetVersions(EntityCharacters, "char1")
		require.NoError(t, err)
		assert.Len(t, versions, 1) // Should have one create version from migration

		version := versions[0]
		assert.Equal(t, OperationCreate, version.Operation)
		assert.True(t, version.Active)
		assert.Equal(t, "char1", version.EntityID)
	})

	// Verify storage stats
	t.Run("Storage Stats", func(t *testing.T) {
		stats, err := fs.GetStorageStats()
		require.NoError(t, err)

		assert.Equal(t, 2, stats.EntitiesByType[EntityCharacters])
		assert.Equal(t, 1, stats.EntitiesByType[EntityLocations])
		assert.Equal(t, 1, stats.EntitiesByType[EntityCodex])
		assert.Equal(t, 4, stats.TotalFiles) // 2 chars + 1 loc + 1 codex
	})
}

func TestMigration_WithBackup(t *testing.T) {
	// Setup old JSON format directory
	oldDir := filepath.Join(os.TempDir(), "old_backup_test_"+time.Now().Format("20060102_150405"))
	
	defer os.RemoveAll(filepath.Dir(oldDir)) // Clean up parent directory

	err := os.MkdirAll(oldDir, 0755)
	require.NoError(t, err)

	// Create test data
	testData := []models.Character{
		{ID: "test1", Name: "Test Character", Description: "Test"},
	}
	writeTestJSON(t, oldDir, "characters.json", testData)

	// Create migration instance
	newDir := filepath.Join(os.TempDir(), "new_backup_test_"+time.Now().Format("20060102_150405"))
	migration := NewMigration(oldDir, newDir)

	// Create backup
	err = migration.CreateBackup()
	require.NoError(t, err)

	// Verify backup was created
	backupDirs, err := filepath.Glob(oldDir + "_backup_*")
	require.NoError(t, err)
	assert.Len(t, backupDirs, 1)

	// Verify backup contains original file
	backupFile := filepath.Join(backupDirs[0], "characters.json")
	assert.FileExists(t, backupFile)

	// Verify backup content matches original
	originalContent, err := os.ReadFile(filepath.Join(oldDir, "characters.json"))
	require.NoError(t, err)
	
	backupContent, err := os.ReadFile(backupFile)
	require.NoError(t, err)
	
	assert.Equal(t, originalContent, backupContent)
}

func TestMigration_MissingFiles(t *testing.T) {
	// Setup directories
	oldDir := filepath.Join(os.TempDir(), "missing_files_test_"+time.Now().Format("20060102_150405"))
	newDir := filepath.Join(os.TempDir(), "new_missing_test_"+time.Now().Format("20060102_150405"))
	
	defer func() {
		os.RemoveAll(oldDir)
		os.RemoveAll(newDir)
	}()

	err := os.MkdirAll(oldDir, 0755)
	require.NoError(t, err)

	// Only create characters.json, leave other files missing
	testCharacters := []models.Character{
		{ID: "char1", Name: "Only Character", Description: "Test"},
	}
	writeTestJSON(t, oldDir, "characters.json", testCharacters)

	// Create folder storage and migrate
	fs := NewFolderStorage(newDir)
	err = fs.MigrateFromJSON(oldDir)
	require.NoError(t, err)

	// Verify only characters were migrated
	characters, err := fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 1)

	locations, err := fs.GetLocations()
	require.NoError(t, err)
	assert.Len(t, locations, 0)
}

func TestMigration_EmptyFiles(t *testing.T) {
	// Setup directories
	oldDir := filepath.Join(os.TempDir(), "empty_files_test_"+time.Now().Format("20060102_150405"))
	newDir := filepath.Join(os.TempDir(), "new_empty_test_"+time.Now().Format("20060102_150405"))
	
	defer func() {
		os.RemoveAll(oldDir)
		os.RemoveAll(newDir)
	}()

	err := os.MkdirAll(oldDir, 0755)
	require.NoError(t, err)

	// Create empty JSON files
	writeTestJSON(t, oldDir, "characters.json", []models.Character{})
	writeTestJSON(t, oldDir, "locations.json", []models.Location{})

	// Create folder storage and migrate
	fs := NewFolderStorage(newDir)
	err = fs.MigrateFromJSON(oldDir)
	require.NoError(t, err)

	// Verify empty results
	characters, err := fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 0)

	locations, err := fs.GetLocations()
	require.NoError(t, err)
	assert.Len(t, locations, 0)
}

func TestMigration_ValidationAfterMigration(t *testing.T) {
	// Setup directories
	oldDir := filepath.Join(os.TempDir(), "validation_test_"+time.Now().Format("20060102_150405"))
	newDir := filepath.Join(os.TempDir(), "new_validation_test_"+time.Now().Format("20060102_150405"))
	
	defer func() {
		os.RemoveAll(oldDir)
		os.RemoveAll(newDir)
	}()

	err := os.MkdirAll(oldDir, 0755)
	require.NoError(t, err)

	// Create comprehensive test data
	testCharacters := make([]models.Character, 5)
	for i := 0; i < 5; i++ {
		testCharacters[i] = models.Character{
			ID:          fmt.Sprintf("char%d", i+1),
			Name:        fmt.Sprintf("Character %d", i+1),
			Description: fmt.Sprintf("Description %d", i+1),
		}
	}

	testLocations := make([]models.Location, 3)
	for i := 0; i < 3; i++ {
		testLocations[i] = models.Location{
			ID:          fmt.Sprintf("loc%d", i+1),
			Name:        fmt.Sprintf("Location %d", i+1),
			Description: fmt.Sprintf("Description %d", i+1),
		}
	}

	writeTestJSON(t, oldDir, "characters.json", testCharacters)
	writeTestJSON(t, oldDir, "locations.json", testLocations)

	// Create migration and validate
	migration := NewMigration(oldDir, newDir)
	err = migration.MigrateAll()
	require.NoError(t, err)

	err = migration.ValidateMigration()
	require.NoError(t, err)

	// Additional validation checks
	stats, err := migration.storage.GetStorageStats()
	require.NoError(t, err)

	assert.Equal(t, 5, stats.EntitiesByType[EntityCharacters])
	assert.Equal(t, 3, stats.EntitiesByType[EntityLocations])
	assert.Equal(t, 8, stats.TotalFiles) // 5 chars + 3 locs
}

// Helper function to write test data to JSON files
func writeTestJSON(t *testing.T, dir, filename string, data interface{}) {
	filePath := filepath.Join(dir, filename)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	require.NoError(t, err)
	
	err = os.WriteFile(filePath, jsonData, 0644)
	require.NoError(t, err)
}
