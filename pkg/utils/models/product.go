package models

type ProductResponse struct {
	ID                   int     `json:"id"`
	Movie_Name           string  `json:"movie_name"`
	Genre_Name           string  `json:"genre_name"`
	Director_Name        string  `json:"director_name"`
	Release_Year         string  `json:"release_year"`
	Format               string  `json:"format"`
	Products_Description string  `json:"product_description"`
	Run_Time             float64 `json:"run_time"`
	Movie_Language       string  `json:"movie_language"`
	Quantity             int     `json:"quantity"`
	Price                float64 `json:"price"`
}

type ProductsReceiver struct {
	Movie_Name           string  `json:"movie_name"`
	GenreID              uint    `json:"genre_id"`
	DirectorID           uint    `json:"director_id"`
	Release_Year         string  `json:"release_year"`
	FormatID             uint    `json:"format_id"`
	Products_Description string  `json:"products_discription"`
	Run_time             float64 `json:"runtime"`
	LanguageID           uint    `json:"language_id"`
	Quantity             int     `json:"quantity"`
	Price                float64 `json:"price"`
}

type ProductsBrief struct {
	ID             int    `json:"id"`
	Movie_Name     string `json:"movie_name"`
	Genre          string `json:"genre"`
	Movie_Language string `json:"movie_language"`
}
