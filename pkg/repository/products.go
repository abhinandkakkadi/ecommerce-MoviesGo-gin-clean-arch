package repository

import (
	"context"
	"errors"

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

func (c *productDatabase) ShowAllProducts(ctx context.Context)([]domain.ProductsBrief, error) {
		var productsBrief []domain.ProductsBrief
	// Perform the raw SQL query
	
	err := c.DB.Raw(`
		SELECT products.movie_name, genres.genre_name AS genre, movie_languages.language AS movie_language
		FROM products
		JOIN genres ON products.genre_id = genres.id
		JOIN movie_languages ON products.language_id = movie_languages.id
	`).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	return productsBrief, nil

}

func (c *productDatabase) ShowIndividualProducts(ctx context.Context,id string)(domain.Products,error) {

	var product domain.Products

	err := c.DB.Raw(`
	SELECT *
	FROM products
	WHERE id = ?
`, id).Scan(&product).Error


	if err != nil {
		return domain.Products{},errors.New("Error entering record")
	}

	return product,nil

}