package interfaces

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type ProductUseCase interface {
	
	ShowAllProducts(c context.Context) ([]domain.Products,error)
}