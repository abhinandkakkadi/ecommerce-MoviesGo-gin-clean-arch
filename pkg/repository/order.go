package repository

import (
	"fmt"
	"strconv"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
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

func (cr *orderRepository) AddressExist(orderBody models.OrderIncoming) (bool, error) {
	fmt.Println("user id = ", orderBody.UserID, "address id = ", orderBody.AddressID)
	var count int
	if err := cr.DB.Raw("select count(*) from addresses where user_id = ? and id = ?", orderBody.UserID, orderBody.AddressID).Scan(&count).Error; err != nil {
		return false, err
	}
	fmt.Println(count)
	return count > 0, nil

}

func (cr *orderRepository) OrderItemsFromCart(orderBody models.OrderIncoming, cartItems []models.Cart) (domain.OrderSuccessResponse, error) {

	var orderDetails domain.Order
	var orderItemDetails domain.OrderItem

	// add general order details - that is to be added to orders table
	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	orderDetails.OrderId = str[:8]
	// details being added to the orders table
	orderDetails.AddressID = orderBody.AddressID
	orderDetails.PaymentMethodID = orderBody.PaymentID
	orderDetails.UserID = int(orderBody.UserID)
	orderDetails.Approval = false
	orderDetails.ShipmentStatus = "processing"

	// get grand total iterating through each products in carts
	for _, c := range cartItems {
		orderDetails.GrandTotal += c.TotalPrice
	}



	discount_price, err := helper.GetCouponDiscountPrice(int(orderBody.UserID), orderDetails.GrandTotal, cr.DB)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if discount_price != 0.0 {
		cr.DB.Exec("update used_coupons set used = true where user_id = ?", orderDetails.UserID)
	}
	orderDetails.FinalPrice = orderDetails.GrandTotal - discount_price
	cr.DB.Create(&orderDetails)
	// details being added to the orderItems table - which shows details about the individual products
	for _, c := range cartItems {
		// for each order save details of products and associated details and use order_id as foreign key ( for each order multiple product will be there)
		fmt.Println(c)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		// after creating the order delete all cart items and also update the quantity of the product
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

func (cr *orderRepository) GetOrderDetails(userID int, page int) ([]models.FullOrderDetails, error) {
	// details of order created byt his particular user
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2

	var orderDetails []models.OrderDetails
	cr.DB.Raw("select order_id,grand_total,shipment_status from orders where user_id = ? limit ? offset ? ", userID, 2, offset).Scan(&orderDetails)
	fmt.Println(orderDetails)

	var fullOrderDetails []models.FullOrderDetails
	// for each order select all the associated products and their details
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

func (cr *orderRepository) GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error) {

	var orderProductDetails []models.OrderProducts
	if err := cr.DB.Raw("select product_id,quantity from order_items where order_id = ?", orderID).Scan(&orderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}
	fmt.Println(orderProductDetails)
	return orderProductDetails, nil
}

func (cr *orderRepository) UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error {
	fmt.Println("the code reached update qunatity")
	fmt.Println(orderProducts)
	for _, o := range orderProducts {
		fmt.Println("quantity = ", o.Quantity, "product: ", o.ProductId)
		var quantity int
		if err := cr.DB.Raw("select quantity from products where id = ?", o.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}
		fmt.Println("products quantity = ", quantity)

		o.Quantity += quantity
		fmt.Println("updated quantity = ", o.Quantity)
		if err := cr.DB.Exec("update products set quantity = ? where id = ?", o.Quantity, o.ProductId).Error; err != nil {
			return err
		}
	}

	return nil

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

func (cr *orderRepository) GetOrderDetailsBrief(page int) ([]models.OrderDetails, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2
	var orderDetails []models.OrderDetails
	err := cr.DB.Raw("select order_id,grand_total,shipment_status from orders limit ? offset ?", 2, offset).Scan(&orderDetails).Error
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

func (cr *orderRepository) SavePayment(charge domain.Charge) error {
	if err := cr.DB.Create(&charge).Error; err != nil {
		return err
	}
	return nil

}

func (cr *orderRepository) GetPaymentDetails(OrderID string) (domain.Charge, error) {

	var paymentDetails domain.Charge
	if err := cr.DB.Raw("select orders.order_id,orders.grand_total,users.email from orders inner join users on orders.user_id = users.id where order_id = ?", OrderID).Scan(&paymentDetails).Error; err != nil {
		return domain.Charge{}, err
	}
	// fmt.Println("amount no problem",productDetails)
	return paymentDetails, nil
}
