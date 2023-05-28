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
	// details being added to the orders table
	orderDetails.AddressID = orderBody.AddressID
	orderDetails.PaymentMethodID = orderBody.PaymentID
	orderDetails.UserID = int(orderBody.UserID)
	orderDetails.Approval = false
	orderDetails.ShipmentStatus = "processing"

	for _, c := range cartItems {
		orderDetails.GrandTotal += c.TotalPrice
	}
	cr.DB.Create(&orderDetails)
	// details being added to the orderItems table - which shows details about the individual products
	for _, c := range cartItems {
		fmt.Println(c)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		cr.DB.Omit("id").Create(&orderItemDetails)
		cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", orderDetails.UserID, c.ProductID)
		fmt.Println(c.Quantity)
		fmt.Println(c.ProductID)
		cr.DB.Exec("update products set quantity = quantity - ? where id = ?", c.Quantity, c.ProductID)
	}

	var orderSuccessResponse domain.OrderSuccessResponse

	cr.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderDetails.OrderId).Scan(&orderSuccessResponse)

	return orderSuccessResponse, nil
}

func (cr *orderRepository) GetOrderAddress(userID int) ([]models.FullOrderDetails, error) {

	var orderDetails []models.OrderDetails
	cr.DB.Raw("select order_id,grand_total,shipment_status from orders where user_id = ?", userID).Scan(&orderDetails)
	fmt.Println(orderDetails)

	var fullOrderDetails []models.FullOrderDetails

	for _, o := range orderDetails {

		var orderProductDetails []models.OrderProductDetails
		cr.DB.Raw("select order_items.product_id,products.movie_name,order_items.quantity,order_items.total_price from order_items inner join products on order_items.product_id = products.id where order_items.order_id = ?", o.OrderId).Scan(&orderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.FullOrderDetails{OrderDetails: o, OrderProductDetails: orderProductDetails})

	}

	return fullOrderDetails, nil

}

func (cr *orderRepository) UserOrderRelationship(orderID string, userID int) (int, error) {

	var testUserID int
	err := cr.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&testUserID).Error
	if err != nil {
		return -1, err
	}
	return testUserID, nil

}

func (cr *orderRepository) CancelOrder(orderID string) (string, error) {

	var shipmentStatus string
	fmt.Println(orderID)
	err := cr.DB.Raw("select shipment_status from orders where order_id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}

	fmt.Println(shipmentStatus)
	fmt.Println("ok")
	if shipmentStatus == "delivered" {
		return "Item already delivered, cannot cancel", nil
	}

	if shipmentStatus == "cancelled" {
		return "Order already cancelled", nil
	}

	shipmentStatus = "cancelled"
	err = cr.DB.Exec("update orders set shipment_status = ?  where order_id = ?", shipmentStatus, orderID).Error
	if err != nil {
		return "", err
	}

	return "Order successfully cancelled", nil
}

func (cr *orderRepository) GetOrderDetailsBrief() ([]models.OrderDetails, error) {

	var orderDetails []models.OrderDetails
	err := cr.DB.Raw("select order_id,grand_total,shipment_status from orders").Scan(&orderDetails).Error
	if err != nil {
		return []models.OrderDetails{}, nil
	}

	return orderDetails, nil
}

func (cr *orderRepository) CheckOrderID(orderID string) (bool, error) {

	var count int
	err := cr.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (cr *orderRepository) GetShipmentStatus(orderID string) (string, error) {

	var shipmentStatus string
	err := cr.DB.Raw("select shipment_status from orders where order_id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}

	return shipmentStatus, nil

}

func (cr *orderRepository) ApproveOrder(orderID string) error {

	err := cr.DB.Exec("update orders set shipment_status = 'order placed',approval = true where order_id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}
