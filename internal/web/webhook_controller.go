package web

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merbinr/catcher/internal/logs/vpc"
	"github.com/merbinr/catcher/internal/models"
)

func AwsVpcLogWebhookHandler(c *gin.Context) {
	// checking authentication
	headers := c.Request.Header
	authentication_success := CheckAuthentication(headers)

	if !authentication_success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
	}

	// reading body
	var request_body models.AwsVpcLogWebhookModel
	if err := c.ShouldBindBodyWithJSON(&request_body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Schema!"})
		return
	}

	// passing to queue
	err := vpc.AwsVpcLogProcessing(request_body)
	if err != nil {
		slog.Error(fmt.Sprintf("error occured on processing %s log, error: %s",
			request_body.RequestId, err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unexpected error occuered!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
