package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// EntityInfo holds common information extracted from entities
type EntityInfo struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Helper methods for FolderStorage

// extractEntityInfo extracts common information from any entity
func (fs *FolderStorage) extractEntityInfo(entity interface{}) (EntityInfo, error) {
	info := EntityInfo{}
	
	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		return info, fmt.Errorf("entity must be a struct")
	}
	
	// Extract ID
	if idField := v.FieldByName("ID"); idField.IsValid() && idField.Kind() == reflect.String {
		info.ID = idField.String()
	}
	
	// Extract Name (try multiple possible field names)
	nameFields := []string{"Name", "Title", "Label"}
	for _, fieldName := range nameFields {
		if nameField := v.FieldByName(fieldName); nameField.IsValid() && nameField.Kind() == reflect.String {
			info.Name = nameField.String()
			break
		}
	}
	
	// Extract timestamps
	if createdField := v.FieldByName("CreatedAt"); createdField.IsValid() {
		if createdField.Type() == reflect.TypeOf(time.Time{}) {
			info.CreatedAt = createdField.Interface().(time.Time)
		}
	}
	
	if updatedField := v.FieldByName("UpdatedAt"); updatedField.IsValid() {
		if updatedField.Type() == reflect.TypeOf(time.Time{}) {
			info.UpdatedAt = updatedField.Interface().(time.Time)
		}
	}
	
	// Fallback for name if not found
	if info.Name == "" {
		info.Name = fmt.Sprintf("entity_%s", info.ID)
	}
	
	return info, nil
}

// setEntityID sets the ID field of an entity
func (fs *FolderStorage) setEntityID(entity interface{}, id string) {
	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	if idField := v.FieldByName("ID"); idField.IsValid() && idField.CanSet() && idField.Kind() == reflect.String {
		idField.SetString(id)
	}
}

// setEntityTimestamps sets the timestamp fields of an entity
func (fs *FolderStorage) setEntityTimestamps(entity interface{}, createdAt, updatedAt time.Time) {
	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	if createdField := v.FieldByName("CreatedAt"); createdField.IsValid() && createdField.CanSet() {
		if createdField.Type() == reflect.TypeOf(time.Time{}) {
			createdField.Set(reflect.ValueOf(createdAt))
		}
	}
	
	if updatedField := v.FieldByName("UpdatedAt"); updatedField.IsValid() && updatedField.CanSet() {
		if updatedField.Type() == reflect.TypeOf(time.Time{}) {
			updatedField.Set(reflect.ValueOf(updatedAt))
		}
	}
}

// writeEntityToFile writes an entity to a JSON file
func (fs *FolderStorage) writeEntityToFile(filePath string, entity interface{}) error {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	
	// Create temporary file for atomic write
	tempPath := filePath + ".tmp"
	
	// Marshal entity to JSON
	data, err := json.MarshalIndent(entity, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal entity: %v", err)
	}
	
	// Write to temporary file
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write temporary file: %v", err)
	}
	
	// Atomic rename
	if err := os.Rename(tempPath, filePath); err != nil {
		os.Remove(tempPath) // Cleanup temp file
		return fmt.Errorf("failed to rename temporary file: %v", err)
	}
	
	return nil
}

