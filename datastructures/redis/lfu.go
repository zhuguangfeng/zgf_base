package redis

import (
	"container/list"
	"errors"
)

type Lfu struct {
	cap      int                      //容量
	cache    map[string]*list.Element //缓存的键值对
	freqList *list.List               //双向链表 按频率排序存储缓存项
}

type Entry struct {
	key   string
	val   any
	count int //访问次数
}

func NewLfu(cap int) *Lfu {
	return &Lfu{
		cap:      cap,
		cache:    make(map[string]*list.Element),
		freqList: list.New(),
	}
}

func (l *Lfu) Get(key string) (any, error) {
	v, ok := l.cache[key]
	if !ok {
		return nil, errors.New("cache nil")
	}
	entry := v.Value.(*Entry)
	entry.count++
	l.UpdateFreqList(v)
	return entry.val, nil
}

func (l *Lfu) Put(key string, val any) {
	e, ok := l.cache[key]
	if ok {
		entry := e.Value.(*Entry)
		entry.val = val
		entry.count++
		l.UpdateFreqList(e)
	} else {
		if len(l.cache) >= l.cap {
			l.RemoveLeastFrequentlyUsed()
		}
		newEntry := Entry{key: key, val: val, count: 1}
		e := l.freqList.PushFront(&newEntry)
		l.cache[key] = e
	}
}

func (l *Lfu) UpdateFreqList(e *list.Element) {
	for e != nil {
		prev := e.Prev()
		if prev == nil || e.Value.(*Entry).count <= prev.Value.(*Entry).count {
			break // 退出循环，因为e已经是第一个元素或频率不比前一个元素高
		}
		l.freqList.MoveBefore(e, prev)
		e = prev
	}
}

func (l *Lfu) RemoveLeastFrequentlyUsed() {
	last := l.freqList.Back()
	if last != nil {
		delete(l.cache, last.Value.(*Entry).key)
		l.freqList.Remove(last)
	}
}
