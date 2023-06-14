package handler

import (
	"fmt"
	"net/http"
	"strconv"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	models "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}

// @Summary Order Items from cart
// @Description Order all products inside the cart
// @Tags Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param orderBody body models.OrderFromCart true "Order details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /order [post]
func (o *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	id, _ := c.Get("user_id")
	userID := id.(int)

	var orderFromCart models.OrderFromCart
	if err := c.BindJSON(&orderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orderSuccessResponse, err := o.orderUseCase.OrderItemsFromCart(orderFromCart, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Order Details to user side
// @Description Order all order details done by user
// @Tags Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "page number"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/orders/{id} [get]
func (o *OrderHandler) GetOrderDetails(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, _ := c.Get("user_id")
	userID := id.(int)

	fullOrderDetails, err := o.orderUseCase.GetOrderDetails(userID, page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Cancel order
// @Description Cancel order by the user using order ID
// @Tags Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/cancel-order/{id} [put]
func (o *OrderHandler) CancelOrder(c *gin.Context) {

	orderID := c.Param("id")
	fmt.Println(orderID)
	id, _ := c.Get("user_id")
	userID := id.(int)

	message, err := o.orderUseCase.CancelOrder(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", message, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get All order details for admin
// @Description Get all order details to the admin side
// @Tags Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/{id} [get]
func (o *OrderHandler) GetAllOrderDetailsForAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	allOrderDetails, err := o.orderUseCase.GetAllOrderDetailsForAdmin(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order Details Retrieved successfully", allOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Approve Order
// @Description Approve Order from admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/approve-order/{id} [get]
func (o *OrderHandler) ApproveOrder(c *gin.Context) {

	orderId := c.Param("order_id")
	message, err := o.orderUseCase.ApproveOrder(orderId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not approve the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order approved successfully", message, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Cancel Order Admin
// @Description Cancel Order from admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/cancel-order/{id} [get]
func (o *OrderHandler) CancelOrderFromAdminSide(c *gin.Context) {

	orderID := c.Param("order_id")
	fmt.Println(orderID)

	message, err := o.orderUseCase.CancelOrderFromAdminSide(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", message, nil)
	c.JSON(http.StatusOK, successRes)
}

func (o *OrderHandler) OrderDelivered(c *gin.Context) {

	orderID := c.Param("order_id")
	fmt.Println(orderID)
	err := o.orderUseCase.OrderDelivered(orderID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "order could not be delivered", nil, err)
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully delivered the product", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) ReturnOrder(c *gin.Context) {

	orderID := c.Param("order_id")
	fmt.Println(orderID)
	err := o.orderUseCase.ReturnOrder(orderID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "order could not be returned", nil, err)
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully delivered the returned", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) RefundUser(c *gin.Context) {

	orderID := c.Param("order_id")
	fmt.Println(orderID)
	err := o.orderUseCase.RefundOrder(orderID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "refund was not possible", nil, err)
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Refunded the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
