package redis

import (
	"container/list"
	"errors"
	"sync"
)

//所有设置了过期时间的键值中，删除最久未使用的

type Lru struct {
	Size     int                      //当前容量
	Cap      int                      //最大容量
	list     *list.List               //链表
	cacheMap map[string]*list.Element //缓存数据
	lock     sync.RWMutex
}

// 键值对
type Data struct {
	Key string
	Val any
}

func NewLru(cap int, list *list.List, cacheMap map[string]*list.Element) *Lru {
	return &Lru{
		Size:     0,
		Cap:      cap,
		list:     list,
		cacheMap: cacheMap,
		lock:     sync.RWMutex{},
	}
}

func (l *Lru) Get(key string) (any, error) {
	//加读锁
	l.lock.RLock()
	val, ok := l.cacheMap[key]
	l.lock.RUnlock()
	if ok {
		//把获取的元素移动到链表头部
		l.lock.Lock()
		l.list.MoveToFront(val)
		l.lock.Unlock()
		return val.Value.(Data).Val, nil
	}
	return nil, errors.New("key nil")
}

func (l *Lru) Set(key string, val any) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	v, ok := l.cacheMap[key]
	if ok {
		l.list.MoveToFront(v)
		v.Value.(*Data).Val = val
	}
	item := Data{Key: key, Val: val}

	if l.Size >= l.Cap {
		e := l.list.Back()
		l.list.Remove(e)
		delKey := e.Value.(Data).Key
		delete(l.cacheMap, delKey)
		l.Size--
	}
	l.list.PushFront(item)
	ne := l.list.Front()
	l.cacheMap[key] = ne
	l.Size++
	return nil
}

func (l *Lru) Delete(key string) error {
	l.lock.Unlock()
	defer l.lock.Unlock()
	v, ok := l.cacheMap[key]
	if ok {
		l.list.Remove(v)
		k := v.Value.(Data).Key
		delete(l.cacheMap, k)
		l.Size--
		return nil
	}
	return errors.New("key nil")
}
