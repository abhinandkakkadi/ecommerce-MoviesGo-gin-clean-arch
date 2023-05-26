package usecase

import (
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
