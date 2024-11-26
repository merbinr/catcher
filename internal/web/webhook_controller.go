package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AwsVpcLogWebhookHandler(c *gin.Context) {
	// checking authentication
	headers := c.Request.Header
	authentication_success := CheckAuthentication(headers)

	if !authentication_success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
	}

	// reading body
	var request_body AwsVpcLogWebhookModel
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
