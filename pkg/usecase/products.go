package usecase

import (
	"context"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)


type productUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
	}
}

func (cr *productUseCase) ShowAllProducts(c context.Context) ([]domain.Products,error) {

	products, err := cr.productRepo.ShowAllProducts(c)
	return products, err
}