// Package app wires concrete dependencies together.
//
// Keep the container boring. Its job is dependency injection, not business
// logic. If this file starts growing decision branches, some behavior
// probably belongs in a lower layer.
package app

import (
	"time"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/cache"
	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/db"
	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/handler"
	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/repository"
	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/service"
)

type Container struct {
	Handler *handler.ItemHandler
}

func NewContainer() *Container {
	itemDB := db.NewMemoryDB()
	itemCache := cache.NewMemoryCache(30 * time.Second)
	itemRepository := repository.NewItemRepository(itemDB)
	itemService := service.NewItemService(itemRepository, itemCache)
	itemHandler := handler.NewItemHandler(itemService)

	// TODO: Make the lifetime of each dependency explicit once configuration is
	// introduced. For example: how long should cache entries live, and where
	// should that value come from?
	return &Container{
		Handler: itemHandler,
	}
}
