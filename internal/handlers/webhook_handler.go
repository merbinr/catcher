package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	// Read the request headers
	headers := c.Request.Header

	// Read the request body
	body, err := io.ReadAll(c.Request.Body) // Using io.ReadAll
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Log the headers and body
	fmt.Println("Headers:", headers)
	fmt.Println("Body:", string(body))

	// Respond with 201 Created
	c.JSON(http.StatusCreated, gin.H{
		"message": "Webhook received",
	})
}
