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
	AddGenre(genre domain.Genre) error
	AddDirector(director domain.Directors) error
	AddFormat(format domain.Movie_Format) error
	AddLanguage(language domain.Movie_Language) error
	Delete(genre_id string) error
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
}
