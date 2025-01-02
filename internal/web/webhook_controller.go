package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/merbinr/catcher/internal/logs/vpc"
	"github.com/merbinr/catcher/internal/models"
	"github.com/merbinr/catcher/pkg/logger"
)

func AwsVpcLogWebhookHandler(c *gin.Context) {
	logger := logger.GetLogger()
	// checking authentication
	headers := c.Request.Header
	authentication_success := CheckAuthentication(headers)

	if !authentication_success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// reading body
	var request_body models.AwsVpcLogWebhookModel

	err := c.ShouldBind(&request_body)
	if err != nil {
		logger.Error(fmt.Sprintf("error occured on binding request body, error: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	// passing to queue
	err = vpc.AwsVpcLogProcessing(request_body)
	if err != nil {
		if strings.Contains(err.Error(), "illegal base64 data at input") {
			// error on base64 decoding, should return unprocessable entity 422
			logger.Error(fmt.Sprintf("invalid base64 data on request %s, error: %s",
				request_body.RequestId, err))

			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unprocessable entity"})
			return
		} else {
			// unknown error, should return internal server error 500
			logger.Error(fmt.Sprintf("error occured on processing %s log error: %s",
				request_body.RequestId, err))

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
