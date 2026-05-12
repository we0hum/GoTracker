package repository

import "GoTracker/internal/order"

type Repository interface {
	Add(order order.Order) error
	GetByID(id int) (order.Order, error)
	GetAll() []order.Order
	Update(order order.Order) error
}
