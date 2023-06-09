package models

type ProductResponse struct {
	ID                   int     `json:"id"`
	Sku                 string   `json:"sku"`
	MovieName           string   `json:"movie_name"`
	GenreName           string   `json:"genre_name"`
	Director       		  string   `json:"director"`
	ReleaseYear         string  `json:"release_year"`
	Format               string  `json:"format"`
	ProductsDescription string  `json:"product_description"`
	RunTime             float64 `json:"run_time"`
	Language      			 string  `json:"language"`
	Studio               string  `json:"studio"`
	Quantity             int     `json:"quantity"`
	Price                float64 `json:"price"`
}

type ProductsReceiver struct {
	MovieName            string    `json:"movie_name"`
	GenreID              uint      `json:"genre_id"`
	ReleaseYear         string    `json:"release_year"`
	Format             	 string    `json:"format"`
	Director       		  string   `json:"director"`
	ProductsDescription string    `json:"products_description"`
	Runtime              float64   `json:"run_time"`
	Language          	 string    	 `json:"language"`
	StudioID             uint      `json:"studio_id"`
	Quantity             int       `json:"quantity"`
	Price                float64   `json:"price"`
}

type ProductsBrief struct {
	ID             int      `json:"id"`
	MovieName      string   `json:"movie_name"`
	Sku                 string   `json:"sku"`
	Genre          string   `json:"genre"`
	Language 			 string   `json:"language"`
	Price          float64  `json:"price"`
	Quantity       int      `json:"quantity"`
	ProductStatus string   `json:"product_status"`
}

type CategoryUpdate struct {
	Genre    string `json:"genre"`
	Director string `json:"director"`
	Format   string `json:"format"`
	Language string `json:"language"`
}

type CategoryUpdateCheck struct {
	GenreCount    int
	DirectorCount int
	FormatCount   int
	LanguageCount int
}

type UpdateProduct struct {
	Quantity  int `json:"quantity"`
	ProductID int `json:"product-id"`
}

type SearchItems struct {
	Name string `json:"name"`
}
