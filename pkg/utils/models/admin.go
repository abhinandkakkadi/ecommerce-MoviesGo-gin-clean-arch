package models


type AdminLogin struct {
	Email				string		`json:"email,omitempty" validate:"required"`
	Password		string 		`json:"password" validate:"min=8,max=20"`
}

type AdminDetails struct {
	ID 		uint		`json:"id" gorm:"uniquekey; not null"`
	Name 	string	`json:"name" gorm:"validate:required"`
	Email		string	`json:"email" gorm:"validate:required"`
}