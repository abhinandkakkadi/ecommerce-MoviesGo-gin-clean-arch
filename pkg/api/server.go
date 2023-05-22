package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/cmd/api/docs"
	handler "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler,otpHandler *handler.OtpHandler,adminHandler *handler.AdminHandler) *ServerHTTP {
	router := gin.New()

	// Use logger from Gin
	router.Use(gin.Logger())

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT

// USER SIDE
	//sign up functionality for user
	router.POST("/signup", userHandler.GenerateUser)
	//login functionality for user
	router.POST("/login", userHandler.LoginHandler)
	//send and verify otp along with user autherisation
	router.POST("/send-otp",otpHandler.SendOTP)
	router.POST("/verify-otp",otpHandler.VerifyOTP)

	product := router.Group("/products")
{
	product.GET("",productHandler.ShowAllProducts)
	product.GET("/:id",productHandler.ShowIndividualProducts)
}

// ADMIN SIDE
	// admin login
	router.POST("/adminsignup",adminHandler.SignupHandler)
	router.POST("/adminlogin",adminHandler.LoginHandler)

	// admin signup
	// router.POST("/adminsignup",adminHandler.SignupHandler)



	// Auth middleware
	api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)

	api.GET("users", adminHandler.GetUsers)
	api.GET("genres",adminHandler.GetGenres)
	api.POST("genres/add_genre",adminHandler.AddGenre)
	api.GET("genres/delete_genre/:id",adminHandler.DeleteGenre)
	api.POST("/products/add_product",productHandler.AddProduct)
	api.DELETE("/products/delete_product/:id",productHandler.DeleteProduct)
	api.GET("/users/block-users/:id",adminHandler.BlockUser)
	// api.GET("/users/unblock-users/:id",productHandler.DeleteProduct)
	
	
	
	
	// api.GET("users/:id", userHandler.FindByID)
	// api.POST("users", userHandler.Save)
	// api.DELETE("users/:id", userHandler.Delete)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
