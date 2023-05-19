package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/thnkrn/go-gin-clean-arch/cmd/api/docs"
	handler "github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
	middleware "github.com/thnkrn/go-gin-clean-arch/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler) *ServerHTTP {
	router := gin.New()

	// Use logger from Gin
	router.Use(gin.Logger())

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT


	//the edit i made
	router.POST("/signup", userHandler.GenerateUser)
	//ends here 
	router.POST("/login", middleware.LoginHandler)

	product := router.Group("/products")
{
	product.GET("",productHandler.ShowAllProducts)
}

	// Auth middleware
	// api := router.Group("/moviesgo", middleware.AuthorizationMiddleware)

	// api.GET("users", userHandler.FindAll)
	// api.GET("users/:id", userHandler.FindByID)
	// api.POST("users", userHandler.Save)
	// api.DELETE("users/:id", userHandler.Delete)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
