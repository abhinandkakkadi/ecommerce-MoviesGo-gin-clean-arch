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

// @Summary Get Products
// @Description Retrieve products with pagination
// @Tags Admin
// @Accept json
// @Produce json
// @Param page path string true "Page number"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/page/{page} [get]
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

// show detailed details about  the product including product decription
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

// handler to add a new product by authenticated admin
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

// handler to delete an existing product by admin
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

func (cr *ProductHandler) UpdateProduct(c *gin.Context) {

	type product struct {
		Quantity  int `json:"quantity"`
		ProductID int `json:"product-id"`
	}
	var p product

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	err := cr.productUseCase.UpdateProduct(p.ProductID, p.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "could not update the product",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Successfully updated the item",
		Data:       nil,
		Error:      nil,
	})

}

// handler to filter category

func (cr *ProductHandler) FilterCategory(c *gin.Context) {

	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	productCategory, err := cr.productUseCase.FilterCategory(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "could not retrieve products by category",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully filtered the category",
		Data:       productCategory,
		Error:      nil,
	})

}

// search for a product with given prefix
func (cr *ProductHandler) SearchProduct(c *gin.Context) {

	type SearchItems struct {
		Name string `json:"name"`
	}
	var prefix SearchItems
	if err := c.ShouldBindJSON(&prefix); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "fields provided are in wrong format",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	productDetails, err := cr.productUseCase.SearchItemBasedOnPrefix(prefix.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "could not retrieve products by prefix seae",
			Data:       nil,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully filtered the category",
		Data:       productDetails,
		Error:      nil,
	})

}
