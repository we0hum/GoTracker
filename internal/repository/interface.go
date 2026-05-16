package repository

import "GoTracker/internal/order"

type Repository interface {
	Add(order order.Order) (order.Order, error)
	GetByID(id int) (order.Order, error)
	GetAll() ([]order.Order, error)
	Update(order order.Order) error
	Delete(id int) error
}
