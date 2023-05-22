package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(c context.Context,adminDetails domain.Admin) (domain.TokenAdmin,error)
	SignupHandler(c context.Context,admin domain.Admin) (domain.TokenAdmin,error)
	GetUsers(c context.Context) ([]models.UserDetails,error)
	GetGenres(c context.Context) ([]domain.Genre,error)
	AddGenre(c context.Context,genre domain.Genre) (domain.Genre,error)
	Delete(c context.Context,genre_id string) error
	BlockUser(c context.Context,id string) error
	
}