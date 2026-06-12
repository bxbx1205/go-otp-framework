package services

import (
	"fmt"
	"os"

	"github.com/bxbx1205/go-otp-framework/config"

	"github.com/twilio/twilio-go"
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

func SendSMSTwilioDynamic(phone string, otp string, accountSID string, authToken string, fromPhone string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	message := fmt.Sprintf("Your OTP is %s", otp)

	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(fromPhone)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	return err
}
