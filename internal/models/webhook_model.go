package models

type AwsVpcLogWebhookModel struct {
	RequestId string                 `json:"requestId" binding:"required"`
	Timestamp int                    `json:"timestamp" binding:"required"`
	Records   []AwsVpcLogRecordsData `json:"records" binding:"required"`
}

type AwsVpcLogRecordsData struct {
	Data string `json:"data"`
}
