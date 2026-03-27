package cache

import (
	"context"
	"sync"
	"time"

	"github.com/thien/backend-learning-go/02-build-crud-api/internal/model"
)

type entry struct {
	item      model.Item
	expiresAt time.Time
}

type MemoryCache struct {
	mu      sync.RWMutex
	ttl     time.Duration
	entries map[string]entry
}

func NewMemoryCache(ttl time.Duration) *MemoryCache {
	return &MemoryCache{
		ttl:     ttl,
		entries: make(map[string]entry),
	}
}

func (c *MemoryCache) Get(ctx context.Context, key string) (model.Item, bool, error) {
	select {
	case <-ctx.Done():
		return model.Item{}, false, ctx.Err()
	default:
	}

	c.mu.RLock()
	cached, ok := c.entries[key]
	c.mu.RUnlock()
	if !ok {
		return model.Item{}, false, nil
	}

	if time.Now().After(cached.expiresAt) {
		c.mu.Lock()
		delete(c.entries, key)
		c.mu.Unlock()
		return model.Item{}, false, nil
	}

	return cached.item, true, nil
}

func (c *MemoryCache) Set(ctx context.Context, key string, item model.Item) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if c.ttl <= 0 {
		return nil
	}

	c.mu.Lock()
	c.entries[key] = entry{item: item, expiresAt: time.Now().Add(c.ttl)}
	c.mu.Unlock()
	return nil
}

func (c *MemoryCache) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	c.mu.Lock()
	delete(c.entries, key)
	c.mu.Unlock()
	return nil
}
