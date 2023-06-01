package handler

import (
	"context"
	"fmt"
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

// sign up application handler for user sign up
func (cr *UserHandler) UserSignUp(c *gin.Context) {

	var user models.UserDetails
	// bind the user details to the struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "constraints not satisfied",
				Data:       nil,
				Error:      err.Error(),
			})
		return
	}

	// business logic goes inside this function
	userCreated, err := cr.userUseCase.UserSignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "User could not signed up",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    "User successfully signed up",
		Data:       userCreated,
		Error:      nil,
	})

}

// login func
func (cr *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "constraints not satisfied",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	user_details, err := cr.userUseCase.LoginHandler(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "User could not be logged in",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "User successfully logged in",
		Data:       user_details,
		Error:      nil,
	})

}

// handler to add a new address for the user
func (cr *UserHandler) AddAddress(c *gin.Context) {

	userID, _ := c.Get("user_id")

	var address models.AddressInfo

	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err := validator.New().Struct(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "constraints does not match",
			Data:       nil,
			Error:      err.Error(),
		})
	}

	addressResponse, err := cr.userUseCase.AddAddress(address, userID.(int))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed adding address",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "address added successfully",
		Data:       addressResponse,
		Error:      nil,
	})

}

// update existing address using address id
func (cr *UserHandler) UpdateAddress(c *gin.Context) {

	id := c.Param("id")
	addressId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "address id not in the right format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	userID, _ := c.Get("user_id")
	user_id := userID.(int)
	var address models.AddressInfo
	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	fmt.Println(address)
	address.UserID = uint(user_id)
	fmt.Println(address)
	updatedAddress, err := cr.userUseCase.UpdateAddress(address, addressId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed updating address",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "address updated successfully",
		Data:       updatedAddress,
		Error:      nil,
	})

}

// checkout section for users after adding items to the cart and adding address
func (cr *UserHandler) CheckOut(c *gin.Context) {

	userID, _ := c.Get("user_id")
	checkoutDetails, err := cr.userUseCase.Checkout(userID.(int))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to retrieve details",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Checkout Page loaded successfully",
		Data:       checkoutDetails,
		Error:      nil,
	})
}

func (cr *UserHandler) UserDetails(c *gin.Context) {

	userID, _ := c.Get("user_id")

	userDetails, err := cr.userUseCase.UserDetails(userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to retrieve details",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "User Details",
		Data:       userDetails,
		Error:      nil,
	})

}

// get all the address added by the user
func (cr *UserHandler) GetAllAddress(c *gin.Context) {

	userID, _ := c.Get("user_id")
	userAddress, err := cr.userUseCase.GetAllAddress(userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to retrieve details",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "User Address",
		Data:       userAddress,
		Error:      nil,
	})

}

// update details of uer in user profile section (can update optional details)
func (cr *UserHandler) UpdateUserDetails(c *gin.Context) {

	user_id, _ := c.Get("user_id")
	ctx := context.Background()

	ctx = context.WithValue(ctx, "userID", user_id.(int))

	var body models.UsersProfileDetails
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	updatedDetails, err := cr.userUseCase.UpdateUserDetails(body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed update user",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated User Details",
		Data:       updatedDetails,
		Error:      nil,
	})
}

// update password of the user
func (cr *UserHandler) UpdatePassword(c *gin.Context) {

	user_id, _ := c.Get("user_id")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", user_id.(int))

	var body models.UpdatePassword
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}
	// fmt.Printf(body.NewPassword)
	fmt.Println(body.ConfirmNewPassword)

	err := cr.userUseCase.UpdatePassword(ctx, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed updating password",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Password updated successfully ",
		Data:       nil,
		Error:      nil,
	})
}
