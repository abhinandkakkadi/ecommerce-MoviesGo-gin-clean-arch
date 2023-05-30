package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductUseCase interface {
	ShowAllProducts(page int) ([]models.ProductsBrief, error)
	ShowIndividualProducts(id string) (models.ProductResponse, error)
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
	DeleteProduct(product_id string) error
	UpdateProduct(productID int,quantity int) error
	
}
