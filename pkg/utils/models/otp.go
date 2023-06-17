package models

type OTPData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}


type Email struct {
	Email string `json:"email" validate:"required"`
}

type OTPCode struct {
	Code string `json:"code" validate:"required"` 
}