package services

import (
	"encoding/json"
	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/models"
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
