package pool

import (
	"fmt"
	"sync"
)

type Pool struct {
	capacity int
	jobs     chan func()
	shutdown chan struct{}
	count    sync.WaitGroup
}

func InitPool(cap int) *Pool {
	//var cnt sync.WaitGroup
	var p = &Pool{
		capacity: cap,
		jobs:     make(chan func(), cap),
		shutdown: make(chan struct{}),
		count:    sync.WaitGroup{},
	}
	return p
}

func (p *Pool) Submit(f func()) {
	p.count.Add(1)
	p.jobs <- f
}

func (p *Pool) Run() {
	for i := 0; i < p.capacity; i++ {
		go func() {
			for {
				select {
				case job, _ := <-p.jobs:
					//if ok {
					job()
					p.count.Done()
					//}
				case <-p.shutdown:
					fmt.Println("Shutting down")
					return
				}
			}
		}()
	}

}

func (p *Pool) Stop() {
	for i := 0; i < p.capacity; i++ {
		p.shutdown <- struct{}{}
	}
	defer close(p.jobs)
}

func (p *Pool) Wait() {
	p.count.Wait()
	//for {
	//	mutex.Lock()
	//	if len(p.jobs) == 0 {
	//		close(p.jobs)
	//		fmt.Println("closed")
	//		return
	//	} else {
	//		time.Sleep(500 * time.Microsecond)
	//	}
	//	mutex.Unlock()
	//}
}
