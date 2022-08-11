package cache

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	tasert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_Default_Client(t *testing.T) {
	assert := tasert.New(t)

	// initialize a client
	client := NewNoOP()

	// get the client name
	cl := client.ID()
	assert.Equal(NOOP, cl)

	// set/modify some options for the client
	client = client.WithOptions(WithMetrics(true))

	// set some options
	client = client.WithOptions(WithMetrics(true))

	ctx := context.Background()

	Convey("Driver Test", t, func() {

		Convey("Default Driver Set Test", func() {
			for i := 0; i < 10; i++ {
				key := fmt.Sprintf("testm-%d", i)
				value := []byte(fmt.Sprintf("value-%d", i))

				err := client.Set(ctx, key, value, 0)
				assert.NoError(err)
			}
			keyb := fmt.Sprintf("test-%d", 77)
			valueb := []byte(fmt.Sprintf("value-%d", 77))
			err := client.Set(ctx, keyb, valueb, 0)
			assert.NoError(err)
		})

		Convey("Default Driver Get Test", func() {
			var rez []byte
			err := client.Get(ctx, "test-77", &rez)
			assert.NoError(err)

		})

		Convey("Has Test", func() {
			assert.False(client.Has(ctx, "testm-1"), "Key testm-1 not in cache")
			assert.False(client.Has(ctx, "hasnotintest-1"), "Key hasnotintest-1 in cache")
		})

		Convey("Delete Test", func() {
			assert.Nil(client.Del(ctx, "testm-1"), "Should not have thrown error")
			assert.Nil(client.Del(ctx, "testm-2"), "Should not have thrown error")
			assert.Nil(client.Del(ctx, "testm-3"), "Should not have thrown error")
		})

	})
}
