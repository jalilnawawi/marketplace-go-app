package repository_impl

import (
	"context"
	"database/sql"

	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/domain"
	"github.com/jalilnawawi/marketplace-app/repository"
)

type OrderRepositoryImpl struct{}

func NewOrderRepository() repository.OrderRepository {
	return &OrderRepositoryImpl{}
}

func (o OrderRepositoryImpl) FindAll(ctx context.Context, db *sql.Tx) ([]domain.Orders, error) {
	SQL := "select o.id, o.product_id, o.quantity, o.total_price, o.created_at from marketplace.orders o"
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Orders
	for rows.Next() {
		order := domain.Orders{}
		err := rows.Scan(&order.ID, &order.ProductId, &order.Quantity, &order.TotalPrice, &order.CreatedAt)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}

	return orders, err
}

func (o OrderRepositoryImpl) FindById(ctx context.Context, db *sql.Tx, id int64) (domain.Orders, error) {
	SQL := "select o.id, o.product_id, o.quantity, o.total_price, o.created_at from marketplace.orders o where o.id=$1"
	result, err := db.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer result.Close()

	order := domain.Orders{}
	if result.Next() {
		err := result.Scan(&order.ID, &order.ProductId, &order.Quantity, &order.TotalPrice, &order.CreatedAt)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, nil
	}
}

func (o OrderRepositoryImpl) Create(ctx context.Context, db *sql.Tx, order domain.Orders, product domain.Product, seller domain.Seller) (domain.Orders, error) {
	SQL := `
			insert into marketplace.orders (seller_id, product_id, quantity, total_price)
			values ($1, $2, $3, $4)
			RETURNING id
		   `

	var id int
	err := db.QueryRowContext(ctx, SQL, order.SellerID, order.ProductId, order.Quantity, order.TotalPrice).Scan(&id)
	helper.PanicIfError(err)

	if err != nil {
		return order, err
	}
	order.ID = int64(id)
	return order, nil
}

func (o OrderRepositoryImpl) Update(ctx context.Context, db *sql.Tx, order domain.Orders) (domain.Orders, error) {
	SQL := `
			update marketplace.orders
			set quantity=$1
			where id=$2
		   `
	result, err := db.ExecContext(ctx, SQL, order.Quantity, order.ID)
	helper.PanicIfError(err)
	rowCount, err := result.RowsAffected()
	_ = rowCount
	return order, err
}

func (o OrderRepositoryImpl) Delete(ctx context.Context, db *sql.Tx, id int64) {
	SQL := `
			delete from marketplace.orders
			where id=$1
			`
	_, err := db.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}
