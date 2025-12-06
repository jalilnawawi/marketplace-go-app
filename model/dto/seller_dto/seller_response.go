package seller_dto

import (
	"time"

	"github.com/jalilnawawi/marketplace-app/model/domain"
)

type SellerResponse struct {
	SellerId    int64     `json:"sellerId"`
	SellerName  string    `json:"sellerName"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"createAt"`
}

func ToSellerResponse(seller domain.Seller) SellerResponse {
	return SellerResponse{
		SellerId:    seller.Id,
		SellerName:  seller.Name,
		Category:    seller.Category,
		Description: seller.Description,
		CreateAt:    seller.CreatedAt,
	}
}

func ToSellerResponses(sellers []domain.Seller) []SellerResponse {
	var sellerResponses []SellerResponse
	for _, seller := range sellers {
		sellerResponses = append(sellerResponses, ToSellerResponse(seller))
	}
	return sellerResponses
}
