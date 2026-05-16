package utils

import (
	"GoTracker/internal/order"
	"encoding/json"
	"errors"
	"net/http"
)

func ValidateCreateOrder(o order.Order) error {
	if o.Customer == "" {
		return errors.New("поле customer не может быть пустым")
	}
	if o.Address == "" {
		return errors.New("поле address не может быть пустым")
	}
	return nil
}

func ValidateUpdateOrder(address string) error {
	if address == "" {
		return errors.New("поле address не может быть пустым")
	}
	return nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, code int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, code int, msg string) {
	_ = WriteJSON(w, code, ErrorResponse{Error: msg})
}
