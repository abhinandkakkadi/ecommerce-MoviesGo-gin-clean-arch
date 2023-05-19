package domain

type Products struct {
	ID						uint 			`json:"id" gorm:"unique;not null"`
	Movie_Name		string		`json:"movie_name"`
	GenreID				uint				`json:"genre_id"`
	Genre         Genre			`json:"-" gorm:"foreignkey:GenreID"`
	
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