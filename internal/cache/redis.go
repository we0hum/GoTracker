package cache

import (
	"GoTracker/internal/order"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(addr string, ttlSeconds int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{
		client: rdb,
		ttl:    time.Duration(ttlSeconds) * time.Second,
	}
}

func (c *RedisCache) key(id int) string {
	return fmt.Sprintf("order:%d", id)
}

func (c *RedisCache) Get(id int) (order.Order, error) {
	var o order.Order
	ctx := context.Background()
	data, err := c.client.Get(ctx, c.key(id)).Result()
	if err == redis.Nil {
		return o, order.ErrOrderNotFound
	}
	if err != nil {
		return o, err
	}
	err = json.Unmarshal([]byte(data), &o)
	return o, err
}

func (c *RedisCache) Set(o order.Order) error {
	ctx := context.Background()
	data, err := json.Marshal(o)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, c.key(o.ID), data, c.ttl).Err()
}

func (c *RedisCache) Delete(id int) error {
	ctx := context.Background()
	return c.client.Del(ctx, c.key(id)).Err()
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}
