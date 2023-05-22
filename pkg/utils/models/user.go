package models

type UserDetails struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Email		string `json:"email" validate:"email"`
	Phone			string	`json:"phone"`
}