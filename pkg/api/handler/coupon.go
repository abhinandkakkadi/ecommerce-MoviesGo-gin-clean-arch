package handler

import (
	"fmt"
	"net/http"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponUseCase services.CouponUseCase
}

func NewCouponHandler(useCase services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		couponUseCase: useCase,
	}
}

func (cr *CouponHandler) AddCoupon(c *gin.Context) {

	var coupon models.Coupon
	if err := c.BindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
			Data:       nil,
			Message:    "could not bind the coupon details",
		})
		return
	}
	fmt.Println(coupon)
	message, err := cr.couponUseCase.AddCoupon(coupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not add coupon",
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Error:      nil,
		Data:       message,
		Message:    "Coupon Added",
	})

}

func (cr *CouponHandler) GetCoupon(c *gin.Context) {

	coupons, err := cr.couponUseCase.GetCoupon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "Could not get coupon details",
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Error:      nil,
		Data:       coupons,
		Message:    "Coupon Retrieved successfully",
	})

}

func (cr *CouponHandler) ExpireCoupon(c *gin.Context) {

	id := c.Param("id")

}
