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
	couponRepo  interfaces.CouponRepository
}

func NewProductUseCase(repo interfaces.ProductRepository, cartRepo interfaces.CartRepository, couponRepo interfaces.CouponRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
		cartRepo:    cartRepo,
		couponRepo:  couponRepo,
	}
}

func (pr *productUseCase) ShowAllProducts(page int, count int) ([]models.ProductOfferBriefResponse, error) {

	productsBrief, err := pr.productRepo.ShowAllProducts(page, count)
	fmt.Println(productsBrief)
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		fmt.Println("the code reached here")
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	var combinedProductsOffer []models.ProductOfferBriefResponse
	for _, p := range productsBrief {
		var productOffer models.ProductOfferBriefResponse
		OfferDetails, err := pr.couponRepo.OfferDetails(p.ID, p.Genre)
		if err != nil {
			return []models.ProductOfferBriefResponse{}, err
		}
		fmt.Println(OfferDetails)
		productOffer.ProductsBrief = p
		productOffer.OfferResponse = OfferDetails
		combinedProductsOffer = append(combinedProductsOffer, productOffer)
	}

	return combinedProductsOffer, err

}

func (pr *productUseCase) ShowIndividualProducts(id string) (models.ProductOfferLongResponse, error) {

	product, err := pr.productRepo.ShowIndividualProducts(id)
	if err != nil {
		return models.ProductOfferLongResponse{}, err
	}
	if product.MovieName == "" {
		err = errors.New("record not available")
		return models.ProductOfferLongResponse{}, err
	}

	var productOfferResponse models.ProductOfferLongResponse
	offerDetails, err := pr.couponRepo.OfferDetails(product.ID, product.GenreName)
	if err != nil {
		return models.ProductOfferLongResponse{}, err
	}

	productOfferResponse.ProductsResponse = product
	productOfferResponse.OfferResponse = offerDetails

	return productOfferResponse, nil

}

func (pr *productUseCase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {
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

	productResponse, err := pr.productRepo.AddProduct(product)

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

func (pr *productUseCase) DeleteProduct(product_id string) error {

	err := pr.productRepo.DeleteProduct(product_id)
	if err != nil {
		return err
	}
	return nil

}

func (pr *productUseCase) UpdateProduct(productID int, quantity int) error {

	ok,genre, err := pr.cartRepo.CheckProduct(productID)
	_ = genre
	
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("error does not exist")
	}

	return pr.productRepo.UpdateQuantity(productID, quantity)

}

func (pr *productUseCase) FilterCategory(data map[string]int) ([]models.ProductsBrief, error) {

	err := pr.productRepo.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductsBrief{}, err
	}

	productByCategory, err := pr.productRepo.GetProductFromCategory(data)
	if err != nil {
		return []models.ProductsBrief{}, err
	}
	fmt.Println("products By Category: ", productByCategory)
	return productByCategory, nil
}

func (cr *productUseCase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, error) {

	return cr.productRepo.SearchItemBasedOnPrefix(prefix)
}
