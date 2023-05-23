package domain

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type Users struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
	Phone    string `json:"phone"`
	Blocked  bool   `json:"blocked" gorm:"default:false"`
}

type TokenUsers struct {
	Users models.UserDetails
	Token string
}
