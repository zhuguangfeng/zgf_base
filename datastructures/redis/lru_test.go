package redis

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestLru(t *testing.T) {
	cap := 5
	lru := NewLru(cap)

	for i := 1; i <= 5; i++ {
		err := lru.Put("k"+strconv.Itoa(i), "v"+strconv.Itoa(i))
		assert.NoError(t, err)
	}

	val, err := lru.Get("2")
	assert.NoError(t, err)
	assert.Equal(t, val, "v1")

	err = lru.Put("k6", "v6")
	assert.NoError(t, err)

	t.Log(lru.cache)

	err = lru.Delete("k")
	assert.EqualError(t, err, "key nil")

}
