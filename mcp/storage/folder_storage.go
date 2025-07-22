package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

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
		indexCache: make(map[string][]Version), // Still needed for versioning operations
	}
	
	// Ensure base directory exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		// Log error but don't fail initialization
	}
	
	// Create entity directories
	fs.initializeDirectories()
	
	// Note: No longer rebuilding cache on startup - using direct file reading for GetAll
	
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
			// Log error but continue
		}
	}
	
	// Create metadata directory
	metadataPath := filepath.Join(fs.basePath, ".metadata")
	if err := os.MkdirAll(metadataPath, 0755); err != nil {
		// Log error but continue
	}
}

// generateFileName creates a timestamped filename for an entity using directory-per-entity format
func (fs *FolderStorage) generateFileName(entityID string, operation string) string {
	// Use RFC3339 format with timezone offset (matching desktop app format)
	timestamp := time.Now().Format("2006-01-02T15-04-05.000-07-00")
	return fmt.Sprintf("%s.json", timestamp)
}

// getEntityPath creates entity path using directory-per-entity structure
func (fs *FolderStorage) getEntityPath(entityType, entityID string) string {
	return filepath.Join(fs.basePath, entityType, entityID)
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
	
	// Create entity directory
	entityDir := fs.getEntityPath(entityType, entityInfo.ID)
	if err := os.MkdirAll(entityDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create entity directory: %v", err)
	}
	
	// Generate filename using desktop app format
	filename := fs.generateFileName(entityInfo.ID, OperationCreate)
	filePath := filepath.Join(entityDir, filename)
	
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
	
	// Create entity directory if not exists
	entityDir := fs.getEntityPath(entityType, id)
	if err := os.MkdirAll(entityDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create entity directory: %v", err)
	}
	
	// Generate filename using desktop app format
	filename := fs.generateFileName(id, OperationUpdate)
	filePath := filepath.Join(entityDir, filename)
	
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
	
	// Create entity directory if it doesn't exist
	entityDir := fs.getEntityPath(entityType, id)
	if err := os.MkdirAll(entityDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create entity directory: %v", err)
	}
	
	filePath := filepath.Join(entityDir, filename)
	
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
	
	log.Printf("[STORAGE] GetAll called for entityType: %s", entityType)
	
	result, err := fs.scanDirectoryForEntities(entityType)
	
	if err != nil {
		log.Printf("[STORAGE] GetAll failed for %s: %v", entityType, err)
	} else {
		log.Printf("[STORAGE] GetAll completed for %s: returned %d entities", entityType, len(result))
	}
	
	return result, err
}

