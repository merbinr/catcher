package web

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Webhook route
	r.POST("aws/logs/vpc/webhook", AwsVpcLogWebhookHandler)

	return r
}
