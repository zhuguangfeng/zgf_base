package cache

import (
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBuildInMapCache_Get(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		cache   func() *BuildInMapCache
		wantVal any
		wantErr error
	}{
		{
			name: "key not found",
			key:  "not exist key",
			cache: func() *BuildInMapCache {
				return NewBuildInMapCache(10 * time.Second)
			},
			wantErr: fmt.Errorf("%w,key:%s", errKeyNotFound, "not exist key"),
		},
		{
			name: "get key",
			key:  "key1",
			cache: func() *BuildInMapCache {
				res := NewBuildInMapCache(10 * time.Second)
				err := res.Set(context.Background(), "key1", "val1", time.Minute)
				require.NoError(t, err)
				return res
			},
			wantVal: "val1",
		},

		{
			name: "expired",
			key:  "expired key",
			cache: func() *BuildInMapCache {
				res := NewBuildInMapCache(10 * time.Second)
				err := res.Set(context.Background(), "expired key", "val1", time.Second)
				require.NoError(t, err)
				time.Sleep(time.Second * 2)
				return res
			},
			wantErr: fmt.Errorf("%w,key:%s", errKeyNotFound, "expired key"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.cache().Get(context.Background(), tc.key)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestBuildInMapCache_Loop(t *testing.T) {
	cnt := 0
	c := NewBuildInMapCache(time.Second, BuildInMapCacheWithEvictedCallback(func(key string, val any) {
		cnt++
	}))
	err := c.Set(context.Background(), "key1", 123, time.Second)
	require.NoError(t, err)
	time.Sleep(time.Second * 3)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.data["key1"]
	require.False(t, ok)
	require.Equal(t, 1, cnt)
}
