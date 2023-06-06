package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	SignUpHandler(admin models.AdminSignUp) (domain.TokenAdmin, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetFullCategory() (domain.CategoryResponse, error)
	AddCategory(genre models.CategoryUpdate) (domain.CategoryManagement, error)
	Delete(genre_id string) error
	BlockUser(id string) error
	UnBlockUser(id string) error
}
