package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/merbinr/catcher/internal/web"
)

func LocalRun() {
	local_data := readLocalData("example_body.json")
	fmt.Printf("%+v\n", local_data)
}

func readLocalData(filename string) web.AwsVpcLogWebhookModel {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var TestRequestBody web.AwsVpcLogWebhookModel
	err = json.Unmarshal(byteValue, &TestRequestBody)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}
	return TestRequestBody
}
