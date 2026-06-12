package repositories

import (
	"context"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers() ([]models.User, error) {

	var users []models.User

	cursor, err :=
		config.UserCollection.Find(
			context.Background(),
			bson.M{},
		)

	if err != nil {
		return nil, err
	}

	err = cursor.All(
		context.Background(),
		&users,
	)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user models.User) error {
	_, err := config.UserCollection.InsertOne(context.Background(), user)
	return err
}

func FindByEmail(email string) (models.User, error) {
	var user models.User
	err := config.UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func FindByID(id string) (models.User, error) {
	var user models.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = config.UserCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	return user, err
}
