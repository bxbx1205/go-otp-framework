package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/metrics"
	"github.com/bxbx1205/go-otp-framework/models"
	"github.com/bxbx1205/go-otp-framework/repositories"
	"github.com/bxbx1205/go-otp-framework/services"
	"github.com/bxbx1205/go-otp-framework/utils"
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

		fmt.Println("Parsed Job:", job)
		fmt.Println("Calling SMS Provider dynamically...")

		providerConfig, err := repositories.GetProviderByUserID(job.UserID.Hex(), "twilio")
		if err != nil {
			providerConfig, err = repositories.GetProviderByUserID(job.UserID.Hex(), "aws")
		}

		if err == nil {
			if providerConfig.Provider == "twilio" {
				decryptedToken, decErr := utils.Decrypt(providerConfig.EncryptedAuthToken)
				if decErr == nil {
					err = services.SendSMSTwilioDynamic(job.Phone, job.OTP, providerConfig.AccountSID, decryptedToken, providerConfig.PhoneNumber)
				} else {
					err = decErr
				}
			} else if providerConfig.Provider == "aws" {
				decryptedSecret, decErr := utils.Decrypt(providerConfig.EncryptedSecretKey)
				if decErr == nil {
					err = services.SendSMSAWSDynamic(job.Phone, job.OTP, providerConfig.AccessKey, decryptedSecret, providerConfig.Region)
				} else {
					err = decErr
				}
			} else {
				err = fmt.Errorf("unknown provider configuration found")
			}
		} else {
			// Fallback to default if no provider configured
			err = services.SendSMS(job.Phone, job.OTP)
		}

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
