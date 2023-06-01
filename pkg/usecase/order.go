package usecase

import (
	"errors"
	"fmt"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/razorpay/razorpay-go"
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

	addressExist, err := cr.orderRepository.AddressExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}
	// get all items a slice of carts
	cartItems, err := cr.cartRepository.GetAllItemsFromCart(int(orderBody.UserID))

	if err != nil {
		return domain.OrderSuccessResponse{}, nil
	}

	orderSuccessResponse, err := cr.orderRepository.OrderItemsFromCart(orderBody, cartItems)
	if err != nil {
		return domain.OrderSuccessResponse{}, nil
	}
	fmt.Println("order id ",orderSuccessResponse.OrderID)
	// if payment is via card
	var paymentDetails domain.Charge
	if orderBody.PaymentID == 2 {
		paymentDetails,err = cr.orderRepository.GetPaymentDetails(orderSuccessResponse.OrderID)
		if err != nil {
			return domain.OrderSuccessResponse{},err
		}
	}
	fmt.Println("order id ",orderSuccessResponse.OrderID)
	client := razorpay.NewClient("rzp_test_kUBAXm7sKjPa0a", "KCkWzEkoKIY8hdWa0Lp8xIbo")

	data := map[string]interface{}{
 	  "amount": int64(paymentDetails.GrandTotal),
  	"currency": "INR",
  	"receipt": "abhinand", 
	}

	payment, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.OrderSuccessResponse{},err
	}
	fmt.Println("order id ",orderSuccessResponse.OrderID)
	value :=  payment["id"]
	fmt.Println("orderid by razor pay : ",value.(string))
	fmt.Println("razorpay sent back details : ",payment)
	err = cr.orderRepository.SavePayment(paymentDetails)
  if err != nil {
		return domain.OrderSuccessResponse{},err
	}
	fmt.Println("order id ",orderSuccessResponse.OrderID)
	return orderSuccessResponse, nil

}

// get order details
func (cr *orderUseCase) GetOrderDetails(userID int,page int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := cr.orderRepository.GetOrderDetails(userID,page)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}

	return fullOrderDetails, nil

}

func (cr *orderUseCase) CancelOrder(orderID string, userID int) (string, error) {

	// check whether the orderID corresponds to the given user (other user with token may try to send orderID as path variables) (have to add this logic to so many places)
	userTest, err := cr.orderRepository.UserOrderRelationship(orderID, userID)
	if err != nil {
		return "", err
	}

	if userTest != userID {
		return "", errors.New("the order is not done by this user")
	}

	orderProducts, err := cr.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return "",err
	}

	// update the quantity to products since th order is cancelled
	err = cr.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return "", err
	}

	return cr.orderRepository.CancelOrder(orderID)

}

func (cr *orderUseCase) CancelOrderFromAdminSide(orderID string) (string, error) {

	orderProducts, err := cr.orderRepository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return "",err
	}

	// update the quantity to products since th order is cancelled
	err = cr.orderRepository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return "", err
	}

	return cr.orderRepository.CancelOrder(orderID)

}

func (cr *orderUseCase) GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error) {

	orderDetails, err := cr.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}

	var allCombinedOrderDetails []models.CombinedOrderDetails

	// we will take the order details from orders table,  and for each order we combine it with the corresponding user details and address of that particular user that made the order
	for _, o := range orderDetails {

		// get users details who made that order
		userDetails, err := cr.userRepository.FindUserByOrderID(o.OrderId)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}
		//  get shipping address for that particular order
		userAddress, err := cr.userRepository.FindUserAddressByOrderID(o.OrderId)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}
		// combine all the three details
		combinedOrderDetails, err := helper.CombinedOrderDetails(o, userDetails, userAddress)
		if err != nil {
			return []models.CombinedOrderDetails{}, err
		}
		// combine all of these details and append it to a slice
		allCombinedOrderDetails = append(allCombinedOrderDetails, combinedOrderDetails)

	}

	return allCombinedOrderDetails, nil
}

func (cr *orderUseCase) ApproveOrder(orderID string) (string, error) {

	// check whether the specified orderID exist
	ok, err := cr.orderRepository.CheckOrderID(orderID)
	fmt.Println(ok)
	if !ok {
		return "Order ID does not exist", err
	}

	// check the shipment status - if the status cancelled, don't approve it
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

	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return "order already approved", nil

}
