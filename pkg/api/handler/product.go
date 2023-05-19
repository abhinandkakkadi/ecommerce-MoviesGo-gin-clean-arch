package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
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

	products,err  := cr.productUseCase.ShowAllProducts(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
	}

	c.JSON(http.StatusCreated,products)
}
