package order

import (
	"errors"
	"time"
)

type Order struct {
	ID          int       `json:"id,omitempty" db:"id"`
	Customer    string    `json:"customer" db:"customer"`
	Address     string    `json:"address" db:"address"`
	IsDelivered bool      `json:"is_delivered" db:"is_delivered"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
}

var ErrOrderNotFound = errors.New("order not found")
