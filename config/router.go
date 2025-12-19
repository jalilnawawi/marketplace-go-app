package config

import (
	"github.com/jalilnawawi/marketplace-app/controller"
	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(sellerController controller.SellerController) *httprouter.Router {

	router := httprouter.New()

	// endpoint swagger
	router.Handler("GET", "/swagger/*any", httpSwagger.WrapHandler)

	router.POST("/api/seller", sellerController.Create)
	router.GET("/api/seller", sellerController.GetAll)
	router.GET("/api/seller/:sellerId", sellerController.GetById)
	router.PUT("/api/seller/:sellerId", sellerController.Update)
	router.DELETE("/api/seller/:sellerId", sellerController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
