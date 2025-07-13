package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestStorage(t *testing.T) (*FolderStorage, string) {
	tempDir := filepath.Join(os.TempDir(), "ainovelprompter_test_"+time.Now().Format("20060102_150405"))
	fs := NewFolderStorage(tempDir)
	
	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})
	
	return fs, tempDir
}

func TestFolderStorage_CharacterOperations(t *testing.T) {
	fs, _ := setupTestStorage(t)

	// Test character creation
	t.Run("Create Character", func(t *testing.T) {
		character := &models.Character{
			Name:        "Test Character",
			Description: "A test character",
			Traits:      map[string]string{"personality": "brave"},
			Notes:       "Test notes",
		}

		version, err := fs.Create(EntityCharacters, character)
		require.NoError(t, err)
		assert.NotEmpty(t, version.ID)
		assert.NotEmpty(t, character.ID)
		assert.Equal(t, OperationCreate, version.Operation)
		assert.True(t, version.Active)
		assert.NotZero(t, version.Timestamp)
	})

	// Test character retrieval
	t.Run("Get Character", func(t *testing.T) {
		// Create character first
		character := &models.Character{
			Name:        "Retrieve Test",
			Description: "Character for retrieval test",
		}
		
		_, err := fs.Create(EntityCharacters, character)
		require.NoError(t, err)

		// Retrieve character
		retrieved, err := fs.GetLatest(EntityCharacters, character.ID)
		require.NoError(t, err)
		
		retrievedChar := retrieved.(*models.Character)
		assert.Equal(t, character.Name, retrievedChar.Name)
		assert.Equal(t, character.Description, retrievedChar.Description)
		assert.Equal(t, character.ID, retrievedChar.ID)
	})

	// Test character update
	t.Run("Update Character", func(t *testing.T) {
		// Create character first
		character := &models.Character{
			Name:        "Update Test",
			Description: "Original description",
		}
		
		createVersion, err := fs.Create(EntityCharacters, character)
		require.NoError(t, err)

		// Update character
		character.Description = "Updated description"
		character.Notes = "Added notes"
		
		updateVersion, err := fs.Update(EntityCharacters, character.ID, character)
		require.NoError(t, err)
		
		assert.Equal(t, OperationUpdate, updateVersion.Operation)
		assert.True(t, updateVersion.Active)
		assert.True(t, updateVersion.Timestamp.After(createVersion.Timestamp))

		// Verify previous version is inactive
		versions, err := fs.GetVersions(EntityCharacters, character.ID)
		require.NoError(t, err)
		assert.Len(t, versions, 2)
		
		// Find the create version and verify it's inactive
		for _, v := range versions {
			if v.Operation == OperationCreate {
				assert.False(t, v.Active)
			}
		}
	})

	// Test character deletion (soft delete)
	t.Run("Delete Character", func(t *testing.T) {
		// Create character first
		character := &models.Character{
			Name:        "Delete Test",
			Description: "Character for deletion test",
		}
		
		_, err := fs.Create(EntityCharacters, character)
		require.NoError(t, err)

		// Delete character
		deleteVersion, err := fs.Delete(EntityCharacters, character.ID)
		require.NoError(t, err)
		
		assert.Equal(t, OperationDelete, deleteVersion.Operation)
		assert.False(t, deleteVersion.Active) // Delete markers are not active

		// Verify character cannot be retrieved
		_, err = fs.GetLatest(EntityCharacters, character.ID)
		assert.Error(t, err)
	})
}

