package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/ZayneLiu/get_iplayer/pkg/models"
)

// DefaultConfig returns the default configuration
func DefaultConfig() *models.Config {
	homeDir, _ := os.UserHomeDir()
	cacheDir := filepath.Join(homeDir, ".get_iplayer")

	return &models.Config{
		CacheDir:    cacheDir,
		OutputDir:   ".",
		DefaultType: "tv",
		Download: models.DownloadOptions{
			TVQuality:    []string{"hd", "sd", "web", "mobile"},
			RadioQuality: []string{"high", "std", "med", "low"},
			Subtitles:    false,
			OutputDir:    ".",
			OutputFormat: "mp4",
		},
	}
}

// LoadConfig loads configuration from file or returns default
func LoadConfig(configPath string) (*models.Config, error) {
	config := DefaultConfig()

	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, ".get_iplayer", "config.json")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, nil // Return default config if file doesn't exist
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return config, err
	}

	return config, nil
}

// SaveConfig saves configuration to file
func SaveConfig(config *models.Config, configPath string) error {
	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = filepath.Join(homeDir, ".get_iplayer", "config.json")
	}

	// Ensure directory exists
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
