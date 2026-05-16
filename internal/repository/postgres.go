package repository

import (
	"GoTracker/internal/order"

	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PostgresOrderRepo struct {
	db *sqlx.DB
}

func NewPostgresOrderRepo(db *sqlx.DB) *PostgresOrderRepo {
	return &PostgresOrderRepo{db: db}
}

func (r *PostgresOrderRepo) Add(o order.Order) (order.Order, error) {
	query := `INSERT INTO orders (customer, address, is_delivered) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.QueryRow(query, o.Customer, o.Address, o.IsDelivered).Scan(&o.ID, &o.CreatedAt)
	return o, err
}

func (r *PostgresOrderRepo) GetByID(id int) (order.Order, error) {
	var o order.Order
	query := `SELECT id, customer, address, is_delivered, created_at FROM orders WHERE id = $1`
	err := r.db.Get(&o, query, id)
	if err == sql.ErrNoRows {
		return o, order.ErrOrderNotFound
	}
	return o, err
}

func (r *PostgresOrderRepo) GetAll() ([]order.Order, error) {
	var orders []order.Order
	query := `SELECT id, customer, address, is_delivered, created_at FROM orders`
	err := r.db.Select(&orders, query)
	return orders, err
}

func (r *PostgresOrderRepo) Update(o order.Order) error {
	query := `UPDATE orders SET customer = $1, address = $2, is_delivered = $3 WHERE id = $4`
	result, err := r.db.Exec(query, o.Customer, o.Address, o.IsDelivered, o.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return order.ErrOrderNotFound
	}

	return nil
}

func (r *PostgresOrderRepo) Delete(id int) error {
	query := `DELETE FROM orders WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return order.ErrOrderNotFound
	}

	return nil
}
