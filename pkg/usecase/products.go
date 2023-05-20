package usecase

import (
	"context"
	"errors"

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

func (cr *productUseCase) ShowAllProducts(c context.Context) ([]domain.ProductsBrief,error) {

	productsBrief, err := cr.productRepo.ShowAllProducts(c)
	return productsBrief, err
}

func (cr *productUseCase) ShowIndividualProducts(c context.Context,id string) (domain.Products,error) {

	product, err := cr.productRepo.ShowIndividualProducts(c,id)
	
	if product.Movie_Name == "" {
		err = errors.New("Record not avaiable")
	}
	return product,err

}