package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/merbinr/catcher/internal/web"
)

func main() {
	// Loading config yml data so helpers.ConfigData will be accessible
	initialize()

	DEPLOYMENT_MODE := os.Getenv("DEPLOYMENT_MODE")
	if DEPLOYMENT_MODE == "" {
		DEPLOYMENT_MODE = "dev"
	}

	if strings.ToLower(DEPLOYMENT_MODE) == "prod" {
		r := web.SetupRouter() // Setup routes
		if err := r.Run(":8080"); err != nil {
			slog.Error(fmt.Sprintf("Failed to start server: %v", err))
			os.Exit(1)
		}
	} else {
		LocalRun()
	}
}
