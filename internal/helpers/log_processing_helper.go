package helpers

import (
	"fmt"
	"os"

	"github.com/cloudflare/cfssl/log"
	"github.com/merbinr/catcher/internal/helpers/aws"
	"github.com/merbinr/catcher/internal/models"
)

func AwsVpcLogProcessing(WebhookData models.AwsVpcLogWebhookModel) error {
	for _, each_log_records := range WebhookData.Records {
		processed_data, err := aws.AwsVpcLogFlowLogParsing(each_log_records)
		if err != nil {
			log.Errorf("unable to process the aws vpc log, request id: %s, record: %s, error: %s",
				WebhookData.RequestId, each_log_records.Data, err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", processed_data)
	}
	return nil
}
