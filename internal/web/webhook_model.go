package web

type AwsVpcLogWebhookModel struct {
	RequestId string                 `json:"requestId" binding:"required"`
	Timestamp int64                  `json:"timestamp" binding:"required"`
	Records   []AwsVpcLogRecordsData `json:"records" binding:"required"`
}

type AwsVpcLogRecordsData struct {
	Data string `json:"data"`
}
