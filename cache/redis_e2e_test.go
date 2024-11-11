//go:build e2e

package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedisCache_e2e_Set(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	c := NewRedisCache(rdb)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := c.Set(ctx, "key1", "val1", time.Minute)
	assert.NoError(t, err)

	val, err := c.Get(ctx, "key1")
	assert.NoError(t, err)
	assert.Equal(t, "val1", val)
}

func TestRedisCache_e2e_SetV1(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	testCases := []struct {
		name string
		//before func()
		after func(t *testing.T)

		key        string
		val        any
		expiration time.Duration

		wantErr error
	}{
		{
			name: "set value",
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				defer cancel()

				res, err := rdb.Get(ctx, "key1").Result()
				require.NoError(t, err)
				assert.Equal(t, "val1", res)

				_, err = rdb.Del(ctx, "key1").Result()

				require.NoError(t, err)
			},
			key:        "key1",
			val:        "val1",
			expiration: time.Minute,
			wantErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewRedisCache(rdb)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()
			err := c.Set(ctx, tc.key, tc.val, tc.expiration)
			assert.NoError(t, err)
			tc.after(t)
		})
	}
}
