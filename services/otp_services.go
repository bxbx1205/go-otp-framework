package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/metrics"
	"github.com/bxbx1205/go-otp-framework/models"
	"github.com/bxbx1205/go-otp-framework/repositories"
	"github.com/bxbx1205/go-otp-framework/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	otpTTL = 5 * time.Minute

	cooldownTTL = 30 * time.Second

	maxOTPRequest = 5

	rateLimitTTL = 1 * time.Minute

	maxVerifyAttempts = 5
)

func SendOTP(userID string, phone string) error {
	phone = strings.TrimSpace(phone)

	rateKey := fmt.Sprintf(
		"rate:%s:%s",
		userID,
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
		"cooldown:%s:%s",
		userID,
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
		"otp:%s:%s",
		userID,
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

	objID, _ := primitive.ObjectIDFromHex(userID)
	smsJob := models.SMSJob{
		UserID: objID,
		Phone:  phone,
		OTP:    otp,
	}

	err = PushSMSJob(smsJob)

	if err != nil {
		return err
	}

	metrics.IncrementOTPSent()

	return nil
}

func VerifyOTP(
	userID string,
	phone string,
	enteredOTP string,
) (string, error) {

	phone = strings.TrimSpace(phone)
	enteredOTP = strings.TrimSpace(enteredOTP)

	key := fmt.Sprintf(
		"otp:%s:%s",
		userID,
		phone,
	)

	storedHash, err := config.RedisClient.Get(
		config.Ctx,
		key,
	).Result()

	if err != nil {

		return "", errors.New(
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

		return "", err
	}

	if attempts >= maxVerifyAttempts {

		return "", errors.New(
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

		return "", errors.New(
			"invalid OTP",
		)
	}

	err = config.RedisClient.Del(
		config.Ctx,
		key,
	).Err()

	if err != nil {
		return "", err
	}

	config.RedisClient.Del(
		config.Ctx,
		attemptKey,
	)

	metrics.IncrementOTPVerified()

	successLog := models.OTPLog{
		Phone:     phone,
		Status:    "verified",
		Timestamp: time.Now(),
	}

	repositories.CreateOTPLog(successLog)

	token := "valid"

	return token, nil
}
