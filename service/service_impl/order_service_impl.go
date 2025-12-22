package service_impl

import (
	"context"
	"database/sql"
	"time"

	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/domain"
	"github.com/jalilnawawi/marketplace-app/model/dto/order_dto"
	"github.com/jalilnawawi/marketplace-app/repository"
	"github.com/jalilnawawi/marketplace-app/service"
)

type OrderServiceImpl struct {
	OrderRepository   repository.OrderRepository
	ProductRepository repository.ProductRepository
	SellerRepository  repository.SellerRepository
	DB                *sql.DB
}

func NewOrderService(orderRepository repository.OrderRepository, productRepository repository.ProductRepository, sellerRepository repository.SellerRepository, DB *sql.DB) service.OrderService {
	return &OrderServiceImpl{
		OrderRepository:   orderRepository,
		ProductRepository: productRepository,
		SellerRepository:  sellerRepository,
		DB:                DB,
	}
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []order_dto.OrderResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	orders, err := service.OrderRepository.FindAll(ctx, db)
	if orders == nil {
		orders = []domain.Orders{}
	}
	helper.PanicIfError(err)

	return order_dto.ToOrderResponses(orders)
}

func (service *OrderServiceImpl) FindById(ctx context.Context, id int64) order_dto.OrderResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	order, err := service.OrderRepository.FindById(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return order_dto.ToOrderResponse(order)
}

func (service *OrderServiceImpl) Create(ctx context.Context, req order_dto.CreateOrderRequest) (order_dto.OrderResponse, error) {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	// check seller if exist
	seller, err := service.SellerRepository.FindByID(ctx, db, req.SellerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// check product if exist
	product, err := service.ProductRepository.FindByID(ctx, db, req.ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// get product price
	totalPrice := product.Price * req.Quantity

	order := domain.Orders{
		SellerID:   req.SellerId,
		ProductId:  req.ProductId,
		Quantity:   req.Quantity,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}

	// create order
	createOrder, err := service.OrderRepository.Create(ctx, db, order, product, seller)
	if err != nil {
		return order_dto.OrderResponse{}, err
	}

	return order_dto.ToOrderResponse(createOrder), nil
}

func (service *OrderServiceImpl) Update(ctx context.Context, req order_dto.UpdateOrderRequest, id int64) order_dto.OrderResponse {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	order, err := service.OrderRepository.FindById(ctx, db, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	order.ProductId = req.ProductId
	order.Quantity = req.Quantity

	updatedOrder, err := service.OrderRepository.Update(ctx, db, order)
	if err != nil {
		return order_dto.OrderResponse{}
	}

	return order_dto.ToOrderResponse(updatedOrder)
}

func (service *OrderServiceImpl) Delete(ctx context.Context, id int64) {
	db, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(db)

	service.OrderRepository.Delete(ctx, db, id)
}
