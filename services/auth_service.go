package services

import (
	"errors"
	"time"

	"github.com/bxbx1205/go-otp-framework/models"
	"github.com/bxbx1205/go-otp-framework/repositories"
	"github.com/bxbx1205/go-otp-framework/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req models.RegisterRequest) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	return repositories.CreateUser(user)
}

func LoginUser(req models.LoginRequest) (string, error) {
	// Find user by email
	user, err := repositories.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT based on user ID instead of phone
	return utils.GenerateJWT(user.ID.Hex())
}
