package services

import (
	"errors"
	"fmt"
	
	"time"
	"otp-service/metrics"
	"otp-service/config"
	"otp-service/models"
	"otp-service/repositories"
	"otp-service/utils"
)

const (
	otpTTL = 5 * time.Minute

	cooldownTTL = 30 * time.Second

	maxOTPRequest = 5

	rateLimitTTL = 1 * time.Minute

	maxVerifyAttempts = 5
)

func SendOTP(phone string) error {

	rateKey := fmt.Sprintf(
		"rate:%s",
		phone,
	)

	requestCount, err := config.RedisClient.Incr(
		config.Ctx,
		rateKey,
	).Result()

	if err != nil {
		return err
	}

	if requestCount == 1 {

		err = config.RedisClient.Expire(
			config.Ctx,
			rateKey,
			rateLimitTTL,
		).Err()

		if err != nil {
			return err
		}
	}

	if requestCount > maxOTPRequest {

		return errors.New(
			"too many OTP requests",
		)
	}

	cooldownKey := fmt.Sprintf(
		"cooldown:%s",
		phone,
	)

	exists, err := config.RedisClient.Exists(
		config.Ctx,
		cooldownKey,
	).Result()

	if err != nil {
		return err
	}

	if exists == 1 {

		return errors.New(
			"please wait before requesting another OTP",
		)
	}

	otp := utils.GenerateOTP()

	hashedOTP, err := utils.HashOTP(otp)

	if err != nil {
		return err
	}

	key := fmt.Sprintf(
		"otp:%s",
		phone,
	)

	err = config.RedisClient.Set(
		config.Ctx,
		key,
		hashedOTP,
		otpTTL,
	).Err()

	if err != nil {
		return err
	}

	err = config.RedisClient.Set(
		config.Ctx,
		cooldownKey,
		"true",
		cooldownTTL,
	).Err()

	if err != nil {
		return err
	}

	smsJob := models.SMSJob{

		Phone: phone,

		OTP: otp,
	}

	err = PushSMSJob(smsJob)

	if err != nil {
		return err
	}

	
	metrics.IncrementOTPSent()

	return nil
}

func VerifyOTP(
	phone string,
	enteredOTP string,
) error {

	key := fmt.Sprintf(
		"otp:%s",
		phone,
	)

	storedHash, err := config.RedisClient.Get(
		config.Ctx,
		key,
	).Result()

	if err != nil {

		return errors.New(
			"OTP expired or not found",
		)
	}

	attemptKey := fmt.Sprintf(
		"attempts:%s",
		phone,
	)

	attempts, err := config.RedisClient.Get(
		config.Ctx,
		attemptKey,
	).Int()

	if err != nil &&
		err.Error() != "redis: nil" {

		return err
	}

	if attempts >= maxVerifyAttempts {

		return errors.New(
			"too many failed attempts",
		)
	}

	err = utils.CompareOTP(
		storedHash,
		enteredOTP,
	)

	if err != nil {

		failedCount, redisErr :=
			config.RedisClient.Incr(
				config.Ctx,
				attemptKey,
			).Result()

		if redisErr == nil &&
			failedCount == 1 {

			config.RedisClient.Expire(
				config.Ctx,
				attemptKey,
				5*time.Minute,
			)
		}

		failedLog := models.OTPLog{
			Phone:     phone,
			Status:    "failed",
			Timestamp: time.Now(),
		}

		repositories.CreateOTPLog(failedLog)

		return errors.New(
			"invalid OTP",
		)
	}

	err = config.RedisClient.Del(
		config.Ctx,
		key,
	).Err()

	if err != nil {
		return err
	}

	config.RedisClient.Del(
		config.Ctx,
		attemptKey,
	)

	user := models.User{
		Phone:     phone,
		Verified:  true,
		CreatedAt: time.Now(),
	}

	err = repositories.CreateUser(user)

	if err != nil {
		return err
	}
	metrics.IncrementOTPVerified()

	successLog := models.OTPLog{
		Phone:     phone,
		Status:    "verified",
		Timestamp: time.Now(),
	}

	repositories.CreateOTPLog(successLog)

	return nil
}
