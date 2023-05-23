package handler

import (
	"net/http"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
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

func (cr *AdminHandler) LoginHandler(c *gin.Context) {

	// var adminDetails models.AdminLogin
	var adminDetails domain.Admin
	
	if err := c.BindJSON(&adminDetails); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"details not in correct format",
		})
		return
	}


	admin,err := cr.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK,admin)

}

func (cr *AdminHandler) SignUpHandler(c *gin.Context) {

	var admin domain.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest,"fields provided are wrong")
	}

	adminDetails, err := cr.adminUseCase.SignUpHandler(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,adminDetails)
	
}

func (cr *AdminHandler) GetUsers(c *gin.Context) {

	users,err := cr.adminUseCase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err,
		})
		return
	}

	c.JSON(http.StatusOK,users)

}

func (cr *AdminHandler) GetGenres(c *gin.Context) {
	
	genres,err  := cr.adminUseCase.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err,
		})
		return
	}

	c.JSON(http.StatusOK,genres)
}

func (cr *AdminHandler) AddCategory(c *gin.Context) {

	var category domain.CategoryManagement
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"body data's not in the right format",
		})
		return
	}

	added_genre,err := cr.adminUseCase.AddCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	c.JSON(http.StatusOK,added_genre)

}

func (cr *AdminHandler) DeleteGenre(c *gin.Context) {
	
	genre_id := c.Param("id")
	err := cr.adminUseCase.Delete(genre_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
	}

	c.JSON(http.StatusOK,"genre successfully deleted")

}



func (cr *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Param("id")
	err := cr.adminUseCase.BlockUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"err":err,
		})
		return
	}

	c.JSON(http.StatusOK,"Blocked the user")

}