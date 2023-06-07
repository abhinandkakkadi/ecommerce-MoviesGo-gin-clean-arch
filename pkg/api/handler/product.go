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
func (pr *ProductHandler) ShowAllProducts(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := pr.productUseCase.ShowAllProducts(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Products To Admin
// @Description Retrieve products with pagination to Admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products [get]
func (pr *ProductHandler) SeeAllProductToAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)

	count, _ := strconv.Atoi(c.Query("count"))

	products, err := pr.productUseCase.ShowAllProducts(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products to admin side", products, nil)
	c.JSON(http.StatusOK, successRes)

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
func (pr *ProductHandler) ShowIndividualProducts(c *gin.Context) {

	id := c.Param("id")
	product, err := pr.productUseCase.ShowIndividualProducts(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Product details retrieved successfully", product, nil)
	c.JSON(http.StatusOK, successRes)

}

func (pr *ProductHandler) AddProduct(c *gin.Context) {

	var product models.ProductsReceiver
	if err := c.BindJSON(&product); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	productResponse, err := pr.productUseCase.AddProduct(product)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added products", productResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add Products
// @Description Add a new product from the admin side
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/delete-product/{id} [delete]
func (pr *ProductHandler) DeleteProduct(c *gin.Context) {

	product_id := c.Param("id")
	err := pr.productUseCase.DeleteProduct(product_id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the item", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Update Products
// @Description Update quantity of already existing product
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param p body models.UpdateProduct true "Product details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/update-product/ [post]
func (pr *ProductHandler) UpdateProduct(c *gin.Context) {

	var p models.UpdateProduct

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := pr.productUseCase.UpdateProduct(p.ProductID, p.Quantity)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Show Products of specified category
// @Description Show Products of specified category
// @Tags Users
// @Accept json
// @Produce json
// @Param data body map[string]int true "Category IDs and quantities"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/filter [post]
func (pr *ProductHandler) FilterCategory(c *gin.Context) {

	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	productCategory, err := pr.productUseCase.FilterCategory(data)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve products by category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully filtered the category", productCategory, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Show Products of specified category
// @Description Show Products of specified category
// @Tags Users
// @Accept json
// @Produce json
// @Param prefix body models.SearchItems true "Name prefix to search"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/search [post]
func (pr *ProductHandler) SearchProduct(c *gin.Context) {

	var prefix models.SearchItems
	if err := c.ShouldBindJSON(&prefix); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	productDetails, err := pr.productUseCase.SearchItemBasedOnPrefix(prefix.Name)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve products by prefix search", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully filtered the category", productDetails, nil)
	c.JSON(http.StatusOK, successRes)

}
