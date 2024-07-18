package redis

import (
	"container/list"
	"errors"
	"sync"
)

//所有设置了过期时间的键值中，删除最久未使用的

type Lru struct {
	cap   int                      //容量
	cache map[string]*list.Element //键值对数据
	list  *list.List               //链表
	lock  sync.RWMutex             //读写锁
}
type data struct {
	key string
	val any
}

func NewLru(cap int) *Lru {
	return &Lru{
		cap:   cap,
		cache: make(map[string]*list.Element),
		list:  list.New(),
		lock:  sync.RWMutex{},
	}
}

func (l *Lru) Get(key string) (any, error) {
	//加读锁
	l.lock.RLock()
	defer l.lock.RUnlock()
	//获取该key的val是否存在
	val, ok := l.cache[key]
	if !ok {
		return nil, errors.New("cache nil")
	}
	//将获取到的元素移动到表头
	l.list.MoveToFront(val)
	return val.Value.(data).val, nil
}

func (l *Lru) Put(key string, val any) error {
	//加读锁
	l.lock.RLock()
	//获取k v
	v, ok := l.cache[key]
	l.lock.RUnlock()

	//加写锁
	l.lock.Lock()
	defer l.lock.Unlock()
	if ok {
		//存在 将该元素移动到表头
		l.list.MoveToFront(v)
		//更新对应的val值
		v.Value.(*data).val = val
		return nil
	}

	//不存在且 大于了最大容量
	if len(l.cache) >= l.cap {
		//获取表尾的元素
		backE := l.list.Back()
		//删除链表和键值对
		delete(l.cache, backE.Value.(data).key)
		l.list.Remove(backE)
		//容量-1
	}
	//将新增加的元素推到表头
	item := data{key: key, val: val}
	e := l.list.PushFront(item)
	l.cache[key] = e
	return nil
}

func (l *Lru) Delete(key string) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	val, ok := l.cache[key]
	if !ok {
		return errors.New("key nil")
	}
	l.list.Remove(val)
	delete(l.cache, val.Value.(data).key)
	return nil
}
