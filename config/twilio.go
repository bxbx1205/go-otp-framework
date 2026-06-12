package config

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
)

var TwilioClient *twilio.RestClient

func ConnectTwilio(sid string, token string) {

	fmt.Println("Creating Twilio Client...")
	fmt.Println("SID:", sid)

	TwilioClient = twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: sid,
			Password: token,
		},
	)
}
