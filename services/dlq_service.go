package services

import (
	"encoding/json"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/models"
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
