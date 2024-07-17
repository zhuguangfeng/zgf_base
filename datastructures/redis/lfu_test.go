package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLft(t *testing.T) {
	lfu := NewLfu(2)
	lfu.Put("k1", "v1")
	lfu.Put("k2", "v2")

	v1, err1 := lfu.Get("k1")
	assert.NoError(t, err1)
	t.Log(v1)
	v2, err2 := lfu.Get("k2")
	assert.NoError(t, err2)
	t.Log(v2)
	v3, err3 := lfu.Get("k1")
	assert.NoError(t, err3)
	t.Log(v3)

	lfu.Put("k3", "v3")
	t.Log(lfu.cache["k3"].Value)
	t.Log(lfu.cache)
}
