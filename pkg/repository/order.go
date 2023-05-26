package repository

import (
	"fmt"
	"strconv"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (cr *orderRepository) OrderItemsFromCart(orderBody models.OrderIncoming, cartItems []models.Cart) (domain.OrderSuccessResponse, error) {

	fmt.Println(cartItems)

	var orderDetails domain.Order
	var orderItemDetails domain.OrderItem

	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	orderDetails.OrderId = str[:8]

	orderDetails.AddressID = orderBody.AddressID
	orderDetails.PaymentMethodID = orderBody.PaymentID
	orderDetails.UserID = int(orderBody.UserID)
	orderDetails.Approval = false
	orderDetails.ShipmentStatus = "processing"

	for _, c := range cartItems {
		orderDetails.GrandTotal += c.TotalPrice
	}
	cr.DB.Create(&orderDetails)

	for _, c := range cartItems {
		fmt.Println(c)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		cr.DB.Omit("id").Create(&orderItemDetails)
		cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", orderDetails.UserID, c.ProductID)
	}

	var orderSuccessResponse domain.OrderSuccessResponse

	cr.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderDetails.OrderId).Scan(&orderSuccessResponse)

	return orderSuccessResponse, nil
}
