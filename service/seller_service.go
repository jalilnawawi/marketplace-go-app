package service

import (
	"context"

	"github.com/jalilnawawi/marketplace-app/model/dto/seller_dto"
)

type SellerService interface {
	FindAll(ctx context.Context) []seller_dto.SellerResponse
	FindById(ctx context.Context, id int64) seller_dto.SellerResponse
	Create(ctx context.Context, req seller_dto.CreateSellerRequest) seller_dto.SellerResponse
	Update(ctx context.Context, req seller_dto.CreateSellerRequest, id int64) seller_dto.SellerResponse
	Delete(ctx context.Context, id int64)
}
