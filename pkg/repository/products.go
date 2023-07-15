package repository

import (
	"errors"
	"strconv"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

func (p *productDatabase) ShowAllProducts(page int, count int) ([]models.ProductsBrief, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productsBrief []models.ProductsBrief
	err := p.DB.Raw(`
		SELECT products.id, products.movie_name,products.sku,products.language,genres.genre_name AS genre, products.language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		 limit ? offset ?
	`, count, offset).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	return productsBrief, nil

}

// detailed product details
func (p *productDatabase) ShowIndividualProducts(sku string) (models.ProductResponse, error) {

	var product models.ProductResponse
	err := p.DB.Raw(`
	SELECT
		p.id,
		p.sku,
		p.movie_name,
		g.genre_name,
		p.director,
		p.release_year,
		p.format,
		p.products_description,
		p.run_time,
		p.language,
		s.studio,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			genres g ON p.genre_id = g.id
		JOIN
			movie_studios s ON p.studio_id = s.id 
		WHERE
			p.sku = ?
			`, sku).Scan(&product).Error

	if err != nil {
		return models.ProductResponse{}, errors.New("error retrieved record")
	}
	return product, nil

}

func (p *productDatabase) UpdateQuantity(productID int, quantity int) error {

	var currentQuantity int
	err := p.DB.Raw("select quantity from products where id = ?", productID).Scan(&currentQuantity).Error
	if err != nil {
		return err
	}

	finalQuantity := currentQuantity + quantity
	err = p.DB.Exec("update products set quantity = ? where id = ?", finalQuantity, productID).Error
	if err != nil {
		return err
	}
	return nil

}

func (p *productDatabase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {

	var id int

	sku := product.MovieName + product.Format + product.Director
	err := p.DB.Raw("INSERT INTO products (movie_name, genre_id,language,director,release_year,format,products_description,run_time,studio_id,quantity,price,sku) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?) RETURNING id", product.MovieName, product.GenreID, product.Language, product.Director, product.ReleaseYear, product.Format, product.ProductsDescription, product.Runtime, product.StudioID, product.Quantity, product.Price, sku).Scan(&id).Error
	if err != nil {
		return models.ProductResponse{}, err
	}

	var productResponse models.ProductResponse
	err = p.DB.Raw(`
	SELECT
		p.id,
		p.sku,
		p.movie_name,
		g.genre_name,
		p.director,
		p.release_year,
		p.format,
		p.products_description,
		p.run_time,
		p.language,
		s.studio,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			genres g ON p.genre_id = g.id
		JOIN
			movie_studios s ON p.studio_id = s.id 
		WHERE
			p.id = ?
			`, id).Scan(&productResponse).Error

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

func (p *productDatabase) DeleteProduct(product_id string) error {

	id, _ := strconv.Atoi(product_id)
	result := p.DB.Exec("delete from products where id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records were of that id exists")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (p *productDatabase) DoesProductExist(productID int) (bool, error) {

	var count int
	err := p.DB.Raw("select count(*) from products where id = ?", productID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (p *productDatabase) CheckValidityOfCategory(data map[string]int) error {

	for _, id := range data {
		var count int
		err := p.DB.Raw("select count(*) from genres where id = ?", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count < 1 {
			return errors.New("genre does not exist")
		}
	}
	return nil
}

func (p *productDatabase) GetProductFromCategory(id int) (models.ProductsBrief, error) {

	var product models.ProductsBrief
	err := p.DB.Raw(`
		SELECT products.id, products.movie_name,products.sku,products.language, genres.genre_name AS genre,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		 where genres.id = ?
	`, id).Scan(&product).Error

	if err != nil {
		return models.ProductsBrief{}, err
	}

	return product, nil

}

func (p *productDatabase) GetQuantityFromProductID(id int) (int, error) {

	var quantity int
	err := p.DB.Raw("select quantity from products where id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	return quantity, nil

}

func (p *productDatabase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, int, error) {

	// find length of prefix
	lengthOfPrefix := len(prefix)
	var productsBrief []models.ProductsBrief
	err := p.DB.Raw(`
		SELECT products.id, products.movie_name,products.sku,genres.genre_name AS genre,products.language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
	`).Scan(&productsBrief).Error

	if err != nil {
		return nil, 0, err
	}

	return productsBrief, lengthOfPrefix, nil

}

func (pr *productDatabase) GetGenres() ([]domain.Genre, error) {

	var genres []domain.Genre
	if err := pr.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

func (pr *productDatabase) GetPriceOfProductFromID(productID int) (float64, error) {

	var productPrice float64
	if err := pr.DB.Raw("select price from products where id = ?", productID).Scan(&productPrice).Error; err != nil {
		pr.DB.Rollback()
		return 0.0, err
	}

	return productPrice, nil

}
