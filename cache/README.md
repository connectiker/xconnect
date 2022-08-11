# Cache

Cache interface with cache manager & cache statistics. The cache manager will manage all the cache instances.

The idea is to have multiple caching drivers and use them when and how you want.
For this to happen there is a generic interface next to which you can register all types of cache drivers and this will allow you to add any new custom cache driver. You just register the new cache driver inside the cache manager and use the cache interface.

The cache library has also a manager that allows you to register as many cache drivers you want, define a default one and also allows you to choose any of the caches you've registered.

## Cache Interface

All cache drivers implement the cache.Cache interface. So, You can add any custom driver.

```
// Cache interface definition
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
```
## Usage example

Please check the unit tests for exact usage