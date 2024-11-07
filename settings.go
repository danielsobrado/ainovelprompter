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
	return filepath.Join(homeDir, ".ai-novel-prompter")
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

// Characters
func (a *App) ReadCharactersFile() (string, error) {
	charactersPath := filepath.Join(a.getAppDataDir(), "characters.json")
	content, err := os.ReadFile(charactersPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading characters file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteCharactersFile(content string) error {
	charactersPath := filepath.Join(a.getAppDataDir(), "characters.json")
	err := os.MkdirAll(filepath.Dir(charactersPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating characters directory: %v", err)
	}
	err = os.WriteFile(charactersPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing characters file: %v", err)
	}
	return nil
}

// Locations
func (a *App) ReadLocationsFile() (string, error) {
	locationsPath := filepath.Join(a.getAppDataDir(), "locations.json")
	content, err := os.ReadFile(locationsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading locations file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteLocationsFile(content string) error {
	locationsPath := filepath.Join(a.getAppDataDir(), "locations.json")
	err := os.MkdirAll(filepath.Dir(locationsPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating locations directory: %v", err)
	}
	err = os.WriteFile(locationsPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing locations file: %v", err)
	}
	return nil
}

// Codex
func (a *App) ReadCodexFile() (string, error) {
	codexPath := filepath.Join(a.getAppDataDir(), "codex.json")
	content, err := os.ReadFile(codexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading codex file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteCodexFile(content string) error {
	codexPath := filepath.Join(a.getAppDataDir(), "codex.json")
	err := os.MkdirAll(filepath.Dir(codexPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating codex directory: %v", err)
	}
	err = os.WriteFile(codexPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing codex file: %v", err)
	}
	return nil
}

// Sample Chapters
func (a *App) ReadSampleChaptersFile() (string, error) {
	chaptersPath := filepath.Join(a.getAppDataDir(), "sample_chapters.json")
	content, err := os.ReadFile(chaptersPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading sample chapters file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteSampleChaptersFile(content string) error {
	chaptersPath := filepath.Join(a.getAppDataDir(), "sample_chapters.json")
	err := os.MkdirAll(filepath.Dir(chaptersPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating sample chapters directory: %v", err)
	}
	err = os.WriteFile(chaptersPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing sample chapters file: %v", err)
	}
	return nil
}

// Fix the typo in ReadRulessFile
func (a *App) ReadRulesFile() (string, error) {
	rulesPath := filepath.Join(a.getAppDataDir(), "rules.json")
	content, err := os.ReadFile(rulesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "[]", nil
		}
		return "", fmt.Errorf("error reading rules file: %v", err)
	}
	return string(content), nil
}

func (a *App) WriteRulesFile(content string) error {
	rulesPath := filepath.Join(a.getAppDataDir(), "rules.json")
	err := os.MkdirAll(filepath.Dir(rulesPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating rules directory: %v", err)
	}
	err = os.WriteFile(rulesPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing rules file: %v", err)
	}
	return nil
}
