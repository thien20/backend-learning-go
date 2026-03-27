// Package service holds business rules.
//
// The service should decide:
// - what inputs are valid
// - when to use the cache
// - which errors should reach the handler
package service

import (
	"context"
	"strings"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/model"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (model.Item, error)
}

type Cache interface {
	Get(ctx context.Context, key string) (model.Item, bool, error)
	Set(ctx context.Context, key string, item model.Item) error
}

type ItemService struct {
	repo  Repository
	cache Cache
}

func NewItemService(repo Repository, cache Cache) *ItemService {
	return &ItemService{
		repo:  repo,
		cache: cache,
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

	// TODO: Decide whether cache read errors should be ignored, logged, or
	// returned. There is no single correct answer; the choice depends on how
	// critical freshness and availability are.

	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Item{}, "", err
	}

	_ = s.cache.Set(ctx, id, item)

	// TODO: Add cache invalidation rules once write operations exist.
	return item, "db", nil
}
