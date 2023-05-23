package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

type UserRepository interface {
	UserSignUp(user domain.Users) (domain.Users,error)
	FindUserByEmail(user domain.Users) (domain.Users,error)
	CheckUserAvailability(user domain.Users) bool
	LoginHandler(user domain.Users) (domain.Users,error)
}
