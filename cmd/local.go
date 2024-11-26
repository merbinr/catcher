package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/merbinr/catcher/internal/models"
)

func LocalRun() {
	local_data := readLocalData("example_body.json")
	fmt.Printf("%+v\n", local_data)
}

func readLocalData(filename string) models.AwsVpcLogWebhookModel {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var TestRequestBody models.AwsVpcLogWebhookModel
	err = json.Unmarshal(byteValue, &TestRequestBody)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}
	return TestRequestBody
}
