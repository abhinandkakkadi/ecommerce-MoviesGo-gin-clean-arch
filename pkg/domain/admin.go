package domain

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"

type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	Email    string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}

type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
}

type CategoryManagement struct {
	Genre    Genre          `json:"genre"`
	Director Directors      `json:"director"`
	Format   Movie_Format   `json:"format"`
	Language Movie_Language `json:"language"`
}

type CategoryResponse struct {
	Genre          []Genre
	Directors      []Directors
	Movie_Format   []Movie_Format
	Movie_Language []Movie_Language
}
