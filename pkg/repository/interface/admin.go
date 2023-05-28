package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails domain.Admin) (domain.Admin, error)
	SignUpHandler(admin models.AdminSignUp) (models.AdminDetailsResponse, error)
	CheckAdminAvailability(admin models.AdminSignUp) bool
	GetUsers(page int) ([]models.UserDetailsResponse, error)
	GetGenres() ([]domain.Genre, error)
	GetDirectors() ([]domain.Directors, error)
	GetMovieFormat() ([]domain.Movie_Format, error)
	GetMovieLanguages() ([]domain.Movie_Language, error)
	AddGenre(genre string) (domain.Genre, error)
	AddDirector(director string) (domain.Directors, error)
	AddFormat(format string) (domain.Movie_Format, error)
	AddLanguage(language string) (domain.Movie_Language, error)
	Delete(genre_id string) error
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	CategoryCount(category models.CategoryUpdate) (models.CategoryUpdateCheck, error)
}
