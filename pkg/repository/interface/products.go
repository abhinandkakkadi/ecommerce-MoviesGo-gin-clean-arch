package interfaces

import (
	"context"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)





type ProductRepository interface {
	ShowAllProducts(ctx context.Context) ([]domain.Products, error)
}