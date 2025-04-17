package channel

import (
	"fmt"
	"sync"
)

// 生产者方法：生成数据并发送到通道

func Producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // 数据发送完成后关闭通道
}

// 消费者方法：从通道接收并处理数据

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Println("接收:", num)
	}
}
