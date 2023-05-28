package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails domain.Admin) (domain.TokenAdmin, error)
	SignUpHandler(admin models.AdminSignUp) (domain.TokenAdmin, error)
	GetUsers(page int) ([]models.UserDetailsResponse, error)
	GetFullCategory() (domain.CategoryResponse, error)
	AddCategory(genre models.CategoryUpdate) (domain.CategoryManagement, error)
	Delete(genre_id string) error
	BlockUser(id string) error
}
