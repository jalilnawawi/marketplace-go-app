package service

import (
	"context"

	"github.com/jalilnawawi/marketplace-app/model/dto/order_dto"
)

type OrderService interface {
	FindAll(ctx context.Context) []order_dto.OrderResponse
	FindById(ctx context.Context, id int64) order_dto.OrderResponse
	Create(ctx context.Context, req order_dto.CreateOrderRequest) (order_dto.OrderResponse, error)
	Update(ctx context.Context, req order_dto.UpdateOrderRequest, id int64) order_dto.OrderResponse
	Delete(ctx context.Context, id int64)
}
