package web

import (
	"os"
)

// firehose sends X-Amz-Firehose-Access-Key in response header with user provider API KEY
// Check whether the header exists and validate the API key
func CheckAuthentication(headers map[string][]string) bool {
	api_token_header := headers["X-Amz-Firehose-Access-Key"]
	if len(api_token_header) == 0 {
		println("X-Amz-Firehose-Access-Key header not present in the header")
		return false
	}
	api_token_from_request := api_token_header[0]

	valid_api_token := os.Getenv("CATCHER_HTTP_WEBHOOK_TOKEN")

	return api_token_from_request == valid_api_token
}

func PassLogsToQueue(requestBody AwsVpcLogWebhookModel) {
	records := requestBody.Records
	for _, record := range records {
		base64_log_data := record.Data
		println(base64_log_data)

	}

}
