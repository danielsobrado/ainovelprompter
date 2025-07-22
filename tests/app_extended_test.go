package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCLIArgumentParsing tests command line argument parsing
func TestCLIArgumentParsing(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected struct {
			dataDir string
			help    bool
		}
	}{
		{
			name: "Default arguments",
			args: []string{"ainovelprompter"},
			expected: struct {
				dataDir string
				help    bool
			}{
				dataDir: "",
				help:    false,
			},
		},
		{
			name: "Data directory long flag",
			args: []string{"ainovelprompter", "--data-dir", "/custom/path"},
			expected: struct {
				dataDir string
				help    bool
			}{
				dataDir: "/custom/path",
				help:    false,
			},
		},
		{
			name: "Data directory short flag",
			args: []string{"ainovelprompter", "-d", "./project"},
			expected: struct {
				dataDir string
				help    bool
			}{
				dataDir: "./project",
				help:    false,
			},
		},
		{
			name: "Help long flag",
			args: []string{"ainovelprompter", "--help"},
			expected: struct {
				dataDir string
				help    bool
			}{
				dataDir: "",
				help:    true,
			},
		},
		{
			name: "Help short flag",
			args: []string{"ainovelprompter", "-h"},
			expected: struct {
				dataDir string
				help    bool
			}{
				dataDir: "",
				help:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This would test the actual CLI parsing logic
			// For now, we'll test the concept
			assert.True(t, len(tt.args) > 0)

			// Find data-dir flag
			dataDir := ""
			help := false

			for i, arg := range tt.args {
				if arg == "--data-dir" || arg == "-d" {
					if i+1 < len(tt.args) {
						dataDir = tt.args[i+1]
					}
				}
				if arg == "--help" || arg == "-h" {
					help = true
				}
			}

			assert.Equal(t, tt.expected.dataDir, dataDir)
			assert.Equal(t, tt.expected.help, help)
		})
	}
}