func TestFolderStorage_VersionManagement(t *testing.T) {
	fs, _ := setupTestStorage(t)

	character := &models.Character{
		Name:        "Version Test",
		Description: "Original description",
	}

	// Create initial version
	_, err := fs.Create(EntityCharacters, character)
	require.NoError(t, err)
	
	// Wait a moment to ensure different timestamps
	time.Sleep(10 * time.Millisecond)

	// Update to create second version
	character.Description = "Updated description"
	v2, err := fs.Update(EntityCharacters, character.ID, character)
	require.NoError(t, err)
	
	// Wait a moment to ensure different timestamps
	time.Sleep(10 * time.Millisecond)

	// Update again to create third version
	character.Description = "Final description"
	character.Notes = "Added notes"
	_, err = fs.Update(EntityCharacters, character.ID, character)
	require.NoError(t, err)

	t.Run("Get Version History", func(t *testing.T) {
		versions, err := fs.GetVersions(EntityCharacters, character.ID)
		require.NoError(t, err)
		assert.Len(t, versions, 3)

		// Verify versions are sorted by timestamp (newest first)
		assert.True(t, versions[0].Timestamp.After(versions[1].Timestamp))
		assert.True(t, versions[1].Timestamp.After(versions[2].Timestamp))

		// Verify only latest version is active
		activeCount := 0
		for _, v := range versions {
			if v.Active {
				activeCount++
			}
		}
		assert.Equal(t, 1, activeCount)
	})

	t.Run("Get Specific Version", func(t *testing.T) {
		// Get the second version (v2)
		entity, err := fs.GetVersion(EntityCharacters, character.ID, v2.Timestamp)
		require.NoError(t, err)
		
		versionChar := entity.(*models.Character)
		assert.Equal(t, "Updated description", versionChar.Description)
		assert.Empty(t, versionChar.Notes) // Notes were added in v3
	})

	t.Run("Restore Version", func(t *testing.T) {
		// Restore to v2 (should create v4)
		restoreVersion, err := fs.RestoreVersion(EntityCharacters, character.ID, v2.Timestamp)
		require.NoError(t, err)
		
		assert.Equal(t, OperationUpdate, restoreVersion.Operation)
		assert.True(t, restoreVersion.Active)

		// Verify content matches v2
		current, err := fs.GetLatest(EntityCharacters, character.ID)
		require.NoError(t, err)
		
		currentChar := current.(*models.Character)
		assert.Equal(t, "Updated description", currentChar.Description)
		assert.Empty(t, currentChar.Notes)

		// Verify we now have 4 versions
		versions, err := fs.GetVersions(EntityCharacters, character.ID)
		require.NoError(t, err)
		assert.Len(t, versions, 4)
	})
}

func TestFolderStorage_LegacyInterface(t *testing.T) {
	fs, _ := setupTestStorage(t)

	// Create test characters using new interface
	char1 := &models.Character{Name: "Character 1", Description: "First character"}
	char2 := &models.Character{Name: "Character 2", Description: "Second character"}
	
	_, err := fs.Create(EntityCharacters, char1)
	require.NoError(t, err)
	
	_, err = fs.Create(EntityCharacters, char2)
	require.NoError(t, err)

	t.Run("GetCharacters", func(t *testing.T) {
		characters, err := fs.GetCharacters()
		require.NoError(t, err)
		assert.Len(t, characters, 2)
		
		names := []string{characters[0].Name, characters[1].Name}
		assert.Contains(t, names, "Character 1")
		assert.Contains(t, names, "Character 2")
	})

	t.Run("GetCharacterByID", func(t *testing.T) {
		character, err := fs.GetCharacterByID(char1.ID)
		require.NoError(t, err)
		assert.Equal(t, char1.Name, character.Name)
		assert.Equal(t, char1.ID, character.ID)
	})

	t.Run("SearchCharacters", func(t *testing.T) {
		results, err := fs.SearchCharacters("First")
		require.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "Character 1", results[0].Name)
	})

	t.Run("UpdateCharacter", func(t *testing.T) {
		char1.Description = "Updated via legacy interface"
		err := fs.UpdateCharacter(char1)
		require.NoError(t, err)

		updated, err := fs.GetCharacterByID(char1.ID)
		require.NoError(t, err)
		assert.Equal(t, "Updated via legacy interface", updated.Description)
	})

	t.Run("DeleteCharacter", func(t *testing.T) {
		err := fs.DeleteCharacter(char1.ID)
		require.NoError(t, err)

		_, err = fs.GetCharacterByID(char1.ID)
		assert.Error(t, err)

		// Verify only one character remains
		characters, err := fs.GetCharacters()
		require.NoError(t, err)
		assert.Len(t, characters, 1)
		assert.Equal(t, "Character 2", characters[0].Name)
	})
}

func TestFolderStorage_FileNaming(t *testing.T) {
	fs, tempDir := setupTestStorage(t)

	character := &models.Character{
		Name:        "Test Character With Spaces & Symbols!",
		Description: "Testing file naming",
	}

	_, err := fs.Create(EntityCharacters, character)
	require.NoError(t, err)

	// Check that file was created with proper naming
	charDir := filepath.Join(tempDir, EntityCharacters)
	files, err := os.ReadDir(charDir)
	require.NoError(t, err)
	assert.Len(t, files, 1)

	filename := files[0].Name()
	
	// Should contain slugified name
	assert.Contains(t, filename, "test_character_with_spaces")
	assert.Contains(t, filename, "_create.json")
	
	// Should not contain special characters
	assert.NotContains(t, filename, " ")
	assert.NotContains(t, filename, "&")
	assert.NotContains(t, filename, "!")
}

