package interfaces

import (
	"context"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {

	GenerateUser(ctx context.Context,user domain.Users) (domain.TokenUsers,error)


	LoginHandler(ctx context.Context,user domain.Users) (domain.TokenUsers,error)
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
