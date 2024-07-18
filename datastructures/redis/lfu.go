package redis

import (
	"container/list"
	"errors"
	"sync"
)

type Lfu struct {
	cap   int                      //容量
	cache map[string]*list.Element //键值对
	list  *list.List               //双向链表
	lock  sync.RWMutex
}

type Data struct {
	key   string
	val   any
	count int
}

func NewLfu(cap int) *Lfu {
	return &Lfu{
		cap:   cap,
		cache: make(map[string]*list.Element),
		list:  list.New(),
		lock:  sync.RWMutex{},
	}
}

func (l *Lfu) Get(key string) (any, error) {
	e, ok := l.cache[key]
	if !ok {
		return nil, errors.New("cache nil")
	}
	data := e.Value.(Data)
	data.count++
	l.UpdateList(e)
	return data.val, nil
}

func (l *Lfu) Put(key string, val any) error {
	e, ok := l.cache[key]
	if ok {
		data := e.Value.(Data)
		data.count++
		data.val = val
		l.UpdateList(e)
		return nil
	}
	if len(l.cache) >= l.cap {
		l.RemoveLeastFrequentlyUsed()
	}
	newData := Data{key: key, val: val, count: 1}
	frontE := l.list.PushBack(newData)
	l.cache[key] = frontE
	return nil
}

func (l *Lfu) UpdateList(e *list.Element) {
	for e != nil {
		//获取e的前一个元素
		prev := e.Prev()
		//如果前一个元素为空 或者 e的访问次数 小于 前面元素的 访问次数 停止循环
		if prev == nil || e.Value.(Data).count <= prev.Value.(Data).count {
			break
		}
		//将当前元素 移动到 前一个元素 之前
		l.list.MoveBefore(e, prev)
		//将e的上一个元素 赋值为当前元素 继续循环
		e = prev
	}
}

// 删除最不常用的元素
func (l *Lfu) RemoveLeastFrequentlyUsed() {
	back := l.list.Back()
	if back != nil {
		delete(l.cache, back.Value.(Data).key)
		l.list.Remove(back)
	}
}
