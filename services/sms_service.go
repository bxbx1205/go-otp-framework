package services

import (
	"errors"
	"os"
)

func SendSMS(
	phone string,
	otp string,
) error {

	provider := os.Getenv(
		"SMS_PROVIDER",
	)

	switch provider {

	case "aws":

		return SendSMSAWS(
			phone,
			otp,
		)

	case "twilio":

		return SendSMSTwilio(
			phone,
			otp,
		)

	default:

		return errors.New(
			"invalid SMS provider",
		)
	}
}

