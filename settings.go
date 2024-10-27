package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// ReadSettingsFile reads the settings from the settings.json file
func (a *App) ReadSettingsFile() (string, error) {
	settingsPath := filepath.Join(a.getAppDataDir(), "settings.json")
	content, err := os.ReadFile(settingsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "{}", nil
		}
		return "", fmt.Errorf("error reading settings file: %v", err)
	}
	return string(content), nil
}

// WriteSettingsFile writes the settings to the settings.json file
func (a *App) WriteSettingsFile(content string) error {
	settingsPath := filepath.Join(a.getAppDataDir(), "settings.json")
	err := os.MkdirAll(filepath.Dir(settingsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating settings directory: %v", err)
	}
	err = os.WriteFile(settingsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing settings file: %v", err)
	}
	return nil
}

// getAppDataDir returns the path to the application data directory
func (a *App) getAppDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("Error getting user home directory: %v", err))
		return ""
	}
	return filepath.Join(homeDir, ".code-prompter")
}

func (a *App) ReadTaskTypesFile() (string, error) {
	taskTypesPath := filepath.Join(a.getAppDataDir(), "task_types.json")
	content, err := os.ReadFile(taskTypesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading task types file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteTaskTypesFile(content string) error {
	taskTypesPath := filepath.Join(a.getAppDataDir(), "task_types.json")
	err := os.MkdirAll(filepath.Dir(taskTypesPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating task types directory: %v", err)
	}
	err = os.WriteFile(taskTypesPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing task types file: %v", err)
	}
	return nil
}

func (a *App) ReadRulessFile() (string, error) {
    RulessPath := filepath.Join(a.getAppDataDir(), "custom_instructions.json")
    content, err := os.ReadFile(RulessPath)
    if err != nil {
        if os.IsNotExist(err) {
            return "[]", nil
        }
        return "", fmt.Errorf("error reading custom instructions file: %v", err)
    }
    return string(content), nil
}

func (a *App) WriteRulessFile(content string) error {
	RulessPath := filepath.Join(a.getAppDataDir(), "custom_instructions.json")
    err := os.MkdirAll(filepath.Dir(RulessPath), 0755)
    if err != nil {
        return fmt.Errorf("error creating custom instructions directory: %v", err)
    }
    err = os.WriteFile(RulessPath, []byte(content), 0644)
    if err != nil {
        return fmt.Errorf("error writing custom instructions file: %v", err)
    }
    return nil
}