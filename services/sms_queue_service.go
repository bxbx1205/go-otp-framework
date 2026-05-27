package services

import (
	"encoding/json"
	"otp-service/config"
	"otp-service/models"
)

const smsQueue = "sms_queue"

func PushSMSJob(job models.SMSJob) error{

	jobJson,err:= json.Marshal(job)

	if err!=nil {
		return err
	}

	err=config.RedisClient.RPush(config.Ctx,smsQueue,jobJson).Err()

	if err != nil {
		return err
	}


	return nil
}