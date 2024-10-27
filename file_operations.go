package main

import (
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) SelectFile() (string, error) {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "All Files",
				Pattern:     "*.*",
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("error selecting file: %v", err)
	}
	return file, nil
}

func (a *App) SelectDirectory() (string, error) {
	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Folder",
	})
	if err != nil {
		return "", fmt.Errorf("error selecting directory: %v", err)
	}
	return directory, nil
}

func (a *App) ReadFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file content: %v", err)
	}
	return string(content), nil
}
