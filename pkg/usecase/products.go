package usecase

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
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

func (cr *productUseCase) ShowIndividualProducts(c context.Context,id string) (models.IndividualProduct,error) {

	product, err := cr.productRepo.ShowIndividualProducts(c,id)
	
	if product.Movie_Name == "" {
		err = errors.New("Record not avaiable")
	}
	return product,err

}

func (cr *productUseCase) AddProduct(c context.Context,product domain.Products) (error) {

	alreadyPresent,err := cr.productRepo.CheckIfAlreadyPresent(c,product)

	if err != nil {
		return err
	}

	if alreadyPresent {
		fmt.Println("it came here")
		err := cr.productRepo.UdateQuantity(c,product)
		if err != nil {
			return err
		}

		return nil
	}

	err = cr.productRepo.AddProduct(c,product)

	if err != nil {
		return err
	}

	return nil

}

func (cr *productUseCase) DeleteProduct(c context.Context, product_id string) error {

	err := cr.productRepo.DeleteProduct(c,product_id)
	if err != nil {
		return err
	}
	return nil

}