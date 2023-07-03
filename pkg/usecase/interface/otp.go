package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OtpUseCase interface {
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
	VerifyOTPtoReset(code models.VerifyData) (string, error)
	SendOTP(phone string) error
	SendOTPtoReset(email string) (string, error)
}
