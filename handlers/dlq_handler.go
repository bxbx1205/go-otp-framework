package handlers

import (
	"net/http"

	"otp-service/config"
	"otp-service/services"

	"github.com/gin-gonic/gin"
)

func GetDLQ(c *gin.Context) {

	jobs, err := config.RedisClient.LRange(
		config.Ctx,
		services.SMSDLQ,
		0,
		-1,
	).Result()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"jobs": jobs,
		},
	)
}