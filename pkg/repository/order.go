package repository

import (
	"errors"
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

func (o *orderRepository) AddressExist(orderBody models.OrderIncoming) (bool, error) {
	fmt.Println("user id = ", orderBody.UserID, "address id = ", orderBody.AddressID)
	var count int
	if err := o.DB.Raw("select count(*) from addresses where user_id = ? and id = ?", orderBody.UserID, orderBody.AddressID).Scan(&count).Error; err != nil {
		return false, err
	}
	fmt.Println(count)
	return count > 0, nil

}

func (o *orderRepository) OrderItemsFromCart(orderBody models.OrderIncoming, cartItems []models.Cart) (domain.OrderSuccessResponse, error) {

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
	orderDetails.PaymentStatus = "not paid"

	// get grand total iterating through each products in carts
	for _, c := range cartItems {
		orderDetails.GrandTotal += c.TotalPrice
	}

	discount_price, err := helper.GetCouponDiscountPrice(int(orderBody.UserID), orderDetails.GrandTotal, o.DB)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if discount_price != 0.0 {
		o.DB.Exec("update used_coupons set used = true where user_id = ?", orderDetails.UserID)
	}
	orderDetails.FinalPrice = orderDetails.GrandTotal - discount_price
	fmt.Println("payment ID = ", orderBody.PaymentID)
	if orderBody.PaymentID == 2 {
		orderDetails.PaymentStatus = "not paid"
		orderDetails.ShipmentStatus = "pending"
	}
	// if the payment method is wallet
	if orderBody.PaymentID == 3 {

		var walletAvailable float64
		o.DB.Raw("select wallet_amount from wallets where user_id = ?", orderBody.UserID).Scan(&walletAvailable)

		// if wallet amount is less than finalamount - make paymentstatus - not paid and shipment status pending
		if walletAvailable < orderDetails.FinalPrice {
			orderDetails.PaymentStatus = "not paid"
			orderDetails.ShipmentStatus = "pending"
		} else {
			o.DB.Exec("update wallets set wallet_amount = ? where user_id = ? ", walletAvailable-orderDetails.FinalPrice, orderBody.UserID)
			orderDetails.PaymentStatus = "paid"
		}

	}

	o.DB.Create(&orderDetails)

	// details being added to the orderItems table - which shows details about the individual products
	for _, c := range cartItems {
		// for each order save details of products and associated details and use order_id as foreign key ( for each order multiple product will be there)
		fmt.Println(c)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		// after creating the order delete all cart items and also update the quantity of the product
		o.DB.Omit("id").Create(&orderItemDetails)
		o.DB.Exec("delete from carts where user_id = ? and product_id = ?", orderDetails.UserID, c.ProductID)
		fmt.Println(c.Quantity)
		fmt.Println(c.ProductID)
		o.DB.Exec("update products set quantity = quantity - ? where id = ?", c.Quantity, c.ProductID)
	}

	var orderSuccessResponse domain.OrderSuccessResponse
	o.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderDetails.OrderId).Scan(&orderSuccessResponse)
	return orderSuccessResponse, nil
}

func (o *orderRepository) GetOrderDetails(userID int, page int,count int) ([]models.FullOrderDetails, error) {
	// details of order created byt his particular user
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count

	var orderDetails []models.OrderDetails
	o.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where user_id = ? limit ? offset ? ", userID, count, offset).Scan(&orderDetails)
	fmt.Println(orderDetails)

	var fullOrderDetails []models.FullOrderDetails
	// for each order select all the associated products and their details
	for _, od := range orderDetails {

		var orderProductDetails []models.OrderProductDetails
		o.DB.Raw("select order_items.product_id,products.movie_name,order_items.quantity,order_items.total_price from order_items inner join products on order_items.product_id = products.id where order_items.order_id = ?", od.OrderId).Scan(&orderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.FullOrderDetails{OrderDetails: od, OrderProductDetails: orderProductDetails})

	}

	return fullOrderDetails, nil

}

func (o *orderRepository) UserOrderRelationship(orderID string, userID int) (int, error) {

	var testUserID int
	err := o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&testUserID).Error
	if err != nil {
		return -1, err
	}
	return testUserID, nil

}

func (o *orderRepository) GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error) {

	var orderProductDetails []models.OrderProducts
	if err := o.DB.Raw("select product_id,quantity from order_items where order_id = ?", orderID).Scan(&orderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}
	fmt.Println(orderProductDetails)
	return orderProductDetails, nil
}

func (o *orderRepository) UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error {
	fmt.Println("the code reached update qunatity")
	fmt.Println(orderProducts)
	for _, od := range orderProducts {
		fmt.Println("quantity = ", od.Quantity, "product: ", od.ProductId)
		var quantity int
		if err := o.DB.Raw("select quantity from products where id = ?", od.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}
		fmt.Println("products quantity = ", quantity)

		od.Quantity += quantity
		fmt.Println("updated quantity = ", od.Quantity)
		if err := o.DB.Exec("update products set quantity = ? where id = ?", od.Quantity, od.ProductId).Error; err != nil {
			return err
		}
	}

	return nil

}

