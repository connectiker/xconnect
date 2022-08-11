package cache

import (
	tasert "github.com/stretchr/testify/assert"
	"testing"
)

func TestStats(t *testing.T) {
	assert := tasert.New(t)

	// initialize a stats instance
	instance := NewMetrics()
	assert.NotNil(instance)

	instance.SetQuery(true)
	instance.SetQuery(false)

	instance.SetMutation(true)
	instance.SetMutation(false)

	instance.SetSize(23)
	instance.SetItems(500)

	stats := instance.GetStats()

	totalItems := stats.Items()
	assert.Equal(int64(500), totalItems)

	size := stats.Size()
	assert.Equal(int64(23), size)

	qHits := stats.QueryHits()
	assert.Equal(int64(1), qHits)

	qMiss := stats.QueryMiss()
	assert.Equal(int64(1), qMiss)

	mHits := stats.MutationHits()
	assert.Equal(int64(1), mHits)

	mMiss := stats.MutationMiss()
	assert.Equal(int64(1), mMiss)

	instance.IncStats(true, true)
	h := stats.QueryHits()
	assert.Equal(int64(2), h)
}
