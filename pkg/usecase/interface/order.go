package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderBody models.OrderFromCart, userId int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error)
	CancelOrder(orderID string, userID int) error
	CancelOrderFromAdminSide(orderID string) error
	GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error)
	ApproveOrder(orderId string) error
	OrderDelivered(orderID string) error
	ReturnOrder(orderID string) error
	RefundOrder(orderID string) error
}
