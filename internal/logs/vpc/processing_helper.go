package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/merbinr/catcher/internal/logs/vpc/aws"
	"github.com/merbinr/catcher/internal/models"
	"github.com/merbinr/catcher/internal/rabbitmq"
)

func AwsVpcLogProcessing(WebhookData models.AwsVpcLogWebhookModel) error {
	for _, each_log_records := range WebhookData.Records {
		normalized_data, err := aws.AwsVpcLogFlowLogParsing(each_log_records)
		if err != nil {
			return fmt.Errorf("unable to process the aws vpc log, request id: %s, record: %s, error: %s",
				WebhookData.RequestId, each_log_records.Data, err)
		}
		err = logProcessing(normalized_data)
		if err != nil {
			return fmt.Errorf("unable to process the AWS VPC log, error: %s, request_id: %s",
				err, WebhookData.RequestId)
		}
	}
	return nil
}

func logProcessing(normalized_data models.VpcNormalizedData) error {

	log_data_json, err := json.Marshal(normalized_data)
	if err != nil {
		return err
	}
	// Send message to queue
	err = rabbitmq.Rabbitmq_conn.SendLogMessage(log_data_json)
	if err != nil {
		return fmt.Errorf("unable to send log message to queue, error: %s, request", err)
	}
	return nil
}
