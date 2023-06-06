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
func (cr *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	id, _ := c.Get("user_id")
	userID := id.(int)

	var orderFromCart models.OrderFromCart
	if err := c.BindJSON(&orderFromCart); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
			Message:    "bad request",
		})
		return
	}

	

	orderSuccessResponse, err := cr.orderUseCase.OrderItemsFromCart(orderFromCart,userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not do the order",
		})
		return
	}

	// if orderBody.PaymentID == 2 {
	// 	c.HTML(http.StatusOK,"index.html")
	// }
	// if orderBody.PaymentID == 2 {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"content": "This is an index page...",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       orderSuccessResponse,
		Message:    "Successfully created the order",
	})

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
func (cr *OrderHandler) GetOrderDetails(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	id, _ := c.Get("user_id")
	userID := id.(int)

	fullOrderDetails, err := cr.orderUseCase.GetOrderDetails(userID, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not do the order",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       fullOrderDetails,
		Message:    "Full Order Details",
	})

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
func (cr *OrderHandler) CancelOrder(c *gin.Context) {

	orderID := c.Param("id")
	fmt.Println(orderID)
	id, _ := c.Get("user_id")
	userID := id.(int)

	message, err := cr.orderUseCase.CancelOrder(orderID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not cancel the order",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       message,
		Message:    "Cancel Successfull",
	})

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
func (cr *OrderHandler) GetAllOrderDetailsForAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	allOrderDetails, err := cr.orderUseCase.GetAllOrderDetailsForAdmin(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not retrieve order details",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       allOrderDetails,
		Message:    "Order Details Retrieved successfully",
	})

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
func (cr *OrderHandler) ApproveOrder(c *gin.Context) {

	orderId := c.Param("order_id")
	message, err := cr.orderUseCase.ApproveOrder(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "could not approve the order",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       message,
		Message:    "Order Details Retrieved successfully",
	})

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
func (cr *OrderHandler) CancelOrderFromAdminSide(c *gin.Context) {

	orderID := c.Param("order_id")
	fmt.Println(orderID)

	message, err := cr.orderUseCase.CancelOrderFromAdminSide(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not cancel the order",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       message,
		Message:    "Cancel Successfull",
	})
}
