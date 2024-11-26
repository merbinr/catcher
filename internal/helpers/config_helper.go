package helpers

import (
	"log"
	"os"

	"github.com/merbinr/catcher/internal/models"
	"gopkg.in/yaml.v3"
)

var ConfigData models.Config

// LoadConfig reads the YAML configuration file once and returns the config instance.
func LoadConfig() {
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(data, &ConfigData)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
}
