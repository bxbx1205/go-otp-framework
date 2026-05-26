package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Phone string `bson:"phone" json:"phone"`

	Verified bool `bson:"verified" json:"verified"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}