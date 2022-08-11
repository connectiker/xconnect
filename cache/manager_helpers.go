package cache

import (
	"context"
	"time"
)

var manager = New()

// Register driver to manager instance
func Register(name string, driver Cache) *Manager {
	manager.UseDefault(name)
	manager.Register(name, driver)
	return manager
}

// DefaultUse set default driver name
func DefaultUse(driverName string) {
	manager.UseDefault(driverName)
}

// Use driver object by name and set it as default driver.
func Use(driverName string) Cache {
	return manager.Use(driverName)
}

// Driver returns a cache instance by its name
func Driver(driverName string) Cache {
	return manager.Driver(driverName)
}

// DefManager get default cache manager instance
func DefManager() *Manager {
	return manager
}

// Default get default cache driver instance
func Default() Cache {
	return manager.Default()
}

// Has checks if key is available in cache
func Has(ctx context.Context, key string) (ok bool) {
	return manager.Default().Has(ctx, key)
}

// Get retrieve value at key from cache
func Get(ctx context.Context, key string, value interface{}) (err error) {
	return manager.Default().Get(ctx, key, value)
}

// Set stores a key with a given lifetime. 0 for permanent
func Set(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error) {
	return manager.Default().Set(ctx, key, value, ttl)
}

// Del remove a key by name
func Del(ctx context.Context, key string) (err error) {
	return manager.Default().Del(ctx, key)
}

// Keys list all available cache keys
func Keys(ctx context.Context, pattern string) (available []string, err error) {
	return manager.Default().Keys(ctx, pattern)
}

// Clear removes all keys and closes the client
func Clear(ctx context.Context) {
	manager.Default().Clear(ctx)
}
