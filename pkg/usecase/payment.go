package usecase

import (
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
)

type paymentUseCase struct {
	paymentReposioty interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentReposioty: repo,
	}
}
