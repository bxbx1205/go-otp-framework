package config

import (
	"fmt"
	"os"

	twilio "github.com/twilio/twilio-go"
)

var TwilioClient *twilio.RestClient

func ConnectTwilio() {

	sid := os.Getenv("TWILIO_ACCOUNT_SID")
	token := os.Getenv("TWILIO_AUTH_TOKEN")

	fmt.Println("Creating Twilio Client...")
	fmt.Println("SID:", sid)

	TwilioClient = twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: sid,
			Password: token,
		},
	)
}
