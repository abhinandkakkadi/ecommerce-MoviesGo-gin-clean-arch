package interfaces

import (
	"time"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error)
	CheckAdminAvailability(admin models.AdminSignUp) bool
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetGenres() ([]domain.Genre, error)
	AddGenre(genre models.CategoryUpdate) error
	Delete(genre_id string) error
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error)
	TotalRevenue() (models.DashboardRevenue, error)
	DashBoardOrder() (models.DashboardOrder, error)
	AmountDetails() (models.DashboardAmount, error)
	DashboardUserDetails() (models.DashboardUser, error)
	DashBoardProductDetails() (models.DashBoardProduct, error)
}