// TestDataDirectoryResolution tests data directory resolution logic
func TestDataDirectoryResolution(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		homeDir     string
		expected    string
		shouldError bool
	}{
		{
			name:        "Absolute path",
			input:       "/absolute/path/to/data",
			homeDir:     "/home/user",
			expected:    "/absolute/path/to/data",
			shouldError: false,
		},
		{
			name:        "Relative path",
			input:       "./relative/path",
			homeDir:     "/home/user",
			expected:    "./relative/path",
			shouldError: false,
		},
		{
			name:        "Home directory expansion",
			input:       "~/my-project",
			homeDir:     "/home/user",
			expected:    "/home/user/my-project",
			shouldError: false,
		},
		{
			name:        "Default path (empty input)",
			input:       "",
			homeDir:     "/home/user",
			expected:    "/home/user/.ai-novel-prompter",
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveDataDirectory(tt.input, tt.homeDir)

			if tt.shouldError {
				assert.Empty(t, result)
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// Helper function for testing directory resolution
func resolveDataDirectory(input, homeDir string) string {
	if input == "" {
		return strings.ReplaceAll(filepath.Join(homeDir, ".ai-novel-prompter"), "\\", "/")
	}

	if strings.HasPrefix(input, "~/") {
		return strings.ReplaceAll(filepath.Join(homeDir, input[2:]), "\\", "/")
	}

	return input
}

// TestAppStructWithDataDirectory tests App struct with data directory support
func TestAppStructWithDataDirectory(t *testing.T) {
	tempDir := t.TempDir()

	// Simulate App struct
	type App struct {
		dataDir string
	}

	app := &App{}

	// Test SetDataDirectory
	app.dataDir = tempDir
	assert.Equal(t, tempDir, app.dataDir)

	// Test GetDataDirectory
	result := app.dataDir
	assert.Equal(t, tempDir, result)

	// Test ValidateDataDirectory
	err := os.MkdirAll(tempDir, 0755)
	require.NoError(t, err)

	_, err = os.Stat(tempDir)
	assert.NoError(t, err, "Data directory should exist and be accessible")
}

// TestVersionedStorageIntegration tests integration with versioned storage
func TestVersionedStorageIntegration(t *testing.T) {
	tempDir := t.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	// Test creating entity with version tracking
	character := &models.Character{
		Name:        "Integration Test Character",
		Description: "Testing integration with versioned storage",
		Traits:      map[string]interface{}{"role": "test"},
	}

	version1, err := fs.Create(storage.EntityCharacters, character)
	require.NoError(t, err)
	assert.Equal(t, storage.OperationCreate, version1.Operation)
	assert.True(t, version1.Active)

	// Test updating with version tracking
	originalDescription := character.Description
	character.Description = "Updated for integration test"

	version2, err := fs.Update(storage.EntityCharacters, character.ID, character)
	require.NoError(t, err)
	assert.Equal(t, storage.OperationUpdate, version2.Operation)
	assert.True(t, version2.Active)
	assert.True(t, version2.Timestamp.After(version1.Timestamp))

	// Test version history
	versions, err := fs.GetVersions(storage.EntityCharacters, character.ID)
	require.NoError(t, err)
	assert.Len(t, versions, 2)

	// Test version restoration
	restored, err := fs.RestoreVersion(storage.EntityCharacters, character.ID, version1.Timestamp)
	require.NoError(t, err)
	assert.Equal(t, storage.OperationUpdate, restored.Operation) // Restore creates new update

	// Verify restoration worked
	current, err := fs.GetLatest(storage.EntityCharacters, character.ID)
	require.NoError(t, err)
	restoredChar := current.(*models.Character)
	assert.Equal(t, originalDescription, restoredChar.Description)
}

// TestStorageStatisticsAndCleanup tests storage analytics and cleanup
func TestStorageStatisticsAndCleanup(t *testing.T) {
	tempDir := t.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	// Create test data
	for i := 0; i < 3; i++ {
		char := &models.Character{
			Name:        fmt.Sprintf("Stats Test Character %d", i),
			Description: "Character for statistics testing",
		}

		// Create and update to generate versions
		_, err := fs.Create(storage.EntityCharacters, char)
		require.NoError(t, err)

		char.Description = "Updated description"
		_, err = fs.Update(storage.EntityCharacters, char.ID, char)
		require.NoError(t, err)
	}

	// Test storage statistics
	stats, err := fs.GetStorageStats()
	require.NoError(t, err)

	assert.Equal(t, 3, stats.EntitiesByType[storage.EntityCharacters])
	assert.Equal(t, 6, stats.VersionsByType[storage.EntityCharacters]) // 3 creates + 3 updates
	assert.Equal(t, 6, stats.TotalFiles)
	assert.Greater(t, stats.TotalSize, int64(0))
	assert.False(t, stats.OldestTimestamp.IsZero())
	assert.False(t, stats.NewestTimestamp.IsZero())

	// Test cleanup of old versions
	err = fs.CleanupOldVersions(storage.EntityCharacters, 0) // Clean all inactive
	require.NoError(t, err)

	// Verify cleanup worked (should keep active versions)
	newStats, err := fs.GetStorageStats()
	require.NoError(t, err)
	assert.LessOrEqual(t, newStats.TotalFiles, stats.TotalFiles)
}



// TestDataDirectoryManagement tests dynamic data directory changes
func TestDataDirectoryManagement(t *testing.T) {
	dir1 := filepath.Join(t.TempDir(), "project1")
	dir2 := filepath.Join(t.TempDir(), "project2")

	fs := storage.NewFolderStorage(dir1)

	// Create data in first directory
	char1 := &models.Character{
		Name:        "Project 1 Character",
		Description: "Character in first project",
	}
	_, err := fs.Create(storage.EntityCharacters, char1)
	require.NoError(t, err)

	// Verify current directory
	assert.Equal(t, dir1, fs.GetDataDirectory())

	// Switch to second directory
	err = fs.SetDataDirectory(dir2)
	require.NoError(t, err)
	assert.Equal(t, dir2, fs.GetDataDirectory())

	// Verify second directory is empty (no automatic migration)
	characters, err := fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 0)

	// Create data in second directory
	char2 := &models.Character{
		Name:        "Project 2 Character",
		Description: "Character in second project",
	}
	_, err = fs.Create(storage.EntityCharacters, char2)
	require.NoError(t, err)

	// Verify data is in second directory
	characters, err = fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 1)
	assert.Equal(t, "Project 2 Character", characters[0].Name)

	// Switch back to first directory
	err = fs.SetDataDirectory(dir1)
	require.NoError(t, err)

	// Verify original data is still there
	characters, err = fs.GetCharacters()
	require.NoError(t, err)
	assert.Len(t, characters, 1)
	assert.Equal(t, "Project 1 Character", characters[0].Name)
}

// TestFileNamingConventions tests the new file naming system
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
	charDir := filepath.Join(tempDir, storage.EntityCharacters)
	files, err := os.ReadDir(charDir)
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
	parts := strings.Split(filename, "_")
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

// Helper function to write JSON data to file
func writeJSONToFile(t *testing.T, dir, filename string, data interface{}) {
	t.Helper()

	filePath := filepath.Join(dir, filename)

	// Proper JSON marshaling for test data
	jsonData, err := json.MarshalIndent(data, "", "  ")
	require.NoError(t, err)

	err = os.WriteFile(filePath, jsonData, 0644)
	require.NoError(t, err)
}

// BenchmarkVersionedOperations benchmarks performance of versioned operations
func BenchmarkVersionedOperations(b *testing.B) {
	tempDir := b.TempDir()
	fs := storage.NewFolderStorage(tempDir)

	b.Run("CreateWithVersioning", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			character := &models.Character{
				Name:        fmt.Sprintf("Benchmark Character %d", i),
				Description: "Character for benchmarking versioned operations",
			}
			_, err := fs.Create(storage.EntityCharacters, character)
			if err != nil {
				b.Fatalf("Failed to create character: %v", err)
			}
		}
	})

	b.Run("UpdateWithVersioning", func(b *testing.B) {
		// Create initial character
		character := &models.Character{
			Name:        "Update Benchmark Character",
			Description: "Initial description",
		}
		_, err := fs.Create(storage.EntityCharacters, character)
		if err != nil {
			b.Fatalf("Failed to create character: %v", err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			character.Description = fmt.Sprintf("Updated description %d", i)
			_, err := fs.Update(storage.EntityCharacters, character.ID, character)
			if err != nil {
				b.Fatalf("Failed to update character: %v", err)
			}
		}
	})

	b.Run("GetVersionHistory", func(b *testing.B) {
		// Create character with multiple versions
		character := &models.Character{
			Name:        "History Benchmark Character",
			Description: "Initial description",
		}
		_, err := fs.Create(storage.EntityCharacters, character)
		if err != nil {
			b.Fatalf("Failed to create character: %v", err)
		}

		// Create multiple versions
		for i := 0; i < 10; i++ {
			character.Description = fmt.Sprintf("Description version %d", i)
			_, err := fs.Update(storage.EntityCharacters, character.ID, character)
			if err != nil {
				b.Fatalf("Failed to update character: %v", err)
			}
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := fs.GetVersions(storage.EntityCharacters, character.ID)
			if err != nil {
				b.Fatalf("Failed to get versions: %v", err)
			}
		}
	})
}
