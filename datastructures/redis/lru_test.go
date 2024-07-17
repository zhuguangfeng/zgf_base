package redis

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLru(t *testing.T) {
	maxCap := 2
	l := list.New()
	m := make(map[string]*list.Element, maxCap)

	lru := NewLru(maxCap, l, m)

	err := lru.Set("k1", "v1")
	assert.NoError(t, err)
	err2 := lru.Set("k2", "v2")
	assert.NoError(t, err2)
	err3 := lru.Set("k3", "v3")
	assert.NoError(t, err3)
	a, b := lru.Get("k1")
	fmt.Println(a, b)
}
