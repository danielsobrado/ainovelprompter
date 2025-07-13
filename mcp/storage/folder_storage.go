package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/models"
	"github.com/google/uuid"
)

// FolderStorage implements VersionedStorage using folder-based storage with versioning
type FolderStorage struct {
	basePath   string
	mu         sync.RWMutex
	indexCache map[string][]Version // Entity type -> versions cache
}

// NewFolderStorage creates a new folder-based storage instance
func NewFolderStorage(basePath string) *FolderStorage {
	fs := &FolderStorage{
		basePath:   basePath,
		indexCache: make(map[string][]Version),
	}
	
	// Ensure base directory exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		// Log error but don't fail initialization
		fmt.Printf("Warning: failed to create base directory %s: %v\n", basePath, err)
	}
	
	// Create entity directories
	fs.initializeDirectories()
	
	// Load index cache
	fs.rebuildIndexCache()
	
	return fs
}

// initializeDirectories creates the folder structure for all entity types
func (fs *FolderStorage) initializeDirectories() {
	entityTypes := []string{
		EntityCharacters, EntityLocations, EntityCodex, EntityRules,
		EntityChapters, EntityStoryBeats, EntityFutureNotes,
		EntitySampleChapters, EntityTaskTypes, EntityProsePrompts,
	}
	
	for _, entityType := range entityTypes {
		dirPath := filepath.Join(fs.basePath, entityType)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Printf("Warning: failed to create directory %s: %v\n", dirPath, err)
		}
	}
	
	// Create metadata directory
	metadataPath := filepath.Join(fs.basePath, ".metadata")
	if err := os.MkdirAll(metadataPath, 0755); err != nil {
		fmt.Printf("Warning: failed to create metadata directory: %v\n", err)
	}
}

// generateFileName creates a timestamped filename for an entity
func (fs *FolderStorage) generateFileName(entityName, operation string) string {
	// Slugify entity name (lowercase, spaces to underscores, remove special chars)
	slug := strings.ToLower(entityName)
	slug = strings.ReplaceAll(slug, " ", "_")
	slug = strings.ReplaceAll(slug, "-", "_")
	// Remove special characters but keep underscores and alphanumeric
	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			result.WriteRune(r)
		}
	}
	slug = result.String()
	
	// Ensure slug is not empty
	if slug == "" {
		slug = "unnamed"
	}
	
	// Generate timestamp
	timestamp := time.Now().Format("20060102_150405")
	
	return fmt.Sprintf("%s_%s_%s.json", slug, timestamp, operation)
}

// Create implements VersionedStorage.Create
func (fs *FolderStorage) Create(entityType string, entity interface{}) (*Version, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	// Extract entity information
	entityInfo, err := fs.extractEntityInfo(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to extract entity info: %v", err)
	}
	
	// Ensure entity has an ID
	if entityInfo.ID == "" {
		entityInfo.ID = uuid.New().String()
		fs.setEntityID(entity, entityInfo.ID)
	}
	
	// Set timestamps
	now := time.Now()
	fs.setEntityTimestamps(entity, now, now)
	
	// Generate filename
	filename := fs.generateFileName(entityInfo.Name, OperationCreate)
	filePath := filepath.Join(fs.basePath, entityType, filename)
	
	// Write entity to file
	if err := fs.writeEntityToFile(filePath, entity); err != nil {
		return nil, fmt.Errorf("failed to write entity to file: %v", err)
	}
	
	// Create version record
	version := &Version{
		ID:        uuid.New().String(),
		EntityID:  entityInfo.ID,
		Timestamp: now,
		Operation: OperationCreate,
		FilePath:  filePath,
		Active:    true,
	}
	
	// Update cache
	fs.addVersionToCache(entityType, *version)
	
	return version, nil
}

