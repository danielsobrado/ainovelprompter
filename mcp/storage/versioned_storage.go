package storage

import (
	"time"
)

// Entity types constants
const (
	EntityCharacters     = "characters"
	EntityLocations      = "locations"
	EntityCodex          = "codex"
	EntityRules          = "rules"
	EntityChapters       = "chapters"
	EntityStoryBeats     = "story_beats"
	EntityFutureNotes    = "future_notes"
	EntitySampleChapters = "sample_chapters"
	EntityTaskTypes      = "task_types"
	EntityProsePrompts   = "prose_prompts"
)

// Operations constants
const (
	OperationCreate = "create"
	OperationUpdate = "update"
	OperationDelete = "delete"
)

// Version represents a versioned file entry
type Version struct {
	ID        string    `json:"id"`
	EntityID  string    `json:"entityId"`
	Timestamp time.Time `json:"timestamp"`
	Operation string    `json:"operation"` // create, update, delete
	FilePath  string    `json:"filePath"`
	Active    bool      `json:"active"`
}

// VersionedStorage provides CRUD operations with versioning support
type VersionedStorage interface {
	// CRUD with versioning
	Create(entityType string, entity interface{}) (*Version, error)
	Update(entityType string, id string, entity interface{}) (*Version, error)
	Delete(entityType string, id string) (*Version, error)
	GetLatest(entityType string, id string) (interface{}, error)
	GetAll(entityType string) ([]interface{}, error)

	// Versioning operations
	GetVersions(entityType string, id string) ([]Version, error)
	GetVersion(entityType string, id string, timestamp time.Time) (interface{}, error)
	RestoreVersion(entityType string, id string, timestamp time.Time) (*Version, error)

	// Directory management
	SetDataDirectory(path string) error
	GetDataDirectory() string

	// Cleanup operations
	CleanupOldVersions(entityType string, retentionDays int) error
	GetStorageStats() (StorageStats, error)
}

// StorageStats provides information about storage usage
type StorageStats struct {
	TotalFiles      int            `json:"totalFiles"`
	TotalSize       int64          `json:"totalSize"`
	EntitiesByType  map[string]int `json:"entitiesByType"`
	VersionsByType  map[string]int `json:"versionsByType"`
	OldestTimestamp time.Time      `json:"oldestTimestamp"`
	NewestTimestamp time.Time      `json:"newestTimestamp"`
}

// VersionedCombinedStorage combines all storage interfaces with versioning
type VersionedCombinedStorage interface {
	VersionedStorage
	StoryContextStorage
	ChapterStorage
	Storage
}
