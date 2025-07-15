package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/danielsobrado/ainovelprompter/mcp/storage"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	dataDir string
}

// NewApp creates a new App application struct
func NewApp(dataDir string) *App {
	// Debug logging for app creation
	debugLog := fmt.Sprintf("DEBUG: NewApp called with dataDir='%s'\n", dataDir)
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}

	return &App{
		dataDir: dataDir,
	}
}

func (a *App) LogInfo(message string) {
	if a.ctx != nil {
		runtime.LogInfo(a.ctx, message)
	}
	// In tests or when context is nil, this becomes a no-op
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Debug logging for startup
	debugLog := fmt.Sprintf("DEBUG: startup() called, a.dataDir='%s'\n", a.dataDir)
	if file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		file.WriteString(debugLog)
		file.Close()
	}
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Add startup tasks here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform cleanup tasks here
}

// getCurrentDirectory returns the current working directory
func (a *App) GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		if a.ctx != nil {
			runtime.LogError(a.ctx, fmt.Sprintf("Error getting current directory: %v", err))
		}
		return "Error getting current directory"
	}
	return dir
}

// GetDataDirectory returns the current data directory
func (a *App) GetDataDirectory() string {
	return a.dataDir
}

// SetDataDirectory changes the data directory and creates it if necessary
func (a *App) SetDataDirectory(newDataDir string) error {
	// Expand relative paths to absolute paths
	absPath, err := filepath.Abs(newDataDir)
	if err != nil {
		return fmt.Errorf("failed to resolve data directory path: %v", err)
	}

	// Ensure directory exists
	if err := os.MkdirAll(absPath, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	a.dataDir = absPath

	// Log the change
	if a.ctx != nil {
		runtime.LogInfo(a.ctx, fmt.Sprintf("Data directory changed to: %s", absPath))
	}

	return nil
}

// GetRecentDataDirectories returns a list of recently used data directories
func (a *App) GetRecentDataDirectories() []string {
	// TODO: Implement reading from config file
	// For now, return default directory
	homeDir, _ := os.UserHomeDir()
	defaultDir := filepath.Join(homeDir, ".ai-novel-prompter")
	return []string{defaultDir}
}

// AddRecentDataDirectory adds a directory to the recent list
func (a *App) AddRecentDataDirectory(path string) error {
	// TODO: Implement saving to config file
	// For now, just validate the path
	return a.ValidateDataDirectoryPath(path)
}

// ValidateDataDirectoryPath checks if a directory path is valid for use as data directory
func (a *App) ValidateDataDirectoryPath(path string) error {
	// Check if path can be created/accessed
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("invalid path: %v", err)
	}

	// Check if we can create the directory
	if err := os.MkdirAll(absPath, 0755); err != nil {
		return fmt.Errorf("cannot create directory: %v", err)
	}

	// Check if we can write to the directory
	testFile := filepath.Join(absPath, ".write_test")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return fmt.Errorf("cannot write to directory: %v", err)
	}

	// Clean up test file
	os.Remove(testFile)

	return nil
}

// ValidateDataDirectory checks if the current data directory is valid
func (a *App) ValidateDataDirectory() error {
	if a.dataDir == "" {
		return fmt.Errorf("data directory not set")
	}
	return a.ValidateDataDirectoryPath(a.dataDir)
}

// GetStorageStats returns storage statistics for the current data directory
func (a *App) GetStorageStats() (*storage.StorageStats, error) {
	if a.dataDir == "" {
		return nil, fmt.Errorf("data directory not set")
	}

	// Create folder storage instance
	fs := storage.NewFolderStorage(a.dataDir)

	// Get storage stats
	stats, err := fs.GetStorageStats()
	if err != nil {
		return nil, fmt.Errorf("failed to get storage stats: %v", err)
	}

	return &stats, nil
}

// CleanupOldVersions cleans up old versions based on retention policy
func (a *App) CleanupOldVersions(entityType string, retentionDays int) error {
	if a.dataDir == "" {
		return fmt.Errorf("data directory not set")
	}

	// Create folder storage instance
	fs := storage.NewFolderStorage(a.dataDir)

	// Cleanup old versions
	err := fs.CleanupOldVersions(entityType, retentionDays)
	if err != nil {
		return fmt.Errorf("failed to cleanup old versions: %v", err)
	}

	if a.ctx != nil {
		runtime.LogInfo(a.ctx, fmt.Sprintf("Cleaned up old versions for %s with %d days retention", entityType, retentionDays))
	}

	return nil
}

