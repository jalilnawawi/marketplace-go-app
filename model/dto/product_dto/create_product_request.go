package product_dto

type CreateProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
	SellerId    int64  `json:"sellerId"`
}
