// Package handler owns HTTP concerns.
//
// Handlers should stay thin: parse the request, call the service, map the
// result to HTTP, and write the response.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/model"
	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/service"
)

type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.handleHealth)
	mux.HandleFunc("GET /items/{id}", h.handleGetItem)
	return mux
}

func (h *ItemHandler) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *ItemHandler) handleGetItem(w http.ResponseWriter, r *http.Request) {
	item, source, err := h.service.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		switch {
		case errors.Is(err, model.ErrInvalidItemID):
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		case errors.Is(err, model.ErrItemNotFound):
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
		return
	}

	w.Header().Set("X-Item-Source", source)

	// TODO: Decide what a stable API response should look like. Is returning the
	// domain model directly good enough, or should there be a dedicated response
	// shape?

	writeJSON(w, http.StatusOK, GetItemResponse(item))
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
