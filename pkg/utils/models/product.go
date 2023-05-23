package models

type IndividualProduct struct {
	Movie_Name  string		`json:"movie_name"`
	Genre_Name  string		`json:"genre_name"`
	Director_Name  string		`json:"director_name"`
	Release_Year  string		`json:"release_year"`
	Format  string					`json:"format"`
	Products_Description	string	`json:"product_description"`
	Run_Time	float64				`json:"run_time"`
	Movie_Language	 string				`json:"movie_language"`
	Quantity	 int				`json:"quantity"`
	Price	 		 float64				`json:"price"`
}


