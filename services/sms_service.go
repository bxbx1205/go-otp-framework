package services

import (
	"context"
	"fmt"
	"otp-service/config"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func SendSMS(
	phone string, otp string,
) error {

	message := fmt.Sprint("Your OTP is : %s",otp)

	input:=&sns.PublishInput{
		Message: &message,
		PhoneNumber: &phone,
	}

	_,err:= config.SNSClient.Publish(
		context.Background(),
		input,
	)

	if err!=nil {
		return err
	}

	return nil
}