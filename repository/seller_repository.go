package repository

import (
	"context"
	"database/sql"

	"github.com/jalilnawawi/marketplace-app/model/domain"
)

type SellerRepository interface {
	FindAll(ctx context.Context, db *sql.Tx) []domain.Seller
	FindByID(ctx context.Context, db *sql.Tx, id int64) (domain.Seller, error)
	Create(ctx context.Context, db *sql.Tx, seller domain.Seller) domain.Seller
	Update(ctx context.Context, db *sql.Tx, seller domain.Seller) domain.Seller
	Delete(ctx context.Context, db *sql.Tx, id int64)
}
