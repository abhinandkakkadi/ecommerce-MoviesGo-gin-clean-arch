package models

type UserDetails struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
	Password	string	`json:"password"`
	ConfirmPassword	string	`json:"confirm-password"`
}

type UserDetailsResponse struct {
	Id 	 int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string	`json:"phone"`
}


type TokenUsers struct {
	Users UserDetailsResponse
	Token string
}

type UserSignInResponse struct {
	Id 	 uint `json:"id"`
	UserID	 uint  	`json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
	Password	string	`json:"password"`
}

type AddressInfo struct {

	UserID	 uint  	`json:"user_id"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string	`json:"pin" validate:"required"`
	Street    string  `json:"street"`
	City      string  `json:"city"`

}

type AddressInfoResponse struct {
	ID			 uint   `json:"id"`
	UserID	 uint  	`json:"user_id"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string	`json:"pin" validate:"required"`
	Street    string  `json:"street"`
	City      string  `json:"city"`

}

