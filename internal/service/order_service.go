package service

import (
	"GoTracker/internal/cache"
	"GoTracker/internal/order"
	"GoTracker/internal/queue"
	"GoTracker/internal/repository"
	"fmt"
)

type OrderService struct {
	repo  repository.Repository
	cache *cache.RedisCache
}

func NewOrderService(repo repository.Repository, cache *cache.RedisCache) *OrderService {
	return &OrderService{repo: repo, cache: cache}
}

func (os *OrderService) AddOrder(o order.Order) (order.Order, error) {
	err := os.repo.Add(o)
	if err != nil {
		return order.Order{}, err
	}
	_ = os.cache.Delete(o.ID)
	_ = queue.SendOrderCreated(o)
	return o, nil
}

func (os *OrderService) GetAll() ([]order.Order, error) {
	return os.repo.GetAll(), nil
}

func (os *OrderService) GetOrderByID(id int) (order.Order, error) {
	if os.cache != nil {
		if o, err := os.cache.Get(id); err == nil {
			fmt.Println("[Cache] Найден заказ в Redis")
			return o, nil
		}
	}

	o, err := os.repo.GetByID(id)
	if err != nil {
		return o, err
	}

	if os.cache != nil {
		_ = os.cache.Set(o)
	}
	return o, nil
}

func (os *OrderService) Update(o order.Order) (order.Order, error) {
	err := os.repo.Update(o)
	if err != nil {
		return o, err
	}
	_ = os.cache.Delete(o.ID)
	return o, nil
}
