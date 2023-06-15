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
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("user_id")
	fmt.Println(product_id, userID)

	cartResponse, err := cr.cartUseCase.AddToCart(product_id, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not add product to the cart", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "successfully added product to the cart", cartResponse, nil)
	c.JSON(http.StatusCreated, successRes)

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
		errorRes := response.ClientResponse(http.StatusBadRequest, "product not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("user_id")

	updatedCart, err := cr.cartUseCase.RemoveFromCart(product_id, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not delete cart items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cart item deleted", updatedCart, nil)
	c.JSON(http.StatusOK, successRes)
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
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not display cart items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cart items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, successRes)
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
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not display cart items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "cart items displayed successfully", emptyCart, nil)
	c.JSON(http.StatusOK, successRes)
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
func (cr *CartHandler) ApplyCoupon(c *gin.Context) {

	userID, _ := c.Get("user_id")
	var couponDetails models.CouponAddUser

	if err := c.BindJSON(&couponDetails); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not bind the coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := cr.cartUseCase.ApplyCoupon(couponDetails.CouponName, userID.(int))
	if err != nil {
		if err != nil {
			errorRes := response.ClientResponse(http.StatusInternalServerError, "coupon could not be added", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		}
	}

	successRes := response.ClientResponse(http.StatusCreated, "Coupon added successfully", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}
