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
// @Param coupon body models.AddCoupon true "Add new Coupon"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon/addcoupon [post]
func (co *CouponHandler) AddCoupon(c *gin.Context) {

	var coupon models.AddCoupon
	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest,"could not bind the coupon details",nil,err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println(coupon)
	message, err := co.couponUseCase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError,"Could not add coupon",nil,err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK,"Coupon Added",message,nil)
	c.JSON(http.StatusCreated, successRes)

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
func (co *CouponHandler) GetCoupon(c *gin.Context) {

	coupons, err := co.couponUseCase.GetCoupon()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError,"Could not get coupon details",nil,err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK,"Coupon Retrieved successfully",coupons,nil)
	c.JSON(http.StatusOK, successRes)

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
func (co *CouponHandler) ExpireCoupon(c *gin.Context) {

	id := c.Param("id")
	couponID, err := strconv.Atoi(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest,"coupon id not in correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = co.couponUseCase.ExpireCoupon(couponID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError,"could not expire coupon",nil,err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK,"Coupon expired successfully",nil,nil)
	c.JSON(http.StatusNoContent, successRes)

}
