package interfaces

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/pkg/utils/models"
)


type OtpUseCase interface {
	VerifyOTP(ctx context.Context,code models.VerifyData) error
	SendOTP(ctx context.Context,phone string) error
	// VerifyMobileNumberAlreadyPresent(ctx context.Context,phone string) error
}