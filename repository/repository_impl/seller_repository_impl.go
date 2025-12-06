package repository_impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/domain"
	"github.com/jalilnawawi/marketplace-app/repository"
)

type SellerRepositoryImpl struct {
}

func NewSellerRepository() repository.SellerRepository {
	return &SellerRepositoryImpl{}
}

func (s *SellerRepositoryImpl) FindAll(ctx context.Context, db *sql.Tx) []domain.Seller {
	SQL := "SELECT id, name, category, description FROM sellers"
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var sellers []domain.Seller
	for rows.Next() {
		seller := domain.Seller{}
		err := rows.Scan(&seller.Id, &seller.Name, &seller.Category, &seller.Description, &seller.CreatedAt)
		helper.PanicIfError(err)
		sellers = append(sellers, seller)
	}

	return sellers
}

func (s *SellerRepositoryImpl) FindByID(ctx context.Context, db *sql.Tx, id int64) (domain.Seller, error) {
	SQL := "SELECT id, name, category, description FROM sellers WHERE id=$1"
	result, err := db.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer result.Close()

	seller := domain.Seller{}
	if result.Next() {
		err := result.Scan(&seller.Id, &seller.Name, &seller.Category, &seller.Description)
		helper.PanicIfError(err)
		return seller, nil
	} else {
		return seller, errors.New(fmt.Sprintf("seller with id %d not found", id))
	}
}

func (s *SellerRepositoryImpl) Create(ctx context.Context, db *sql.Tx, seller domain.Seller) domain.Seller {
	SQL := "insert into marketplace.seller (name, category, description) values ($1, $2, $3) returning id"
	result, err := db.ExecContext(ctx, SQL, seller.Name, seller.Category, seller.Description)
	helper.PanicIfError(err)
	log.Printf("result is %v", result)
	return seller
}

func (s *SellerRepositoryImpl) Update(ctx context.Context, db *sql.Tx, seller domain.Seller) domain.Seller {
	SQL := "update seller set name=$1, category=$2, description=$3 where id=$4"
	result, err := db.ExecContext(ctx, SQL, seller.Name, seller.Category, seller.Description, seller.Id)
	helper.PanicIfError(err)

	log.Printf("result is %v", result)
	return seller
}

func (s *SellerRepositoryImpl) Delete(ctx context.Context, db *sql.Tx, id int64) {
	SQL := "delete from seller where id=$1"
	result, err := db.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)

	log.Printf("result is %v", result)
}
