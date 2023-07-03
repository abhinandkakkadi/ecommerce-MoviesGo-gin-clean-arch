package usecase

import (
	"errors"

	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type cartUseCase struct {
	cartRepository   interfaces.CartRepository
	couponRepository interfaces.CouponRepository
}

func NewCartUseCase(repository interfaces.CartRepository, couponRepo interfaces.CouponRepository) services.CartUseCase {

	return &cartUseCase{
		cartRepository:   repository,
		couponRepository: couponRepo,
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
		return models.CartResponse{},err
	}

	offerDetails := helper.OfferHelper(combinedOfferDetails)


	// Now check if the offer is already used by the user
	if offerDetails.OfferType != "no offer" {
		offerDetails, err = cr.couponRepository.CheckIfOfferAlreadyUsed(offerDetails, product_id, userID)
		if err != nil {
			return models.CartResponse{},err
		}
	}


	// if offer id ! =0 that means some kind of offer exist - do the complete things inside this
	err = cr.couponRepository.OfferUpdate(offerDetails, userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartDetails, err := cr.cartRepository.AddToCart(product_id, userID, offerDetails)

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

	emptyCart, err := cr.cartRepository.EmptyCart(userID)

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
		Cart:       emptyCart,
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

	couponStatus, err := cr.cartRepository.CouponValidity(coupon, userID)
	if err != nil {
		return err
	}

	if couponStatus {
		return nil
	}
	return errors.New("could not add the coupon")
}
