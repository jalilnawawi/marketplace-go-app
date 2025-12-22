package domain

import "time"

type Orders struct {
	ID         int64     `db:"id" json:"id"`
	SellerID   int64     `db:"seller_id" json:"sellerId"`
	ProductId  int64     `db:"product_id" json:"productId"`
	Quantity   int       `db:"quantity" json:"quantity"`
	TotalPrice int       `db:"total_price" json:"totalPrice"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
}
