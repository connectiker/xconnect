package cache

import (
	"context"
	"time"
)

const NOOP = "noop"

type NoOp struct{}

func NewNoOP() *NoOp                                                               { return &NoOp{} }
func (n *NoOp) ID() string                                                         { return NOOP }
func (n *NoOp) WithOptions(opts ...Option) Cache                                   { return n }
func (n *NoOp) Get(ctx context.Context, key string, value interface{}) (err error) { return nil }
func (n *NoOp) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error) {
	return nil
}
func (n *NoOp) Has(ctx context.Context, key string) (ok bool)   { return }
func (n *NoOp) Del(ctx context.Context, key string) (err error) { return nil }
func (n *NoOp) Keys(ctx context.Context, pattern string) (available []string, err error) {
	return nil, err
}
func (n *NoOp) Clear(ctx context.Context) {}
