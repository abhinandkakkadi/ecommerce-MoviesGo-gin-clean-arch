package domain

type Products struct {
	ID						uint 			`json:"id" gorm:"unique;not null"`
	Movie_Name		string		`json:"movie_name"`
	GenreID				uint				`json:"genre_id"`
	Genre         Genre			`json:"-" gorm:"foreignkey:GenreID;constraint:OnDelete:CASCADE"`
	DirectorID		uint	    `json:"director_id"`
	Directors			Directors			`json:"-" gorm:"foreignkey:DirectorID;constraint:OnDelete:CASCADE"`
	Release_Year	string		`json:"release_year"`
	FormatID			uint				`json:"format_id"`
	Movie_Format	Movie_Format		`json:"-" gorm:"foreignkey:FormatID;constraint:OnDelete:CASCADE"`
	Products_Description  string	`json:"products_discription"`
	Run_time			float64				`json:"runtime"`
	LanguageID		uint				`json:"language_id"`	
	Movie_Language	Movie_Language		`json:"-" gorm:"foreignkey:LanguageID;constraint:OnDelete:CASCADE"`
	Quantity			int		`json:"quantity"`
	Price					float64		`json:"price"`
}

type ProductsBrief struct {
	Movie_Name		string		`json:"movie_name"`
	Genre					string		`json:"genre"`
	Movie_Language	string		`json:"movie_language"`
}

type Genre struct {
	ID					uint			`json:"id" gorm:"unique; not null"`
	Genre_Name				string		`json:"genre_name"`
}

type Directors struct {
	ID 					uint 			`json:"id" gorm:"unique; not null"`
	Director_Name			string		`json:"director_name"`
}

type Movie_Format struct {
	ID 				uint 			`json:"id" gorm:"unique; not null"`
	Format		string	  `json:"movie_format"`
}

type Rating struct {
	ID uint 	`json:"id" gorm:"unique; not null"`
	ProductID		uint		`json:"product_id"`
	Products 	Products	`json:"-" gorm:"foreignkey:ProductID"`
	Rating		int		`json:"rating"`
}

type Movie_Language	struct {
	ID 	uint  `json:"id" gorm:"unique; not null"`
	Language	string		`json:"language"`
}