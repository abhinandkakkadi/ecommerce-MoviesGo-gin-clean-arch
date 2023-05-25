package handler

import (
	"fmt"
	"strconv"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
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

	cr.cartUseCase.AddToCart(product_id,userID.(int))

}