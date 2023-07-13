package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
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

func (o *orderRepository) DoesCartExist(userID int) (bool, error) {

	var exist bool
	err := o.DB.Raw("select exists(select 1 from carts where user_id = ?)", userID).Scan(&exist).Error
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (o *orderRepository) AddressExist(orderBody models.OrderIncoming) (bool, error) {

	var count int
	if err := o.DB.Raw("select count(*) from addresses where user_id = ? and id = ?", orderBody.UserID, orderBody.AddressID).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil

}

func (o *orderRepository) UpdateCouponDetails(discount_price float64, UserID int) error {

	if discount_price != 0.0 {
		err := o.DB.Exec("update used_coupons set used = true where user_id = ?", UserID).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *orderRepository) GetWalletAmount(UserID uint) (float64, error) {

	var walletAvailable float64
	err := o.DB.Raw("select wallet_amount from wallets where user_id = ?", UserID).Scan(&walletAvailable).Error
	if err != nil {
		return 0.0, err
	}

	return walletAvailable, nil
}

func (o *orderRepository) GetCouponDiscountPrice(UserID int, GrandTotal float64) (float64, error) {

	discountPrice, err := helper.GetCouponDiscountPrice(UserID, GrandTotal, o.DB)
	if err != nil {
		return 0.0, err
	}

	return discountPrice, nil

}

func (o *orderRepository) UpdateWalletAmount(walletAmount float64, UserID uint) error {

	err := o.DB.Exec("update wallets set wallet_amount = ? where user_id = ? ", walletAmount, UserID).Error
	if err != nil {
		return err
	}
	return nil

}

func (o *orderRepository) CreateOrder(orderDetails domain.Order) error {

	err := o.DB.Create(&orderDetails).Error
	if err != nil {
		return err
	}
	return nil

}

func (o *orderRepository) AddOrderItems(orderItemDetails domain.OrderItem, UserID int, ProductID uint, Quantity float64) error {

	// after creating the order delete all cart items and also update the quantity of the product
	err := o.DB.Omit("id").Create(&orderItemDetails).Error
	if err != nil {
		return err
	}

	err = o.DB.Exec("delete from carts where user_id = ? and product_id = ?", UserID, ProductID).Error
	if err != nil {
		return err
	}

	err = o.DB.Exec("update products set quantity = quantity - ? where id = ?", Quantity, ProductID).Error
	if err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) UpdateUsedOfferDetails(userID uint) error {

	o.DB.Exec("update category_offer_useds set used = true where user_id = ?", userID)
	o.DB.Exec("update product_offer_useds set used = true where user_id = ?", userID)

	return nil
}

func (o *orderRepository) GetBriefOrderDetails(orderID string) (domain.OrderSuccessResponse, error) {

	var orderSuccessResponse domain.OrderSuccessResponse
	o.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderID).Scan(&orderSuccessResponse)
	return orderSuccessResponse, nil

}

func (o *orderRepository) GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error) {
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

	return orderProductDetails, nil
}

func (o *orderRepository) UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error {

	for _, od := range orderProducts {

		var quantity int
		if err := o.DB.Raw("select quantity from products where id = ?", od.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}

		od.Quantity += quantity
		if err := o.DB.Exec("update products set quantity = ? where id = ?", od.Quantity, od.ProductId).Error; err != nil {
			return err
		}
	}

	return nil

}

