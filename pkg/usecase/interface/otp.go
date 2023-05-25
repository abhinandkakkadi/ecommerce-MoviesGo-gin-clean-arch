package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OtpUseCase interface {
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
	SendOTP(phone string) error

	// VerifyMobileNumberAlreadyPresent(ctx context.Context,phone string) error
}
