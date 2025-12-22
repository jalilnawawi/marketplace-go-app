package order_dto

type UpdateOrderRequest struct {
	ProductId int64 `json:"productId"`
	Quantity  int   `json:"quantity"`
}
