package repository_impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/jalilnawawi/marketplace-app/exception"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/domain"
	"github.com/jalilnawawi/marketplace-app/repository"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() repository.ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) FindAll(ctx context.Context, db *sql.Tx) []domain.Product {
	SQL := "select id, name, category, price, description, seller_id, created_at from marketplace.products"
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description, &product.SellerId, &product.CreatedAt)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}

func (p ProductRepositoryImpl) FindByID(ctx context.Context, db *sql.Tx, id int64) (domain.Product, error) {
	SQL := "select id, name, category, price, description, seller_id, created_at from marketplace.products where id=$1"
	result, err := db.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer result.Close()

	product := domain.Product{}
	if result.Next() {
		err := result.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description, &product.SellerId, &product.CreatedAt)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, nil
	}
}

func (p ProductRepositoryImpl) Create(ctx context.Context, db *sql.Tx, product domain.Product) (domain.Product, error) {
	SQL := `
			insert into marketplace.products (name, description, price, category, seller_id)
			values ($1, $2, $3, $4, $5)
			RETURNING id
		   `

	var id int64
	err := db.QueryRowContext(ctx, SQL, product.Name, product.Description, product.Price, product.Category, product.SellerId).Scan(&id)
	helper.PanicIfError(err)

	if err != nil {
		return product, err
	}

	product.ID = id
	return product, nil
}

func (p ProductRepositoryImpl) Update(ctx context.Context, db *sql.Tx, product domain.Product) domain.Product {
	SQL := "update marketplace.products set name=$1, description=$2, price=$3, category=$4 where id=$5"
	result, err := db.ExecContext(ctx, SQL, product.Name, product.Description, product.Price, product.Category, product.ID)
	helper.PanicIfError(err)

	rowCount, err := result.RowsAffected()
	if rowCount == 0 {
		panic(exception.NewNotFoundError("product not found"))
	}

	log.Printf("row affected %d", rowCount)
	return product
}

func (p ProductRepositoryImpl) Delete(ctx context.Context, db *sql.Tx, id int64) {
	SQL := "delete from marketplace.products where id=$1"
	result, err := db.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
	rowCount, err := result.RowsAffected()
	helper.PanicIfError(err)

	log.Printf("row affected %d", rowCount)
}

func (p ProductRepositoryImpl) FindPriceByProductId(ctx context.Context, db *sql.Tx, id int64) (float64, error) {
	SQL := "select price from marketplace.products where id=$1"
	result, err := db.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	var productPrice float64
	if result.Next() {
		err := result.Scan(&productPrice)
		helper.PanicIfError(err)
		return productPrice, nil
	} else {
		return 0, nil
	}
}
