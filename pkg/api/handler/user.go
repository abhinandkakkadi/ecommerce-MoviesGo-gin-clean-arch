package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
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

	var user domain.Users
	// bind the user details to the struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Message:    "Error while binding",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "send data in correct format",
			})
		return
	}

	// business logic goes inside this function
	userCreated, err := cr.userUseCase.UserSignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error signing up",
		})
		return
	}

	c.JSON(http.StatusCreated, userCreated)

}

// login func
func (cr *UserHandler) LoginHandler(c *gin.Context) {

	var user domain.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "send data's in the correct format")
		return
	}

	user_details, err := cr.userUseCase.LoginHandler(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user_details)

}
