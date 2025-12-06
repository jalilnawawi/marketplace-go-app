package seller_dto

type CreateSellerRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
