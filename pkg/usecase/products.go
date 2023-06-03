package usecase

import (
	"errors"
	"fmt"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
	cartRepo    interfaces.CartRepository
}

func NewProductUseCase(repo interfaces.ProductRepository, cartRepo interfaces.CartRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
		cartRepo:    cartRepo,
	}
}

func (cr *productUseCase) ShowAllProducts(page int) ([]models.ProductsBrief, error) {

	productsBrief, err := cr.productRepo.ShowAllProducts(page)

	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		fmt.Println("the code reached here")
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.Product_Status = "out of stock"
		} else {
			p.Product_Status = "in stock"
		}
	}

	return productsBrief, err

}

func (cr *productUseCase) ShowIndividualProducts(id string) (models.ProductResponse, error) {

	product, err := cr.productRepo.ShowIndividualProducts(id)
	if product.Movie_Name == "" {
		err = errors.New("record not available")
	}
	return product, err

}

func (cr *productUseCase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {
	// this logic is to add the quantity of product if admin try to add duplicate product (have to work on this in the future)
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

	productResponse, err := cr.productRepo.AddProduct(product)

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

func (cr *productUseCase) DeleteProduct(product_id string) error {

	err := cr.productRepo.DeleteProduct(product_id)
	if err != nil {
		return err
	}
	return nil

}

func (cr *productUseCase) UpdateProduct(productID int, quantity int) error {

	ok, err := cr.cartRepo.CheckProduct(productID)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("error does not exist")
	}

	return cr.productRepo.UpdateQuantity(productID, quantity)

}

func (cr *productUseCase) FilterCategory(data map[string]int) ([]models.ProductsBrief, error) {

	err := cr.productRepo.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductsBrief{}, err
	}

	productByCategory, err := cr.productRepo.GetProductFromCategory(data)
	if err != nil {
		return []models.ProductsBrief{}, err
	}
	fmt.Println("products By Category: ", productByCategory)
	return productByCategory, nil
}

func (cr *productUseCase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief,error) {

	return cr.productRepo.SearchItemBasedOnPrefix(prefix)
}