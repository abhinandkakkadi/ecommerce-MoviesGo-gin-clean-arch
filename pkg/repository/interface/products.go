package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductRepository interface {
	ShowAllProducts(page int) ([]models.ProductsBrief, error)
	ShowIndividualProducts(id string) (models.ProductResponse, error)
	// CheckIfAlreadyPresent(c context.Context,product domain.Products) (bool,error)
	UpdateQuantity(productID int,quantity int) error
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
	DeleteProduct(product_id string) error
}
