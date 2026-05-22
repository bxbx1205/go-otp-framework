package services

import (
	"fmt"
	"otp-service/config"
	"otp-service/utils"
	"time"
)

const otpTTL = time.Minute*5

func SendOTP(phone string) error{

	otp:= utils.GenerateOTP()

	hashedOTP, err := utils.HashOTP(otp)

	if err != nil {
		return err
	}

	key:= fmt.Sprintf("otp:%s",phone)

	err = config.RedisClient.Set(config.Ctx,key,hashedOTP,otpTTL).Err()

	if err != nil {
		return err
	}


	// Print OTP temporarily
	// Later SMS provider will send this
	fmt.Println("Generated OTP:", otp)


	return nil
}