func TestFolderStorage_StorageStats(t *testing.T) {
	fs, _ := setupTestStorage(t)

	// Create test data
	char := &models.Character{Name: "Stats Test", Description: "Test character"}
	loc := &models.Location{Name: "Test Location", Description: "Test location"}
	
	_, err := fs.Create(EntityCharacters, char)
	require.NoError(t, err)
	
	_, err = fs.Create(EntityLocations, loc)
	require.NoError(t, err)

	// Update character to create more versions
	char.Description = "Updated description"
	_, err = fs.Update(EntityCharacters, char.ID, char)
	require.NoError(t, err)

	stats, err := fs.GetStorageStats()
	require.NoError(t, err)

	assert.Equal(t, 3, stats.TotalFiles) // 2 creates + 1 update
	assert.Equal(t, 2, stats.EntitiesByType[EntityCharacters]) // 1 character (despite 2 versions)
	assert.Equal(t, 1, stats.EntitiesByType[EntityLocations])  // 1 location
	assert.Equal(t, 2, stats.VersionsByType[EntityCharacters]) // 2 versions for character
	assert.Equal(t, 1, stats.VersionsByType[EntityLocations])  // 1 version for location
	assert.Greater(t, stats.TotalSize, int64(0))
}

func TestFolderStorage_CleanupOldVersions(t *testing.T) {
	fs, _ := setupTestStorage(t)

	character := &models.Character{
		Name:        "Cleanup Test",
		Description: "Test character for cleanup",
	}

	// Create and update character multiple times
	_, err := fs.Create(EntityCharacters, character)
	require.NoError(t, err)

	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Millisecond) // Ensure different timestamps
		character.Description = fmt.Sprintf("Update %d", i+1)
		_, err = fs.Update(EntityCharacters, character.ID, character)
		require.NoError(t, err)
	}

	// Verify we have 4 versions
	versions, err := fs.GetVersions(EntityCharacters, character.ID)
	require.NoError(t, err)
	assert.Len(t, versions, 4)

	// Cleanup old versions (retain 0 days = cleanup all inactive)
	err = fs.CleanupOldVersions(EntityCharacters, 0)
	require.NoError(t, err)

	// Should still have at least the active version
	stats, err := fs.GetStorageStats()
	require.NoError(t, err)
	assert.GreaterOrEqual(t, stats.VersionsByType[EntityCharacters], 1)
}

func TestFolderStorage_SetDataDirectory(t *testing.T) {
	fs, _ := setupTestStorage(t)

	// Create character in original directory
	character := &models.Character{Name: "Migration Test", Description: "Test character"}
	_, err := fs.Create(EntityCharacters, character)
	require.NoError(t, err)

	// Set new data directory
	newDir := filepath.Join(os.TempDir(), "new_data_dir_"+time.Now().Format("20060102_150405"))
	defer os.RemoveAll(newDir)

	err = fs.SetDataDirectory(newDir)
	require.NoError(t, err)

	// Verify directory changed
	assert.Equal(t, newDir, fs.GetDataDirectory())

	// New directory should be empty (no automatic migration)
	characters, err := fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 0)
}

// Benchmark tests
func BenchmarkFolderStorage_CreateCharacter(b *testing.B) {
	fs, tempDir := setupBenchStorage(b)
	defer os.RemoveAll(tempDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		character := &models.Character{
			Name:        fmt.Sprintf("Benchmark Character %d", i),
			Description: "Benchmark test character",
		}
		_, err := fs.Create(EntityCharacters, character)
		if err != nil {
			b.Fatalf("Failed to create character: %v", err)
		}
	}
}

func BenchmarkFolderStorage_GetCharacters(b *testing.B) {
	fs, tempDir := setupBenchStorage(b)
	defer os.RemoveAll(tempDir)

	// Create test data
	for i := 0; i < 100; i++ {
		character := &models.Character{
			Name:        fmt.Sprintf("Character %d", i),
			Description: "Test character",
		}
		_, err := fs.Create(EntityCharacters, character)
		if err != nil {
			b.Fatalf("Failed to create character: %v", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fs.GetCharacters()
		if err != nil {
			b.Fatalf("Failed to get characters: %v", err)
		}
	}
}

func setupBenchStorage(b *testing.B) (*FolderStorage, string) {
	tempDir := filepath.Join(os.TempDir(), "ainovelprompter_bench_"+time.Now().Format("20060102_150405"))
	fs := NewFolderStorage(tempDir)
	return fs, tempDir
}