// Update implements VersionedStorage.Update
func (fs *FolderStorage) Update(entityType string, id string, entity interface{}) (*Version, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	// Check if entity exists
	existing, err := fs.getLatestEntity(entityType, id)
	if err != nil {
		return nil, fmt.Errorf("entity not found: %v", err)
	}
	
	// Extract entity information
	entityInfo, err := fs.extractEntityInfo(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to extract entity info: %v", err)
	}
	
	// Ensure ID matches
	if entityInfo.ID != id {
		fs.setEntityID(entity, id)
	}
	
	// Set timestamps (preserve created, update modified)
	existingInfo, _ := fs.extractEntityInfo(existing)
	now := time.Now()
	fs.setEntityTimestamps(entity, existingInfo.CreatedAt, now)
	
	// Generate filename
	filename := fs.generateFileName(entityInfo.Name, OperationUpdate)
	filePath := filepath.Join(fs.basePath, entityType, filename)
	
	// Write entity to file
	if err := fs.writeEntityToFile(filePath, entity); err != nil {
		return nil, fmt.Errorf("failed to write entity to file: %v", err)
	}
	
	// Mark previous versions as inactive
	fs.deactivateVersions(entityType, id)
	
	// Create version record
	version := &Version{
		ID:        uuid.New().String(),
		EntityID:  id,
		Timestamp: now,
		Operation: OperationUpdate,
		FilePath:  filePath,
		Active:    true,
	}
	
	// Update cache
	fs.addVersionToCache(entityType, *version)
	
	return version, nil
}

// Delete implements VersionedStorage.Delete
func (fs *FolderStorage) Delete(entityType string, id string) (*Version, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	// Check if entity exists
	_, err := fs.getLatestEntity(entityType, id)
	if err != nil {
		return nil, fmt.Errorf("entity not found: %v", err)
	}
	
	// Create delete marker file
	now := time.Now()
	filename := fs.generateFileName(fmt.Sprintf("deleted_%s", id), OperationDelete)
	filePath := filepath.Join(fs.basePath, entityType, filename)
	
	// Write delete marker
	deleteMarker := map[string]interface{}{
		"id":        id,
		"deletedAt": now,
		"operation": OperationDelete,
	}
	
	if err := fs.writeEntityToFile(filePath, deleteMarker); err != nil {
		return nil, fmt.Errorf("failed to write delete marker: %v", err)
	}
	
	// Mark all versions as inactive
	fs.deactivateVersions(entityType, id)
	
	// Create version record
	version := &Version{
		ID:        uuid.New().String(),
		EntityID:  id,
		Timestamp: now,
		Operation: OperationDelete,
		FilePath:  filePath,
		Active:    false, // Delete markers are not active
	}
	
	// Update cache
	fs.addVersionToCache(entityType, *version)
	
	return version, nil
}

// GetLatest implements VersionedStorage.GetLatest
func (fs *FolderStorage) GetLatest(entityType string, id string) (interface{}, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	return fs.getLatestEntity(entityType, id)
}

// GetAll implements VersionedStorage.GetAll
func (fs *FolderStorage) GetAll(entityType string) ([]interface{}, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	// Get all active entities of this type
	versions := fs.getActiveVersionsByType(entityType)
	var entities []interface{}
	
	// Group by entity ID and get latest version
	entityMap := make(map[string]Version)
	for _, version := range versions {
		if existing, ok := entityMap[version.EntityID]; !ok || version.Timestamp.After(existing.Timestamp) {
			entityMap[version.EntityID] = version
		}
	}
	
	// Load entities from latest versions
	for _, version := range entityMap {
		if version.Operation != OperationDelete {
			entity, err := fs.loadEntityFromFile(version.FilePath, entityType)
			if err != nil {
				continue // Skip corrupted files
			}
			entities = append(entities, entity)
		}
	}
	
	return entities, nil
}

