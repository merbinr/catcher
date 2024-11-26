package routes

import (
	"github.com/merbinr/catcher/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Webhook route
	r.POST("aws/logs/vpc/webhook", handlers.AwsVpcLogWebhookHandler)

	return r
}
