package services

import (
	"context"
	"fmt"
	"github.com/bxbx1205/go-otp-framework/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func SendSMSAWS(
	phone string, otp string,
) error {

	message := fmt.Sprintf("Your OTP is : %s", otp)

	input := &sns.PublishInput{
		Message:     &message,
		PhoneNumber: &phone,
	}

	_, err := config.SNSClient.Publish(
		context.Background(),
		input,
	)

	if err != nil {
		return err
	}

	return nil
}

func SendSMSAWSDynamic(phone string, otp string, accessKey string, secretKey string, region string) error {
	cfg := aws.Config{
		Region:      region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}

	client := sns.NewFromConfig(cfg)

	message := fmt.Sprintf("Your OTP is : %s", otp)

	input := &sns.PublishInput{
		Message:     &message,
		PhoneNumber: &phone,
	}

	_, err := client.Publish(context.Background(), input)
	return err
}
