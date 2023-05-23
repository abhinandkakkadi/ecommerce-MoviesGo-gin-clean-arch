package domain

type OTPData struct {
	PhoneNumber string `json:"phon_number,omitempty" validate:"required"`
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
