package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(useCase services.ProductUseCase) *ProductHandler {
		return &ProductHandler{
			productUseCase: useCase,
		}
}

func (cr *ProductHandler) ShowAllProducts(c *gin.Context) {

	products,err  := cr.productUseCase.ShowAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
	}

	c.JSON(http.StatusCreated,products)

}

func (cr *ProductHandler) ShowIndividualProducts(c *gin.Context) {
	
	id := c.Param("id")
	product,err := cr.productUseCase.ShowIndividualProducts(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,product)
	
}


func (cr *ProductHandler) AddProduct(c *gin.Context) {

	var product domain.Products
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest,"error while binding")
		return
	}

	err := cr.productUseCase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err,
		})
		return
	}

	c.JSON(http.StatusOK,"success updating the product details")

}



func (cr *ProductHandler) DeleteProduct(c *gin.Context) {

	product_id := c.Param("id")
	err := cr.productUseCase.DeleteProduct(product_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,"successfully deleted")

}