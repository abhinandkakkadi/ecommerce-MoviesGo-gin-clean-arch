package domain

type Products struct {
	ID                  uint           `json:"id" gorm:"unique;not null"`
	MovieName           string         `json:"movie_name"`
	GenreID             uint           `json:"genre_id"`
	Genre               Genre          `json:"-" gorm:"foreignkey:GenreID;constraint:OnDelete:CASCADE"`
	DirectorID          uint           `json:"director_id"`
	Directors           Directors      `json:"-" gorm:"foreignkey:DirectorID;constraint:OnDelete:CASCADE"`
	ReleaseYear         string         `json:"release_year"`
	FormatID            uint           `json:"format_id"`
	MovieFormat         Movie_Format   `json:"-" gorm:"foreignkey:FormatID;constraint:OnDelete:CASCADE"`
	ProductsDescription string         `json:"products_discription"`
	RunTime             float64        `json:"runtime"`
	LanguageID          uint           `json:"language_id"`
	MovieLanguage       Movie_Language `json:"-" gorm:"foreignkey:LanguageID;constraint:OnDelete:CASCADE"`
	Quantity            int            `json:"quantity"`
	Price               float64        `json:"price"`
}

type Genre struct {
	ID        uint   `json:"id" gorm:"unique; not null"`
	GenreName string `json:"genre_name"`
}

type Directors struct {
	ID           uint   `json:"id" gorm:"unique; not null"`
	DirectorName string `json:"director_name"`
}

type Movie_Format struct {
	ID     uint   `json:"id" gorm:"unique; not null"`
	Format string `json:"movie_format"`
}

type Rating struct {
	ID        uint     `json:"id" gorm:"unique; not null"`
	ProductID uint     `json:"product_id"`
	Products  Products `json:"-" gorm:"foreignkey:ProductID"`
	Rating    int      `json:"rating"`
}

type Movie_Language struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Language string `json:"language"`
}
