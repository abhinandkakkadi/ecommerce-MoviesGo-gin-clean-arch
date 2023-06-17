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
	// GetDirectors() ([]domain.Directors, error)
	// GetMovieFormat() ([]domain.Movie_Format, error)
	// GetMovieLanguages() ([]domain.Movie_Language, error)
	AddGenre(genre models.CategoryUpdate) error
	// AddDirector(director string) (domain.Directors, error)
	// AddFormat(format string) (domain.Movie_Format, error)
	// AddLanguage(language string) (domain.Movie_Language, error)
	Delete(genre_id string) error
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error

	FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error)
	// CategoryCount(category models.CategoryUpdate) (models.CategoryUpdateCheck, error)

	TotalRevenue() (models.DashboardRevenue, error)
	DashBoardOrder() (models.DashboardOrder, error)
	AmountDetails() (models.DashboardAmount, error)
	DashboardUserDetails() (models.DashboardUser, error)
	DashBoardProductDetails() (models.DashBoardProduct, error)
}
