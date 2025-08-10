package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config == nil {
		t.Error("DefaultConfig should return a valid config")
	}
	if config.DefaultType != "tv" {
		t.Errorf("Expected default type 'tv', got '%s'", config.DefaultType)
	}
	if len(config.Download.TVQuality) == 0 {
		t.Error("Default config should have TV quality settings")
	}
}

func TestLoadConfigNonExistent(t *testing.T) {
	// Test loading a non-existent config file
	config, err := LoadConfig("/nonexistent/path/config.json")
	if err != nil {
		t.Errorf("LoadConfig should not error on non-existent file: %v", err)
	}
	if config == nil {
		t.Error("LoadConfig should return default config when file doesn't exist")
	}
}

func TestSaveAndLoadConfig(t *testing.T) {
	// Create temporary directory for test
	tempDir := os.TempDir()
	configPath := filepath.Join(tempDir, "test_config.json")

	// Clean up after test
	defer os.Remove(configPath)

	// Create and save config
	config := DefaultConfig()
	config.DefaultType = "radio"

	err := SaveConfig(config, configPath)
	if err != nil {
		t.Errorf("SaveConfig should not error: %v", err)
	}

	// Load config back
	loadedConfig, err := LoadConfig(configPath)
	if err != nil {
		t.Errorf("LoadConfig should not error: %v", err)
	}

	if loadedConfig.DefaultType != "radio" {
		t.Errorf("Expected loaded type 'radio', got '%s'", loadedConfig.DefaultType)
	}
}
