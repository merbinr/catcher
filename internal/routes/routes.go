package routes

import (
	"github.com/merbinr/catcher/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Example route
	r.GET("/test", handlers.TestHandler)

	// Webhook route
	r.POST("/webhook", handlers.WebhookHandler)

	return r
}
