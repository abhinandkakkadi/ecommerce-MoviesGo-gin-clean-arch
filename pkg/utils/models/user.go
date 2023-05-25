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
	Id 	 int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
	Password	string	`json:"password"`
}