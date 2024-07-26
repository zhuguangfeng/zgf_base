package redis

import (
	"container/list"
	"time"
)

type LFUCache struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

func Constructor(cap int) *LFUCache {
	return &LFUCache{
		cap:   cap,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

type Data struct {
	key   int
	val   int
	count int
	time  int64 //Timestamp
}

func (this *LFUCache) Get(key int) int {
	e, ok := this.cache[key]
	if !ok {
		return -1
	}
	data := e.Value.(*Data)
	data.count++
	data.time = time.Now().UnixNano()
	this.UpdateList(e)
	this.cache[key] = e
	return data.val
}

func (this *LFUCache) UpdateList(e *list.Element) {
	for e != nil {
		prev := e.Prev()
		if prev == nil || e.Value.(*Data).count < prev.Value.(*Data).count || (e.Value.(*Data).count == prev.Value.(*Data).count && e.Value.(*Data).time < prev.Value.(*Data).time) {
			break
		}
		//if e.Value.(*Data).count == prev.Value.(*Data).count {
		//	// 如果访问次数相同，则根据时间戳比较，时间戳晚的排在前面
		//	if e.Value.(*Data).time > prev.Value.(*Data).time {
		//		break // 当前元素的时间戳更晚，不需要移动
		//	}
		//}
		// 移动当前元素到前一个元素的位置
		this.list.MoveBefore(e, prev)
		e = prev // 更新为前一个元素，继续比较
	}
}

func (this *LFUCache) Put(key int, val int) {
	e, ok := this.cache[key]
	if ok {
		e.Value.(*Data).count++
		e.Value.(*Data).val = val
		e.Value.(*Data).time = time.Now().UnixNano()
		this.UpdateList(e)
		return
	}
	if len(this.cache) >= this.cap {
		this.RemoveLeastFrequentlyUsed()
	}
	item := &Data{key: key, val: val, count: 1, time: time.Now().UnixNano()}
	backE := this.list.PushBack(item)
	this.cache[key] = backE
	return
}

func (this *LFUCache) RemoveLeastFrequentlyUsed() {
	back := this.list.Back()
	if back != nil {
		delete(this.cache, back.Value.(*Data).key)
		this.list.Remove(back)

	}
}

//type Lfu struct {
//	cap   int                      //容量
//	cache map[string]*list.Element //键值对
//	list  *list.List               //双向链表
//	lock  sync.RWMutex
//}
//
//type Data struct {
//	key   string
//	val   any
//	count int
//}
//
//func NewLfu(cap int) *Lfu {
//	return &Lfu{
//		cap:   cap,
//		cache: make(map[string]*list.Element),
//		list:  list.New(),
//		lock:  sync.RWMutex{},
//	}
//}
//
//func (l *Lfu) Get(key string) (any, error) {
//	e, ok := l.cache[key]
//	if !ok {
//		return nil, errors.New("cache nil")
//	}
//	data := e.Value.(Data)
//	data.count++
//	l.UpdateList(e)
//	return data.val, nil
//}
//
//func (l *Lfu) Put(key string, val any) error {
//	e, ok := l.cache[key]
//	if ok {
//		data := e.Value.(Data)
//		data.count++
//		data.val = val
//		l.UpdateList(e)
//		return nil
//	}
//	if len(l.cache) >= l.cap {
//		l.RemoveLeastFrequentlyUsed()
//	}
//	newData := Data{key: key, val: val, count: 1}
//	frontE := l.list.PushBack(newData)
//	l.cache[key] = frontE
//	return nil
//}
//
//func (l *Lfu) UpdateList(e *list.Element) {
//	for e != nil {
//		//获取e的前一个元素
//		prev := e.Prev()
//		//如果前一个元素为空 或者 e的访问次数 小于 前面元素的 访问次数 停止循环
//		if prev == nil || e.Value.(Data).count <= prev.Value.(Data).count {
//			break
//		}
//		//将当前元素 移动到 前一个元素 之前
//		l.list.MoveBefore(e, prev)
//		//将e的上一个元素 赋值为当前元素 继续循环
//		e = prev
//	}
//}
//
//// 删除最不常用的元素
//func (l *Lfu) RemoveLeastFrequentlyUsed() {
//	back := l.list.Back()
//	if back != nil {
//		delete(l.cache, back.Value.(Data).key)
//		l.list.Remove(back)
//	}
//}
