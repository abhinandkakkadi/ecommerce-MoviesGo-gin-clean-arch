package handler

import (
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCase
}

func NewPaymentHandler(useCase services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: useCase,
	}
}
