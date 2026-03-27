package db

import (
	"context"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/model"
)

type ItemStore interface {
	GetItem(ctx context.Context, id string) (model.Item, error)
	ListItems(ctx context.Context) ([]model.Item, error)
	ListItemIDs(ctx context.Context) ([]string, error)
	ListTagsByItemID(ctx context.Context, id string) ([]string, error)
}
