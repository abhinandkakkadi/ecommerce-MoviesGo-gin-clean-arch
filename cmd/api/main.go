package main

import (
	"log"

	_ "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/cmd/api/docs"
	config "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
	di "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/di"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@title			moviesGo eCommerce API
//	@version		1.0
//	@description	API for ecommerce website
//	@in							header
//	@name						token
//	@host			www.abhinandkakkadi.com
//	@BasePath		/

//	@schemes	http
func main() {

	// swagger 2.0 Meta Information
	// docs.SwaggerInfo.Title = "MoviesGo - E-commerce"
	// docs.SwaggerInfo.Description = "MoviesGo - E-commerce"
	// docs.SwaggerInfo.Version = "1.0"
	// //docs.SwaggerInfo.Host = "localhost:3000"
	// docs.SwaggerInfo.BasePath = ""
	// docs.SwaggerInfo.Schemes = []string{"http"}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
