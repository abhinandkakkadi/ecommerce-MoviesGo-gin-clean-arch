package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderBody models.OrderFromCart, userId int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int, page int) ([]models.FullOrderDetails, error)
	CancelOrder(orderID string, userID int) (string, error)
	CancelOrderFromAdminSide(orderID string) (string, error)
	GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error)
	ApproveOrder(orderId string) (string, error)
}
