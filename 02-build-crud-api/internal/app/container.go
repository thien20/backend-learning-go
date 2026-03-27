package app

import (
	"log/slog"
	"os"
	"time"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/cache"
	"github.com/thien/backend-learning-go/02-build-crud-api/internal/db"
	"github.com/thien/backend-learning-go/02-build-crud-api/internal/handler"
	"github.com/thien/backend-learning-go/02-build-crud-api/internal/repository"
	"github.com/thien/backend-learning-go/02-build-crud-api/internal/service"
)

type Container struct {
	Handler *handler.ItemHandler
	Logger  *slog.Logger
}

func NewContainer() *Container {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	itemDB := db.NewMemoryDB()
	itemCache := cache.NewMemoryCache(1 * time.Minute)
	itemRepository := repository.NewItemRepository(itemDB)
	itemService := service.NewItemService(itemRepository, itemCache, logger)
	itemHandler := handler.NewItemHandler(itemService, logger)

	// TODO: Move addresses, TTLs, and feature flags into configuration once the
	// student is comfortable with the dependency graph.
	return &Container{
		Handler: itemHandler,
		Logger:  logger,
	}
}
