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
// @Description Add A new Coupon which can be used by the users from the checkout section
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.AddCoupon true "Add new Coupon"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/coupons/addcoupon [post]
func (co *CouponHandler) AddCoupon(c *gin.Context) {

	var coupon models.AddCoupon
	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not bind the coupon details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println(coupon)
	message, err := co.couponUseCase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not add coupon", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Coupon Added", message, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Get coupon details
// @Description Get Available coupon details for admin side
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/coupons [get]
func (co *CouponHandler) GetCoupon(c *gin.Context) {

	coupons, err := co.couponUseCase.GetCoupon()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not get coupon details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Coupon Retrieved successfully", coupons, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Expire Coupon
// @Description Expire Coupon by admin which are already present by passing coupon id
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Coupon id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/coupons/expire/{id} [get]
func (co *CouponHandler) ExpireCoupon(c *gin.Context) {

	id := c.Param("id")
	couponID, err := strconv.Atoi(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coupon id not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = co.couponUseCase.ExpireCoupon(couponID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not expire coupon", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Coupon expired successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add  Product Offer
// @Description Add a new Offer for a product by specifying a limit
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.ProductOfferReceiver true "Add new Product Offer"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/product-offer [post]
func (co *CouponHandler) AddProdcutOffer(c *gin.Context) {

	var productOffer models.ProductOfferReceiver
	if err := c.BindJSON(&productOffer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "request fields in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := co.couponUseCase.AddProductOffer(productOffer)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add offer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully added offer", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Add  Category Offer
// @Description Add a new Offer for a Category by specifying a limit
// @Tags Coupon
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.CategoryOfferReceiver true "Add new Category Offer"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/category-offer [post]
func (co *CouponHandler) AddCategoryOffer(c *gin.Context) {

	var categoryOffer models.CategoryOfferReceiver
	if err := c.BindJSON(&categoryOffer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "request fields in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := co.couponUseCase.AddCategoryOffer(categoryOffer)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add offer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully added offer", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}
