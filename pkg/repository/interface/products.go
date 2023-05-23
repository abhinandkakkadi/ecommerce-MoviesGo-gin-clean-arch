package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductRepository interface {
	ShowAllProducts() ([]domain.ProductsBrief, error)
	ShowIndividualProducts(id string) (models.IndividualProduct, error)
	// CheckIfAlreadyPresent(c context.Context,product domain.Products) (bool,error)
	UpdateQuantity(product domain.Products) error
	AddProduct(product domain.Products) error
	DeleteProduct(product_id string) error
}
