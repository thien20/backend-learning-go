// Package cache contains cache implementations.
//
// This cache is intentionally simple. The student should treat it as a place to
// practice cache-aside and cache miss behavior, not as production-ready code.
package cache

import (
	"context"
	"sync"
	"time"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/model"
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

	// TODO: Decide whether expired entries should be treated as ordinary misses
	// or whether the caller should be told that an expiration happened.
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
	c.entries[key] = entry{
		item:      item,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()

	// TODO: Add instrumentation hooks later so cache hit rate can be measured
	// instead of guessed.
	return nil
}
