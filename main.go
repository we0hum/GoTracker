package main

import (
	"errors"
	"fmt"
)

type OrderRepository interface {
	Add(order Order)
	GetByID(id int) (Order, error)
	GetAll() []Order
	Update(order Order) error
}

type InMemoryOrderRepo struct {
	orders map[int]Order
}
type Order struct {
	ID          int
	Customer    string
	Address     string
	IsDelivered bool
}

func (o *Order) MarkDelivered() {
	o.IsDelivered = true
}

func (or *InMemoryOrderRepo) Add(order Order) {
	or.orders[order.ID] = order
}

func (or *InMemoryOrderRepo) GetByID(id int) (Order, error) {
	if order, ok := or.orders[id]; !ok {
		return Order{}, ErrOrderNotFound
	} else {
		return order, nil
	}
}

func (or *InMemoryOrderRepo) GetAll() []Order {
	orders := []Order{}
	for _, order := range or.orders {
		orders = append(orders, order)
	}
	return orders
}

func (or *InMemoryOrderRepo) Update(order Order) error {
	if _, ok := or.orders[order.ID]; !ok {
		return ErrOrderNotFound
	}
	or.orders[order.ID] = order
	return nil
}

func PrintAllOrders(repo *InMemoryOrderRepo) {
	for _, order := range repo.orders {
		fmt.Printf("ID: %d, %s, Delivered: %v\n", order.ID, order.Customer, order.IsDelivered)
	}
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[int]Order),
	}
}

var ErrOrderNotFound = errors.New("order not found")

func main() {
	repo := NewInMemoryOrderRepo()

	repo.Add(Order{
		ID:          1,
		Customer:    "Андрей",
		Address:     "Москва",
		IsDelivered: false,
	})
	repo.Add(Order{
		ID:          2,
		Customer:    "Иван",
		Address:     "Питер",
		IsDelivered: false,
	})

	fmt.Println("[Before]")
	PrintAllOrders(repo)

	order, err := repo.GetByID(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		order.MarkDelivered()
		err := repo.Update(order)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("\nUpdated order", order.ID)
		}
	}

	fmt.Println("\n[After]")
	PrintAllOrders(repo)

	nonExistent := Order{ID: 999, Customer: "Петя", Address: "Питер"}
	err = repo.Update(nonExistent)
	if err != nil {
		fmt.Println("\nОшибка:", err)
	}
}
