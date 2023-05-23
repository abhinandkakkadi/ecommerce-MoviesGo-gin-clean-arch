package interfaces

import (
	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {
	UserSignUp(user domain.Users) (domain.TokenUsers, error)

	LoginHandler(user domain.Users) (domain.TokenUsers, error)
}
