package handler

import (
	"fmt"
	"net/http"
	"strconv"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUseCase services.CartUseCase
}

func NewCartHandler(usecase services.CartUseCase) *CartHandler {

	return &CartHandler{
		cartUseCase: usecase,
	}

}

// AddToCart
// @Summary Add a product to the cart
// @ID add-to-cart
// @Description Add a product to the user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Security bearerAuth
// @Success 200 {object} response.Response
// @Failure 402 {object} response.Response
// @Router /cart/{id} [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {

	id := c.Param("id")
	product_id, _ := strconv.Atoi(id)

	userID, _ := c.Get("user_id")
	fmt.Println(product_id, userID)

	cartResponse, err := cr.cartUseCase.AddToCart(product_id, userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 402,
			Error:      err.Error(),
			Data:       nil,
			Message:    "could not add product to the cart",
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       cartResponse,
		Message:    "successfully added product to the cart",
	})

}

// @Summary Remove items from the cart
// @ID remove-from-cart
// @Description Remove items from the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Security bearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cart/{id} [delete]
func (cr *CartHandler) RemoveFromCart(c *gin.Context) {

	id := c.Param("id")
	product_id, _ := strconv.Atoi(id)

	userID, _ := c.Get("user_id")

	updatedCart, err := cr.cartUseCase.RemoveFromCart(product_id, userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "could not delete cart item",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       updatedCart,
		Message:    "Cart item deleted",
	})
}

// @Summary Display cart items
// @ID display-cart
// @Description Display cart items
// @Tags Cart
// @Accept json
// @Produce json
// @Security bearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cart [get]
func (cr *CartHandler) DisplayCart(c *gin.Context) {

	userID, _ := c.Get("user_id")
	cart, err := cr.cartUseCase.DisplayCart(userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not display cart items",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       cart,
		Message:    "Cart items displayed successfully",
	})
}

/// @Summary Delete all items from the cart
// @ID empty-cart
// @Description Delete all items from the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security bearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cart [delete]
func (cr *CartHandler) EmptyCart(c *gin.Context) {

	userID, _ := c.Get("user_id")
	emptyCart, err := cr.cartUseCase.EmptyCart(userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not display cart items",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       emptyCart,
		Message:    "Cart items displayed successfully",
	})
}
