package domain

type Users struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Email		string `json:"email" validate:"email"`
	Password 	string	`json:"password" validate:"min=8,max=20"`
	Phone			string	`json:"phone"`
}


type TokenUsers struct {
	Users Users
	Token string
}