package domain

type Products struct {
	ID						uint 			`json:"id" gorm:"unique;not null"`
	Movie_Name		string		`json:"movie_name"`
	GenreID				uint				`json:"genre_id"`
	Genre         Genre			`json:"-" gorm:"foreignkey:GenreID"`
	DirectorID		uint	    `json:"director_id"`
	Directors			Directors			`json:"-" gorm:"foreignkey:DirectorID"`
	Release_Year	string		`json:"release_year"`
	FormatID			uint				`json:"format_id"`
	Movie_Format	Movie_Format		`json:"-" gorm:"foreignkey:FormatID"`
	Products_Description  string	`json:"products_discription"`
	Run_time			float64				`json:"runtime"`
	LanguageID		uint				`json:"language_id"`	
	Movie_Language	Movie_Language		`json:"-" gorm:"foreignkey:LanguageID"`
	RatingID		uint		`json:"rating_id"`
	Rating			Rating		`json:"-" gorm:"foreignkey:RatingID"`
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
	Movie_Format	uint	`json:"movie_format"`
	Quantity			int		`json:"quantity"`
	Price					float64		`json:"price"`
}

type Rating struct {
	ID uint 	`json:"id" gorm:"unique; not null"`
	Rating		int		`json:"rating"`
}

type Movie_Language	struct {
	ID 	uint  `json:"id" gorm:"unique; not null"`
	Language	string		`json:"language"`
}