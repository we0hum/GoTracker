package service

import (
	"GoTracker/internal/cache"
	"GoTracker/internal/order"
	"GoTracker/internal/queue"
	"GoTracker/internal/repository"
)

type OrderService struct {
	repo  repository.Repository
	cache *cache.RedisCache
}

func NewOrderService(repo repository.Repository, cache *cache.RedisCache) *OrderService {
	return &OrderService{repo: repo, cache: cache}
}

func (os *OrderService) AddOrder(o order.Order) (order.Order, error) {
	created, err := os.repo.Add(o)
	if err != nil {
		return order.Order{}, err
	}

	if os.cache != nil {
		_ = os.cache.Delete(created.ID)
	}

	if err := queue.SendOrderCreated(created); err != nil {
		return created, err
	}

	return created, nil
}

func (os *OrderService) GetAll() ([]order.Order, error) {
	return os.repo.GetAll()
}

func (os *OrderService) GetOrderByID(id int) (order.Order, error) {
	if os.cache != nil {
		if o, err := os.cache.Get(id); err == nil {
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
	if err := os.repo.Update(o); err != nil {
		return order.Order{}, err
	}

	if os.cache != nil {
		_ = os.cache.Delete(o.ID)
	}

	return os.repo.GetByID(o.ID)
}

func (os *OrderService) Delete(id int) error {
	if err := os.repo.Delete(id); err != nil {
		return err
	}

	if os.cache != nil {
		_ = os.cache.Delete(id)
	}

	return nil
}
