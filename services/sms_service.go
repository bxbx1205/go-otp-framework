package services

import "fmt"

type SMSProvider interface {
	Send(phone string, otp string) error
}

var activeProviders []SMSProvider

func RegisterProvider(p SMSProvider) {
	activeProviders = append(activeProviders, p)
}

func SendSMS(phone string, otp string) error {
	for _, p := range activeProviders {
		err := p.Send(phone, otp)
		if err == nil {
			return nil
		}
		fmt.Printf("Provider failed, attempting next. Error: %v\n", err)
	}
	return fmt.Errorf("all sms providers failed")
}
