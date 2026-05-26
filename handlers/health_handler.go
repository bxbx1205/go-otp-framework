package handlers

import (
	"net/http"
	"otp-service/config"
	"otp-service/utils"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context){

	_,redisErr := config.RedisClient.Ping(config.Ctx).Result()

	mongoErr := config.MongoClient.Ping(
		config.Ctx,
		nil,
	)

	if redisErr != nil || mongoErr != nil {

		utils.ErrorResponse(
			c,
			http.StatusServiceUnavailable,
			"service unavailable",
		)

		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"service healthy",
		nil,
	)

}