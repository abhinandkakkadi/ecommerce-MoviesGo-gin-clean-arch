package handler

import (
	"fmt"
	"net/http"
	"strconv"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
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

// @Summary Add to Cart
// @Description Add items to the cart
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "product-id"
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/addtocart/{id} [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {

	id := c.Param("id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Error:      err.Error(),
			Data:       nil,
			Message:    "product id not in correct format",
		})
		return
	}

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

// @Summary Remove Items from cart
// @Description Remove specified product of quantity 1 from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/removefromcart/{id} [delete]
func (cr *CartHandler) RemoveFromCart(c *gin.Context) {

	id := c.Param("id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Error:      err.Error(),
			Data:       nil,
			Message:    "product id not in correct format",
		})
		return
	}

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

// @Summary Display Cart
// @Description Display all items of the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
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

// @Summary Empty Items from cart
// @Description Remove all product from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
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

// @Summary Add coupon to out order
// @Description Add Coupon to get discount
// @Tags Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Param couponDetails body models.CouponAddUser true "Add coupon to order"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /coupon/add [post]
func (cr *CartHandler) AddCoupon(c *gin.Context) {

	userID, _ := c.Get("user_id")
	var couponDetails models.CouponAddUser

	if err := c.BindJSON(&couponDetails); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not bind the coupon",
		})
		return
	}

	err := cr.cartUseCase.AddCoupon(couponDetails.CouponName, userID.(int))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Response{
				StatusCode: http.StatusInternalServerError,
				Error:      err.Error(),
				Data:       nil,
				Message:    "Coupon could not be added",
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.Response{
		StatusCode: http.StatusOK,
		Error:      nil,
		Data:       nil,
		Message:    "Coupon added successfully",
	})

}
