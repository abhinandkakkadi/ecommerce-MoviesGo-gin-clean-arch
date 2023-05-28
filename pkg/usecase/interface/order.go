package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderBody models.OrderIncoming) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int) ([]models.FullOrderDetails, error)
	CancelOrder(orderID string, userID int) (string, error)
	GetAllOrderDetailsForAdmin() ([]models.CombinedOrderDetails, error)
	ApproveOrder(orderId string) (string, error)
}
