package workers

import (
	"encoding/json"
	"fmt"
	"log"
	"otp-service/config"
	"otp-service/models"
)

const smsQueue = "sms_queue"

func StartSMSWorker() {

	fmt.Println("SMS working started")

	for{
		result,err:= config.RedisClient.BLPop(config.Ctx,0,smsQueue).Result()

		if err!=nil {
			log.Println(
				"Worker error",
				err,
			)

			continue
		}


		jobData:=result[1]

		var job models.SMSJob

		err=json.Unmarshal(
			[]byte(jobData),
			&job,
		)

		if err!=nil {
			log.Println(
				"json parse error: ",
				err,
			)

			continue
		}

		fmt.Printf("Sending OTP %s to %s \n",job.OTP,job.Phone)
	}
}