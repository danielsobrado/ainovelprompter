package main

import (
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) LogInfo(message string) {
	runtime.LogInfo(a.ctx, message)
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
func (a *App) getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("Error getting current directory: %v", err))
		return "Error getting current directory"
	}
	return dir
}
