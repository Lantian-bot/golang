package channel

//1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
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

// 在 Go 语言中，defer wg.Done() 是一个常见的并发控制模式，它的作用与 sync.WaitGroup 配合使用
/*
核心概念分解
1. sync.WaitGroup 的作用
用途：用于等待一组协程（goroutine）完成工作。
关键方法：
	Add(n)：增加等待的协程数量（计数器 +n）。
	Done()：标记一个协程完成（计数器 -1）。
	Wait()：阻塞主协程，直到计数器归零。
2. defer 的作用
用途：延迟执行一个函数调用，确保在 当前函数返回前 执行。
特点：无论函数是正常返回，还是因 panic 异常退出，defer 语句都会执行。
作用解释：
1.确保协程完成时标记完成：
	当协程函数 worker 结束时（无论正常结束还是异常退出），defer wg.Done() 会调用 wg.Done()，将 WaitGroup 的计数器减 1。
2.防止遗漏调用 Done()：
	如果协程中有多个可能的退出路径（例如多个 return 语句或可能触发 panic），直接在代码末尾写 wg.Done() 容易遗漏。
	用 defer 可以保证 无论从哪个路径退出，wg.Done() 都会被调用。

*/
