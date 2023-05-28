package usecase

import (
	"errors"
	"fmt"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	cartRepository  interfaces.CartRepository
	userRepository  interfaces.UserRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository, cartRepo interfaces.CartRepository, userRepo interfaces.UserRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: orderRepo,
		cartRepository:  cartRepo,
		userRepository:  userRepo,
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

func (cr *orderUseCase) GetOrderDetails(userID int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := cr.orderRepository.GetOrderAddress(userID)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}

	return fullOrderDetails, nil

}

func (cr *orderUseCase) CancelOrder(orderID string, userID int) (string, error) {

	// check whether the orderID corresponds to the given user (other user with token may try to send orderID as path variables)
	userTest, err := cr.orderRepository.UserOrderRelationship(orderID, userID)
	if err != nil {
		return "", err
	}

	if userTest != userID {
		return "", errors.New("the order is not done by this user")
	}

	return cr.orderRepository.CancelOrder(orderID)

}

func (cr *orderUseCase) CancelOrderFromAdminSide(orderID string) (string, error) {

	return cr.orderRepository.CancelOrder(orderID)

}

func (cr *orderUseCase) GetAllOrderDetailsForAdmin() ([]models.CombinedOrderDetails, error) {

	orderDetails, err := cr.orderRepository.GetOrderDetailsBrief()
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}

	var allCombinedOrderDetails []models.CombinedOrderDetails
	for _, o := range orderDetails {

		userDetails, err := cr.userRepository.FindUserByOrderID(o.OrderId)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}

		userAddress, err := cr.userRepository.FindUserAddressByOrderID(o.OrderId)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}

		combinedOrderDetails, err := helper.CombinedOrderDetails(o, userDetails, userAddress)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}

		allCombinedOrderDetails = append(allCombinedOrderDetails, combinedOrderDetails)

	}

	return allCombinedOrderDetails, nil
}

func (cr *orderUseCase) ApproveOrder(orderID string) (string, error) {
	fmt.Println(orderID)
	ok, err := cr.orderRepository.CheckOrderID(orderID)
	fmt.Println(ok)
	if !ok {
		return "Order ID does not exist", err
	}

	shipmentStatus, err := cr.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return "", err
	}
	fmt.Println(shipmentStatus)
	if shipmentStatus == "cancelled" {

		return "The order is cancelled, cannot approve it", nil
	}

	if shipmentStatus == "processing" {
		fmt.Println("reached here")
		err := cr.orderRepository.ApproveOrder(orderID)

		if err != nil {
			return "", err
		}

		return "order approved successfully", nil
	}

	return "order already approved", nil

}
