package order

import "errors"

type Order struct {
	ID       int    `json:"id"`
	Customer string `json:"name"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}

func (o *Order) MarkDelivered() {
	o.Status = "delivered"
}

var ErrOrderNotFound = errors.New("order not found")
