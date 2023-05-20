package interfaces

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)





type ProductRepository interface {
	ShowAllProducts(ctx context.Context) ([]domain.ProductsBrief, error)
	ShowIndividualProducts(ctx context.Context,id string)(domain.Products,error)
}