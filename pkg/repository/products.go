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

func (c *productDatabase) ShowAllProducts(page int) ([]models.ProductsBrief, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2
	var productsBrief []models.ProductsBrief
	err := c.DB.Limit(1).Raw(`
		SELECT products.id, products.movie_name, genres.genre_name AS genre, movie_languages.language AS movie_language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		JOIN movie_languages ON products.language_id = movie_languages.id limit ? offset ?
	`, 2, offset).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	return productsBrief, nil

}

// detailed product details
func (c *productDatabase) ShowIndividualProducts(id string) (models.ProductResponse, error) {

	var product models.ProductResponse
	product_id, _ := strconv.Atoi(id)

	err := c.DB.Raw(`
	SELECT
		p.id,
		p.movie_name,
		g.genre_name,
		d.director_name,
		p.release_year,
		mf.format,
		p.products_description,
		p.run_time,
		ml.language AS movie_language,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			genres g ON p.genre_id = g.id
		JOIN
			directors d ON p.director_id = d.id
		JOIN
			movie_formats mf ON p.format_id = mf.id
		JOIN
			movie_languages ml ON p.language_id = ml.id 
		WHERE
			p.id = ?
			`, product_id).Scan(&product).Error

	if err != nil {
		return models.ProductResponse{}, errors.New("error entering record")
	}
	return product, nil

}

func (cr *productDatabase) UpdateQuantity(productID int, quantity int) error {

	var currentQuantity int
	err := cr.DB.Raw("select quantity from products where id = ?", productID).Scan(&currentQuantity).Error
	if err != nil {
		return err
	}
	finalQuantity := currentQuantity + quantity
	err = cr.DB.Exec("update products set quantity = ? where id = ?", finalQuantity, productID).Error
	if err != nil {
		return err
	}
	return nil

}

func (cr *productDatabase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {

	var id int
	err := cr.DB.Raw("INSERT INTO products (movie_name, genre_id, director_id, release_year,format_id,products_description,run_time,language_id,quantity,price) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id", product.Movie_Name, product.GenreID, product.DirectorID, product.Release_Year, product.FormatID, product.Products_Description, product.Run_time, product.LanguageID, product.Quantity, product.Price).Scan(&id).Error
	if err != nil {
		return models.ProductResponse{}, err
	}

	var productResponse models.ProductResponse
	err = cr.DB.Raw(`
	SELECT
		p.id,
		p.movie_name,
		g.genre_name,
		d.director_name,
		p.release_year,
		mf.format,
		p.products_description,
		p.run_time,
		ml.language AS movie_language,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			genres g ON p.genre_id = g.id
		JOIN
			directors d ON p.director_id = d.id
		JOIN
			movie_formats mf ON p.format_id = mf.id
		JOIN
			movie_languages ml ON p.language_id = ml.id 
		WHERE
			p.id = ?
			`, id).Scan(&productResponse).Error

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

func (cr *productDatabase) DeleteProduct(product_id string) error {

	id, _ := strconv.Atoi(product_id)
	fmt.Println(id)
	result := cr.DB.Exec("delete from products where id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records were of that id exists")
	}

	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (cr *productDatabase) DoesProductExist(productID int) (bool, error) {

	var count int
	err := cr.DB.Raw("select count(*) from products where id = ?", productID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (cr *productDatabase) CheckValidityOfCategory(data map[string]int) error {

	for _, id := range data {
		var count int
		err := cr.DB.Raw("select count(*) from genres where id = ?", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count < 1 {
			return errors.New("one or some of the category does not exist")
		}
	}
	return nil
}

func (cr *productDatabase) GetProductFromCategory(data map[string]int) ([]models.ProductsBrief, error) {

	var productFromCategory []models.ProductsBrief
	for _, id := range data {
		var product models.ProductsBrief
		err := cr.DB.Raw(`
		SELECT products.id, products.movie_name, genres.genre_name AS genre, movie_languages.language AS movie_language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		JOIN movie_languages ON products.language_id = movie_languages.id where genres.id = ?
	`, id).Scan(&product).Error

		if err != nil {
			return nil, err
		}
		fmt.Println("individual product details : ", product)

		// if a product exist for that genre. Then only append it
		if product.ID != 0 {
			productFromCategory = append(productFromCategory, product)
		}

	}
	fmt.Println("complete product details")

	return productFromCategory, nil
}

func (cr *productDatabase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductsBrief, error) {

	// find length of prefix
	lengthOfPrefix := len(prefix)
	var productsBrief []models.ProductsBrief
	err := cr.DB.Raw(`
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
		length := len(p.Movie_Name)
		if length >= lengthOfPrefix {
			// slice the movie name to length of prefix
			moviePrefix := p.Movie_Name[:lengthOfPrefix]
			// if they are equal - append that movie to the returning slice
			if moviePrefix == prefix {
				fmt.Println("got the condition right")
				filteredProductBrief = append(filteredProductBrief, p)
			}
		}
	}

	return filteredProductBrief, nil

}
