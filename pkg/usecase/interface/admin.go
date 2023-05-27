package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails domain.Admin) (domain.TokenAdmin, error)
	SignUpHandler(admin models.AdminSignUp) (domain.TokenAdmin, error)
	GetUsers(page int) ([]models.UserDetailsResponse, error)
	GetGenres() ([]domain.Genre, error)
	AddCategory(genre domain.CategoryManagement) (domain.CategoryManagement, error)
	Delete(genre_id string) error
	BlockUser(id string) error
}
