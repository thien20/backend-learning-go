package db

import (
	"context"
	"sync"
	"time"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/model"
)

type MemoryDB struct {
	mu    sync.RWMutex
	items map[string]model.Item
}

func NewMemoryDB() *MemoryDB {
	now := time.Date(2026, time.March, 1, 12, 0, 0, 0, time.UTC)
	return &MemoryDB{
		items: map[string]model.Item{
			"item-1": {
				ID:          "item-1",
				Name:        "Foundation item",
				Description: "Start here before adding writes",
				CreatedAt:   now,
				UpdatedAt:   now,
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
	return item, nil
}

func (db *MemoryDB) ListItems(ctx context.Context) ([]model.Item, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	items := make([]model.Item, 0, len(db.items))
	for _, item := range db.items {
		items = append(items, item)
	}
	return items, nil
}

func (db *MemoryDB) CreateItem(ctx context.Context, item model.Item) (model.Item, error) {
	select {
	case <-ctx.Done():
		return model.Item{}, ctx.Err()
	default:
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.items[item.ID] = item
	return item, nil
}

func (db *MemoryDB) UpdateItem(ctx context.Context, item model.Item) (model.Item, error) {
	select {
	case <-ctx.Done():
		return model.Item{}, ctx.Err()
	default:
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item.ID]; !ok {
		return model.Item{}, model.ErrItemNotFound
	}

	db.items[item.ID] = item
	return item, nil
}

func (db *MemoryDB) DeleteItem(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[id]; !ok {
		return model.ErrItemNotFound
	}

	delete(db.items, id)
	return nil
}
