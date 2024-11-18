package main

import (
	"log"

	"github.com/merbinr/catcher/internal/routes"
)

func main() {
	r := routes.SetupRouter() // Setup routes
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