func (o *orderRepository) CancelOrder(orderID string) error {

	shipmentStatus := "cancelled"
	err := o.DB.Exec("update orders set shipment_status = ? where order_id = ?", shipmentStatus, orderID).Error
	if err != nil {
		return err
	}

	var paymentMethod int
	err = o.DB.Raw("select payment_method_id from orders where order_id = ?", orderID).Scan(&paymentMethod).Error
	if err != nil {
		return err
	}

	if paymentMethod == 3 || paymentMethod == 2 {

		err = o.DB.Exec("update orders set payment_status = 'refunded'  where order_id = ?", orderID).Error
		if err != nil {
			return err
		}

		type AmountDetails struct {
			FinalPrice float64
			UserID     int
		}

		var amountDetails AmountDetails
		err = o.DB.Raw("select final_price,user_id from orders where order_id = ?", orderID).Scan(&amountDetails).Error
		if err != nil {
			return err
		}

		// check if a user have a uer have a wallet record if not create on
		result := o.DB.Exec("update wallets set wallet_amount = wallet_amount + ? where user_id = ?", amountDetails.FinalPrice, amountDetails.UserID)
		if result.Error != nil {
			return err
		}

		// if update didn't effect any row that means the record is not present
		if result.RowsAffected == 0 {
			result := o.DB.Exec("insert into wallets (user_id,wallet_amount) values(?,?)", amountDetails.UserID, amountDetails.FinalPrice)
			if result.Error != nil {
				return err
			}
		}

	}

	return nil
}

func (o *orderRepository) GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2
	var orderDetails []models.CombinedOrderDetails

	err := o.DB.Raw("select orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.name,users.email,users.phone,addresses.house_name,addresses.state,addresses.pin,addresses.street,addresses.city from orders inner join users on orders.user_id = users.id inner join addresses on users.id = addresses.user_id limit ? offset ?", 2, offset).Scan(&orderDetails).Error

	if err != nil {
		return []models.CombinedOrderDetails{}, nil
	}

	return orderDetails, nil
}

func (o *orderRepository) GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error) {

	var orderDetails models.CombinedOrderDetails

	err := o.DB.Raw("select orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.name,users.email,users.phone,addresses.house_name,addresses.state,addresses.pin,addresses.street,addresses.city from orders inner join users on orders.user_id = users.id inner join addresses on users.id = addresses.user_id where order_id = ?", orderID).Scan(&orderDetails).Error

	if err != nil {
		return models.CombinedOrderDetails{}, nil
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

func (o *orderRepository) CheckPaymentStatus(razorID string, orderID string) (string, error) {

	var paymentStatus string
	err := o.DB.Raw("select payment_status from orders where order_id = ?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}

	return paymentStatus, nil

}

func (o *orderRepository) UpdateShipmentAndPaymentByOrderID(shipmentStatus string, paymentStatus string, orderID string) error {

	err := o.DB.Exec("update orders set payment_status = ?, shipment_status = ? where order_id = ?", paymentStatus, shipmentStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) UpdatePaymentDetails(orderID string, paymentID string) error {

	err := o.DB.Exec("update razer_pays set payment_id = ? where order_id = ?", paymentID, orderID).Error
	if err != nil {
		return err
	}
	return nil

}

func (o *orderRepository) UpdateShipmentStatus(shipmentStatus string, orderID string) error {

	currentTime := time.Now()
	err := o.DB.Exec("update orders set shipment_status = ?, payment_status = 'paid',delivery_time = ? where order_id = ?", shipmentStatus, currentTime, orderID).Error
	if err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) GetDeliveredTime(orderID string) (time.Time, error) {

	var deliveryTime time.Time
	err := o.DB.Raw("select delivery_time from orders where order_id = ?", orderID).Scan(&deliveryTime).Error
	if err != nil {
		return deliveryTime, err
	}

	return deliveryTime, nil

}

func (o *orderRepository) ReturnOrder(shipmentStatus string, orderID string) error {

	err := o.DB.Exec("update orders set shipment_status = ?, payment_status = 'refund-init' where order_id = ?", shipmentStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil

}

func (o *orderRepository) GetPaymentStatus(orderID string) (string, error) {

	var paymentStatus string
	err := o.DB.Raw("select payment_status from orders where order_id = ?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}

	return paymentStatus, nil

}

func (o *orderRepository) RefundOrder(paymentStatus string, orderID string) error {

	err := o.DB.Exec("update orders set payment_status = ?,shipment_status = 'returned' where order_id = ?", paymentStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil
}