// loadEntityFromFile loads an entity from a JSON file
func (fs *FolderStorage) loadEntityFromFile(filePath string, entityType string) (interface{}, error) {
	log.Printf("[STORAGE] Loading entity from file: %s (type: %s)", filePath, entityType)
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("[STORAGE] ERROR: Failed to read file %s: %v", filePath, err)
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	
	log.Printf("[STORAGE] Successfully read %d bytes from file %s", len(data), filePath)
	
	// Create appropriate entity type
	entity := fs.createEntityByType(entityType)
	if entity == nil {
		log.Printf("[STORAGE] ERROR: Unknown entity type: %s", entityType)
		return nil, fmt.Errorf("unknown entity type: %s", entityType)
	}
	
	log.Printf("[STORAGE] Created entity instance for type: %s", entityType)
	
	// Unmarshal JSON
	if err := json.Unmarshal(data, entity); err != nil {
		log.Printf("[STORAGE] ERROR: Failed to unmarshal JSON for file %s: %v", filePath, err)
		log.Printf("[STORAGE] JSON data preview (first 200 chars): %s", string(data[:min(200, len(data))]))
		return nil, fmt.Errorf("failed to unmarshal entity: %v", err)
	}
	
	log.Printf("[STORAGE] Successfully unmarshaled entity from file: %s", filePath)
	return entity, nil
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// createEntityByType creates a new entity instance based on type
func (fs *FolderStorage) createEntityByType(entityType string) interface{} {
	switch entityType {
	case EntityCharacters:
		return &models.Character{}
	case EntityLocations:
		return &models.Location{}
	case EntityCodex:
		return &models.CodexEntry{}
	case EntityRules:
		return &models.Rule{}
	case EntityChapters:
		return &models.Chapter{}
	case EntityStoryBeats:
		return &models.StoryBeats{}
	case EntityFutureNotes:
		return &models.FutureNotes{}
	case EntitySampleChapters:
		return &models.SampleChapter{}
	case EntityTaskTypes:
		return &models.TaskType{}
	case EntityProsePrompts:
		return &models.ProseImprovementPrompt{}
	default:
		return nil
	}
}

// getLatestEntity gets the latest version of an entity (internal, no locking)
func (fs *FolderStorage) getLatestEntity(entityType string, id string) (interface{}, error) {
	versions := fs.getVersionsInternal(entityType, id)
	if len(versions) == 0 {
		return nil, fmt.Errorf("entity not found")
	}
	
	// Find latest active version
	for _, version := range versions {
		if version.Active && version.Operation != OperationDelete {
			return fs.loadEntityFromFile(version.FilePath, entityType)
		}
	}
	
	return nil, fmt.Errorf("entity not found or deleted")
}

// getActiveVersionsByType gets all active versions for an entity type
func (fs *FolderStorage) getActiveVersionsByType(entityType string) []Version {
	if versions, ok := fs.indexCache[entityType]; ok {
		var active []Version
		for _, version := range versions {
			if version.Active {
				active = append(active, version)
			}
		}
		return active
	}
	return []Version{}
}

// addVersionToCache adds a version to the index cache
func (fs *FolderStorage) addVersionToCache(entityType string, version Version) {
	if fs.indexCache[entityType] == nil {
		fs.indexCache[entityType] = []Version{}
	}
	fs.indexCache[entityType] = append(fs.indexCache[entityType], version)
}

// deactivateVersions marks all versions of an entity as inactive
func (fs *FolderStorage) deactivateVersions(entityType string, entityID string) {
	if versions, ok := fs.indexCache[entityType]; ok {
		for i := range versions {
			if versions[i].EntityID == entityID {
				versions[i].Active = false
			}
		}
	}
}

// rebuildIndexCache rebuilds the index cache by scanning all files
func (fs *FolderStorage) rebuildIndexCache() {
	fs.indexCache = make(map[string][]Version)
	
	entityTypes := []string{
		EntityCharacters, EntityLocations, EntityCodex, EntityRules,
		EntityChapters, EntityStoryBeats, EntityFutureNotes,
		EntitySampleChapters, EntityTaskTypes, EntityProsePrompts,
	}
	
	for _, entityType := range entityTypes {
		fs.scanEntityDirectory(entityType)
	}
}

// scanEntityDirectory scans a directory and builds version index
func (fs *FolderStorage) scanEntityDirectory(entityType string) {
	dirPath := filepath.Join(fs.basePath, entityType)
	
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return // Directory doesn't exist or can't be read
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			// This is an entity directory (desktop app format)
			entityID := entry.Name()
			entityDirPath := filepath.Join(dirPath, entityID)
			
			// Scan files in entity directory
			entityFiles, err := os.ReadDir(entityDirPath)
			if err != nil {
				continue
			}
			
			for _, file := range entityFiles {
				if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
					continue
				}
				
				// Parse timestamp from filename: 2025-07-15T07-44-53.706+04-00.json
				timestampStr := strings.TrimSuffix(file.Name(), ".json")
				
				// Parse timestamp from desktop app format: 2025-07-15T07-44-53.706+04-00
				var timestamp time.Time
				var err error
				
				// Convert desktop app format to proper RFC3339
				// Replace the last "-" in timezone with ":"
				rfc3339Str := timestampStr
				if len(timestampStr) >= 6 {
					// Find timezone part (last 6 characters should be +04-00 or -05-00)
					timezoneStart := len(timestampStr) - 6
					if timestampStr[timezoneStart-1] == '+' || timestampStr[timezoneStart-1] == '-' {
						timezonePart := timestampStr[timezoneStart:]
						if len(timezonePart) == 6 && timezonePart[2] == '-' {
							// Replace the dash in timezone with colon: 04-00 -> 04:00
							fixedTimezone := timezonePart[:2] + ":" + timezonePart[3:]
							rfc3339Str = timestampStr[:timezoneStart] + fixedTimezone
						}
					}
				}
				
				// Parse RFC3339 format
				timestamp, err = time.Parse(time.RFC3339Nano, rfc3339Str)
				if err != nil {
					// Try without nanoseconds
					timestamp, err = time.Parse(time.RFC3339, rfc3339Str)
					if err != nil {
						continue // Invalid timestamp
					}
				}
				
				filePath := filepath.Join(entityDirPath, file.Name())
				
				// Determine operation based on file content or default to update
				operation := OperationUpdate
				if fs.isFirstVersionForEntity(entityType, entityID, timestamp) {
					operation = OperationCreate
				}
				
				version := Version{
					ID:        fmt.Sprintf("%s_%s_%s", entityID, timestampStr, operation),
					EntityID:  entityID,
					Timestamp: timestamp,
					Operation: operation,
					FilePath:  filePath,
					Active:    true, // Will be corrected by deactivation logic
				}
				
				fs.addVersionToCache(entityType, version)
			}
		}
	}
	
	// Fix active status - only latest non-delete version should be active
	fs.fixActiveStatus(entityType)
}

