package http

import (
	"GoTracker/internal/order"
	"GoTracker/internal/service"
	"GoTracker/internal/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.OrderService
}

type createOrderRequest struct {
	Customer    string `json:"customer"`
	Address     string `json:"address"`
	IsDelivered bool   `json:"is_delivered"`
}

type updateOrderRequest struct {
	Address     string `json:"address"`
	IsDelivered bool   `json:"is_delivered"`
}

func NewRouter(service *service.OrderService) http.Handler {
	router := mux.NewRouter()
	h := NewHandler(service)

	router.HandleFunc("/orders", h.handleGetAll).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id:[0-9]+}", h.handleGetByID).Methods(http.MethodGet)
	router.HandleFunc("/orders", h.handleCreate).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id:[0-9]+}", h.handleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/orders/{id:[0-9]+}", h.handleDelete).Methods(http.MethodDelete)

	return router
}

func NewHandler(service *service.OrderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *Handler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный ID")
		return
	}

	result, err := h.service.GetOrderByID(id)
	if err != nil {
		if errors.Is(err, order.ErrOrderNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный JSON")
		return
	}

	if err := utils.ValidateCreateOrder(order.Order{Customer: req.Customer, Address: req.Address}); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.service.AddOrder(order.Order{Customer: req.Customer, Address: req.Address, IsDelivered: req.IsDelivered})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, created)
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный ID")
		return
	}

	var req updateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный JSON")
		return
	}

	if err := utils.ValidateUpdateOrder(req.Address); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := h.service.Update(order.Order{ID: id, Address: req.Address, IsDelivered: req.IsDelivered})
	if err != nil {
		if errors.Is(err, order.ErrOrderNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, updated)
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "Невалидный ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, order.ErrOrderNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
