package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/metrics"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/model"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/service"
)

type ItemHandler struct {
	service   *service.ItemService
	collector *metrics.Collector
	logger    *slog.Logger
}

func NewItemHandler(service *service.ItemService, collector *metrics.Collector, logger *slog.Logger) *ItemHandler {
	return &ItemHandler{
		service:   service,
		collector: collector,
		logger:    logger,
	}
}

func (h *ItemHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.handleHealth)
	mux.HandleFunc("GET /metrics", h.handleMetrics)
	mux.HandleFunc("GET /items", h.handleListItems)
	mux.HandleFunc("GET /items/{id}", h.handleGetItem)
	registerProfiler(mux)
	return withMiddleware(h.logger, mux)
}

func (h *ItemHandler) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *ItemHandler) handleMetrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(h.collector.RenderPrometheus()))
}

func (h *ItemHandler) handleGetItem(w http.ResponseWriter, r *http.Request) {
	item, source, err := h.service.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("X-Item-Source", source)
	writeJSON(w, http.StatusOK, model.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Tags:        item.Tags,
		CreatedAt:   item.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   item.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (h *ItemHandler) handleListItems(w http.ResponseWriter, r *http.Request) {
	detailed := r.URL.Query().Get("view") == "detailed"
	items, err := h.service.ListItems(r.Context(), detailed)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, model.ListItemsResponse{Items: items})
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, model.ErrInvalidItemID), errors.Is(err, model.ErrInvalidItem):
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, model.ErrItemNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
