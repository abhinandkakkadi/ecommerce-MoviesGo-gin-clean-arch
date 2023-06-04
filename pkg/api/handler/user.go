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

// @Summary SignUp functionality for user
// @Description SignUp functionality at the user side
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserDetails true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /signup [post]
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

// @Summary LogIn functionality for user
// @Description LogIn functionality at the user side
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserDetails true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /login [post]
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

// @Summary AddAddress functionality for user
// @Description AddAddress functionality at the user side
// @Tags Users
// @Accept json
// @Produce json
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address [post]
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

// @Summary Update User Address
// @Description Update User address by sending in address id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "address id"
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address/{id} [put]
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

// @Summary Checkout Order
// @Description Checkout at the user side
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /checkout [get]
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

// @Summary User Details
// @Description User Details from User Profile
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users [get]
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

// @Summary Get all address for the user
// @Description Display all the added user addresses
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address [get]
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

func (cr *UserHandler) UpdateUserDetails(c *gin.Context) {

	user_id, _ := c.Get("user_id")
	ctx := context.Background()

	ctx = context.WithValue(ctx, "userID", user_id.(int))

	var user models.UsersProfileDetails
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	updatedDetails, err := cr.userUseCase.UpdateUserDetails(user, ctx)
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

// @Summary Update User Password
// @Description Update User Password
// @Tags Users
// @Accept json
// @Produce json
// @Param body body models.UpdatePassword true "User Password update"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/update-password [post]
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

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wishlist/add/{id} [get]
func (cr *UserHandler) AddToWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "product id is in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err = cr.userUseCase.AddToWishList(productID, userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to item to the wishlist",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "SuccessFully added product to the wishlist",
		Data:       nil,
		Error:      nil,
	})

}

// @Summary Display Wishlist
// @Description Display wish List
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wishlist [get]
func (cr *UserHandler) GetWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	wishList, err := cr.userUseCase.GetWishList(userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to retrieve wishlist details",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "SuccessFully retrieved wishlist",
		Data:       wishList,
		Error:      nil,
	})

}

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wishlist/remove/{id} [get]
func (cr *UserHandler) RemoveFromWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "product id is in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err = cr.userUseCase.RemoveFromWishList(productID, userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to remove item from wishlist",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "SuccessFully deleted product from wishlist",
		Data:       nil,
		Error:      nil,
	})

}
