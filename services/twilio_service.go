package services

import (
	"fmt"
	"os"

	"otp-service/config"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMSTwilio(
	phone string,
	otp string,
) error {

	message := fmt.Sprintf(
		"Your OTP is %s",
		otp,
	)

	params := &openapi.CreateMessageParams{}

	params.SetTo(phone)

	params.SetFrom(
		os.Getenv("TWILIO_PHONE_NUMBER"),
	)

	params.SetBody(message)

	_, err := config.TwilioClient.
		Api.
		CreateMessage(params)

	if err != nil {
		return err
	}

	return nil
}