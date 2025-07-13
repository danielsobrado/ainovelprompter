package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Parse command line arguments
	var dataDir string
	var showHelp bool
	
	flag.StringVar(&dataDir, "data-dir", "", "Data directory path (defaults to ~/.ai-novel-prompter)")
	flag.StringVar(&dataDir, "d", "", "Data directory path (short form)")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.BoolVar(&showHelp, "h", false, "Show help message (short form)")
	
	flag.Parse()
	
	if showHelp {
		showHelpMessage()
		os.Exit(0)
	}
	
	// Resolve data directory
	resolvedDataDir := resolveDataDirectory(dataDir)
	
	// Create an instance of the app structure with data directory
	app := NewApp(resolvedDataDir)

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "AI Novel Prompter",
		Width:             1200,
		Height:            1380,
		MinWidth:          1024,
		MinHeight:         768,
		MaxWidth:          1280,
		MaxHeight:         1600,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			WebviewUserDataPath:  "",
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "AI Novel Prompter",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

// showHelpMessage displays the command line help
func showHelpMessage() {
	fmt.Fprintf(os.Stderr, `AI Novel Prompter

Usage: %s [OPTIONS]

Options:
  -d, --data-dir PATH    Data directory path (default: ~/.ai-novel-prompter)
  -h, --help            Show this help message

Examples:
  %s                                        # Use default data directory
  %s -d ./my-story                         # Use relative path
  %s --data-dir /path/to/story/data       # Use absolute path
  %s -d "C:\My Stories\Novel Data"        # Windows path with spaces

The application will create the data directory if it doesn't exist.
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

// resolveDataDirectory resolves the data directory path
func resolveDataDirectory(dataDir string) string {
	if dataDir == "" {
		// Use default location in user home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get user home directory: %v", err)
		}
		return filepath.Join(homeDir, ".ai-novel-prompter")
	}
	
	// Expand relative paths to absolute paths
	absPath, err := filepath.Abs(dataDir)
	if err != nil {
		log.Fatalf("Failed to resolve data directory path: %v", err)
	}
	
	// Ensure directory exists
	if err := os.MkdirAll(absPath, 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}
	
	return absPath
}