// === MCP COMPATIBILITY LAYER ===
// Functions to align MCP and Server file formats for compatibility

// MCPCharacterData represents the enhanced MCP character format
type MCPCharacterData struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Traits      map[string]interface{} `json:"traits,omitempty"`
	Notes       string                 `json:"notes,omitempty"`
	CreatedAt   string                 `json:"createdAt,omitempty"`
	UpdatedAt   string                 `json:"updatedAt,omitempty"`
}

// ConvertMCPCharacterToServer converts MCP character to server format
func (a *App) ConvertMCPCharacterToServer(mcpChar MCPCharacterData) map[string]interface{} {
	return map[string]interface{}{
		"id":          mcpChar.ID,
		"label":       mcpChar.Name, // MCP uses 'name', Server uses 'label'
		"description": mcpChar.Description,
	}
}

// ConvertServerCharacterToMCP converts server character to MCP format
func (a *App) ConvertServerCharacterToMCP(serverChar map[string]interface{}) MCPCharacterData {
	mcpChar := MCPCharacterData{
		Traits: make(map[string]interface{}),
	}

	if id, ok := serverChar["id"].(string); ok {
		mcpChar.ID = id
	}
	if label, ok := serverChar["label"].(string); ok {
		mcpChar.Name = label // Server uses 'label', MCP uses 'name'
	}
	if desc, ok := serverChar["description"].(string); ok {
		mcpChar.Description = desc
	}

	// Set timestamps if not present
	now := time.Now().Format(time.RFC3339)
	mcpChar.CreatedAt = now
	mcpChar.UpdatedAt = now

	return mcpChar
}

// SetDataDirectoryMCP sets the data directory for both Server and MCP compatibility
func (a *App) SetDataDirectoryMCP(newPath string) error {
	// Validate the path
	if err := a.ValidateDataDirectoryPath(newPath); err != nil {
		return err
	}

	// Set the server data directory
	a.dataDir = newPath

	// Ensure directory exists
	if err := os.MkdirAll(newPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Create both Server and MCP directory structures
	return a.initializeMCPCompatibleDirectories(newPath)
}

// initializeMCPCompatibleDirectories creates directory structure for both formats
func (a *App) initializeMCPCompatibleDirectories(basePath string) error {
	// MCP format directories (entity subdirectories)
	mcpDirs := []string{
		"characters", "locations", "codex", "rules",
		"chapters", "story_beats", "future_notes",
		"sample_chapters", "task_types", "prose_prompts",
	}

	for _, dir := range mcpDirs {
		dirPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create MCP directory %s: %v", dir, err)
		}
	}

	return nil
}

// GetDataDirectoryInfo returns information about data directory compatibility
func (a *App) GetDataDirectoryInfo() (map[string]interface{}, error) {
	dataDir := a.dataDir
	if dataDir == "" {
		// Try to get from settings
		dataDir = a.getAppDataDir()
	}

	info := map[string]interface{}{
		"path":        dataDir,
		"serverFiles": make(map[string]bool),
		"mcpDirs":     make(map[string]bool),
		"compatible":  false,
	}

	// Check for Server format files
	serverFiles := []string{
		"characters.json", "locations.json", "codex.json", "rules.json",
		"sample_chapters.json", "task_types.json", "prose_prompts.json",
	}

	serverFileCount := 0
	for _, file := range serverFiles {
		filePath := filepath.Join(dataDir, file)
		if _, err := os.Stat(filePath); err == nil {
			info["serverFiles"].(map[string]bool)[file] = true
			serverFileCount++
		}
	}

	// Check for MCP format directories
	mcpDirs := []string{
		"characters", "locations", "codex", "rules",
		"sample_chapters", "task_types", "prose_prompts",
	}

	mcpDirCount := 0
	for _, dir := range mcpDirs {
		dirPath := filepath.Join(dataDir, dir)
		if _, err := os.Stat(dirPath); err == nil {
			info["mcpDirs"].(map[string]bool)[dir] = true
			mcpDirCount++
		}
	}

	// Determine compatibility
	info["compatible"] = serverFileCount > 0 || mcpDirCount > 0
	info["hasServerFormat"] = serverFileCount > 0
	info["hasMCPFormat"] = mcpDirCount > 0

	return info, nil
}
