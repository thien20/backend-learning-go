// Package repository translates storage operations into domain-friendly methods.
//
// Keep repositories focused on data access. Validation and HTTP decisions
// belong elsewhere.
package repository

import (
	"context"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/model"
)

type ItemStore interface {
	GetItem(ctx context.Context, id string) (model.Item, error)
}

type ItemRepository struct {
	store ItemStore
}

func NewItemRepository(store ItemStore) *ItemRepository {
	return &ItemRepository{store: store}
}

func (r *ItemRepository) GetByID(ctx context.Context, id string) (model.Item, error) {
	item, err := r.store.GetItem(ctx, id)
	if err != nil {
		// TODO: Translate low-level storage errors into sentinel errors that the
		// service can reason about. For now the in-memory DB already returns
		// `model.ErrItemNotFound`, but a SQL implementation would not.
		return model.Item{}, err
	}

	return item, nil
}
