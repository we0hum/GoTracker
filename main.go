package main

import "fmt"

type OrderRepository interface {
	Add(order Order)
	GetByID(id int) (Order, bool)
	GetAll() []Order
	Update(order Order)
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

func (or *InMemoryOrderRepo) GetByID(id int) (Order, bool) {
	order, ok := or.orders[id]
	return order, ok
}

func (or *InMemoryOrderRepo) GetAll() []Order {
	orders := []Order{}
	for _, order := range or.orders {
		orders = append(orders, order)
	}
	return orders
}

func (or *InMemoryOrderRepo) Update(order Order) {
	or.orders[order.ID] = order
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

	order, found := repo.GetByID(1)
	if found {
		order.MarkDelivered()
		repo.Update(order)
	}

	fmt.Println("\n[After]")
	PrintAllOrders(repo)
}
