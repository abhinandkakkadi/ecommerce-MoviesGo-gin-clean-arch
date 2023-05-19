package repository

import (
	"context"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
) 


type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB,
	}
}

func (c *productDatabase) ShowAllProducts(ctx context.Context)([]domain.Products, error) {

	var products []domain.Products
	err := c.DB.Find(&products).Error

	return products, err

}