func (o *orderRepository) CancelOrder(orderID string) (string, error) {

	var shipmentStatus string
	fmt.Println(orderID)
	err := o.DB.Raw("select shipment_status from orders where order_id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}

	fmt.Println(shipmentStatus)
	fmt.Println("ok")
	if shipmentStatus == "delivered" {
		return "Item already delivered, cannot cancel", nil
	}

	if shipmentStatus == "pending" {
		return "The order was not placed, so no point in cancelling", nil
	}

	if shipmentStatus == "cancelled" {
		return "Order already cancelled", nil
	}

	shipmentStatus = "cancelled"
	err = o.DB.Exec("update orders set shipment_status = ?  where order_id = ?", shipmentStatus, orderID).Error
	if err != nil {
		return "", err
	}

	var paymentMethod int
	err = o.DB.Raw("select payment_method_id from orders where order_id = ?", orderID).Scan(&paymentMethod).Error
	if err != nil {
		return "", err
	}

	if paymentMethod == 3 || paymentMethod == 2 {
		fmt.Println("the code reached here since this order is done by wallet")
		type AmountDetails struct {
			FinalPrice float64
			UserID     int
		}
		var amountDetails AmountDetails
		err = o.DB.Raw("select final_price,user_id from orders where order_id = ?", orderID).Scan(&amountDetails).Error
		if err != nil {
			return "", err
		}
		fmt.Println("amount details = ", amountDetails)
		err = o.DB.Exec("update wallets set wallet_amount = wallet_amount + ? where user_id = ?", amountDetails.FinalPrice, amountDetails.UserID).Error
		if err != nil {
			return "", err
		}

	}

	return "Order successfully cancelled", nil
}

func (o *orderRepository) GetOrderDetailsBrief(page int) ([]models.OrderDetails, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2
	var orderDetails []models.OrderDetails
	err := o.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders limit ? offset ?", 2, offset).Scan(&orderDetails).Error
	if err != nil {
		return []models.OrderDetails{}, nil
	}

	return orderDetails, nil
}

func (o *orderRepository) CheckOrderID(orderID string) (bool, error) {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (o *orderRepository) GetShipmentStatus(orderID string) (string, error) {

	var shipmentStatus string
	err := o.DB.Raw("select shipment_status from orders where order_id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}

	return shipmentStatus, nil

}

func (o *orderRepository) ApproveOrder(orderID string) error {

	err := o.DB.Exec("update orders set shipment_status = 'order placed',approval = true where order_id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) SavePayment(charge domain.Charge) error {
	if err := o.DB.Create(&charge).Error; err != nil {
		return err
	}
	return nil

}

func (o *orderRepository) GetPaymentDetails(OrderID string) (domain.Charge, error) {

	var paymentDetails domain.Charge
	if err := o.DB.Raw("select orders.order_id,orders.grand_total,users.email from orders inner join users on orders.user_id = users.id where order_id = ?", OrderID).Scan(&paymentDetails).Error; err != nil {
		return domain.Charge{}, err
	}
	// fmt.Println("amount no problem",productDetails)
	return paymentDetails, nil
}

func (o *orderRepository) CheckOrder(orderID string, userID int) error {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 0 {
		return errors.New("no such order exist")
	}
	var checkUser int
	err = o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&checkUser).Error
	if err != nil {
		return err
	}

	if userID != checkUser {
		return errors.New("the order is not done by this user")
	}

	return nil
}

func (o *orderRepository) GetOrderDetail(orderID string) (models.OrderDetails, error) {

	var orderDetails models.OrderDetails
	err := o.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where order_id = ?", orderID).Scan(&orderDetails).Error
	if err != nil {
		return models.OrderDetails{}, err
	}

	return orderDetails, nil

}

func (o *orderRepository) AddRazorPayDetails(orderID string, razorPayOrderID string) error {

	err := o.DB.Exec("insert into razer_pays (order_id,razor_id) values (?,?)", orderID, razorPayOrderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) CheckPaymentStatus(razorID string) error {

	fmt.Println(razorID)
	var OrderID string
	err := o.DB.Raw("select order_id from razer_pays where razor_id = ?", razorID).Scan(&OrderID).Error
	if err != nil {
		return err
	}

	fmt.Print("order id corresponding to razor id := ", OrderID)
	var paymentStatus string
	err = o.DB.Raw("select payment_status from orders where order_id = ?", OrderID).Scan(&paymentStatus).Error
	if err != nil {
		return err
	}

	if paymentStatus == "not paid" {
		fmt.Println("have to reach here")
		err = o.DB.Exec("update orders set payment_status = paid, shipment_status = processing from orders where order_id = ?", OrderID).Scan(&paymentStatus).Error
		if err != nil {
			return err
		}
		return nil
	}
	fmt.Println("should not reach here")
	return errors.New("already paid")

}

func (o *orderRepository) UpdatePaymentDetails(razorID string, paymentID string) error {

	err := o.DB.Exec("update razer_pays set payment_id = ? where razor_id = ?", paymentID, razorID).Error
	if err != nil {
		return err
	}
	return nil

}
