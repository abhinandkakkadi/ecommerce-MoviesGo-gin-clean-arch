package repository

import (
	"errors"
	"fmt"
	"strconv"

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
		SELECT products.id, products.movie_name,products.sku,genres.genre_name AS genre, products.language AS movie_language,products.price,products.quantity
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
func (p *productDatabase) ShowIndividualProducts(id string) (models.ProductResponse, error) {

	var product models.ProductResponse
	product_id, _ := strconv.Atoi(id)

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
			p.id = ?
			`, product_id).Scan(&product).Error

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
	fmt.Println(id)
	result := p.DB.Exec("delete from products where id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records were of that id exists")
	}

	fmt.Println(result.Error)
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
			return errors.New("one or some of the category does not exist")
		}
	}
	return nil
}

func (p *productDatabase) GetProductFromCategory(data map[string]int) ([]models.ProductsBrief, error) {

	var productFromCategory []models.ProductsBrief
	for _, id := range data {
		var product models.ProductsBrief
		err := p.DB.Raw(`
		SELECT products.id, products.movie_name, genres.genre_name AS genre, movie_languages.language AS movie_language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		JOIN movie_languages ON products.language_id = movie_languages.id where genres.id = ?
	`, id).Scan(&product).Error

		if err != nil {
			return nil, err
		}
		fmt.Println("individual product details : ", product)
		var quantity int
		err = p.DB.Raw("select quantity from products where id = ?", product.ID).Scan(&quantity).Error
		if err != nil {
			return nil, err
		}

		if quantity == 0 {
			product.ProductStatus = "out of stock"
		} else {
			product.ProductStatus = "in stock"
		}
		// if a product exist for that genre. Then only append it
		if product.ID != 0 {
			productFromCategory = append(productFromCategory, product)
		}

	}
	fmt.Println("complete product details")

	return productFromCategory, nil
}

func (p *productDatabase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, error) {

	// find length of prefix
	lengthOfPrefix := len(prefix)
	var productsBrief []models.ProductsBrief
	err := p.DB.Raw(`
		SELECT products.id, products.movie_name, genres.genre_name AS genre, movie_languages.language AS movie_language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		JOIN movie_languages ON products.language_id = movie_languages.id
	`).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}
	// Create a slice to add the products which have the given prefix
	var filteredProductBrief []models.ProductsBrief
	for _, p := range productsBrief {
		// If length of the movie name is greater than prefix - continue the logic
		length := len(p.MovieName)
		if length >= lengthOfPrefix {
			// slice the movie name to length of prefix
			moviePrefix := p.MovieName[:lengthOfPrefix]
			// if they are equal - append that movie to the returning slice
			if moviePrefix == prefix {
				fmt.Println("got the condition right")
				filteredProductBrief = append(filteredProductBrief, p)
			}
		}
	}

	return filteredProductBrief, nil

}
