package handler

import (
	"fmt"
	"net/http"

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

func (cr *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	id, _ := c.Get("user_id")
	userID := id.(int)

	var orderBody models.OrderIncoming
	if err := c.BindJSON(&orderBody); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
			Message:    "bad request",
		})
		return
	}

	orderBody.UserID = uint(userID)

	orderSuccessResponse, err := cr.orderUseCase.OrderItemsFromCart(orderBody)
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
		Data:       orderSuccessResponse,
		Message:    "Successfully created the order",
	})

}

func (cr *OrderHandler) GetOrderDetails(c *gin.Context) {

	id, _ := c.Get("user_id")
	userID := id.(int)

	fullOrderDetails, err := cr.orderUseCase.GetOrderDetails(userID)
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

func (cr *OrderHandler) GetAllOrderDetailsForAdmin(c *gin.Context) {

	allOrderDetails, err := cr.orderUseCase.GetAllOrderDetailsForAdmin()
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
