package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary SignUp functionality for user
// @Description SignUp functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.UserDetails true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /signup [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {

	var user models.UserDetails

	// bind the user details to the struct
	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest,
			errRes)
		return
	}

	// business logic goes inside this function
	userCreated, err := u.userUseCase.UserSignUp(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary LogIn functionality for user
// @Description LogIn functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /login [post]
func (u *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserLogin

	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user_details, err := u.userUseCase.LoginHandler(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", user_details, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary AddAddress functionality for user
// @Description AddAddress functionality at the user side
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address [post]
func (u *UserHandler) AddAddress(c *gin.Context) {

	userID, _ := c.Get("user_id")

	var address models.AddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(address)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints does not match", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = u.userUseCase.AddAddress(address, userID.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed adding address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "address added successfully", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Update User Address
// @Description Update User address by sending in address id
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "address id"
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address/{id} [put]
func (u *UserHandler) UpdateAddress(c *gin.Context) {

	id := c.Param("id")
	addressId, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "address id not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("user_id")
	user_id := userID.(int)

	var address models.AddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	updatedAddress, err := u.userUseCase.UpdateAddress(address, addressId, user_id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed updating address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "address updated successfully", updatedAddress, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Checkout Order
// @Description Checkout at the user side
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /checkout [get]
func (u *UserHandler) CheckOut(c *gin.Context) {

	userID, _ := c.Get("user_id")
	checkoutDetails, err := u.userUseCase.Checkout(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Checkout Page loaded successfully", checkoutDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary User Details
// @Description User Details from User Profile
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users [get]
func (u *UserHandler) UserDetails(c *gin.Context) {

	userID, _ := c.Get("user_id")

	userDetails, err := u.userUseCase.UserDetails(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "user Details", userDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get all address for the user
// @Description Display all the added user addresses
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/address [get]
func (u *UserHandler) GetAllAddress(c *gin.Context) {

	userID, _ := c.Get("user_id")
	userAddress, err := u.userUseCase.GetAllAddress(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User Address", userAddress, nil)
	c.JSON(http.StatusOK, successRes)

}

func (u *UserHandler) UpdateUserDetails(c *gin.Context) {

	user_id, _ := c.Get("user_id")
	ctx := context.Background()

	ctx = context.WithValue(ctx, "userID", user_id.(int))

	var user models.UsersProfileDetails

	if err := c.ShouldBindJSON(&user); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	updatedDetails, err := u.userUseCase.UpdateUserDetails(user, ctx)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed update user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Updated User Details", updatedDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Update User Password
// @Description Update User Password
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body models.UpdatePassword true "User Password update"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/update-password [put]
func (u *UserHandler) UpdatePassword(c *gin.Context) {

	user_id, _ := c.Get("user_id")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", user_id.(int))

	var body models.UpdatePassword

	if err := c.ShouldBindJSON(&body); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := u.userUseCase.UpdatePassword(ctx, body)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed updating password", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Password updated successfully", nil, nil)
	c.JSON(http.StatusCreated, successRes)
}

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wishlist/add/{id} [get]
func (u *UserHandler) AddToWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	id := c.Param("id")
	productID, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = u.userUseCase.AddToWishList(productID, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to item to the wishlist", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully added product to the wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Display Wishlist
// @Description Display wish List
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wishlist [get]
func (u *UserHandler) GetWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	wishList, err := u.userUseCase.GetWishList(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve wishlist detailss", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully retrieved wishlist", wishList, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wishlist/remove/{id} [delete]
func (u *UserHandler) RemoveFromWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	id := c.Param("id")
	productID, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = u.userUseCase.RemoveFromWishList(productID, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to remove item from wishlist", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully deleted product from wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Apply referrals
// @Description Apply referrals amount to order
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /referral/apply [get]
func (u *UserHandler) ApplyReferral(c *gin.Context) {

	userID, _ := c.Get("user_id")
	message, err := u.userUseCase.ApplyReferral(userID.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add referral amount", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	if message != "" {
		errRes := response.ClientResponse(http.StatusOK, message, nil, nil)
		c.JSON(http.StatusOK, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added referral amount", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Reset Password Using OTP
// @Description Reset Password using token Received from confirming OTP
// @Tags User Authentication
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body models.ResetPassword true "User Password Reset"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /forgot-password/reset [put]
func (u *UserHandler) ResetPassword(c *gin.Context) {

	userID, _ := c.Get("user_id")

	var pass models.ResetPassword

	if err := c.ShouldBindJSON(&pass); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := u.userUseCase.ResetPassword(userID.(int), pass)

	if err != nil {
		errRes := response.ClientResponse(http.StatusOK, "could not update the user", nil, err.Error())
		c.JSON(http.StatusOK, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully updated the password", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
