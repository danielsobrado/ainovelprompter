package config

import (
	"os" // Required for os.Getenv and to check if .env exists
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus" // Assuming your logger is initialized and available
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Attempt to load .env file from the current working directory of the Go app.
	// For Wails, this might be the root of your project during `wails dev`,
	// or alongside the executable when built.
	// Ensure your .env file is in the `server` directory if running `go run ./cmd/main.go` from there.
	// If running `wails dev`, it looks from the project root.
	// For consistency, you might want to specify a path or check multiple common locations.
	envPath := ".env" // Assumes .env is in the same directory as the executable or where Go app is run
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		// Try one level up if running from a subdirectory like `cmd`
		// This is a common pattern but might need adjustment based on your build/run process.
		// For Wails, the CWD is usually the project root.
		// envPath = "../.env"
	}

	err := godotenv.Load(envPath) // Load .env file
	if err != nil {
		// Logrus might not be initialized yet if this is the first thing called.
		// Using fmt.Println for early config errors might be safer, or ensure logger is up.
		// For now, assuming logger might be available or this is informational.
		logrus.Infof("No .env file found at %s or error loading: %v. Relying on system env vars and config.yaml.", envPath, err)
	} else {
		logrus.Info(".env file loaded successfully.")
	}

	viper.SetConfigName("config")   // Name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./server") // Path to look for the config file in (relative to where app runs)
	viper.AddConfigPath(".")        // Also look in the current directory

	// For environment variables
	viper.SetEnvPrefix("APP") // Environment variables will be APP_YOURKEY
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // e.g., openrouter.api_key becomes APP_OPENROUTER_API_KEY

	// Set defaults before reading config file or env vars
	// These are the ultimate fallbacks.
	viper.SetDefault("openrouter.api_key", "")
	viper.SetDefault("openrouter.default_model", "anthropic/claude-3-haiku")
	viper.SetDefault("lmstudio.api_url", "http://localhost:1234/v1/chat/completions")
	viper.SetDefault("lmstudio.default_model", "local-model")
	// Keep existing ollama defaults if they are there
	viper.SetDefault("ollama.generate_url", "http://localhost:11434/api/generate")
	viper.SetDefault("ollama.chat_url", "http://localhost:11434/api/chat")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("config.yaml not found, relying on .env, environment variables, and defaults.")
		} else {
			logrus.Warnf("Error reading config.yaml: %v. Relying on .env, environment variables, and defaults.", err)
		}
	} else {
		logrus.Info("config.yaml loaded successfully.")
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}
