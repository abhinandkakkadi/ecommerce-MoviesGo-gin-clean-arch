package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// LoginHandler handles the login functionality for admin
// @Summary Admin login
// @Description Authenticates an admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param adminDetails body AdminRequest true "Admin credentials"
// @Success 200 {object} AdminResponse
// @Failure 400 {object} ErrorResponse
// @Router /admin/login [post]
func (cr *AdminHandler) LoginHandler(c *gin.Context) {  // login handler for the admin

	// var adminDetails models.AdminLogin
	var adminDetails domain.Admin
	fmt.Println("it is here")
	if err := c.BindJSON(&adminDetails); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "details not in the correct format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	admin, err := cr.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "cannot authenticate user",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Admin authenticated successfully",
		Data:       admin,
		Error:      nil,
	})

}
// signup handler for admin
func (cr *AdminHandler) SignUpHandler(c *gin.Context) {

	var admin models.AdminSignUp
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are wrong",
			Data:       nil,
			Error:      err.Error(),
		})
	}

	adminDetails, err := cr.adminUseCase.SignUpHandler(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "cannot authenticate user",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully signed up the user ",
		Data:       adminDetails,
		Error:      nil,
	})

}

// handler to get user details 
func (cr *AdminHandler) GetUsers(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	users, err := cr.adminUseCase.GetUsers(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "could not retrieve records",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved the users",
		Data:       users,
		Error:      nil,
	})

}
// handler to get teh categories (have to rename it as GetCategory)
func (cr *AdminHandler) GetGenres(c *gin.Context) {

	genres, err := cr.adminUseCase.GetFullCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved the genres",
		Data:       genres,
		Error:      nil,
	})
}

// handler to add new category (genres,language,directors,movie formats)
func (cr *AdminHandler) AddCategory(c *gin.Context) {

	var category models.CategoryUpdate
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	addedCategory, err := cr.adminUseCase.AddCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "The category could not be added",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully added the record",
		Data:       addedCategory,
		Error:      nil,
	})

}

// handler to delete genre ( have to modify this to delete the whole category)
func (cr *AdminHandler) DeleteGenre(c *gin.Context) {

	genre_id := c.Param("id")
	err := cr.adminUseCase.Delete(genre_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "could not delete the specified genre",
			Data:       nil,
			Error:      err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully deleted the product",
		Data:       nil,
		Error:      nil,
	})

}
// handler to block the user
func (cr *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Param("id")
	err := cr.adminUseCase.BlockUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "user could not be blocked",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Successfully blocked the user",
		Data:       nil,
		Error:      nil,
	})

}

// handler to unblock the user
func (cr *AdminHandler) UnBlockUser(c *gin.Context) {

	id := c.Param("id")
	err := cr.adminUseCase.UnBlockUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "user could not be unblocked",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Successfully unblocked the user",
		Data:       nil,
		Error:      nil,
	})
}



func (cr *UserHandler) AddNewUsers(c *gin.Context) {
	fmt.Println("add users")
	var userDetails models.UserDetails
  if err := c.BindJSON(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: http.StatusBadRequest,
			Message: "could not bind the user details",
			Data: nil,
			Error: err,
		})
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(userDetails)
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
	userCreated, err := cr.userUseCase.UserSignUp(userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "User could not be created up",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    "User successfully created",
		Data:       userCreated,
		Error:      nil,
	})

	



}	