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
func (cr *AdminHandler) LoginHandler(c *gin.Context) {

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