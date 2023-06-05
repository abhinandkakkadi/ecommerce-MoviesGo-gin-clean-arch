package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderRepository interface {
	OrderItemsFromCart(orderBody models.OrderIncoming, cartItems []models.Cart) (domain.OrderSuccessResponse, error)
	AddressExist(orderBody models.OrderIncoming) (bool, error)
	GetOrderDetails(userID int, page int) ([]models.FullOrderDetails, error)
	CancelOrder(orderID string) (string, error)
	GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error)
	UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error
	UserOrderRelationship(orderID string, userID int) (int, error)
	GetOrderDetailsBrief(page int) ([]models.OrderDetails, error)
	GetShipmentStatus(orderID string) (string, error)
	ApproveOrder(orderID string) error
	CheckOrderID(orderID string) (bool, error)

	SavePayment(charge domain.Charge) error
	GetPaymentDetails(OrderID string) (domain.Charge, error)

	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (models.OrderDetails, error)

	AddRazorPayDetails(orderID string,razorPayOrderID string)
}
