package cache

import (
	"context"
	"time"
)

// Cache is an interface abstraction for caching
type Cache interface {
	// ID returns the name of the cache storage
	ID() string
	// WithOptions applies the supplied Options, and returns the resulting Cache.
	WithOptions(opts ...Option) Cache
	// Get retrieve value at key from cache
	Get(ctx context.Context, key string, value interface{}) (err error)
	// Set stores a key with a given lifetime. 0 for permanent
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error)
	// Has checks if key is available in cache
	Has(ctx context.Context, key string) (ok bool)
	// Del removes a key by name
	Del(ctx context.Context, key string) (err error)
	// Keys list all available keys in cache
	Keys(ctx context.Context, pattern string) (available []string, err error)
	// Clear removes all keys
	Clear(ctx context.Context)
}
