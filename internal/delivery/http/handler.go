package http

import (
	"encoding/json"
	"github.com/AZRV17/Course_client/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(r chi.Router) {
	r.Route("/coins", func(r chi.Router) {
		r.Get("/", h.getAllCoins)
		r.Get("/{id}", h.getCoin)
	})
}

func (h *Handler) getCoin(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	coin, ok := h.service.CoinService.Get(id)
	if !ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": "coin not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(coin)
}

func (h *Handler) getAllCoins(w http.ResponseWriter, r *http.Request) {
	coins := h.service.CoinService.GetAll()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(coins)
}
