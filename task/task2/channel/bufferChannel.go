package channel

//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制。

import (
	"fmt"
	"sync"
)

// 生产者函数：向通道发送100个整数

func Bufferproducer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 确保协程结束前标记完成
	defer close(ch) // 确保最后关闭通道

	for i := 1; i <= 100; i++ {
		ch <- i                   // 发送数据到缓冲通道
		fmt.Printf("生产: %d\n", i) // 可选生产日志
	}
}

// 消费者函数：从通道接收并打印整数

func Bufferconsumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // 自动检测通道关闭
		fmt.Printf("消费: %d\n", num)
	}
}
