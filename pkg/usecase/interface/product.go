package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductUseCase interface {
	
	ShowAllProducts(c context.Context) ([]domain.ProductsBrief,error)
	ShowIndividualProducts(c context.Context,id string) (models.IndividualProduct,error)
	AddProduct(c context.Context,product domain.Products) (error)
	DeleteProduct(c context.Context, product_id string) error
}