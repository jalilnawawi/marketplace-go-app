package main

// @title		Marketplace API
// @version		1.0
// @description	Simple CRUD API using Golang
// @host		localhost:8181
// @Basepath	/api

import (
	"log"
	"net/http"

	"github.com/jalilnawawi/marketplace-app/config"
	"github.com/jalilnawawi/marketplace-app/controller/controller_impl"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/repository/repository_impl"
	"github.com/jalilnawawi/marketplace-app/service/service_impl"
	"github.com/joho/godotenv"

	_ "github.com/jalilnawawi/marketplace-app/docs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	db := config.NewPostgresDB(cfg.DatabaseDSN)
	defer db.Close()

	sellerRepository := repository_impl.NewSellerRepository()
	sellerService := service_impl.NewSellerService(sellerRepository, db)
	sellerController := controller_impl.NewSellerController(sellerService)
	router := config.NewRouter(sellerController)

	server := http.Server{
		Addr:    cfg.HTTPPort,
		Handler: router,
	}
	log.Printf("🚀 Server starting on http://localhost%s", server.Addr)
	log.Printf("🚀 Swagger UI available on http://localhost%s/swagger/index.html", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
