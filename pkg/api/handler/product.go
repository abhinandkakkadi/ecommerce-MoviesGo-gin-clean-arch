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

// @Summary Get Products to users
// @Description Retrieve products with pagination
// @Tags Users
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

// @Summary Get Products To Admin
// @Description Retrieve products with pagination to Admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Param page path string true "Page number"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products [get]
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

// @Summary Get Full Product Details
// @Description Retrieve Complete products details at user side
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/{id} [get]
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

// @Summary Add Products
// @Description Add a new product from the admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/delete-product/{id} [delete]
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

// @Summary Update Products
// @Description Update quantity of already existing product
// @Tags Admin
// @Accept json
// @Produce json
// @Param p body models.UpdateProduct true "Product details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/update-product/ [post]
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {

	var p models.UpdateProduct

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

// @Summary Show Products of specified category
// @Description Show Products of specified category
// @Tags Users
// @Accept json
// @Produce json
// @Param data body map[string]int true "Category IDs and quantities"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /filer [post]
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

// @Summary Show Products of specified category
// @Description Show Products of specified category
// @Tags Users
// @Accept json
// @Produce json
// @Param prefix body models.SearchItems true "Name prefix to search"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /search [post]
func (cr *ProductHandler) SearchProduct(c *gin.Context) {

	var prefix models.SearchItems
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
