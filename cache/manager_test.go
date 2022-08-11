package cache

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	tasert "github.com/stretchr/testify/assert"
	"testing"
)

func TestManager(t *testing.T) {
	assert := tasert.New(t)

	// initialize a client
	client := NewNoOP()

	// set some options
	client = client.WithOptions(WithMetrics(true))

	// register one(or more) cache drivers in the manager
	Register(NOOP, client)

	// select a default cache driver to use.
	DefaultUse(NOOP)

	// switch to a new cache driver to use.
	Use(NOOP)

	//get the cache instance by its type
	ch := Driver(NOOP)
	assert.NotNil(ch)

	ctx := context.Background()

	Convey("Manager Test", t, func() {

		Convey("Default Driver Set Test", func() {
			for i := 0; i < 10; i++ {
				key := fmt.Sprintf("testm-%d", i)
				value := []byte(fmt.Sprintf("value-%d", i))

				err := Set(ctx, key, value, 0)
				assert.NoError(err)
			}
			keyb := fmt.Sprintf("test-%d", 77)
			valueb := []byte(fmt.Sprintf("value-%d", 77))
			err := Set(ctx, keyb, valueb, 0)
			assert.NoError(err)

			err = manager.Set(ctx, keyb, valueb, 0)
			assert.NoError(err)
		})

		Convey("Default Driver Get Test", func() {
			var rez []byte
			err := Get(ctx, "test-77", &rez)
			assert.NoError(err)

			err = manager.Get(ctx, "test-22", &rez)
			assert.NoError(err)

		})

		Convey("Has Test", func() {
			assert.False(Has(ctx, "testm-1"), "Key testm-1 not in cache")
			assert.False(manager.Has(ctx, "hasnotintest-1"), "Key hasnotintest-1 in cache")
		})

		Convey("Delete Test", func() {
			assert.Nil(Del(ctx, "testm-1"), "Should not have thrown error")
			assert.Nil(manager.Del(ctx, "testm-2"), "Should not have thrown error")
		})

		Convey("Keys Test", func() {
			_, err := Keys(ctx, "testm-1")
			assert.Nil(err, "Should not have thrown error")
		})

		Convey("Clear Test", func() {
			Clear(ctx)
		})

	})
}
