package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (domain.Users,error)
}