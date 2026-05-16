package order

import "errors"

type Order struct {
	ID          int    `json:"id"`
	Customer    string `json:"name"`
	Address     string `json:"address"`
	IsDelivered bool   `json:"is_delivered"`
}

var ErrOrderNotFound = errors.New("order not found")
