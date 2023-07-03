package models

type OTPData struct {
	PhoneNumber string `json:"phone" binding:"required" validate:"required"`
}

type VerifyData struct {
	User *OTPData `json:"user" binding:"required" validate:"required"`
	Code string   `json:"code" binding:"required" validate:"required"`
}

type Email struct {
	Email string `json:"email" binding:"required" validate:"required"`
}

type OTPCode struct {
	Code string `json:"code" binding:"required" validate:"required"`
}
