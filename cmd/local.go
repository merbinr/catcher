package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/merbinr/catcher/internal/logs/vpc"
	"github.com/merbinr/catcher/internal/models"
)

func LocalRun() {
	local_data := readLocalData("/app/example_body.json")
	err := vpc.AwsVpcLogProcessing(local_data)
	if err != nil {
		slog.Error(fmt.Sprintf("Error in local run: %s", err))
	}
}

func readLocalData(filename string) models.AwsVpcLogWebhookModel {
	file, err := os.Open(filename)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to open file for local run: %s", err))
		os.Exit(1)
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to read file for local run: %s", err))
		os.Exit(1)
	}

	var TestRequestBody models.AwsVpcLogWebhookModel
	err = json.Unmarshal(byteValue, &TestRequestBody)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to unmarshal JSON for local run: %s", err))
		os.Exit(1)
	}
	return TestRequestBody
}
