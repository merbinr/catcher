package main

import (
	"log"
	"os"
	"strings"

	"github.com/merbinr/catcher/internal/helpers"
	"github.com/merbinr/catcher/internal/routes"
)

func main() {
	// Loading config yml data so helpers.ConfigData will be accessible
	helpers.LoadConfig()

	DEPLOYMENT_MODE := os.Getenv("DEPLOYMENT_MODE")
	if DEPLOYMENT_MODE == "" {
		DEPLOYMENT_MODE = "dev"
	}

	if strings.ToLower(DEPLOYMENT_MODE) == "prod" {
		r := routes.SetupRouter() // Setup routes
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	} else {
		//pass
	}
}
