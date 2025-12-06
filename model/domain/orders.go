package domain

import "time"

type Orders struct {
	ID         int64     `db:"id" json:"id"`
	ProductId  int64     `db:"product_id" json:"product_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
	TotalPrice int       `db:"total_price" json:"total_price"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
