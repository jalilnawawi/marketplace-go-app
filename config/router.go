package config

import (
	"github.com/jalilnawawi/marketplace-app/controller"
	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(
	sellerController controller.SellerController,
	productController controller.ProductController,
	orderController controller.OrderController,
) *httprouter.Router {

	router := httprouter.New()

	// endpoint swagger
	router.Handler("GET", "/swagger/*any", httpSwagger.WrapHandler)

	router.POST("/api/seller", sellerController.Create)
	router.GET("/api/seller", sellerController.GetAll)
	router.GET("/api/seller/:sellerId", sellerController.GetById)
	router.PUT("/api/seller/:sellerId", sellerController.Update)
	router.DELETE("/api/seller/:sellerId", sellerController.Delete)

	router.POST("/api/product", productController.Create)
	router.GET("/api/product", productController.GetAll)
	router.GET("/api/product/:productId", productController.GetById)
	router.PUT("/api/product/:productId", productController.Update)
	router.DELETE("/api/product/:productId", productController.Delete)

	router.POST("/api/order", orderController.Create)
	router.GET("/api/order", orderController.GetAll)
	router.GET("/api/order/:orderId", orderController.GetById)
	router.PUT("/api/order/:orderId", orderController.Update)
	router.DELETE("/api/order/:orderId", orderController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
