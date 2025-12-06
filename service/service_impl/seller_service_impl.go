package service_impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/domain"
	"github.com/jalilnawawi/marketplace-app/model/dto/seller_dto"
	"github.com/jalilnawawi/marketplace-app/repository"
	"github.com/jalilnawawi/marketplace-app/service"
)

type SellerServiceImpl struct {
	SellerRepository repository.SellerRepository
	DB               *sql.DB
}

func NewSellerService(sellerRepository repository.SellerRepository, db *sql.DB) service.SellerService {
	return &SellerServiceImpl{
		SellerRepository: sellerRepository,
		DB:               db,
	}
}

func (service *SellerServiceImpl) FindAll(ctx context.Context) []seller_dto.SellerResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	sellers := service.SellerRepository.FindAll(ctx, db)

	return seller_dto.ToSellerResponses(sellers)
}

func (service *SellerServiceImpl) FindById(ctx context.Context, id int64) seller_dto.SellerResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	seller, err := service.SellerRepository.FindByID(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return seller_dto.ToSellerResponse(seller)
}

func (service *SellerServiceImpl) Create(ctx context.Context, req seller_dto.CreateSellerRequest) seller_dto.SellerResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	seller := domain.Seller{
		Name:        req.Name,
		Category:    req.Category,
		Description: req.Description,
	}

	seller = service.SellerRepository.Create(ctx, db, seller)
	fmt.Println("seller created with id:", seller.Id)

	return seller_dto.ToSellerResponse(seller)
}

func (service *SellerServiceImpl) Update(ctx context.Context, req seller_dto.CreateSellerRequest, id int64) seller_dto.SellerResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	seller, err := service.SellerRepository.FindByID(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	seller.Name = req.Name
	seller.Category = req.Category
	seller.Description = req.Description

	seller = service.SellerRepository.Update(ctx, db, seller)
	return seller_dto.ToSellerResponse(seller)
}

func (service *SellerServiceImpl) Delete(ctx context.Context, id int64) {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	service.SellerRepository.Delete(ctx, db, id)
}
