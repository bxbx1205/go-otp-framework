package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SMSJob struct {
	UserID primitive.ObjectID `json:"user_id"`

	Phone string `json:"phone"`

	OTP string `json:"otp"`

	RetryCount int `json:"retry_count"`
}
