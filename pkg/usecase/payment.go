package usecase

import (
	"fmt"

	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/razorpay/razorpay-go"
)

type paymentUseCase struct {
	paymentReposioty interfaces.PaymentRepository
	orderRepository  interfaces.OrderRepository
	userRepository   interfaces.UserRepository
}

func NewPaymentUseCase(paymentRepo interfaces.PaymentRepository, orderRepo interfaces.OrderRepository, userRepo interfaces.UserRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentReposioty: paymentRepo,
		orderRepository:  orderRepo,
		userRepository:   userRepo,
	}
}

func (p *paymentUseCase) MakePaymentRazorPay(orderID string, userID int) (models.CombinedOrderDetails,string,error) {

	// check whether there is an order given by this order and also check if the current user have made this order
	err := p.orderRepository.CheckOrder(orderID, userID)
	if err != nil {
		panic(err)
	}

	orderDetails, err := p.orderRepository.GetOrderDetail(orderID)
	if err != nil {
		panic(err)
	}

	userDetails, err := p.userRepository.FindUserByOrderID(orderID)
	if err != nil {
		panic(err)
	}
	//  get shipping address for that particular order
	userAddress, err := p.userRepository.FindUserAddressByOrderID(orderID)
	if err != nil {
		panic(err)
	}
	// combine all the three details
	combinedOrderDetails, err := helper.CombinedOrderDetails(orderDetails, userDetails, userAddress)
	if err != nil {
		panic(err)
	}

	client := razorpay.NewClient("rzp_test_6m0J6O6Dngl96V", "F9zSviAWO3DIXnNAtKgrufzT")

  data := map[string]interface{}{
    "amount":   int(combinedOrderDetails.FinalPrice)*100,
    "currency": "INR",
    "receipt":  "some_receipt_id",
  }
  body, err := client.Order.Create(data, nil)
  fmt.Println(body)
  razorPayOrderID := body["id"].(string)

	p.orderRepository.AddRazorPayDetails(orderID,razorPayOrderID)

	return combinedOrderDetails,razorPayOrderID,nil

}
