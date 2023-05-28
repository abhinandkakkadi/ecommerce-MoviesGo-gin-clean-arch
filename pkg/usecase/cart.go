package usecase

import (
	"errors"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type cartUseCase struct {
	cartRepository interfaces.CartRepository
}

func NewCartUseCase(repository interfaces.CartRepository) services.CartUseCase {

	return &cartUseCase{
		cartRepository: repository,
	}

}

func (cr *cartUseCase) AddToCart(product_id int, userID int) (models.CartResponse, error) {

	ok, err := cr.cartRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}

	if !ok {
		return models.CartResponse{}, errors.New("product does not exist")
	}

	cartDetails, err := cr.cartRepository.AddToCart(product_id, userID)

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
		Cart:       cartDetails,
	}

	return cartResponse, nil
}

func (cr *cartUseCase) RemoveFromCart(product_id int, userID int) (models.CartResponse, error) {

	updatedCart, err := cr.cartRepository.RemoveFromCart(product_id, userID)

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
