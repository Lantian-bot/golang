package main

//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ：原子操作、并发数据安全。

import (
	"fmt"
	"sync"
)

// Counter 结构体包含一个共享的计数器和互斥锁

type AtomicCounter struct {
	val int
	mu  sync.Mutex
}

// Increment 方法使用互斥锁保护计数器递增操作

func (c *AtomicCounter) AtomicIncrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func main() {
	var wg sync.WaitGroup
	counter := &AtomicCounter{} // 创建计数器实例

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程执行1000次递增操作
			for j := 0; j < 1000; j++ {
				counter.AtomicIncrement()
			}
		}()
	}
	// 等待所有协程完成
	wg.Wait()
	fmt.Println("Final counter value:", counter.val) // 输出最终结果
}
