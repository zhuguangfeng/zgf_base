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
		delKey := e.Value.(*Data).Key
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

//type Lru struct {
//	size     int                      //当前的容量
//	capCache int                      //限定的容量
//	list     *list.List               //列表
//	storage  map[string]*list.Element //存储的key
//}
//
//func NewLru(capCache int, list *list.List, storage map[string]*list.Element) Lru {
//	return Lru{
//		capCache: capCache,
//		list:     list,
//		storage:  storage,
//	}
//}
//
//type Data struct {
//	key string
//	val any
//}
//
//func (l *Lru) Get(key string) any {
//	v, ok := l.storage[key]
//	if ok {
//		l.list.MoveToFront(v)
//		return v.Value.(Data).val
//	}
//	return nil
//}
//
//func (l *Lru) Set(key string, val string) error {
//	e, ok := l.storage[key]
//	if ok {
//		n := e.Value.(Data)
//		n.val = val
//		e.Value = n
//		l.list.MoveToFront(e)
//		return nil
//	}
//
//	if l.size >= l.capCache {
//		//删除链表尾部元素
//		e = l.list.Back()
//		dk := e.Value.(Data).key
//		l.list.Remove(e)
//		delete(l.storage, dk)
//		l.size--
//	}
//
//	n := Data{key: key, val: val}
//	l.list.PushFront(n)
//	ne := l.list.Front()
//	l.storage[key] = ne
//	l.size++
//	return nil
//}
//
//func (l *Lru) Delete(key string) error {
//	v, ok := l.storage[key]
//	if ok {
//		l.list.Remove(v)
//		k := v.Value.(Data).key
//		delete(l.storage, k)
//	}
//	return errors.New("key nil")
//}
