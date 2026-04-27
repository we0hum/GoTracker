package http

import (
	"GoTracker/internal/order"
	"GoTracker/internal/service"
	"GoTracker/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *service.OrderService
}

func NewHandler(service *service.OrderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		w.Write([]byte("Добро пожаловать в GoTracker API"))
		return
	}

	if r.URL.Path == "/orders" && r.Method == http.MethodGet {
		h.handleGetAll(w, r)
		return
	}

	if r.URL.Path == "/orders" && r.Method == http.MethodPost {
		h.handleCreate(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/orders/") && r.Method == http.MethodPut {
		h.handleUpdate(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request) {
	orders, _ := h.service.ListOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var o order.Order

	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный JSON")
		return
	}

	if err := utils.ValidateBook(o); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if o.Status != "pending" && o.Status != "delivered" {
		utils.WriteError(w, http.StatusBadRequest, "невалидный статус")
		return
	}

	created, _ := h.service.CreateOrder(o)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		utils.WriteError(w, http.StatusBadRequest, "Неправильный URL")
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный ID")
		return
	}

	var req struct {
		Address string `json:"address"`
		Status  string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный JSON")
		return
	}

	if req.Address == "" {
		utils.WriteError(w, http.StatusBadRequest, "address не может быть пустым")
		return
	}

	if req.Status != "pending" && req.Status != "delivered" {
		utils.WriteError(w, http.StatusBadRequest, "невалидный статус")
		return
	}

	updated, err := h.service.UpdateOrder(id, req.Address, req.Status)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}
