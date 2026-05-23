package services

import (
	"errors"
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
	fmt.Println("Generated OTP:", otp)


	return nil
}

func VerifyOTP(phone string,enteredOTP string,) (error){
	key:=fmt.Sprintf("otp:%s",phone)


	storedHash,err:=config.RedisClient.Get(config.Ctx,key,).Result()

	if err != nil {
		return errors.New("OTP expired or not found")
	}

	err = utils.CompareOTP(
		storedHash,
		enteredOTP,
	)

	if err != nil {

		return errors.New("invalid OTP")
	}

	err=config.RedisClient.Del(config.Ctx,key).Err()

	if err != nil {
		return err
	}


	return nil
}