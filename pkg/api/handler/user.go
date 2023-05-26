package handler

import (
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
			Message: "fields provided are in wrong format",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.Response{
				StatusCode: http.StatusBadRequest,
				Message: "constraints not satisfied",
				Data: nil,
				Error: err.Error(),
			})
		return
	}

	// business logic goes inside this function
	userCreated, err := cr.userUseCase.UserSignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message: "User could not signed up",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message: "User successfully signed up",
		Data: userCreated,
		Error: err.Error(),
	})

}

// login func
func (cr *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message: "fields provided are in wrong format",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message: "constraints not satisfied",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	user_details, err := cr.userUseCase.LoginHandler(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message: "User could not be logged in",
			Data: nil,
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message: "User successfully logged in",
		Data: user_details,
		Error: nil,
	})

}



func (cr *UserHandler) AddAddress(c *gin.Context) {

	userID, _ := c.Get("user_id")

	var address models.AddressInfo

	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message: "fields provided are in wrong format",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	err := validator.New().Struct(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message: "constraints does not match",
			Data: nil,
			Error: err.Error(),
		})
	}

	addressResponse, err := cr.userUseCase.AddAddress(address,userID.(int))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message: "failed adding address",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message: "address added successfully",
		Data: addressResponse,
		Error: nil,
	})

}


func (cr *UserHandler) UpdateAddress(c *gin.Context) {

	id := c.Param("id")
	addressId,_ := strconv.Atoi(id)
	userID, _ := c.Get("user_id")
	user_id := userID.(int)
	var address models.AddressInfo
	if err := c.BindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message: "fields provided are in wrong format",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	fmt.Println(address)
	address.UserID = uint(user_id)
	fmt.Println(address)
	updatedAddress,err := cr.userUseCase.UpdateAddress(address,addressId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message: "failed updating address",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message: "address updated successfully",
		Data: updatedAddress,
		Error: nil,
	})

}