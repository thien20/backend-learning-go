package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/model"
	"github.com/thien/backend-learning-go/02-build-crud-api/internal/service"
)

type ItemHandler struct {
	service *service.ItemService
	logger  *slog.Logger
}

func NewItemHandler(service *service.ItemService, logger *slog.Logger) *ItemHandler {
	return &ItemHandler{
		service: service,
		logger:  logger,
	}
}

func (h *ItemHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.handleHealth)
	mux.HandleFunc("GET /metrics", h.handleMetrics)
	mux.HandleFunc("GET /items", h.handleListItems)
	mux.HandleFunc("POST /items", h.handleCreateItem)
	mux.HandleFunc("GET /items/{id}", h.handleGetItem)
	mux.HandleFunc("PUT /items/{id}", h.handleUpdateItem)
	mux.HandleFunc("DELETE /items/{id}", h.handleDeleteItem)
	return withMiddleware(h.logger, mux)
}

func (h *ItemHandler) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *ItemHandler) handleMetrics(w http.ResponseWriter, _ *http.Request) {
	// TODO: Replace this with real Prometheus output. Start with one request
	// counter and one latency-oriented metric so you can practice the difference
	// between counters and histograms.
	writeJSON(w, http.StatusOK, map[string]string{"metrics": "student task: expose counters and latency metrics"})
}

func (h *ItemHandler) handleListItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.ListItems(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	resp := model.ListItemsResponse{Items: make([]model.ItemResponse, 0, len(items))}
	for _, item := range items {
		resp.Items = append(resp.Items, service.MapItemResponse(item))
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *ItemHandler) handleGetItem(w http.ResponseWriter, r *http.Request) {
	item, source, err := h.service.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("X-Item-Source", source)
	writeJSON(w, http.StatusOK, service.MapItemResponse(item))
}

func (h *ItemHandler) handleCreateItem(w http.ResponseWriter, r *http.Request) {
	var req model.CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json body"})
		return
	}

	item, err := h.service.CreateItem(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, service.MapItemResponse(item))
}

func (h *ItemHandler) handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json body"})
		return
	}

	item, err := h.service.UpdateItem(r.Context(), r.PathValue("id"), req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, service.MapItemResponse(item))
}

func (h *ItemHandler) handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteItem(r.Context(), r.PathValue("id")); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, model.ErrInvalidItemID), errors.Is(err, model.ErrInvalidItem):
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, model.ErrItemNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.Is(err, model.ErrNotImplemented):
		writeJSON(w, http.StatusNotImplemented, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
