package interfaces

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type UserRepository interface {
	GenerateUser(user domain.Users) (domain.Users,error)
	FindUserByEmail(user domain.Users) (domain.Users,error)
	CheckUserAvailability(user domain.Users) bool
	LoginHandler(ctx context.Context, user domain.Users) (domain.Users,error)
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
