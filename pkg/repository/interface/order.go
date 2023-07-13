package interfaces

import (
	"time"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type OrderRepository interface {
	GetBriefOrderDetails(orderDetails string) (domain.OrderSuccessResponse, error)
	DoesCartExist(userID int) (bool, error)
	UpdateCouponDetails(discount_price float64, UserID int) error
	GetWalletAmount(UserID uint) (float64, error)
	UpdateWalletAmount(walletAmount float64, UserID uint) error
	CreateOrder(orderDetails domain.Order) error
	UpdateUsedOfferDetails(userID uint) error
	GetCouponDiscountPrice(UserID int, GrandTotal float64) (float64, error)
	AddOrderItems(orderItemDetails domain.OrderItem, UserID int, ProductID uint, Quantity float64) error
	AddressExist(orderBody models.OrderIncoming) (bool, error)
	GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error)
	CancelOrder(orderID string) error
	GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error)
	UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error
	UserOrderRelationship(orderID string, userID int) (int, error)
	GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error)
	GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error)
	GetShipmentStatus(orderID string) (string, error)
	ApproveOrder(orderID string) error
	CheckOrderID(orderID string) (bool, error)
	SavePayment(charge domain.Charge) error
	GetPaymentDetails(OrderID string) (domain.Charge, error)
	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (models.OrderDetails, error)
	AddRazorPayDetails(orderID string, razorPayOrderID string) error
	CheckPaymentStatus(razorID string, orderID string) (string, error)
	UpdatePaymentDetails(orderID string, paymentID string) error
	UpdateShipmentStatus(shipmentStatus string, orderID string) error
	GetDeliveredTime(orderID string) (time.Time, error)
	ReturnOrder(shipmentStatus string, orderID string) error
	GetPaymentStatus(orderID string) (string, error)
	RefundOrder(paymentStatus string, orderID string) error
	UpdateShipmentAndPaymentByOrderID(shipmentStatus string, paymentStatus string, orderID string) error
}
