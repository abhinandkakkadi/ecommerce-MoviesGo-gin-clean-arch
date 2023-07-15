package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductRepository interface {
	ShowAllProducts(page int, count int) ([]models.ProductsBrief, error)
	ShowIndividualProducts(sku string) (models.ProductResponse, error)
	UpdateQuantity(productID int, quantity int) error
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
	DeleteProduct(product_id string) error
	DoesProductExist(productID int) (bool, error)
	CheckValidityOfCategory(data map[string]int) error
	GetProductFromCategory(id int) (models.ProductsBrief, error)
	SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, int, error)
	GetGenres() ([]domain.Genre, error)
	GetQuantityFromProductID(id int) (int, error)
	GetPriceOfProductFromID(productID int) (float64, error)
}
