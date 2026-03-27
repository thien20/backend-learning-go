// Package db holds storage implementations.
//
// This stage uses an in-memory store so the student can focus on layering
// before learning a real database driver.
package db

import (
	"context"
	"sync"
	"time"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/model"
)

type MemoryDB struct {
	mu    sync.RWMutex
	items map[string]model.Item
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		items: map[string]model.Item{
			"item-1": {
				ID:        "item-1",
				Name:      "First teaching item",
				UpdatedAt: time.Date(2026, time.March, 1, 12, 0, 0, 0, time.UTC),
			},
		},
	}
}

func (db *MemoryDB) GetItem(ctx context.Context, id string) (model.Item, error) {
	select {
	case <-ctx.Done():
		return model.Item{}, ctx.Err()
	default:
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	item, ok := db.items[id]
	if !ok {
		return model.Item{}, model.ErrItemNotFound
	}

	// TODO: Add more seeded records and a write path once the basic read flow is
	// comfortable. That is where the mutex choices become easier to reason about.
	return item, nil
}
