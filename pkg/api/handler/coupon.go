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

type CouponHandler struct {
	couponUseCase services.CouponUseCase
}

func NewCouponHandler(useCase services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		couponUseCase: useCase,
	}
}

// @Summary Add  a new coupon by Admin
// @Description Add A new Coupon which can be used by the users
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon/addcoupon [post]
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

// @Summary Get coupon details 
// @Description Get Available coupon details for admin
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon [get]
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


// @Summary Get coupon details 
// @Description Get Available coupon details for admin
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Coupon id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon/expire/{id} [get]
func (cr *CouponHandler) ExpireCoupon(c *gin.Context) {

	id := c.Param("id")
	couponID, _ := strconv.Atoi(id)

	err := cr.couponUseCase.ExpireCoupon(couponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Data:       nil,
			Message:    "could not expire coupon",
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Error:      nil,
		Data:       nil,
		Message:    "Coupon expired successfully",
	})

}
