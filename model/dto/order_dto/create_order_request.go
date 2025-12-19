package order_dto

type CreateOrderRequest struct {
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
