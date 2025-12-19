package service

import (
	"context"

	"github.com/jalilnawawi/marketplace-app/model/dto/product_dto"
)

type ProductService interface {
	FindAll(ctx context.Context) []product_dto.ProductResponse
	FindById(ctx context.Context, id int64) product_dto.ProductResponse
	Create(ctx context.Context, req product_dto.CreateProductRequest) (product_dto.ProductResponse, error)
	Update(ctx context.Context, req product_dto.UpdateProductRequest, id int64) product_dto.ProductResponse
	Delete(ctx context.Context, id int64)
	FindPriceByProductId(ctx context.Context, id int64) (float64, error)
}
