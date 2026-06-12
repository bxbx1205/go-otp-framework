package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTPLog struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Phone string `bson:"phone" json:"phone"`

	Status string `bson:"status" json:"status"`

	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
