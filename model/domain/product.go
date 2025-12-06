package domain

import "time"

type Product struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Category    string    `db:"category" json:"category"`
	Price       int       `db:"price" json:"price"`
	Description string    `db:"description" json:"description"`
	MerchantId  int64     `db:"merchant_id" json:"merchant_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
