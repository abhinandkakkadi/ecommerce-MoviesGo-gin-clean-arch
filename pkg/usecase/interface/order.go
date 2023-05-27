package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderBody models.OrderIncoming) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int) ([]models.FullOrderDetails,error)
	CancelOrder(orderID string) (string,error)
}
