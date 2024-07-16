package redis

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLru(t *testing.T) {
	maxCap := 3
	l := list.New()
	m := make(map[string]*list.Element, maxCap)

	lru := NewLru(maxCap, l, m)

	err := lru.Set("k1", "v1")
	assert.NoError(t, err)
	a, b := lru.Get("k1")
	fmt.Println(a, b)
}
