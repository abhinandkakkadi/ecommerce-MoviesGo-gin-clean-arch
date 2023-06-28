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

// @Summary Get Products Details to users
// @Description Retrieve all product Details with pagination to users
// @Tags User Product
// @Accept json
// @Produce json
// @Param page path string true "Page number"
// @Param count query string true "Page Count"
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

// @Summary Get Product Details To Admin
// @Description Retrieve product Details with pagination to Admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Param count query string true "Products Count Per Page"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products [get]
func (pr *ProductHandler) SeeAllProductToAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)

	count, _ := strconv.Atoi(c.Query("count"))

	products, err := pr.productUseCase.ShowAllProductsToAdmin(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products to admin side", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Individual Product Details
// @Description Get Individual Detailed product details to user side
// @Tags User Product
// @Accept json
// @Produce json
// @Param id path string true "sku"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/{id} [get]
func (pr *ProductHandler) ShowIndividualProducts(c *gin.Context) {

	sku := c.Param("id")
	product, err := pr.productUseCase.ShowIndividualProducts(sku)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Product details retrieved successfully", product, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add Products
// @Description Add product from admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param product body models.ProductsReceiver true "Product details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/add-product/ [post]
func (pr *ProductHandler) AddProduct(c *gin.Context) {

	var product models.ProductsReceiver

	if err := c.ShouldBindJSON(&product); err != nil {
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

// @Summary Delete product
// @Description Delete a product from the admin side
// @Tags Admin Product Management
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

// @Summary Update Products quantity
// @Description Update quantity of already existing product
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param productUpdate body models.UpdateProduct true "Product details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/update-product/ [post]
func (pr *ProductHandler) UpdateProduct(c *gin.Context) {

	var productUpdate models.UpdateProduct

	if err := c.ShouldBindJSON(&productUpdate); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := pr.productUseCase.UpdateProduct(productUpdate.ProductID, productUpdate.Quantity)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated product quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Show Products of specified category
// @Description Show all the Products belonging to a specified category
// @Tags User Product
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

// @Summary Search Products
// @Description Show Products by it's prefix (case insensitive)
// @Tags User Product
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

// @Summary Get genre to user side
// @Description Display genre details on the user side
// @Tags User Product
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/genres [get]
func (ad *ProductHandler) GetGenresToUser(c *gin.Context) {

	genres, err := ad.productUseCase.GetGenres()

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the genres", genres, nil)
	c.JSON(http.StatusOK, successRes)
}
