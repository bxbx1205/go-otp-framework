package models

type VerifyOTPRequest struct{
	Phone string `json:"phone"`

	OTP string `json:"otp"`
}