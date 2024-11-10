package channel

import "sync"

type Task func()
type TaskPool struct {
	tasks chan Task
	//close *atomic.Bool

	close chan struct{}

	closeOnce sync.Once
}

// numG是goroutine的数量，就是你要控制住的
// cap是缓存的容量
func NewTaskPool(numG int, cap int) *TaskPool {
	res := &TaskPool{
		tasks: make(chan Task, cap),
		//close: atomic.NewBool(false),

		close: make(chan struct{}),
	}

	for i := 0; i < numG; i++ {
		go func() {

			//for t := range res.tasks {
			//	if res.close.Load() {
			//		return
			//	}
			//
			//	t()
			//}

			for {
				select {
				case <-res.close:
					return
				case t := <-res.tasks:
					t()
				}
			}
		}()
	}

	return res
}

// 提交任务
func (p *TaskPool) Submit(t Task) {
	p.tasks <- t
}

func (p *TaskPool) Close() error {
	//p.close.Store(true)

	//这种实现有个缺陷 重复调用Close方法 会panic
	close(p.close)

	//p.closeOnce.Do(func() {
	//	close(p.close)
	//})
	return nil
}

// Do执行任务
func (p *TaskPool) Do(t Task) {

}
