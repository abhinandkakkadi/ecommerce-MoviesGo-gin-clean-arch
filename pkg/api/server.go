package http

import (
	handler "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, productHandler *handler.ProductHandler, otpHandler *handler.OtpHandler, adminHandler *handler.AdminHandler,cartHandler *handler.CartHandler) *ServerHTTP {
	router := gin.New()

	// Use logger from Gin
	router.Use(gin.Logger())

	// USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)
	router.POST("/send-otp", otpHandler.SendOTP)
	router.POST("/verify-otp", otpHandler.VerifyOTP)
	product := router.Group("/products")
	product.GET("", productHandler.ShowAllProducts)
	product.GET("/page/:page", productHandler.ShowAllProducts)
	product.GET("/:id", productHandler.ShowIndividualProducts)

	router.Use(middleware.AuthMiddleware())
	{	
		user := router.Group("/cart") 
		{
			user.POST("/addtocart/:id",cartHandler.AddToCart)
		}
	
	}

	

	// ADMIN SIDE
	router.POST("/adminsignup", adminHandler.SignUpHandler)
	router.POST("/adminlogin", adminHandler.LoginHandler)
	api := router.Group("/admin_panel", middleware.AuthorizationMiddleware)
	api.GET("users", adminHandler.GetUsers)
	api.GET("genres", adminHandler.GetGenres)
	api.POST("genres/add_genre", adminHandler.AddCategory)
	api.GET("genres/delete_genre/:id", adminHandler.DeleteGenre)
	api.POST("/products/add_product", productHandler.AddProduct)
	api.DELETE("/products/delete_product/:id", productHandler.DeleteProduct)
	api.GET("/users/block-users/:id", adminHandler.BlockUser)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
