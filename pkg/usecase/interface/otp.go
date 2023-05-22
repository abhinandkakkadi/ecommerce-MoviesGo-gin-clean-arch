package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)


type OtpUseCase interface {
	VerifyOTP(ctx context.Context,code models.VerifyData) (domain.TokenUsers,error)
	SendOTP(ctx context.Context,phone string) error
	
	// VerifyMobileNumberAlreadyPresent(ctx context.Context,phone string) error
}