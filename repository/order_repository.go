package repository

import (
	"context"
	"database/sql"

	"github.com/jalilnawawi/marketplace-app/model/domain"
)

type OrderRepository interface {
	FindAll(ctx context.Context, db *sql.Tx) ([]domain.Orders, error)
	FindById(ctx context.Context, db *sql.Tx, id int) (domain.Orders, error)
	Create(ctx context.Context, db *sql.Tx, order domain.Orders, product domain.Product) (domain.Orders, error)
	Update(ctx context.Context, db *sql.Tx, order domain.Orders) (domain.Orders, error)
	Delete(ctx context.Context, db *sql.Tx, id int)
}
