package app

import (
	"log/slog"
	"os"
	"time"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/cache"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/db"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/handler"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/metrics"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/repository"
	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/service"
)

type Container struct {
	Handler *handler.ItemHandler
	Logger  *slog.Logger
}

func NewContainer() *Container {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	collector := metrics.NewCollector()
	itemDB := db.NewMemoryDB()
	itemCache := cache.NewMemoryCache(1 * time.Minute)
	itemRepository := repository.NewItemRepository(itemDB)
	itemService := service.NewItemService(itemRepository, itemCache, collector, logger)
	itemHandler := handler.NewItemHandler(itemService, collector, logger)

	return &Container{
		Handler: itemHandler,
		Logger:  logger,
	}
}
