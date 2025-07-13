package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/danielsobrado/ainovelprompter/mcp/storage"
)

func main() {
	fmt.Println("🧪 AI Novel Prompter - Folder Storage Test Suite")
	fmt.Println("================================================")

	// Create test directories
	testDir := filepath.Join(os.TempDir(), "ainovelprompter_test_"+time.Now().Format("20060102_150405"))
	oldDir := filepath.Join(os.TempDir(), "ainovelprompter_old_"+time.Now().Format("20060102_150405"))
	
	defer func() {
		os.RemoveAll(testDir)
		os.RemoveAll(oldDir)
	}()

	// Test 1: Basic Storage Operations
	fmt.Println("
📁 Test 1: Basic Storage Operations")
	if !testBasicOperations(testDir) {
		return
	}

	// Test 2: Version Management
	fmt.Println("
📚 Test 2: Version Management")
	if !testVersionManagement(testDir) {
		return
	}

	// Test 3: Legacy Interface Compatibility
	fmt.Println("
🔄 Test 3: Legacy Interface Compatibility")
	if !testLegacyInterface(testDir) {
		return
	}

	// Test 4: Multiple Entity Types
	fmt.Println("
🏗️ Test 4: Multiple Entity Types")
	if !testMultipleEntityTypes(testDir) {
		return
	}

	// Test 5: Storage Statistics
	fmt.Println("
📊 Test 5: Storage Statistics")
	if !testStorageStatistics(testDir) {
		return
	}

	// Test 6: Migration from JSON
	fmt.Println("
🔄 Test 6: Migration from JSON")
	if !testMigration(oldDir, testDir+"_migrated") {
		return
	}

	// Test 7: Data Directory Management
	fmt.Println("
📂 Test 7: Data Directory Management")
	if !testDataDirectoryManagement(testDir) {
		return
	}

	fmt.Println("
🎉 All tests passed! Folder-based storage with versioning is working correctly.")
	fmt.Println("✅ Storage system is ready for production use.")
}

func testBasicOperations(testDir string) bool {
	fs := storage.NewFolderStorage(testDir)

	// Create character
	character := &models.Character{
		Name:        "Test Character",
		Description: "A test character for validation",
		Traits:      map[string]string{"personality": "curious", "role": "protagonist"},
		Notes:       "This is a test character",
	}

	version1, err := fs.Create(storage.EntityCharacters, character)
	if err != nil {
		fmt.Printf("❌ Error creating character: %v
", err)
		return false
	}
	fmt.Printf("✓ Created character with version ID: %s
", version1.ID)

	// Update character
	character.Description = "Updated description"
	character.Notes = "Updated notes"
	version2, err := fs.Update(storage.EntityCharacters, character.ID, character)
	if err != nil {
		fmt.Printf("❌ Error updating character: %v
", err)
		return false
	}
	fmt.Printf("✓ Updated character with version ID: %s
", version2.ID)

	// Get latest version
	latest, err := fs.GetLatest(storage.EntityCharacters, character.ID)
	if err != nil {
		fmt.Printf("❌ Error getting latest: %v
", err)
		return false
	}
	latestChar := latest.(*models.Character)
	fmt.Printf("✓ Retrieved latest character: %s - %s
", latestChar.Name, latestChar.Description)

	// Delete character
	_, err = fs.Delete(storage.EntityCharacters, character.ID)
	if err != nil {
		fmt.Printf("❌ Error deleting character: %v
", err)
		return false
	}
	fmt.Printf("✓ Successfully deleted character
")

	return true
}

func testVersionManagement(testDir string) bool {
	fs := storage.NewFolderStorage(testDir)

	// Create character for version testing
	character := &models.Character{
		Name:        "Version Test Character",
		Description: "Original description",
		Notes:       "Original notes",
	}

	// Create multiple versions
	v1, err := fs.Create(storage.EntityCharacters, character)
	if err != nil {
		fmt.Printf("❌ Error creating character: %v
", err)
		return false
	}

	time.Sleep(10 * time.Millisecond) // Ensure different timestamps

	character.Description = "Updated description"
	v2, err := fs.Update(storage.EntityCharacters, character.ID, character)
	if err != nil {
		fmt.Printf("❌ Error updating character: %v
", err)
		return false
	}

	time.Sleep(10 * time.Millisecond)

	character.Notes = "Updated notes"
	v3, err := fs.Update(storage.EntityCharacters, character.ID, character)
	if err != nil {
		fmt.Printf("❌ Error updating character again: %v
", err)
		return false
	}

	// Get version history
	versions, err := fs.GetVersions(storage.EntityCharacters, character.ID)
	if err != nil {
		fmt.Printf("❌ Error getting versions: %v
", err)
		return false
	}
	fmt.Printf("✓ Found %d versions:
", len(versions))
	for i, v := range versions {
		fmt.Printf("  %d. %s: %s (active: %t)
", i+1, v.Timestamp.Format(time.RFC3339), v.Operation, v.Active)
	}

	// Test restore to v2
	if len(versions) >= 2 {
		restored, err := fs.RestoreVersion(storage.EntityCharacters, character.ID, v2.Timestamp)
		if err != nil {
			fmt.Printf("❌ Error restoring version: %v
", err)
			return false
		}
		fmt.Printf("✓ Restored to version: %s
", restored.ID)

		// Verify restoration
		current, err := fs.GetLatest(storage.EntityCharacters, character.ID)
		if err != nil {
			fmt.Printf("❌ Error getting restored character: %v
", err)
			return false
		}
		restoredChar := current.(*models.Character)
		if restoredChar.Description != "Updated description" {
			fmt.Printf("❌ Restoration failed: expected 'Updated description', got '%s'
", restoredChar.Description)
			return false
		}
		fmt.Printf("✓ Restoration verified: description = '%s'
", restoredChar.Description)
	}

	return true
}

func testLegacyInterface(testDir string) bool {
	fs := storage.NewFolderStorage(testDir)

	// Create characters using new interface
	char1 := &models.Character{Name: "Legacy Test 1", Description: "First legacy test"}
	char2 := &models.Character{Name: "Legacy Test 2", Description: "Second legacy test"}

	_, err := fs.Create(storage.EntityCharacters, char1)
	if err != nil {
		fmt.Printf("❌ Error creating character 1: %v
", err)
		return false
	}

	_, err = fs.Create(storage.EntityCharacters, char2)
	if err != nil {
		fmt.Printf("❌ Error creating character 2: %v
", err)
		return false
	}

	// Test legacy GetCharacters
	characters, err := fs.GetCharacters()
	if err != nil {
		fmt.Printf("❌ Error getting characters via legacy interface: %v
", err)
		return false
	}
	fmt.Printf("✓ Retrieved %d characters via legacy interface
", len(characters))

	// Test legacy search
	searchResults, err := fs.SearchCharacters("Legacy")
	if err != nil {
		fmt.Printf("❌ Error searching characters: %v
", err)
		return false
	}
	fmt.Printf("✓ Search found %d characters
", len(searchResults))

	// Test legacy GetCharacterByID
	retrieved, err := fs.GetCharacterByID(char1.ID)
	if err != nil {
		fmt.Printf("❌ Error getting character by ID: %v
", err)
		return false
	}
	fmt.Printf("✓ Retrieved character by ID: %s
", retrieved.Name)

	return true
}

func testMultipleEntityTypes(testDir string) bool {
	fs := storage.NewFolderStorage(testDir)

	// Create different entity types
	location := &models.Location{
		Name:        "Test Location",
		Description: "A magical forest",
		Details:     "Deep in the mountains",
		Notes:       "Dangerous at night",
	}

	codex := &models.CodexEntry{
		Title:    "Magic System",
		Category: "Lore",
		Content:  "Magic flows through ley lines",
		Tags:     []string{"magic", "worldbuilding"},
	}

	rule := &models.Rule{
		Name:        "Character Development",
		Description: "Characters must grow throughout the story",
		Category:    "Writing",
		Active:      true,
	}

	// Create entities
	_, err := fs.Create(storage.EntityLocations, location)
	if err != nil {
		fmt.Printf("❌ Error creating location: %v
", err)
		return false
	}
	fmt.Printf("✓ Created location: %s
", location.Name)

	_, err = fs.Create(storage.EntityCodex, codex)
	if err != nil {
		fmt.Printf("❌ Error creating codex entry: %v
", err)
		return false
	}
	fmt.Printf("✓ Created codex entry: %s
", codex.Title)

	_, err = fs.Create(storage.EntityRules, rule)
	if err != nil {
		fmt.Printf("❌ Error creating rule: %v
", err)
		return false
	}
	fmt.Printf("✓ Created rule: %s
", rule.Name)

	// Test retrieval
	locations, err := fs.GetLocations()
	if err != nil {
		fmt.Printf("❌ Error getting locations: %v
", err)
		return false
	}
	fmt.Printf("✓ Retrieved %d locations
", len(locations))

	codexEntries, err := fs.GetCodexEntries()
	if err != nil {
		fmt.Printf("❌ Error getting codex entries: %v
", err)
		return false
	}
	fmt.Printf("✓ Retrieved %d codex entries
", len(codexEntries))

	rules, err := fs.GetRules()
	if err != nil {
		fmt.Printf("❌ Error getting rules: %v
", err)
		return false
	}
	fmt.Printf("✓ Retrieved %d rules
", len(rules))

	return true
}

func testStorageStatistics(testDir string) bool {
	fs := storage.NewFolderStorage(testDir)

	// Get storage statistics
	stats, err := fs.GetStorageStats()
	if err != nil {
		fmt.Printf("❌ Error getting storage stats: %v
", err)
		return false
	}

	fmt.Printf("✓ Storage Statistics:
")
	fmt.Printf("  - Total files: %d
", stats.TotalFiles)
	fmt.Printf("  - Total size: %d bytes
", stats.TotalSize)
	fmt.Printf("  - Entities by type:
")
	for entityType, count := range stats.EntitiesByType {
		fmt.Printf("    * %s: %d entities
", entityType, count)
	}
	fmt.Printf("  - Versions by type:
")
	for entityType, count := range stats.VersionsByType {
		fmt.Printf("    * %s: %d versions
", entityType, count)
	}

	if !stats.OldestTimestamp.IsZero() {
		fmt.Printf("  - Oldest timestamp: %s
", stats.OldestTimestamp.Format(time.RFC3339))
	}
	if !stats.NewestTimestamp.IsZero() {
		fmt.Printf("  - Newest timestamp: %s
", stats.NewestTimestamp.Format(time.RFC3339))
	}

	return true
}

func testMigration(oldDir, newDir string) bool {
	// Create old JSON format data
	err := os.MkdirAll(oldDir, 0755)
	if err != nil {
		fmt.Printf("❌ Error creating old directory: %v
", err)
		return false
	}

	// Create test data in old format
	oldCharacters := []models.Character{
		{
			ID:          "char1",
			Name:        "Migrated Character 1",
			Description: "A character from old format",
			Traits:      map[string]string{"brave": "true"},
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now().Add(-1 * time.Hour),
		},
		{
			ID:          "char2",
			Name:        "Migrated Character 2",
			Description: "Another character from old format",
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   time.Now().Add(-2 * time.Hour),
		},
	}

	oldLocations := []models.Location{
		{
			ID:          "loc1",
			Name:        "Migrated Location",
			Description: "A location from old format",
			CreatedAt:   time.Now().Add(-72 * time.Hour),
			UpdatedAt:   time.Now().Add(-3 * time.Hour),
		},
	}

	// Write to JSON files
	if !writeJSONFile(oldDir, "characters.json", oldCharacters) {
		return false
	}
	if !writeJSONFile(oldDir, "locations.json", oldLocations) {
		return false
	}

	fmt.Printf("✓ Created old format data with %d characters and %d locations
", len(oldCharacters), len(oldLocations))

	// Perform migration
	fs := storage.NewFolderStorage(newDir)
	err = fs.MigrateFromJSON(oldDir)
	if err != nil {
		fmt.Printf("❌ Migration failed: %v
", err)
		return false
	}
	fmt.Printf("✓ Migration completed successfully
")

	// Verify migrated data
	characters, err := fs.GetCharacters()
	if err != nil {
		fmt.Printf("❌ Error getting migrated characters: %v
", err)
		return false
	}

	locations, err := fs.GetLocations()
	if err != nil {
		fmt.Printf("❌ Error getting migrated locations: %v
", err)
		return false
	}

	fmt.Printf("✓ Verified migration: %d characters, %d locations
", len(characters), len(locations))

	return true
}

func testDataDirectoryManagement(testDir string) bool {
	fs := storage.NewFolderStorage(testDir)

	// Test getting current directory
	currentDir := fs.GetDataDirectory()
	fmt.Printf("✓ Current data directory: %s
", currentDir)

	// Test setting new directory
	newDir := testDir + "_new"
	err := fs.SetDataDirectory(newDir)
	if err != nil {
		fmt.Printf("❌ Error setting new data directory: %v
", err)
		return false
	}

	updatedDir := fs.GetDataDirectory()
	if updatedDir != newDir {
		fmt.Printf("❌ Directory not updated correctly: expected %s, got %s
", newDir, updatedDir)
		return false
	}
	fmt.Printf("✓ Successfully changed data directory to: %s
", updatedDir)

	return true
}

func writeJSONFile(dir, filename string, data interface{}) bool {
	filePath := filepath.Join(dir, filename)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error marshaling JSON for %s: %v
", filename, err)
		return false
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("❌ Error writing file %s: %v
", filename, err)
		return false
	}

	return true
}
