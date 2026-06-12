package handlers

import (
	"net/http"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/services"

	"github.com/gin-gonic/gin"
)

func QueueStatus(c *gin.Context) {

	queueSize, _ := config.RedisClient.LLen(
		config.Ctx,
		services.SMSQueue,
	).Result()

	dlqSize, _ := config.RedisClient.LLen(
		config.Ctx,
		services.SMSDLQ,
	).Result()

	c.JSON(
		http.StatusOK,
		gin.H{
			"queue_size": queueSize,
			"dlq_size":   dlqSize,
		},
	)
}