// GetVersions implements VersionedStorage.GetVersions
func (fs *FolderStorage) GetVersions(entityType string, id string) ([]Version, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	var versions []Version
	if cached, ok := fs.indexCache[entityType]; ok {
		for _, version := range cached {
			if version.EntityID == id {
				versions = append(versions, version)
			}
		}
	}
	
	// Sort by timestamp (newest first)
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Timestamp.After(versions[j].Timestamp)
	})
	
	return versions, nil
}

// GetVersion implements VersionedStorage.GetVersion
func (fs *FolderStorage) GetVersion(entityType string, id string, timestamp time.Time) (interface{}, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	// Find version with exact or closest timestamp
	versions, err := fs.GetVersions(entityType, id)
	if err != nil {
		return nil, err
	}
	
	var targetVersion *Version
	for _, version := range versions {
		if version.Timestamp.Equal(timestamp) || (targetVersion == nil && version.Timestamp.Before(timestamp)) {
			targetVersion = &version
		}
	}
	
	if targetVersion == nil {
		return nil, fmt.Errorf("no version found for timestamp %v", timestamp)
	}
	
	return fs.loadEntityFromFile(targetVersion.FilePath, entityType)
}

// RestoreVersion implements VersionedStorage.RestoreVersion
func (fs *FolderStorage) RestoreVersion(entityType string, id string, timestamp time.Time) (*Version, error) {
	// Get the version to restore
	entity, err := fs.GetVersion(entityType, id, timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get version: %v", err)
	}
	
	// Create new version with restored data
	return fs.Update(entityType, id, entity)
}

// SetDataDirectory implements VersionedStorage.SetDataDirectory
func (fs *FolderStorage) SetDataDirectory(path string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	// Validate new directory
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	
	fs.basePath = path
	fs.initializeDirectories()
	fs.rebuildIndexCache()
	
	return nil
}

// GetDataDirectory implements VersionedStorage.GetDataDirectory
func (fs *FolderStorage) GetDataDirectory() string {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	return fs.basePath
}



// CleanupOldVersions implements VersionedStorage.CleanupOldVersions
func (fs *FolderStorage) CleanupOldVersions(entityType string, retentionDays int) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	cutoffTime := time.Now().AddDate(0, 0, -retentionDays)
	
	versions := fs.indexCache[entityType]
	for _, version := range versions {
		if version.Timestamp.Before(cutoffTime) && !version.Active {
			// Remove old inactive versions
			if err := os.Remove(version.FilePath); err != nil {
				fmt.Printf("Warning: failed to remove old version file %s: %v\n", version.FilePath, err)
			}
		}
	}
	
	// Rebuild cache after cleanup
	fs.rebuildIndexCache()
	
	return nil
}

// GetStorageStats implements VersionedStorage.GetStorageStats
func (fs *FolderStorage) GetStorageStats() (StorageStats, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	stats := StorageStats{
		EntitiesByType: make(map[string]int),
		VersionsByType: make(map[string]int),
	}
	
	for entityType, versions := range fs.indexCache {
		stats.VersionsByType[entityType] = len(versions)
		
		// Count active entities
		activeEntities := make(map[string]bool)
		for _, version := range versions {
			if version.Active && version.Operation != OperationDelete {
				activeEntities[version.EntityID] = true
			}
		}
		stats.EntitiesByType[entityType] = len(activeEntities)
		
		// Track timestamps
		for _, version := range versions {
			if stats.OldestTimestamp.IsZero() || version.Timestamp.Before(stats.OldestTimestamp) {
				stats.OldestTimestamp = version.Timestamp
			}
			if stats.NewestTimestamp.IsZero() || version.Timestamp.After(stats.NewestTimestamp) {
				stats.NewestTimestamp = version.Timestamp
			}
		}
		
		stats.TotalFiles += len(versions)
	}
	
	// Calculate total size by walking directories
	filepath.Walk(fs.basePath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			stats.TotalSize += info.Size()
		}
		return nil
	})
	
	return stats, nil
}

// Helper methods continue in next file...
