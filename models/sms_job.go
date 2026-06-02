package models


type SMSJob struct {

	Phone string `json:"phone"`

	OTP string `json:"otp"`

	RetryCount int `json:"retry_count"`
}