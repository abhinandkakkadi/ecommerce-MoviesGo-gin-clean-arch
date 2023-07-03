package routes

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {

	// USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)

	router.POST("/send-otp", otpHandler.SendOTP)
	router.POST("/verify-otp", otpHandler.VerifyOTP)

	forgotPassword := router.Group("/forgot-password")
	{
		forgotPassword.POST("", otpHandler.SendOTPtoReset)
		forgotPassword.POST("/verify-otp", otpHandler.VerifyOTPToReset)

		forgotPassword.Use(middleware.AuthMiddlewareReset())
		forgotPassword.PUT("/reset", userHandler.ResetPassword)

	}

	product := router.Group("/products")
	{
		product.GET("", productHandler.ShowAllProducts)
		product.GET("/page/:page", productHandler.ShowAllProducts)
		product.GET("/:id", productHandler.ShowIndividualProducts)
		product.POST("/filter", productHandler.FilterCategory)
		product.POST("/search", productHandler.SearchProduct)
		product.GET("/genres", productHandler.GetGenresToUser)
	}

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

			users.GET("/delivered/:order_id", orderHandler.OrderDelivered)
			users.GET("/return/:order_id", orderHandler.ReturnOrder)

		}

		wishList := router.Group("wishlist")
		{
			wishList.GET("", userHandler.GetWishList)
			wishList.GET("/add/:id", userHandler.AddToWishList)
			wishList.DELETE("/remove/:id", userHandler.RemoveFromWishList)

		}

		coupon := router.Group("/coupon")
		{
			coupon.POST("/apply", cartHandler.ApplyCoupon)

		}

		router.GET("/referral/apply", userHandler.ApplyReferral)

		router.GET("/checkout", userHandler.CheckOut)
		router.POST("/order", orderHandler.OrderItemsFromCart)

		router.GET("/payment/:id", paymentHandler.MakePaymentRazorPay)
		router.GET("/payment-success", paymentHandler.VerifyPayment)

	}

}
