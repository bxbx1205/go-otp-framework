package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func GenerateOTP() string {

	max := big.NewInt(900000)
	n, _ := rand.Int(rand.Reader, max)

	otp := n.Int64() + 100000
	return fmt.Sprintf("%d", otp)
}

func HashOTP(otp string) (string,error){

	otpBytes:=[]byte(otp)

	hashedOTP,err := bcrypt.GenerateFromPassword(
		otpBytes,
		bcrypt.DefaultCost,
	)

	if err!=nil {
		return "", err
	}
	return string(hashedOTP), nil
}

func CompareOTP(hashedOTP string,plainOTP string) (error) {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedOTP),
		[]byte(plainOTP),
	)
}