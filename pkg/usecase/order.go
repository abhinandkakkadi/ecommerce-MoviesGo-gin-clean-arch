package usecase

import (
	"errors"
	"fmt"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/jinzhu/copier"
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

func (o *orderUseCase) OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error) {

	var orderBody models.OrderIncoming
	err := copier.Copy(&orderBody, &orderFromCart)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderBody.UserID = uint(userID)

	addressExist, err := o.orderRepository.AddressExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}
	// get all items a slice of carts
	cartItems, err := o.cartRepository.GetAllItemsFromCart(int(orderBody.UserID))

	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderSuccessResponse, err := o.orderRepository.OrderItemsFromCart(orderBody, cartItems)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	fmt.Println("order id ", orderSuccessResponse.OrderID)

	fmt.Println("order id ", orderSuccessResponse.OrderID)
	return orderSuccessResponse, nil

}

// get order details
func (o *orderUseCase) GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := o.orderRepository.GetOrderDetails(userID, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}

	return fullOrderDetails, nil

}

func (o *orderUseCase) CancelOrder(orderID string, userID int) (string, error) {

	// check whether the orderID corresponds to the given user (other user with token may try to send orderID as path variables) (have to add this logic to so many places)
	userTest, err := o.orderRepository.UserOrderRelationship(orderID, userID)
	if err != nil {
		return "", err
	}

	if userTest != userID {
		return "", errors.New("the order is not done by this user")
	}

	orderProducts, err := o.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return "", err
	}

	// there is an error in this code. even though the order is 
	// update the quantity to products since the order is cancelled
	err = o.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return "", err
	}

	return o.orderRepository.CancelOrder(orderID)

}

func (o *orderUseCase) CancelOrderFromAdminSide(orderID string) (string, error) {

	orderProducts, err := o.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return "", err
	}

	// update the quantity to products since th order is cancelled
	err = o.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return "", err
	}

	return o.orderRepository.CancelOrder(orderID)

}

func (o *orderUseCase) GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error) {

	orderDetails, err := o.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}

	var allCombinedOrderDetails []models.CombinedOrderDetails

	// we will take the order details from orders table,  and for each order we combine it with the corresponding user details and address of that particular user that made the order
	for _, od := range orderDetails {

		// get users details who made that order
		userDetails, err := o.userRepository.FindUserByOrderID(od.OrderId)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}
		//  get shipping address for that particular order
		userAddress, err := o.userRepository.FindUserAddressByOrderID(od.OrderId)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}
		// combine all the three details
		combinedOrderDetails, err := helper.CombinedOrderDetails(od, userDetails, userAddress)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}
		// combine all of these details and append it to a slice
		allCombinedOrderDetails = append(allCombinedOrderDetails, combinedOrderDetails)

	}

	return allCombinedOrderDetails, nil
}

func (o *orderUseCase) ApproveOrder(orderID string) (string, error) {

	// check whether the specified orderID exist
	ok, err := o.orderRepository.CheckOrderID(orderID)
	fmt.Println(ok)
	if !ok {
		return "Order ID does not exist", err
	}

	// check the shipment status - if the status cancelled, don't approve it
	shipmentStatus, err := o.orderRepository.GetShipmentStatus(orderID)
	if err != nil {
		return "", err
	}
	fmt.Println(shipmentStatus)
	if shipmentStatus == "cancelled" {

		return "The order is cancelled, cannot approve it", nil
	}

	if shipmentStatus == "processing" {
		fmt.Println("reached here")
		err := o.orderRepository.ApproveOrder(orderID)

		if err != nil {
			return "", err
		}

		return "order approved successfully", nil
	}

	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return "order already approved", nil

}
