package services

import (
	"encoding/json"

	"otp-service/config"
	"otp-service/models"
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