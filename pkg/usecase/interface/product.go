package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductUseCase interface {
	ShowAllProducts(page int, count int) ([]models.ProductOfferBriefResponse, error)
	ShowAllProductsToAdmin(page int, count int) ([]models.ProductsBrief, error)
	ShowIndividualProducts(sku string) (models.ProductOfferLongResponse, error)
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
	DeleteProduct(product_id string) error
	UpdateProduct(productID int, quantity int) error
	FilterCategory(data map[string]int) ([]models.ProductsBrief, error)
	SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, error)
	GetGenres() ([]domain.Genre, error)
}
