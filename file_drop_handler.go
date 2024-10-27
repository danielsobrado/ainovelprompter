package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) HandleFileDrop(files []string) error {
	runtime.LogDebug(a.ctx, fmt.Sprintf("Handling file drop for files: %v", files))

	// Convert dropped file paths to absolute paths
	for i, file := range files {
		absolutePath, err := filepath.Abs(file)
		if err != nil {
			runtime.LogError(a.ctx, fmt.Sprintf("Error getting absolute path: %v", err))
			return err
		}
		files[i] = absolutePath
	}

	processedFiles, err := a.processDroppedFiles(files)
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("Error handling file drop: %v", err))
		return err
	}

	runtime.LogDebug(a.ctx, fmt.Sprintf("Processed files: %v", processedFiles))
	runtime.EventsEmit(a.ctx, "files-dropped", processedFiles)
	return nil
}

func (a *App) processDroppedFiles(files []string) ([]string, error) {
	var processedFiles []string
	for _, file := range files {
		fullPath, err := filepath.Abs(file)
		if err != nil {
			runtime.LogWarning(a.ctx, fmt.Sprintf("Error getting absolute path for %s: %v", file, err))
			continue
		}

		info, err := os.Stat(fullPath)
		if err != nil {
			runtime.LogWarning(a.ctx, fmt.Sprintf("Error getting file info for %s: %v", fullPath, err))
			continue
		}

		if info.IsDir() {
			folderFiles, err := a.ProcessFolder(fullPath, map[string]interface{}{
				"recursive":      true,
				"ignoreSuffixes": ".env,.log,.json,.gitignore,.npmrc,.prettierrc",
				"ignoreFolders":  ".git,.vscode,.idea,node_modules,venv,build,dist,coverage,out,next",
			})
			if err != nil {
				runtime.LogWarning(a.ctx, fmt.Sprintf("Error processing folder %s: %v", fullPath, err))
				continue
			}
			processedFiles = append(processedFiles, folderFiles...)
		} else {
			if info.Size() <= 500*1024 {
				processedFiles = append(processedFiles, fullPath)
			}
		}
	}
	return processedFiles, nil
}

func (a *App) ProcessFolder(folderPath string, config map[string]interface{}) ([]string, error) {
	var files []string

	recursive, _ := config["recursive"].(bool)
	ignoreSuffixes, _ := config["ignoreSuffixes"].(string)
	ignoreFolders, _ := config["ignoreFolders"].(string)

	ignoreSuffixList := strings.Split(ignoreSuffixes, ",")
	ignoreFolderList := strings.Split(ignoreFolders, ",")

	// Get the absolute path of the folder
	absoluteFolderPath, err := filepath.Abs(folderPath)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path of folder: %v", err)
	}

	err = filepath.Walk(absoluteFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if !recursive && path != folderPath {
				return filepath.SkipDir
			}
			for _, ignoreFolder := range ignoreFolderList {
				if strings.HasSuffix(path, ignoreFolder) {
					return filepath.SkipDir
				}
			}
			return nil
		}

		for _, ignoreSuffix := range ignoreSuffixList {
			if strings.HasSuffix(path, ignoreSuffix) {
				return nil
			}
		}

		if info.Size() > 500*1024 {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error processing folder: %v", err)
	}

	return files, nil
}
