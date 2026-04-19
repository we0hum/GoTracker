package repository

import (
	"GoTracker/internal/order"
	"sync"
)

type InMemoryOrderRepo struct {
	mu     sync.Mutex
	orders map[int]order.Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[int]order.Order),
	}
}

func (or *InMemoryOrderRepo) Add(order order.Order) {
	or.mu.Lock()
	defer or.mu.Unlock()
	or.orders[order.ID] = order
}

func (or *InMemoryOrderRepo) GetByID(id int) (order.Order, error) {
	or.mu.Lock()
	defer or.mu.Unlock()
	if o, ok := or.orders[id]; !ok {
		return order.Order{}, order.ErrOrderNotFound
	} else {
		return o, nil
	}
}

func (or *InMemoryOrderRepo) GetAll() []order.Order {
	or.mu.Lock()
	defer or.mu.Unlock()
	var result []order.Order
	for _, o := range or.orders {
		result = append(result, o)
	}
	return result
}

func (or *InMemoryOrderRepo) Update(o order.Order) error {
	or.mu.Lock()
	defer or.mu.Unlock()
	if _, ok := or.orders[o.ID]; !ok {
		return order.ErrOrderNotFound
	}
	or.orders[o.ID] = o
	return nil
}
