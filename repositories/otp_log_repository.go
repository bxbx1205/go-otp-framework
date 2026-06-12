package repositories

import (
	"context"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllOTPLogs() (
	[]models.OTPLog,
	error,
) {

	var logs []models.OTPLog
	cursor, err :=
		config.OTPLogCollection.Find(
			context.Background(),
			bson.M{},
		)

	if err != nil {
		return nil, err
	}

	err = cursor.All(
		context.Background(),
		&logs,
	)

	if err != nil {
		return nil, err
	}

	return logs, nil
}

func CreateOTPLog(log models.OTPLog) error {
	_, err := config.OTPLogCollection.InsertOne(context.Background(), log)
	return err
}
