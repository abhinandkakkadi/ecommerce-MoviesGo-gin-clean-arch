package handler

import (
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
