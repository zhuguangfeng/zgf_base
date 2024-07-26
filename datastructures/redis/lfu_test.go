package redis

import (
	"testing"
)

//["LFUCache","put","put","get","put","get","get","put","get","get","get"]
//[[2],        [1,1],[2,2],[1],[3,3],  [2],  [3],  [4,4],[1],  [3]   ,[4]]

// 1.2 3.3 4.2
func TestLft(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	a1 := lfu.Get(1)
	t.Log(a1)
	lfu.Put(3, 3)
	t.Log(lfu.cache)
	a2 := lfu.Get(2)
	t.Log(a2)
	a3 := lfu.Get(3)
	t.Log(a3)
	lfu.Put(4, 4)
	t.Log(lfu.cache)
	a4 := lfu.Get(1)
	t.Log(a4)
	a5 := lfu.Get(3)
	t.Log(a5)
	a6 := lfu.Get(4)
	t.Log(a6)

}
