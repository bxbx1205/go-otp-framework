package models

type SendOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
}
