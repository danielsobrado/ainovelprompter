package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_NewApp(t *testing.T) {
	app := NewApp()

	assert.NotNil(t, app)
}

func TestApp_LogInfo(t *testing.T) {
	app := NewApp()

	// This test mainly ensures the method doesn't panic
	assert.NotPanics(t, func() {
		app.LogInfo("Test log message")
	})
}

func TestApp_GetCurrentDirectory(t *testing.T) {
	app := NewApp()

	result := app.GetCurrentDirectory()

	assert.NotEmpty(t, result)
}
