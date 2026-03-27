package service

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/model"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (model.Item, error)
	List(ctx context.Context) ([]model.Item, error)
	Create(ctx context.Context, item model.Item) (model.Item, error)
	Update(ctx context.Context, item model.Item) (model.Item, error)
	Delete(ctx context.Context, id string) error
}

type Cache interface {
	Get(ctx context.Context, key string) (model.Item, bool, error)
	Set(ctx context.Context, key string, item model.Item) error
	Delete(ctx context.Context, key string) error
}

type ItemService struct {
	repo   Repository
	cache  Cache
	logger *slog.Logger
}

func NewItemService(repo Repository, cache Cache, logger *slog.Logger) *ItemService {
	return &ItemService{
		repo:   repo,
		cache:  cache,
		logger: logger,
	}
}

func (s *ItemService) GetItem(ctx context.Context, id string) (model.Item, string, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return model.Item{}, "", model.ErrInvalidItemID
	}

	cached, found, err := s.cache.Get(ctx, id)
	if err == nil && found {
		return cached, "cache", nil
	}

	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Item{}, "", err
	}

	_ = s.cache.Set(ctx, id, item)
	return item, "db", nil
}

func (s *ItemService) ListItems(ctx context.Context) ([]model.Item, error) {
	return s.repo.List(ctx)
}

func (s *ItemService) CreateItem(ctx context.Context, req model.CreateItemRequest) (model.Item, error) {
	// TODO: Validate the request, generate an ID, map DTO -> domain model, and
	// decide which timestamps are owned by the service versus the DB.
	s.logger.Info("create item requested", "name", req.Name)
	return model.Item{}, model.ErrNotImplemented
}

func (s *ItemService) UpdateItem(ctx context.Context, id string, req model.UpdateItemRequest) (model.Item, error) {
	// TODO: Load the current item, apply partial updates carefully, validate the
	// result, and invalidate or refresh cache entries.
	s.logger.Info("update item requested", "id", id)
	return model.Item{}, model.ErrNotImplemented
}

func (s *ItemService) DeleteItem(ctx context.Context, id string) error {
	// TODO: Decide whether deleting a missing item should be idempotent for this
	// API. Many APIs choose 204, some choose 404, but the choice should be
	// intentional.
	s.logger.Info("delete item requested", "id", id)
	return model.ErrNotImplemented
}

func MapItemResponse(item model.Item) model.ItemResponse {
	return model.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
	}
}

func NewItemID() string {
	// TODO: Replace this with a more deliberate ID strategy after reading about
	// UUIDs, ULIDs, sortable IDs, and how ID choice affects indexing.
	return fmt.Sprintf("item-%d", time.Now().UnixNano())
}
