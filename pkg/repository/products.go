package repository

import (
	"errors"
	"fmt"
	"strconv"

	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
) 


type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB,}
}

func (c *productDatabase) ShowAllProducts()([]domain.ProductsBrief, error) {
		
	var productsBrief []domain.ProductsBrief
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

func (c *productDatabase) ShowIndividualProducts(id string)(models.IndividualProduct,error) {

	var product models.IndividualProduct
	product_id,_ := strconv.Atoi(id)

	err := c.DB.Raw(`
	SELECT
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
		return models.IndividualProduct{},errors.New("error entering record")
	}
	return product,nil

}


func (cr *productDatabase) UpdateQuantity(product domain.Products) error {

	var intialQuantity int
	err := cr.DB.Raw("select quantity from products where movie_name = ? and format_id = ?",product.Movie_Name,product.FormatID).Scan(&intialQuantity).Error

	if err != nil {
		return err
	}
	
	finalQuantity := intialQuantity + product.Quantity
	err = cr.DB.Raw("update from products set quantity = ? ",finalQuantity).Error

	if err != nil {
		return err
	}

	return nil

}

func (cr *productDatabase) AddProduct(product domain.Products) error {

	err := cr.DB.Create(&product).Error
	if err != nil {
		return err
	}

	return nil

}

func (cr *productDatabase) DeleteProduct(product_id string) error {
	
	id,_ := strconv.Atoi(product_id)
	fmt.Println(id)
	result := cr.DB.Exec("delete from products where id = ?",id)

	if result.RowsAffected < 1 {
		return errors.New("no records were of that id exists")
	}

	fmt.Println(result.Error)
    if result.Error != nil {
        return result.Error
    }

    return nil

}