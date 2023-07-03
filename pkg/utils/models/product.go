package models

type ProductResponse struct {
	ID                  int     `json:"id"`
	Sku                 string  `json:"sku"`
	MovieName           string  `json:"movie_name"`
	GenreName           string  `json:"genre_name"`
	Director            string  `json:"director"`
	ReleaseYear         string  `json:"release_year"`
	Format              string  `json:"format"`
	ProductsDescription string  `json:"product_description"`
	RunTime             float64 `json:"run_time"`
	Language            string  `json:"language"`
	Studio              string  `json:"studio"`
	Quantity            int     `json:"quantity"`
	Price               float64 `json:"price"`
}

type ProductsReceiver struct {
	MovieName           string  `json:"movie_name binding:required"`
	GenreID             uint    `json:"genre_id binding:required"`
	ReleaseYear         string  `json:"release_year binding:required"`
	Format              string  `json:"format binding:required"`
	Director            string  `json:"director binding:required"`
	ProductsDescription string  `json:"products_description binding:required"`
	Runtime             float64 `json:"run_time binding:required"`
	Language            string  `json:"language binding:required"`
	StudioID            uint    `json:"studio_id binding:required"`
	Quantity            int     `json:"quantity binding:required"`
	Price               float64 `json:"price binding:required"`
}

type ProductsBrief struct {
	ID            int     `json:"id"`
	MovieName     string  `json:"movie_name"`
	Sku           string  `json:"sku"`
	Genre         string  `json:"genre"`
	Language      string  `json:"language"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	ProductStatus string  `json:"product_status"`
}

type CategoryUpdate struct {
	Genre string `json:"genre binding:required"`
}

type CategoryUpdateCheck struct {
	GenreCount    int
	DirectorCount int
	FormatCount   int
	LanguageCount int
}

type UpdateProduct struct {
	Quantity  int `json:"quantity" binding:"required"`
	ProductID int `json:"product-id" binding:"required"`
}

type SearchItems struct {
	Name string `json:"name" binding:"required"`
}
