package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTextFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{
			name:     "text file with .txt extension",
			filename: "document.txt",
			expected: true,
		},
		{
			name:     "markdown file",
			filename: "readme.md",
			expected: true,
		},
		{
			name:     "JSON file",
			filename: "config.json",
			expected: true,
		},
		{
			name:     "Go source file",
			filename: "main.go",
			expected: true,
		},
		{
			name:     "binary file",
			filename: "image.jpg",
			expected: false,
		},
		{
			name:     "executable file",
			filename: "program.exe",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTextFile(tt.filename)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper function to determine if a file is text-based
func isTextFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	textExtensions := map[string]bool{
		".txt": true, ".md": true, ".json": true, ".js": true, ".jsx": true,
		".ts": true, ".tsx": true, ".go": true, ".py": true, ".java": true,
		".c": true, ".cpp": true, ".h": true, ".css": true, ".html": true,
		".xml": true, ".yaml": true, ".yml": true, ".toml": true, ".ini": true,
		".cfg": true, ".conf": true, ".log": true, ".csv": true, ".sql": true,
	}

	if ext == "" {
		return true // Treat extensionless files as text
	}

	return textExtensions[ext]
}
