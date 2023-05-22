package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)


type AdminRepository interface {
	LoginHandler(c context.Context,adminDetails domain.Admin) (domain.Admin,error) 
	SignupHandler(c context.Context,admin domain.Admin) (domain.Admin,error)
	CheckAdminAvailability(admin domain.Admin) bool
	GetUsers(c context.Context) ([]models.UserDetails,error)
	GetGenres(c context.Context) ([]domain.Genre,error)
	AddGenre(c context.Context,genre domain.Genre) (domain.Genre,error)
	Delete(c context.Context,genre_id string) error
	GetUserByID(c context.Context,id string) (domain.Users,error)
	UpdateUserByID(c context.Context,user domain.Users) error
	
}


