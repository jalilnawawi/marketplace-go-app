package order_dto

import "github.com/jalilnawawi/marketplace-app/model/domain"

type OrderResponse struct {
	ID         int64 `json:"id"`
	ProductId  int64 `json:"productId"`
	Quantity   int   `json:"quantity"`
	TotalPrice int   `json:"totalPrice"`
}

func ToOrderResponse(order domain.Orders) OrderResponse {
	return OrderResponse{
		ID:         order.ID,
		ProductId:  order.ProductId,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
	}
}

func ToOrderResponses(orders []domain.Orders) []OrderResponse {
	var orderResponses []OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}
	return orderResponses
}
