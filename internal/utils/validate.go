package utils

import (
	"GoTracker/internal/order"
	"encoding/json"
	"errors"
	"net/http"
)

func ValidateBook(o order.Order) error {
	if o.Customer == "" {
		return errors.New("поле 'customer' не может быть пустым")
	}
	if o.Address == "" {
		return errors.New("поле 'address' не может быть пустым")
	}
	if o.ID <= 0 {
		return errors.New("ID должен быть больше нуля")
	}
	return nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}
