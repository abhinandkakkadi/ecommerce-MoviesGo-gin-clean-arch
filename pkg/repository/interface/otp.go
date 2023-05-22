package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(ctx context.Context,phone string) (domain.Users,error)
	// VerifyMobileNumberAlreadyPresent(ctx context.Context,phone string) bool
}