// extractEntityIDFromFile extracts the entity ID from a file
func (fs *FolderStorage) extractEntityIDFromFile(filePath string, entityType string) string {
	entity, err := fs.loadEntityFromFile(filePath, entityType)
	if err != nil {
		return ""
	}
	
	info, err := fs.extractEntityInfo(entity)
	if err != nil {
		return ""
	}
	
	return info.ID
}

// fixActiveStatus ensures only the latest non-delete version is active
func (fs *FolderStorage) fixActiveStatus(entityType string) {
	if versions, ok := fs.indexCache[entityType]; ok {
		// Group by entity ID
		entityVersions := make(map[string][]int) // entity ID -> indices in cache
		for i, version := range versions {
			entityVersions[version.EntityID] = append(entityVersions[version.EntityID], i)
		}
		
		// For each entity, mark only latest non-delete as active
		for _, indices := range entityVersions {
			// First, mark all versions of this entity as inactive
			for _, idx := range indices {
				fs.indexCache[entityType][idx].Active = false
			}
			
			// Find the latest non-delete version
			latestIdx := -1
			latestTime := time.Time{}
			
			for _, idx := range indices {
				version := fs.indexCache[entityType][idx]
				if version.Operation != OperationDelete {
					if latestIdx == -1 || version.Timestamp.After(latestTime) {
						latestIdx = idx
						latestTime = version.Timestamp
					}
				}
			}
			
			// Mark the latest version as active
			if latestIdx != -1 {
				fs.indexCache[entityType][latestIdx].Active = true
			}
		}
	}
}

// isFirstVersionForEntity checks if this is the first version for an entity
func (fs *FolderStorage) isFirstVersionForEntity(entityType, entityID string, timestamp time.Time) bool {
	if versions, ok := fs.indexCache[entityType]; ok {
		for _, version := range versions {
			if version.EntityID == entityID && version.Timestamp.Before(timestamp) {
				return false
			}
		}
	}
	return true
}

// findLatestVersionFileInDir finds the most recent version file in an entity directory
func (fs *FolderStorage) findLatestVersionFileInDir(entityDirPath string) (string, error) {
	files, err := os.ReadDir(entityDirPath)
	if err != nil {
		return "", fmt.Errorf("failed to read entity directory: %v", err)
	}
	
	var latestFile string
	var latestTime time.Time
	
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		
		// Parse timestamp from filename
		timestampStr := strings.TrimSuffix(file.Name(), ".json")
		timestamp, err := fs.parseTimestampFromFilename(timestampStr)
		if err != nil {
				continue
		}
		
		if latestFile == "" || timestamp.After(latestTime) {
			latestTime = timestamp
			latestFile = filepath.Join(entityDirPath, file.Name())
		}
	}
	
	if latestFile == "" {
		return "", fmt.Errorf("no valid version files found")
	}
	
	return latestFile, nil
}

