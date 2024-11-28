package main

import (
	"fmt"
	"log/slog"

	"github.com/merbinr/catcher/internal/config"
	"github.com/merbinr/catcher/internal/rabbitmq"
)

func initialize() {
	err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("%s", err))
	}

	err = rabbitmq.CreateQueueConn()
	if err != nil {
		slog.Error(fmt.Sprintf("%s", err))
	}
}
