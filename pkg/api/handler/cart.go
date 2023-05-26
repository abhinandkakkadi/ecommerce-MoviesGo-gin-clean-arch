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


func (cr *CartHandler) AddToCart(c *gin.Context) {

	id := c.Param("id")
	product_id,_ := strconv.Atoi(id)

	userID, _ := c.Get("user_id")
	fmt.Println(product_id,userID)

	cartResponse,err := cr.cartUseCase.AddToCart(product_id,userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: 402,
			Error: err.Error(),
			Data: nil,
			Message: "sorry bro",
		})
		return
	}

	c.JSON(http.StatusCreated,response.Response{
		StatusCode: http.StatusOK,
		Error: nil,
		Data: cartResponse,
		Message:"successfully added product to the cart" ,
	})

	

}

func (cr *CartHandler) RemoveFromCart(c *gin.Context) {

	id := c.Param("id")
	product_id,_ := strconv.Atoi(id)

	userID, _ := c.Get("user_id")
	

	updatedCart,err := cr.cartUseCase.RemoveFromCart(product_id,userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusInternalServerError,
			Error: err.Error(),
			Data: nil,
			Message: "could not delete cart item",
		})
		return
	}

	c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusOK,
			Error: nil,
			Data: updatedCart,
			Message: "Cart item deleted",
	})
}


func (cr *CartHandler) DisplayCart(c *gin.Context) {
	
	userID, _ := c.Get("user_id")
	cart,err := cr.cartUseCase.DisplayCart(userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusInternalServerError,
			Error: err.Error(),
			Data: nil,
			Message: "Could not display cart items",
		})
		return
	}

	c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusOK,
			Error: nil,
			Data: cart,
			Message: "Cart items displayed successfully",
	})
}

func (cr *CartHandler) EmptyCart(c *gin.Context) {

	userID, _ := c.Get("user_id")
	emptyCart,err := cr.cartUseCase.EmptyCart(userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusInternalServerError,
			Error: err.Error(),
			Data: nil,
			Message: "Could not display cart items",
		})
		return
	}

	c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusOK,
			Error: nil,
			Data: emptyCart,
			Message: "Cart items displayed successfully",
	})
}