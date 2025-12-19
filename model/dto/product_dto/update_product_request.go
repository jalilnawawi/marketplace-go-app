package product_dto

type UpdateProductRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Price       int    `json:"price" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=10"`
	Category    string `json:"category" validate:"required"`
}
