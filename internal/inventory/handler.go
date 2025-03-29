package inventory

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := h.repo.CreateItem(item); err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
