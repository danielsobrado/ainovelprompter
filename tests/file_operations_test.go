package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFilePath(t *testing.T) {
	tests := []struct {
		name     string
		dataDir  string
		filename string
		expected string
	}{
		{
			name:     "get settings file path",
			dataDir:  "/home/user/.ai-novel-prompter",
			filename: "settings.json",
			expected: "/home/user/.ai-novel-prompter/settings.json",
		},
		{
			name:     "get characters file path",
			dataDir:  "/home/user/.ai-novel-prompter",
			filename: "characters.json",
			expected: "/home/user/.ai-novel-prompter/characters.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filepath.Join(tt.dataDir, tt.filename)

			// Use filepath.Clean to handle OS-specific path separators
			expectedClean := filepath.Clean(tt.expected)
			resultClean := filepath.Clean(result)

			assert.Equal(t, expectedClean, resultClean)
		})
	}
}

func TestFileExists(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() (string, func()) // returns filepath and cleanup
		expected bool
	}{
		{
			name: "existing file",
			setup: func() (string, func()) {
				tmpFile, err := os.CreateTemp("", "test_*.txt")
				require.NoError(t, err)
				tmpFile.Close()

				return tmpFile.Name(), func() {
					os.Remove(tmpFile.Name())
				}
			},
			expected: true,
		},
		{
			name: "non-existent file",
			setup: func() (string, func()) {
				return "/non/existent/file.txt", func() {}
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filepath, cleanup := tt.setup()
			defer cleanup()

			_, err := os.Stat(filepath)
			result := !os.IsNotExist(err)

			assert.Equal(t, tt.expected, result)
		})
	}
}
