package config

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigData Config

// LoadConfig reads the YAML configuration file once and returns the config instance.
func LoadConfig() {
	filename := "config/config.yml"
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to read YAML file: %s, error: %s", filename, err))
	}

	err = yaml.Unmarshal(data, &ConfigData)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to unmarshal config YAML file: %s, error: %s", filename, err))
	}
}
