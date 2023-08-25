package main

import (
	"log"
	"os"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/cmd/api/docs"
	_ "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/cmd/api/docs"
	config "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
	di "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/di"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func main() {

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "MoviesGo - E-commerce"
	docs.SwaggerInfo.Description = "MoviesGo - E-commerce"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	server, productRepo, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start(productRepo, infoLog, errorLog)
	}

}
