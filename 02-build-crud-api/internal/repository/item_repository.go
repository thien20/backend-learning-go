package repository

import (
	"context"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/db"
	"github.com/thien/backend-learning-go/02-build-crud-api/internal/model"
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

func (r *ItemRepository) Create(ctx context.Context, item model.Item) (model.Item, error) {
	// TODO: Add storage-specific translation once a real DB is introduced.
	return r.store.CreateItem(ctx, item)
}

func (r *ItemRepository) Update(ctx context.Context, item model.Item) (model.Item, error) {
	return r.store.UpdateItem(ctx, item)
}

func (r *ItemRepository) Delete(ctx context.Context, id string) error {
	return r.store.DeleteItem(ctx, id)
}
