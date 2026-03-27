package db

import (
	"context"
	"sync"
	"time"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/model"
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
				Name:        "Measured item one",
				Description: "Useful for smoke tests",
				Tags:        []string{"baseline", "teaching"},
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			"item-2": {
				ID:          "item-2",
				Name:        "Measured item two",
				Description: "Useful for list endpoints",
				Tags:        []string{"load", "profiling"},
				CreatedAt:   now.Add(1 * time.Minute),
				UpdatedAt:   now.Add(1 * time.Minute),
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

func (db *MemoryDB) ListItemIDs(ctx context.Context) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	ids := make([]string, 0, len(db.items))
	for id := range db.items {
		ids = append(ids, id)
	}
	return ids, nil
}

func (db *MemoryDB) ListTagsByItemID(ctx context.Context, id string) ([]string, error) {
	item, err := db.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}
	return append([]string(nil), item.Tags...), nil
}
