package handler

import (
	"fmt"
	"net/http"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCase
}

func NewPaymentHandler(useCase services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: useCase,
	}
}

func (p *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {

	orderID := c.Param("id")
	userID, _ := c.Get("user_id")

	orderDetail,razorID,err := p.paymentUseCase.MakePaymentRazorPay(orderID, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError,response.Response{
			StatusCode: http.StatusInternalServerError,
			Message: "could not generate order details",
			Error: err.Error(),
			Data: nil,
			
		})
		return
	}

	c.HTML(http.StatusOK,"index.html",gin.H{
		"final_price":orderDetail.FinalPrice*100,
		"order_id":razorID,
		"user_id":userID,
	})
}

func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	fmt.Println("Hello Abhinand")
	
}
