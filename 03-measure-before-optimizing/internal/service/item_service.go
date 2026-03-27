package service

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/metrics"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/model"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (model.Item, error)
	List(ctx context.Context) ([]model.Item, error)
	ListDetailedSlow(ctx context.Context) ([]model.Item, error)
}

type Cache interface {
	Get(ctx context.Context, key string) (model.Item, bool, error)
	Set(ctx context.Context, key string, item model.Item) error
}

type ItemService struct {
	repo      Repository
	cache     Cache
	collector *metrics.Collector
	logger    *slog.Logger
}

func NewItemService(repo Repository, cache Cache, collector *metrics.Collector, logger *slog.Logger) *ItemService {
	return &ItemService{
		repo:      repo,
		cache:     cache,
		collector: collector,
		logger:    logger,
	}
}

func (s *ItemService) GetItem(ctx context.Context, id string) (model.Item, string, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return model.Item{}, "", model.ErrInvalidItemID
	}

	s.collector.ObserveRequest()

	cached, found, err := s.cache.Get(ctx, id)
	if err == nil && found {
		s.collector.ObserveCacheHit()
		return cached, "cache", nil
	}

	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Item{}, "", err
	}

	_ = s.cache.Set(ctx, id, item)
	return item, "db", nil
}

func (s *ItemService) ListItems(ctx context.Context, detailed bool) ([]model.ItemResponse, error) {
	s.collector.ObserveRequest()

	var (
		items []model.Item
		err   error
	)

	if detailed {
		items, err = s.repo.ListDetailedSlow(ctx)
	} else {
		items, err = s.repo.List(ctx)
	}
	if err != nil {
		return nil, err
	}

	responses := make([]model.ItemResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, model.ItemResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Tags:        item.Tags,
			Digest:      buildDigestSlowly(item),
			CreatedAt:   item.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
		})
	}

	return responses, nil
}

func buildDigestSlowly(item model.Item) string {
	// FIX-ME: This is intentionally CPU-heavy so the student has something easy
	// to find in a CPU profile. Do not "optimize" this before capturing a
	// baseline and confirming it matters.
	digest := item.Name + "|" + item.Description
	for i := 0; i < 250; i++ {
		digest = strings.ToUpper(strings.ToLower(digest))
	}
	return digest
}
