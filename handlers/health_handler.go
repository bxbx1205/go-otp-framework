package handlers

import (
	"net/http"

	"github.com/myusername/otp-framework/config"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {

	redisStatus := "up"

	err := config.RedisClient.Ping(
		config.Ctx,
	).Err()

	if err != nil {
		redisStatus = "down"
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "healthy",
			"redis":   redisStatus,
			"mongodb": "up",
		},
	)
}
