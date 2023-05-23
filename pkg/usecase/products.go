package usecase

import (
	"errors"

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

func (cr *productUseCase) ShowAllProducts() ([]domain.ProductsBrief, error) {

	productsBrief, err := cr.productRepo.ShowAllProducts()
	return productsBrief, err

}

func (cr *productUseCase) ShowIndividualProducts(id string) (models.IndividualProduct, error) {

	product, err := cr.productRepo.ShowIndividualProducts(id)
	if product.Movie_Name == "" {
		err = errors.New("record not available")
	}
	return product, err

}

func (cr *productUseCase) AddProduct(product domain.Products) error {
	// this logic is to add the quantity of product if admin try to add duplicate product
	// alreadyPresent,err := cr.productRepo.CheckIfAlreadyPresent(c,product)

	// if err != nil {
	// 	return err
	// }

	// if alreadyPresent {
	// 	fmt.Println("it came here")
	// 	err := cr.productRepo.UpdateQuantity(c,product)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// }

	err := cr.productRepo.AddProduct(product)

	if err != nil {
		return err
	}

	return nil

}

func (cr *productUseCase) DeleteProduct(product_id string) error {

	err := cr.productRepo.DeleteProduct(product_id)
	if err != nil {
		return err
	}
	return nil

}
