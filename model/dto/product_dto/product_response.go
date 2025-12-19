package product_dto

import "github.com/jalilnawawi/marketplace-app/model/domain"

type ProductResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
	SellerID    int64  `json:"sellerId"`
}

func ToProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category:    product.Category,
		SellerID:    product.SellerId,
	}
}

func ToProductResponses(products []domain.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
