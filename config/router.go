package config

import (
	"github.com/jalilnawawi/marketplace-app/controller"
	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(sellerController controller.SellerController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/seller", sellerController.Create)
	router.GET("/api/seller", sellerController.GetAll)
	router.GET("/api/seller/:id", sellerController.GetById)
	router.PUT("/api/seller/:id", sellerController.Update)
	router.DELETE("/api/seller/:id", sellerController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
