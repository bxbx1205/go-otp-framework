package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"otp-service/config"
	"otp-service/metrics"
	"otp-service/models"
	"otp-service/services"
)

const smsQueue = "sms_queue"

func StartSMSWorker() {

	fmt.Println("SMS worker started")

	for {

		fmt.Println("Waiting for queue job...")

		result, err := config.RedisClient.BLPop(
			context.Background(),
			0,
			smsQueue,
		).Result()

		if err != nil {

			log.Println(
				"BLPOP error:",
				err,
			)

			continue
		}

		fmt.Println(
			"Raw Redis result:",
			result,
		)

		if len(result) < 2 {

			log.Println(
				"Invalid BLPOP response",
			)

			continue
		}

		jobData := result[1]

		var job models.SMSJob

		err = json.Unmarshal(
			[]byte(jobData),
			&job,
		)

		if err != nil {

			log.Println(
				"JSON parse error:",
				err,
			)

			continue
		}

		fmt.Println(
			"Parsed Job:",
			job,
		)

		fmt.Println(
			"Calling SMS Provider...",
		)

		err = services.SendSMS(
			job.Phone,
			job.OTP,
		)

		if err != nil {

			// Metrics
			metrics.IncrementSMSFailed()

			log.Println(
				"SMS sending failed:",
				err,
			)

			// Increment retry count
			job.RetryCount++

			// Retry if limit not reached
			if job.RetryCount <= services.MaxRetries {

				log.Printf(
					"Retrying job (%d/%d)",
					job.RetryCount,
					services.MaxRetries,
				)

				retryErr := services.RetrySMSJob(job)

				if retryErr != nil {

					log.Println(
						"Retry queue push failed:",
						retryErr,
					)
				}

				continue
			}

			// Move to DLQ
			log.Println(
				"Max retries reached. Moving to DLQ",
			)

			dlqErr := services.PushToDLQ(job)

			if dlqErr != nil {

				log.Println(
					"DLQ push failed:",
					dlqErr,
				)
			}

			continue
		}

		// Metrics
		metrics.IncrementSMSSuccess()

		fmt.Printf(
			"OTP sent successfully to %s\n",
			job.Phone,
		)
	}
}