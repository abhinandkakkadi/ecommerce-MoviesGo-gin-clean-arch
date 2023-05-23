package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductUseCase interface {
	
	ShowAllProducts() ([]domain.ProductsBrief,error)
	ShowIndividualProducts(id string) (models.IndividualProduct,error)
	AddProduct(product domain.Products) (error)
	DeleteProduct(product_id string) error
}