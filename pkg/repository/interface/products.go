package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)





type ProductRepository interface {
	ShowAllProducts(ctx context.Context) ([]domain.ProductsBrief, error)
	ShowIndividualProducts(ctx context.Context,id string)(models.IndividualProduct,error)
	CheckIfAlreadyPresent(c context.Context,product domain.Products) (bool,error)
	UdateQuantity(c context.Context,product domain.Products) error
	AddProduct(c context.Context,product domain.Products) error
	DeleteProduct(c context.Context,product_id string) error
}