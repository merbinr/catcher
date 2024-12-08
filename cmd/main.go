package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/merbinr/catcher/internal/config"
	"github.com/merbinr/catcher/internal/rabbitmq"
	"github.com/merbinr/catcher/internal/web"
)

func main() {
	// Loading config yml data so helpers.ConfigData will be accessible
	err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	// Creating a connection to RabbitMQ, so rabbitmq.Rabbitmq_conn will be accessible
	err = rabbitmq.CreateQueueConn()
	if err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	r := web.SetupRouter() // Setup routes
	if err := r.Run(":8080"); err != nil {
		slog.Error(fmt.Sprintf("Failed to start server: %v", err))
		os.Exit(1)
	}
}
