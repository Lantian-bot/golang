package goroutine

// 1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行

import (
	"fmt"
	"sync"
	"time"
)

// 奇数

func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// 偶数

func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// 2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

func Scheduler(tasks []func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Now()
			fmt.Println(task, "执行时间:", end.Sub(start))
		}(task)
	}
	wg.Wait()
}
