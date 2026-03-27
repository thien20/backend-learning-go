package db

import (
	"context"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/model"
)

// ItemStore keeps the DB contract small and swappable.
//
// A future Postgres implementation should be able to satisfy this interface
// without forcing handler or service changes.
type ItemStore interface {
	GetItem(ctx context.Context, id string) (model.Item, error)
	ListItems(ctx context.Context) ([]model.Item, error)
	CreateItem(ctx context.Context, item model.Item) (model.Item, error)
	UpdateItem(ctx context.Context, item model.Item) (model.Item, error)
	DeleteItem(ctx context.Context, id string) error
}
