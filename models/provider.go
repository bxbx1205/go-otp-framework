package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Provider struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID             primitive.ObjectID `bson:"user_id" json:"user_id"`
	Provider           string             `bson:"provider" json:"provider"` // "twilio" or "aws"
	AccountSID         string             `bson:"account_sid,omitempty" json:"account_sid,omitempty"`
	EncryptedAuthToken string             `bson:"encrypted_auth_token,omitempty" json:"-"`
	PhoneNumber        string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	AccessKey          string             `bson:"access_key,omitempty" json:"access_key,omitempty"`
	EncryptedSecretKey string             `bson:"encrypted_secret_key,omitempty" json:"-"`
	Region             string             `bson:"region,omitempty" json:"region,omitempty"`
}

type UpsertProviderRequest struct {
	Provider string `json:"provider" binding:"required"` // "twilio" or "aws"

	// Twilio fields
	AccountSID  string `json:"account_sid"`
	AuthToken   string `json:"auth_token"`
	PhoneNumber string `json:"phone_number"`

	// AWS fields
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}
