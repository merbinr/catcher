package config

import (
	"fmt"
	"os"

	"github.com/merbinr/catcher/internal/models"
	"gopkg.in/yaml.v3"
)

var ConfigData models.Config

// LoadConfig reads the YAML configuration file once and returns the config instance.
func LoadConfig() error {
	filename := "config/config.yml"
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %s, error: %s", filename, err)
	}

	err = yaml.Unmarshal(data, &ConfigData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config YAML file: %s, error: %s", filename, err)
	}
	return nil
}
