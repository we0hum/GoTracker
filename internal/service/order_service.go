package service

import (
	"GoTracker/internal/order"
	"GoTracker/internal/repository"
	"fmt"
	"sync"
)

type OrderService struct {
	repo repository.Repository
}

func NewOrderService(repo repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (os *OrderService) PrintAllOrders() {
	for _, order := range os.repo.GetAll() {
		fmt.Printf("ID: %d, %s, Delivered: %v\n", order.ID, order.Customer, order.Status)
	}
}

func (os *OrderService) DeliverMany(ids []int) {
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		idOrder := id
		go func(id int) {
			defer wg.Done()
			order, err := os.repo.GetByID(id)
			if err == nil {
				order.MarkDelivered()
				err1 := os.repo.Update(order)
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

func (os *OrderService) CreateOrder(o order.Order) (order.Order, error) {
	os.repo.Add(o)
	return o, nil
}

func (os *OrderService) ListOrders() ([]order.Order, error) {
	return os.repo.GetAll(), nil
}

func (os *OrderService) UpdateOrder(id int, address string, status string) (order.Order, error) {
	o, err := os.repo.GetByID(id)
	if err != nil {
		return order.Order{}, err
	}

	o.Address = address
	o.Status = status

	if err := os.repo.Update(o); err != nil {
		return order.Order{}, err
	}

	return o, nil
}
