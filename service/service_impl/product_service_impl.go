package service_impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/domain"
	"github.com/jalilnawawi/marketplace-app/model/dto/product_dto"
	"github.com/jalilnawawi/marketplace-app/repository"
	"github.com/jalilnawawi/marketplace-app/service"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	SellerRepository  repository.SellerRepository
	DB                *sql.DB
}

func NewProductServiceImpl(productRepository repository.ProductRepository, sellerRepository repository.SellerRepository, db *sql.DB) service.ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		SellerRepository:  sellerRepository,
		DB:                db,
	}
}

func (p *ProductServiceImpl) FindAll(ctx context.Context) []product_dto.ProductResponse {
	db, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	products := p.ProductRepository.FindAll(ctx, db)

	return product_dto.ToProductResponses(products)
}

func (p *ProductServiceImpl) FindById(ctx context.Context, id int64) product_dto.ProductResponse {
	db, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	product, err := p.ProductRepository.FindByID(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return product_dto.ToProductResponse(product)
}

func (p *ProductServiceImpl) Create(ctx context.Context, req product_dto.CreateProductRequest) (product_dto.ProductResponse, error) {
	db, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	_, err = p.SellerRepository.FindByID(ctx, db, req.SellerId)
	if err != nil {
		defer helper.CommitOrRollback(db)

		return product_dto.ProductResponse{}, fmt.Errorf("failed to create product: seller with id %d not found", req.SellerId)
	}

	product := domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		SellerId:    req.SellerId,
	}

	created, err := p.ProductRepository.Create(ctx, db, product)
	if err != nil {
		return product_dto.ProductResponse{}, err
	}

	return product_dto.ToProductResponse(created), nil
}

func (p *ProductServiceImpl) Update(ctx context.Context, req product_dto.UpdateProductRequest, id int64) product_dto.ProductResponse {
	db, err := p.DB.Begin()
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
			panic(r)
		}
	}()

	product, err := p.ProductRepository.FindByID(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Category = req.Category

	product = p.ProductRepository.Update(ctx, db, product)

	return product_dto.ToProductResponse(product)
}

func (p *ProductServiceImpl) Delete(ctx context.Context, id int64) {
	db, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	p.ProductRepository.Delete(ctx, db, id)
}

func (p *ProductServiceImpl) FindPriceByProductId(ctx context.Context, id int64) (float64, error) {
	db, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	price, err := p.ProductRepository.FindPriceByProductId(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return price, nil
}
