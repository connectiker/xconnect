package cache

import (
	tasert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_Options(t *testing.T) {
	assert := tasert.New(t)

	// initialize a config
	cfg := NewConfig(WithMetrics(true))
	assert.NotNil(cfg)

	cfg.SetOptions(WithMetrics(true))

	st := cfg.Metrics()
	assert.True(st)
}
