package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merbinr/catcher/internal/helpers"
	"github.com/merbinr/catcher/internal/models"
)

func AwsVpcLogWebhookHandler(c *gin.Context) {
	// checking authentication
	headers := c.Request.Header
	authentication_success := helpers.CheckAuthentication(headers)

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

	fmt.Printf("Parsed Request: %+v\n", request_body)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
