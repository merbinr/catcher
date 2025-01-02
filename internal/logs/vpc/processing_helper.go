package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/merbinr/catcher/internal/logs/vpc/aws"
	"github.com/merbinr/catcher/internal/models"
	deduplicator_queue "github.com/merbinr/catcher/internal/queue/deduplicator"
	"github.com/merbinr/catcher/pkg/logger"
	log_models "github.com/merbinr/log_models/models"
)

func AwsVpcLogProcessing(WebhookData models.AwsVpcLogWebhookModel) error {
	logger := logger.GetLogger()
	logger.Info(fmt.Sprintf("processing AWS VPC log, request id: %s", WebhookData.RequestId))
	for _, each_log_records := range WebhookData.Records {
		normalized_data, err := aws.AwsVpcLogFlowLogParsing(each_log_records)
		if err != nil {
			return fmt.Errorf("unable to parse the aws vpc log, request id: %s, record: %s, error: %s",
				WebhookData.RequestId, each_log_records.Data, err)
		}
		err = logProcessing(normalized_data)
		if err != nil {
			return fmt.Errorf("unable to process the AWS VPC log, error: %s, request_id: %s",
				err, WebhookData.RequestId)
		}
	}
	logger.Info(fmt.Sprintf("AWS VPC log processed successfully, request id: %s", WebhookData.RequestId))
	return nil
}

func logProcessing(normalized_data log_models.VpcNormalizedData) error {
	log_data_json, err := json.Marshal(normalized_data)
	if err != nil {
		return err
	}
	// Send message to dedupe queue
	err = deduplicator_queue.DedupeRabbitmqConn.SendLogMessageToDedupeQueue(&log_data_json)
	if err != nil {
		return fmt.Errorf("unable to send log message to deduplicator queue, error: %s, request", err)
	}
	return nil
}