// scanDirectoryForEntities scans directory directly without cache
func (fs *FolderStorage) scanDirectoryForEntities(entityType string) ([]interface{}, error) {
	dirPath := filepath.Join(fs.basePath, entityType)
	
	log.Printf("[STORAGE] Starting scan for entityType=%s in directory=%s", entityType, dirPath)
	
	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		log.Printf("[STORAGE] Directory does not exist: %s", dirPath)
		return []interface{}{}, nil
	}
	
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Printf("[STORAGE] Failed to read directory %s: %v", dirPath, err)
		return nil, fmt.Errorf("failed to read directory %s: %v", dirPath, err)
	}
	
	log.Printf("[STORAGE] Found %d entries in directory %s", len(entries), dirPath)
	
	var entities []interface{}
	var errors []string // Collect errors for debugging
	var skippedDirs, processedDirs, deletedEntities int
	
	// Process each entity directory
	for _, entry := range entries {
		if !entry.IsDir() {
			skippedDirs++
			continue // Skip non-directory files
		}
		
		entityID := entry.Name()
		entityDirPath := filepath.Join(dirPath, entityID)
		processedDirs++
		
		log.Printf("[STORAGE] Processing entity directory: %s (ID: %s)", entityDirPath, entityID)
		
		// Find latest version file
		files, err := os.ReadDir(entityDirPath)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to read entity directory %s: %v", entityDirPath, err)
			errors = append(errors, errorMsg)
			log.Printf("[STORAGE] ERROR: %s", errorMsg)
			continue
		}
		
		log.Printf("[STORAGE] Found %d files in entity directory %s", len(files), entityDirPath)
		
		var latestFile string
		var latestTime time.Time
		var timestampErrors []string
		var jsonFileCount int
		
		// Find the most recent .json file
		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
				continue
			}
			
			jsonFileCount++
			log.Printf("[STORAGE] Processing JSON file: %s", file.Name())
			
			// Parse timestamp from filename (remove .json extension)
			timestampStr := strings.TrimSuffix(file.Name(), ".json")
			timestamp, err := fs.parseTimestampFromFilename(timestampStr)
			if err != nil {
				errorMsg := fmt.Sprintf("Invalid timestamp in file %s: %v", file.Name(), err)
				timestampErrors = append(timestampErrors, errorMsg)
				log.Printf("[STORAGE] TIMESTAMP ERROR: %s", errorMsg)
				continue
			}
			
			log.Printf("[STORAGE] Successfully parsed timestamp %s -> %v", timestampStr, timestamp)
			
			if latestFile == "" || timestamp.After(latestTime) {
				latestTime = timestamp
				latestFile = filepath.Join(entityDirPath, file.Name())
				log.Printf("[STORAGE] Updated latest file to: %s (timestamp: %v)", latestFile, latestTime)
			}
		}
		
		log.Printf("[STORAGE] Entity %s: processed %d JSON files, latest file: %s", entityID, jsonFileCount, latestFile)
		
		if latestFile == "" {
			errorMsg := fmt.Sprintf("No valid files found for entity %s (processed %d JSON files)", entityID, jsonFileCount)
			if len(timestampErrors) > 0 {
				errorMsg += ". Timestamp errors: " + strings.Join(timestampErrors, ", ")
			}
			errors = append(errors, errorMsg)
			log.Printf("[STORAGE] ERROR: %s", errorMsg)
			continue
		}
		
		// Skip if this is a delete marker
		if fs.isDeleteMarker(latestFile) {
			deletedEntities++
			log.Printf("[STORAGE] Skipping deleted entity %s (file: %s)", entityID, latestFile)
			continue // This is expected behavior, not an error
		}
		
		// Load entity from latest file
		log.Printf("[STORAGE] Loading entity %s from file: %s", entityID, latestFile)
		entity, err := fs.loadEntityFromFile(latestFile, entityType)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to load entity %s from file %s: %v", entityID, latestFile, err)
			errors = append(errors, errorMsg)
			log.Printf("[STORAGE] LOAD ERROR: %s", errorMsg)
			continue
		}
		
		entities = append(entities, entity)
		log.Printf("[STORAGE] Successfully loaded entity %s", entityID)
	}
	
	// Summary logging
	log.Printf("[STORAGE] Scan complete for %s: processed=%d dirs, skipped=%d non-dirs, deleted=%d, loaded=%d entities, errors=%d", 
		entityType, processedDirs, skippedDirs, deletedEntities, len(entities), len(errors))
	
	// Log detailed errors if any
	if len(errors) > 0 {
		log.Printf("[STORAGE] ERRORS encountered during scan of %s:", entityType)
		for i, err := range errors {
			log.Printf("[STORAGE] Error %d: %s", i+1, err)
		}
	}
	
	return entities, nil
}

// GetVersions implements VersionedStorage.GetVersions
func (fs *FolderStorage) GetVersions(entityType string, id string) ([]Version, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	return fs.getVersionsInternal(entityType, id), nil
}

// getVersionsInternal gets versions without locking (internal use)
func (fs *FolderStorage) getVersionsInternal(entityType string, id string) []Version {
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
	
	return versions
}

// GetVersion implements VersionedStorage.GetVersion
func (fs *FolderStorage) GetVersion(entityType string, id string, timestamp time.Time) (interface{}, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	// Find version with exact or closest timestamp
	versions := fs.getVersionsInternal(entityType, id)
	
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
