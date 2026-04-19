package main

import (
	"GoTracker/internal/order"
	"GoTracker/internal/repository"
	"GoTracker/internal/service"
)

func main() {
	repo := repository.NewInMemoryOrderRepo()
	svc := service.NewOrderService(repo)

	repo.Add(order.Order{
		ID:          1,
		Customer:    "Андрей",
		Address:     "Москва",
		IsDelivered: false,
	})
	repo.Add(order.Order{
		ID:          2,
		Customer:    "Иван",
		Address:     "Питер",
		IsDelivered: false,
	})
	repo.Add(order.Order{
		ID:          3,
		Customer:    "Петр",
		Address:     "Казань",
		IsDelivered: false,
	})
	repo.Add(order.Order{
		ID:          4,
		Customer:    "Катя",
		Address:     "ЕКБ",
		IsDelivered: false,
	})

	svc.DeliverMany([]int{1, 2, 3, 99})
	svc.PrintAllOrders()
}
