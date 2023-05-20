package models

type OTPData struct {
	PhoneNumber string	`json:"phone,omitempty" validate:"required"`
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string		`json:"code,omitempty" validate:"required"`
}

