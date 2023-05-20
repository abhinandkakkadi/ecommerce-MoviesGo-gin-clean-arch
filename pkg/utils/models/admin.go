package models


type AdminLogin struct {
	Email				string		`json:"email,omitempty" validate:"required"`
	Password		string 		`json:"password,omitempty" validate:"required"`
}