package services

import (
	"encoding/json"

	"github.com/myusername/otp-framework/config"
	"github.com/myusername/otp-framework/models"
)

func PushToDLQ(
	job models.SMSJob,
) error {

	jobJSON, err := json.Marshal(job)

	if err != nil {
		return err
	}

	return config.RedisClient.RPush(
		config.Ctx,
		SMSDLQ,
		jobJSON,
	).Err()
}
