package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
)

// EntityInfo contains basic information about an entity
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
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	
	// Create appropriate entity type
	entity := fs.createEntityByType(entityType)
	if entity == nil {
		return nil, fmt.Errorf("unknown entity type: %s", entityType)
	}
	
	// Unmarshal JSON
	if err := json.Unmarshal(data, entity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal entity: %v", err)
	}
	
	return entity, nil
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

// getLatestEntity gets the latest version of an entity
func (fs *FolderStorage) getLatestEntity(entityType string, id string) (interface{}, error) {
	versions, err := fs.GetVersions(entityType, id)
	if err != nil {
		return nil, err
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
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		
		// Parse filename: entityname_timestamp_operation.json
		parts := strings.Split(strings.TrimSuffix(entry.Name(), ".json"), "_")
		if len(parts) < 3 {
			continue // Invalid filename format
		}
		
		operation := parts[len(parts)-1]
		timestampStr := parts[len(parts)-2]
		
		// Parse timestamp
		timestamp, err := time.Parse("20060102_150405", timestampStr)
		if err != nil {
			continue // Invalid timestamp
		}
		
		filePath := filepath.Join(dirPath, entry.Name())
		
		// Try to extract entity ID from file content
		entityID := fs.extractEntityIDFromFile(filePath, entityType)
		if entityID == "" {
			continue // Could not determine entity ID
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
		entityVersions := make(map[string][]Version)
		for _, version := range versions {
			entityVersions[version.EntityID] = append(entityVersions[version.EntityID], version)
		}
		
		// For each entity, mark only latest non-delete as active
		for entityID, entityVersionList := range entityVersions {
			// Sort by timestamp (newest first)
			versions := entityVersionList
			for i := range versions {
				versions[i].Active = false // Reset all to inactive
			}
			
			// Sort and find latest non-delete
			for i := 0; i < len(versions); i++ {
				for j := i + 1; j < len(versions); j++ {
					if versions[j].Timestamp.After(versions[i].Timestamp) {
						versions[i], versions[j] = versions[j], versions[i]
					}
				}
			}
			
			// Mark latest non-delete as active
			for i, version := range versions {
				if version.Operation != OperationDelete {
					versions[i].Active = true
					break
				}
			}
			
			// Update cache
			for i, version := range fs.indexCache[entityType] {
				if version.EntityID == entityID {
					for _, updated := range versions {
						if version.ID == updated.ID {
							fs.indexCache[entityType][i] = updated
							break
						}
					}
				}
			}
		}
	}
}
