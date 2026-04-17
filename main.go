package main

import (
	"errors"
	"fmt"
	"sync"
)

type OrderRepository interface {
	Add(order Order)
	GetByID(id int) (Order, error)
	GetAll() []Order
	Update(order Order) error
}

type InMemoryOrderRepo struct {
	mu     sync.Mutex
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
	or.mu.Lock()
	defer or.mu.Unlock()
	or.orders[order.ID] = order
}

func (or *InMemoryOrderRepo) GetByID(id int) (Order, error) {
	or.mu.Lock()
	defer or.mu.Unlock()
	if order, ok := or.orders[id]; !ok {
		return Order{}, ErrOrderNotFound
	} else {
		return order, nil
	}
}

func (or *InMemoryOrderRepo) GetAll() []Order {
	or.mu.Lock()
	defer or.mu.Unlock()
	orders := []Order{}
	for _, order := range or.orders {
		orders = append(orders, order)
	}
	return orders
}

func (or *InMemoryOrderRepo) Update(order Order) error {
	or.mu.Lock()
	defer or.mu.Unlock()
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

func DeliverMany(repo OrderRepository, ids []int) {
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		idOrder := id
		go func(id int) {
			defer wg.Done()
			order, err := repo.GetByID(id)
			if err == nil {
				order.MarkDelivered()
				err1 := repo.Update(order)
				if err1 != nil {
					fmt.Printf("Ошибка при обновлении ID %d: %v\n", id, err1)
				} else {
					fmt.Printf("Order %d delivered\n", order.ID)
				}
			} else {
				fmt.Printf("Ошибка:%v (ID: %d)\n", err, id)
			}
		}(idOrder)
	}

	wg.Wait()
}

var ErrOrderNotFound = errors.New("order not found")

func main() {
	repo := NewInMemoryOrderRepo()
	ids := []int{1, 2, 3, 99}

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
	repo.Add(Order{
		ID:          3,
		Customer:    "Петр",
		Address:     "Казань",
		IsDelivered: false,
	})
	repo.Add(Order{
		ID:          4,
		Customer:    "Катя",
		Address:     "ЕКБ",
		IsDelivered: false,
	})

	fmt.Println("[До доставки]")
	PrintAllOrders(repo)

	fmt.Println("\n[Результаты доставки]")
	DeliverMany(repo, ids)

	fmt.Println("\n[После доставки]")
	PrintAllOrders(repo)
}
