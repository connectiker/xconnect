package cache

import (
	"sync"
)

// Config represents the options for the cache.
type Config struct {
	lock *sync.RWMutex

	withMetrics bool
	stats       Statistics
}

type Option func(config *Config)

// NewConfig create a cache config
func NewConfig(opts ...Option) *Config {
	var c Config
	c.lock = new(sync.RWMutex)

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

//SetOptions allows you to set additional options
func (c *Config) SetOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *Config) Metrics() bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.withMetrics
}

func WithMetrics(activate bool) Option {
	return func(c *Config) {
		c.lock.Lock()
		c.withMetrics = activate

		if c.stats == nil {
			c.stats = NewMetrics()
		}
		c.lock.Unlock()
	}
}
