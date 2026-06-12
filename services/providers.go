package services

func init() {
	
	RegisterProvider(&TwilioProvider{})
	RegisterProvider(&AWSProvider{})
}

type TwilioProvider struct{}

func (t *TwilioProvider) Send(phone string, otp string) error {
	return SendSMSTwilio(phone, otp)
}

type AWSProvider struct{}

func (a *AWSProvider) Send(phone string, otp string) error {
	return SendSMSAWS(phone, otp)
}
