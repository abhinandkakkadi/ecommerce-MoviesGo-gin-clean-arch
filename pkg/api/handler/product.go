package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/response"
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

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)

	products, err := cr.productUseCase.ShowAllProducts(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Could not retrieve products",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully Retrieved all products",
		Data:       products,
		Error:      nil,
	})

}

func (cr *ProductHandler) SeeAllProductToAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)

	products, err := cr.productUseCase.ShowAllProducts(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Could not retrieve products to admin side",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully Retrieved all products to admin side",
		Data:       products,
		Error:      nil,
	})

}

func (cr *ProductHandler) ShowIndividualProducts(c *gin.Context) {

	id := c.Param("id")
	product, err := cr.productUseCase.ShowIndividualProducts(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "path variables in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Product details retrieved successfully",
		Data:       product,
		Error:      nil,
	})

}

func (cr *ProductHandler) AddProduct(c *gin.Context) {

	var product models.ProductsReceiver
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	productResponse, err := cr.productUseCase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Could not add the product",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully added products",
		Data:       productResponse,
		Error:      nil,
	})

}

func (cr *ProductHandler) DeleteProduct(c *gin.Context) {

	product_id := c.Param("id")
	err := cr.productUseCase.DeleteProduct(product_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Successfully deleted the item",
		Data:       nil,
		Error:      nil,
	})

}
