package repository

import (
	"context"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/db"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/model"
)

type ItemRepository struct {
	store db.ItemStore
}

func NewItemRepository(store db.ItemStore) *ItemRepository {
	return &ItemRepository{store: store}
}

func (r *ItemRepository) GetByID(ctx context.Context, id string) (model.Item, error) {
	return r.store.GetItem(ctx, id)
}

func (r *ItemRepository) List(ctx context.Context) ([]model.Item, error) {
	return r.store.ListItems(ctx)
}

func (r *ItemRepository) ListDetailedSlow(ctx context.Context) ([]model.Item, error) {
	ids, err := r.store.ListItemIDs(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]model.Item, 0, len(ids))
	for _, id := range ids {
		item, err := r.store.GetItem(ctx, id)
		if err != nil {
			return nil, err
		}

		tags, err := r.store.ListTagsByItemID(ctx, id)
		if err != nil {
			return nil, err
		}

		item.Tags = tags
		items = append(items, item)
	}

	// FIX-ME: This is intentionally N+1-shaped. On a real DB, fetching IDs and
	// then reloading each item and its tags separately becomes expensive quickly.
	// Measure it before rewriting it.
	return items, nil
}