// parseTimestampFromFilename parses timestamp from filename like "2025-07-19T16-05-17.795+04-00"
func (fs *FolderStorage) parseTimestampFromFilename(timestampStr string) (time.Time, error) {
	log.Printf("[STORAGE] Parsing timestamp from filename: '%s'", timestampStr)
	
	// The desktop app uses format: 2006-01-02T15-04-05.000-07-00
	// We need to convert this to proper RFC3339 format for parsing
	
	if timestampStr == "" {
		log.Printf("[STORAGE] ERROR: Empty timestamp string")
		return time.Time{}, fmt.Errorf("empty timestamp string")
	}
	
	// Convert desktop app format to proper RFC3339
	rfc3339Str := timestampStr
	
	// Handle timezone conversion: +04-00 or -05-00 -> +04:00 or -05:00
	if len(timestampStr) >= 6 {
		// Check if we have a timezone suffix like +04-00 or -05-00
		for i := len(timestampStr) - 6; i >= 0; i-- {
			if timestampStr[i] == '+' || timestampStr[i] == '-' {
				// Found timezone marker, check if it matches pattern XX-XX
				if i+6 <= len(timestampStr) {
					timezonePart := timestampStr[i+1:]
					if len(timezonePart) == 5 && timezonePart[2] == '-' {
						// Replace the dash in timezone with colon: 04-00 -> 04:00
						fixedTimezone := timezonePart[:2] + ":" + timezonePart[3:]
						rfc3339Str = timestampStr[:i+1] + fixedTimezone
						log.Printf("[STORAGE] Converted timezone: %s -> %s", timezonePart, fixedTimezone)
						break
					}
				}
			}
		}
	}
	
	log.Printf("[STORAGE] Converted timestamp: '%s' -> '%s'", timestampStr, rfc3339Str)
	
	// Try multiple parsing strategies
	parseFormats := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02T15:04:05.000Z07:00",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15-04-05.000-07:00", // Original desktop format
		"2006-01-02T15-04-05-07-00",    // Backup format
	}
	
	for i, format := range parseFormats {
		if timestamp, err := time.Parse(format, rfc3339Str); err == nil {
			log.Printf("[STORAGE] Successfully parsed with format %d (%s): %v", i, format, timestamp)
			return timestamp, nil
		}
		// Also try with original string
		if timestamp, err := time.Parse(format, timestampStr); err == nil {
			log.Printf("[STORAGE] Successfully parsed original string with format %d (%s): %v", i, format, timestamp)
			return timestamp, nil
		}
	}
	
	log.Printf("[STORAGE] ERROR: Failed to parse timestamp '%s' (converted to '%s') with any format", timestampStr, rfc3339Str)
	return time.Time{}, fmt.Errorf("failed to parse timestamp '%s' (converted to '%s')", timestampStr, rfc3339Str)
}

// isDeleteMarker checks if a file is a delete marker
func (fs *FolderStorage) isDeleteMarker(filePath string) bool {
	log.Printf("[STORAGE] Checking if file is delete marker: %s", filePath)
	
	// Read the file content to check for delete operation
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("[STORAGE] ERROR: Could not read file %s to check delete marker: %v", filePath, err)
		return false
	}
	
	var content map[string]interface{}
	if err := json.Unmarshal(data, &content); err != nil {
		log.Printf("[STORAGE] ERROR: Could not unmarshal JSON from %s to check delete marker: %v", filePath, err)
		return false
	}
	
	// Check if operation field indicates deletion
	if operation, ok := content["operation"].(string); ok && operation == OperationDelete {
		log.Printf("[STORAGE] File %s is a delete marker (operation field)", filePath)
		return true
	}
	
	// Fallback: check if filename contains "deleted_" pattern (legacy)
	filename := filepath.Base(filePath)
	isDeleted := strings.Contains(filename, "deleted_")
	if isDeleted {
		log.Printf("[STORAGE] File %s is a delete marker (filename pattern)", filePath)
	} else {
		log.Printf("[STORAGE] File %s is NOT a delete marker", filePath)
	}
	return isDeleted
}
