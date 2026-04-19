package order

import "errors"

type Order struct {
	ID          int
	Customer    string
	Address     string
	IsDelivered bool
}

func (o *Order) MarkDelivered() {
	o.IsDelivered = true
}

var ErrOrderNotFound = errors.New("order not found")
