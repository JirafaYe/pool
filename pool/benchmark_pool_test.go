package pool

import (
	"sync"
	"testing"
)

func BenchmarkFunction(b *testing.B) {
	p := InitPool(13000)
	p.Run()
	// 进行需要测试的操作
	for i := 0; i < b.N; i++ {
		//num := i
		p.Submit(func() {
			//fmt.Println("task", num)
		})
		// 执行重复运行的代码
		// 这里可以是你想要测试性能的函数或操作
	}
	p.Wait()
}

func BenchmarkRawFunction(b *testing.B) {
	group := sync.WaitGroup{}
	//p := InitPool(10)
	//p.Run()
	// 进行需要测试的操作
	for i := 0; i < b.N; i++ {
		//num := i
		go func() {
			group.Add(1)
			defer group.Done()
			//fmt.Println("task", num)
		}()
		// 执行重复运行的代码
		// 这里可以是你想要测试性能的函数或操作
	}
	group.Wait()
}
