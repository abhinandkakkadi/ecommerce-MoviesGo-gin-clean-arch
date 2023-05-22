package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
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


	admin,err := cr.adminUseCase.LoginHandler(c.Request.Context(),adminDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK,admin)

}

func (cr *AdminHandler) SignupHandler(c *gin.Context) {

	var admin domain.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest,"fields provided are wrong")
	}

	adminDetails, err := cr.adminUseCase.SignupHandler(c.Request.Context(),admin)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,adminDetails)
	
}

func (cr *AdminHandler) GetUsers(c *gin.Context) {

	var users []models.UserDetails
	users,err := cr.adminUseCase.GetUsers(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err,
		})
		return
	}

	c.JSON(http.StatusOK,users)
}

func (cr *AdminHandler) GetGenres(c *gin.Context) {
	fmt.Println("the code reached here")
	// var genres []domain.Genre
	genres,err  := cr.adminUseCase.GetGenres(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Could not retreive genres",
		})
		return
	}
	fmt.Println("the code reached here")
	c.JSON(http.StatusOK,genres)
}

func (cr *AdminHandler) AddGenre(c *gin.Context) {

	var genre domain.Genre
	if err := c.BindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"body data's not in the right fomrat",
		})
		return
	}
	added_genre,err := cr.adminUseCase.AddGenre(c.Request.Context(),genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	c.JSON(http.StatusOK,added_genre)
}

func (cr *AdminHandler) DeleteGenre(c *gin.Context) {
	fmt.Println("reached here")
	genre_id := c.Param("id")
	fmt.Println(genre_id)
	err := cr.adminUseCase.Delete(c.Request.Context(),genre_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
	}

	c.JSON(http.StatusOK,"Success")
}



func (cr *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Param("id")
	err := cr.adminUseCase.BlockUser(c.Request.Context(),id)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"err":err,
		})
		return
	}

	c.JSON(http.StatusOK,"Blocked the user")
}