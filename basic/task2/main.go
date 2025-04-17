package main

import (
	"basic/task2/goroutine"
	"basic/task2/pointer"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1.题目
	x := 5
	fmt.Println("修改前的值:", x) // 打印初始值
	pointer.Add(&x)          // 传递x的地址给函数
	fmt.Println("修改后的值:", x) // 显示增加10后的结果
	// 2.题目
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", slice)
	pointer.DoubleSlice(&slice)
	fmt.Println("修改后的切片:", slice)
	// 3.题目
	var wg sync.WaitGroup
	wg.Add(2)
	go goroutine.PrintEven(&wg)
	go goroutine.PrintOdd(&wg)
	wg.Wait()
	// 4.题目
	var tasks = []func(){
		func() {
			fmt.Println("task1")
			time.Sleep(2 * time.Second)
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("task2")
		},
		func() {
			fmt.Println("task3")
			time.Sleep(1 * time.Second)
		},
	}
	goroutine.Scheduler(tasks)

}
