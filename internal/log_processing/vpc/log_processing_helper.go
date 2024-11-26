package vpc

import (
	"fmt"

	"github.com/merbinr/catcher/internal/log_processing/models"
	"github.com/merbinr/catcher/internal/log_processing/vpc/aws"
	"github.com/merbinr/catcher/internal/web"
)

func AwsVpcLogProcessing(WebhookData web.AwsVpcLogWebhookModel) error {
	for _, each_log_records := range WebhookData.Records {
		processed_data, err := aws.AwsVpcLogFlowLogParsing(each_log_records)
		if err != nil {
			return fmt.Errorf("unable to process the aws vpc log, request id: %s, record: %s, error: %s",
				WebhookData.RequestId, each_log_records.Data, err)
		}
		fmt.Printf("%+v\n", processed_data)
	}
	return nil
}

func logProcessing(normalized_data models.VpcNormalizedProcessedData) error {
	// Redis_client.Get()
	return nil
}
