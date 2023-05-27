package usecase

import (
	"errors"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	cartRepository  interfaces.CartRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository, cartRepo interfaces.CartRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: orderRepo,
		cartRepository:  cartRepo,
	}
}

func (cr *orderUseCase) OrderItemsFromCart(orderBody models.OrderIncoming) (domain.OrderSuccessResponse, error) {

	cartItems, err := cr.cartRepository.GetAllItemsFromCart(int(orderBody.UserID))

	if err != nil {
		return domain.OrderSuccessResponse{}, nil
	}

	orderSuccessResponse, err := cr.orderRepository.OrderItemsFromCart(orderBody, cartItems)
	if err != nil {
		return domain.OrderSuccessResponse{}, nil
	}

	return orderSuccessResponse, nil

}

func (cr *orderUseCase) GetOrderDetails(userID int) ([]models.FullOrderDetails,error)  {

	fullOrderDetails,err := cr.orderRepository.GetOrderAddress(userID)
	if err != nil {
		return []models.FullOrderDetails{},err
	}

	return fullOrderDetails,nil

}


func (cr *orderUseCase) CancelOrder(orderID string,userID int) (string,error) {

	// check whether the orderID corresponds to the given user (other user with token may try to send orderID as path variables)
	userTest,err := cr.orderRepository.UserOrderRelationship(orderID,userID)
	if err != nil {
		return "",err
	}

	if userTest != userID {
		return "",errors.New("the order is not done by this user")
	}

	return cr.orderRepository.CancelOrder(orderID)

}