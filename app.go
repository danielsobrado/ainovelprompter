package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	dataDir string
}

// NewApp creates a new App application struct
func NewApp(dataDir string) *App {
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

// ValidateDataDirectory checks if a directory path is valid for use as data directory
func (a *App) ValidateDataDirectory(path string) error {
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
	return a.ValidateDataDirectory(path)
}
