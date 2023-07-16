package usecase

import (
	"errors"
	"fmt"

	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type cartUseCase struct {
	cartRepository    interfaces.CartRepository
	couponRepository  interfaces.CouponRepository
	productRepository interfaces.ProductRepository
}

func NewCartUseCase(repository interfaces.CartRepository, couponRepo interfaces.CouponRepository, productRepo interfaces.ProductRepository) services.CartUseCase {

	return &cartUseCase{
		cartRepository:    repository,
		couponRepository:  couponRepo,
		productRepository: productRepo,
	}

}

func (cr *cartUseCase) AddToCart(product_id int, userID int) (models.CartResponse, error) {

	//  to check whether the product exist
	ok, genre, err := cr.cartRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}

	if !ok {
		return models.CartResponse{}, errors.New("product does not exist")
	}

	combinedOfferDetails, err := cr.couponRepository.OfferDetails(product_id, genre)
	if err != nil {
		return models.CartResponse{}, err
	}

	offerDetails := helper.OfferHelper(combinedOfferDetails)

	// Now check if the offer is already used by the user
	if offerDetails.OfferType != "no offer" {

		if offerDetails.OfferType == "product" {

			offerDetails, err = cr.couponRepository.CheckIfProductOfferAlreadyUsed(offerDetails, product_id, userID)
			if err != nil {
				return models.CartResponse{}, err
			}

		} else if offerDetails.OfferType == "category" {

			offerDetails, err = cr.couponRepository.CheckIfCategoryOfferAlreadyUsed(offerDetails, product_id, userID)
			if err != nil {
				return models.CartResponse{}, err
			}
		}
	}

	if offerDetails.OfferType == "product" {
		err = cr.couponRepository.OfferUpdateProduct(offerDetails, userID)
		if err != nil {
			return models.CartResponse{}, err
		}

	} else if offerDetails.OfferType == "category" {
		err = cr.couponRepository.OfferUpdateCategory(offerDetails, userID)
		if err != nil {
			return models.CartResponse{}, err
		}

	}

	quantityOfProductInCart, err := cr.cartRepository.QuantityOfProductInCart(userID, product_id)
	fmt.Println(quantityOfProductInCart)
	if err != nil {
		return models.CartResponse{}, err
	}

	quantityOfProduct, err := cr.productRepository.GetQuantityFromProductID(product_id)
	fmt.Println(quantityOfProduct)
	if err != nil {
		return models.CartResponse{}, err
	}

	if quantityOfProduct == 0 {
		return models.CartResponse{}, errors.New("product out of stock")
	}

	if quantityOfProduct == quantityOfProductInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}

	productPrice, err := cr.productRepository.GetPriceOfProductFromID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}

	if offerDetails.OfferPrice != productPrice {

		if quantityOfProductInCart < offerDetails.OfferLimit {
			productPrice = offerDetails.OfferPrice
		}

	}

	if quantityOfProductInCart == 0 {
		err := cr.cartRepository.AddItemToCart(userID, product_id, 1, productPrice)
		if err != nil {
			return models.CartResponse{}, err
		}
	} else {
		currentTotal, err := cr.cartRepository.TotalPriceForProductInCart(userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}

		err = cr.cartRepository.UpdateCart(quantityOfProductInCart+1, currentTotal+productPrice, userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	cartDetails, err := cr.cartRepository.DisplayCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	// function to get the grand total price
	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}

	return cartResponse, nil
}

// remove items cart (if a product of multiple quantity is present - item will be removed one by one)
func (cr *cartUseCase) RemoveFromCart(product_id int, userID int) (models.CartResponse, error) {

	// check the product to be removed exist
	productExist, err := cr.cartRepository.ProductExist(product_id, userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !productExist {
		return models.CartResponse{}, errors.New("the product does not exist in catt")
	}

	// if offer is applied decrement offer price, else decrement actual price
	priceDecrement, err := cr.couponRepository.GetPriceBasedOnOffer(product_id, userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	var cartDetails struct {
		Quantity   int
		TotalPrice float64
	}

	cartDetails, err = cr.cartRepository.GetQuantityAndTotalPrice(userID, product_id, cartDetails)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartDetails.Quantity = cartDetails.Quantity - 1
	// after decrementing one quantity if the quantity = 0. delete that item from the cart
	if cartDetails.Quantity == 0 {

		err := cr.cartRepository.RemoveProductFromCart(userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	if cartDetails.Quantity != 0 {

		cartDetails.TotalPrice = cartDetails.TotalPrice - priceDecrement

		err := cr.cartRepository.UpdateCartDetails(cartDetails, userID, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}

	}

	updatedCart, err := cr.cartRepository.RemoveFromCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       updatedCart,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) DisplayCart(userID int) (models.CartResponse, error) {

	displayCart, err := cr.cartRepository.DisplayCart(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       displayCart,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) EmptyCart(userID int) (models.CartResponse, error) {

	cartExist, err := cr.cartRepository.DoesCartExist(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	if !cartExist {
		return models.CartResponse{}, errors.New("cart already empty")
	}

	err = cr.cartRepository.EmptyCart(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	// CATEGORY OFFER RESTORED
	categoryOfferIDS, err := cr.cartRepository.GetUnUsedCategoryOfferIDS(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	for _, cOfferId := range categoryOfferIDS {
		err := cr.cartRepository.UpdateUnUsedCategoryOffer(cOfferId, userID)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	// PRODUCT OFFER RESTORED
	productOfferIDS, err := cr.cartRepository.GetUnUsedProductOfferIDS(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	for _, pOfferID := range productOfferIDS {
		err := cr.cartRepository.UpdateUnUsedProductOffer(pOfferID, userID)
		if err != nil {
			return models.CartResponse{}, err
		}
	}

	cartTotal, err := cr.cartRepository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       []models.Cart{},
	}

	return cartResponse, nil
}

func (cr *cartUseCase) ApplyCoupon(coupon string, userID int) error {

	cartExist, err := cr.cartRepository.DoesCartExist(userID)
	if err != nil {
		return err
	}

	if !cartExist {
		return errors.New("cart empty, can't apply coupon")
	}

	couponExist, err := cr.couponRepository.CouponExist(coupon)
	if err != nil {
		return err
	}

	if !couponExist {
		return errors.New("coupon does not exist")
	}

	couponValidity, err := cr.couponRepository.CouponValidity(coupon)
	if err != nil {
		return err
	}

	if !couponValidity {
		return errors.New("coupon expired")
	}

	minDiscountPrice, err := cr.couponRepository.GetCouponMinimumAmount(coupon)
	if err != nil {
		return err
	}

	totalPriceFromCarts, err := cr.cartRepository.GetTotalPriceFromCart(userID)
	if err != nil {
		return err
	}

	// if the total Price is less than minDiscount price don't allow coupon to be added
	if totalPriceFromCarts < minDiscountPrice {
		return errors.New("coupon cannot be added as the total amount is less than minimum amount for coupon")
	}

	userAlreadyUsed, err := cr.couponRepository.DidUserAlreadyUsedThisCoupon(coupon, userID)
	if err != nil {
		return err
	}

	if userAlreadyUsed {
		return errors.New("user already used this coupon")
	}

	couponStatus, err := cr.cartRepository.UpdateUsedCoupon(coupon, userID)
	if err != nil {
		return err
	}

	if couponStatus {
		return nil
	}
	return errors.New("could not add the coupon")

}
