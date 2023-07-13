package usecase

import (
	"errors"
	"fmt"
	"strings"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
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
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
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
		combinedOfferDetails, err := pr.couponRepo.OfferDetails(p.ID, p.Genre)
		if err != nil {
			return []models.ProductOfferBriefResponse{}, err
		}

		offerDetails := helper.OfferHelper(combinedOfferDetails)

		productOffer.ProductsBrief = p
		productOffer.OfferResponse = offerDetails
		combinedProductsOffer = append(combinedProductsOffer, productOffer)
	}

	return combinedProductsOffer, err

}

func (pr *productUseCase) ShowAllProductsToAdmin(page int, count int) ([]models.ProductsBrief, error) {

	productsBrief, err := pr.productRepo.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductsBrief{}, err
	}
	fmt.Println(productsBrief)
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return productsBrief, nil
}

func (pr *productUseCase) ShowIndividualProducts(id string) (models.ProductOfferLongResponse, error) {

	product, err := pr.productRepo.ShowIndividualProducts(id)
	if err != nil {
		return models.ProductOfferLongResponse{}, err
	}
	if product.MovieName == "" {
		return models.ProductOfferLongResponse{}, errors.New("record not available")
	}

	var productOfferResponse models.ProductOfferLongResponse
	combinedOfferDetails, err := pr.couponRepo.OfferDetails(product.ID, product.GenreName)
	if err != nil {
		return models.ProductOfferLongResponse{}, err
	}

	offerDetails := helper.OfferHelper(combinedOfferDetails)

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

	ok, genre, err := pr.cartRepo.CheckProduct(productID)
	_ = genre

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("product does not exist")
	}

	return pr.productRepo.UpdateQuantity(productID, quantity)

}

func (pr *productUseCase) FilterCategory(data map[string]int) ([]models.ProductsBrief, error) {

	err := pr.productRepo.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductsBrief{}, err
	}

	var productFromCategory []models.ProductsBrief
	for _, id := range data {

		product, err := pr.productRepo.GetProductFromCategory(id)
		if err != nil {
			return []models.ProductsBrief{}, err
		}

		quantity, err := pr.productRepo.GetQuantityFromProductID(product.ID)
		if err != nil {
			return []models.ProductsBrief{}, err
		}

		if quantity == 0 {
			product.ProductStatus = "out of stock"
		} else {
			product.ProductStatus = "in stock"
		}

		// if a product exist for that genre. Then only append it
		if product.ID != 0 {
			productFromCategory = append(productFromCategory, product)
		}

	}
	return productFromCategory, nil

}

func (pr *productUseCase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, error) {

	productsBrief, lengthOfPrefix, err := pr.productRepo.SearchItemBasedOnPrefix(prefix)
	if err != nil {
		return []models.ProductsBrief{}, err
	}

	// Create a slice to add the products which have the given prefix
	var filteredProductBrief []models.ProductsBrief
	for _, p := range productsBrief {
		length := len(p.MovieName)
		if length >= lengthOfPrefix {
			moviePrefix := p.MovieName[:lengthOfPrefix]
			if strings.ToLower(moviePrefix) == strings.ToLower(prefix) {
				filteredProductBrief = append(filteredProductBrief, p)
			}
		}
	}

	for i := range filteredProductBrief {
		fmt.Println("the code reached here")
		p := &filteredProductBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return filteredProductBrief, nil
}

func (pr *productUseCase) GetGenres() ([]domain.Genre, error) {

	return pr.productRepo.GetGenres()
}
