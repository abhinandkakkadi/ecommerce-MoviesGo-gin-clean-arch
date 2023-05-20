package domain

type Admin struct {
	ID 		uint		`json:"id" gorm:"uniquekey; not null"`
	Name 	string	`json:"name" gorm:"validate:required"`
	Email		string	`json:"email" gorm:"validate:required"`
	Password  string	`json:"password" gorm:"validate:required"`
}