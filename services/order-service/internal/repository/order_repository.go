package repository

import (
	"context"
	"errors"
	model "order-service/internal/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	ID       int
	UserID   int
	Product  string
	Quantity int
}

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(order model.Order) error {
	query := `INSERT INTO orders (
		user_id,
		product,
		quantity
		)
		VALUES ($1, $2, $3)`

	_, err := r.db.Exec(
		context.Background(),
		query,
		order.UserID,
		order.Product,
		order.Quantity,
	)
	return err
}


func (r *OrderRepository) GetOrderByID(ID string) (model.Order, error) {

	var order model.Order

	query := `
		SELECT id, user_id, product, quantity
		FROM orders
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		ID,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.Product,
		&order.Quantity,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Order{}, errors.New("order not found")
		}
		return model.Order{}, err
	}

	return order, nil
}


func (r *OrderRepository) GetAllOrders() ([]model.Order, error) {

	var orders []model.Order

	query := `
		SELECT id, user_id, product, quantity
		FROM orders
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var order model.Order

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Product,
			&order.Quantity,
		)
		if err != nil {
			return nil, err
		}

		orders = append(
			orders,
			order,
		)
	}

	return orders, nil
}