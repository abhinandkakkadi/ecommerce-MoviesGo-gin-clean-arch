package routes

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler) {

	// USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)
	router.POST("/send-otp", otpHandler.SendOTP)
	router.POST("/verify-otp", otpHandler.VerifyOTP)
	product := router.Group("/products")
	product.GET("", productHandler.ShowAllProducts)
	product.GET("/page/:page", productHandler.ShowAllProducts)
	product.GET("/:id", productHandler.ShowIndividualProducts)

	router.POST("/filter", productHandler.FilterCategory)
	router.POST("/search", productHandler.SearchProduct)

	router.Use(middleware.AuthMiddleware())
	{
		cart := router.Group("/cart")
		{
			cart.POST("/addtocart/:id", cartHandler.AddToCart)
			cart.DELETE("/removefromcart/:id", cartHandler.RemoveFromCart)
			cart.GET("", cartHandler.DisplayCart)
			cart.DELETE("", cartHandler.EmptyCart)
		}

		address := router.Group("/address")
		{
			address.POST("", userHandler.AddAddress)
			address.PUT("/:id", userHandler.UpdateAddress)
		}

		users := router.Group("/users")
		{

			users.GET("", userHandler.UserDetails)
			users.PUT("", userHandler.UpdateUserDetails)
			users.GET("/address", userHandler.GetAllAddress)
			users.POST("/address", userHandler.AddAddress)
			orders := users.Group("orders")
			{
				orders.GET("", orderHandler.GetOrderDetails)
				orders.GET("/:page", orderHandler.GetOrderDetails)
			}

			users.PUT("/cancel-order/:id", orderHandler.CancelOrder)
			users.PUT("/update-password", userHandler.UpdatePassword)

		}

		whishlist := router.Group("wishlist")
		{
			whishlist.GET("", userHandler.GetWishList)
			whishlist.GET("/add/:id", userHandler.AddToWishList)
			whishlist.DELETE("/remove/:id", userHandler.RemoveFromWishList)
		}

		coupon := router.Group("/coupon")
		{
			coupon.POST("/add", cartHandler.AddCoupon)
		}

		router.GET("/checkout", userHandler.CheckOut)

		router.POST("/order", orderHandler.OrderItemsFromCart)

	}

}
