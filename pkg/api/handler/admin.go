package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

// @Summary Admin Login
// @Description Login handler for admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param  admin body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/adminlogin [post]
func (cr *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

	// var adminDetails models.AdminLogin
	var adminDetails models.AdminLogin
	fmt.Println("it is here")
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := cr.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Admin Signup
// @Description Signup handler for admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param  admin body models.AdminSignUp true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/adminsignup [post]
func (cr *AdminHandler) SignUpHandler(c *gin.Context) {

	var admin models.AdminSignUp
	if err := c.BindJSON(&admin); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	adminDetails, err := cr.adminUseCase.SignUpHandler(admin)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully signed up the user", adminDetails, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Get Users
// @Description Retrieve users with pagination
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/users/{page} [get]
func (ad *AdminHandler) GetUsers(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.adminUseCase.GetUsers(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get genre to admin side
// @Description Display genre details on the admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/genres [get]
func (ad *AdminHandler) GetGenres(c *gin.Context) {

	genres, err := ad.adminUseCase.GetFullCategory()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the genres", genres, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Add Category
// @Description Add Category for existing films
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param  category body models.CategoryUpdate true "Update Category"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/genres/add_genre [POST]
func (ad *AdminHandler) AddCategory(c *gin.Context) {

	var category models.CategoryUpdate
	if err := c.BindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	addedCategory, err := ad.adminUseCase.AddCategory(category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "The category could not be added", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "Successfully added the record", addedCategory, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Delete Category
// @Description Delete Category for existing films and films long with it
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "genre-id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/genres/delete_genre/{id} [POST]
func (ad *AdminHandler) DeleteGenre(c *gin.Context) {

	genre_id := c.Param("id")
	err := ad.adminUseCase.Delete(genre_id)
	errorRes := response.ClientResponse(http.StatusBadRequest, "could not delete the specified genre", nil, err.Error())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorRes)
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the product", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Block an existing user
// @Description Block user
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "user-id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/users/block-users/{id} [POST]
func (ad *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Param("id")
	err := ad.adminUseCase.BlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary UnBlock an existing user
// @Description UnBlock user
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "user-id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/users/unblock-users/{id} [POST]
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {

	id := c.Param("id")
	err := ad.adminUseCase.UnBlockUser(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// check this before doing ny operations on this.
func (ad *UserHandler) AddNewUsers(c *gin.Context) {
	fmt.Println("add users")
	var userDetails models.UserDetails
	if err := c.BindJSON(&userDetails); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not bind the user details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(userDetails)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest,
			errorRes)
		return
	}

	// business logic goes inside this function
	userCreated, err := ad.userUseCase.UserSignUp(userDetails)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be created", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully created", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)

}
