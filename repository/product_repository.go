package repository

import (
	"context"
	"database/sql"

	"github.com/jalilnawawi/marketplace-app/model/domain"
)

type ProductRepository interface {
	FindAll(ctx context.Context, db *sql.Tx) []domain.Product
	FindByID(ctx context.Context, db *sql.Tx, id int64) (domain.Product, error)
	Create(ctx context.Context, db *sql.Tx, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, db *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, db *sql.Tx, id int64)
	FindPriceByProductId(ctx context.Context, db *sql.Tx, id int64) (float64, error)
}
