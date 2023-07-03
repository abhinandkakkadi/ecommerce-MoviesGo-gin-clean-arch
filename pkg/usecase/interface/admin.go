package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	CreateAdmin(admin models.AdminSignUp) (domain.TokenAdmin, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetGenres() ([]domain.Genre, error)
	AddGenres(genre models.CategoryUpdate) error
	Delete(genre_id string) error
	BlockUser(id string) error
	UnBlockUser(id string) error
	FilteredSalesReport(timePeriod string) (models.SalesReport, error)
	DashBoard() (models.CompleteAdminDashboard, error)
}
