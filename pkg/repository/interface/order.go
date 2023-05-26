package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderRepository interface {
	OrderItemsFromCart(orderBody models.OrderIncoming, cartItems []models.Cart) (domain.OrderSuccessResponse, error)
}
