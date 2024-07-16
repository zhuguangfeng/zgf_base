package redis

import (
	"container/list"
	"testing"
)

func TestLru(t *testing.T) {
	cap := 2
	l := list.New()
	m := make(map[string]*list.Element, cap)
	lru := NewLru(cap, l, m)

	lru.Set("k1", "v1")
	lru.Set("k2", "v2")
	t.Log(lru.storage)
	lru.Set("k3", "v3")
	t.Log(lru.storage["k3"])
	t.Log(lru.storage["k1"])

}
