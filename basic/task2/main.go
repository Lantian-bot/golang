package main

import (
	"basic/task2/channel"
	"basic/task2/goroutine"
	"basic/task2/object"
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
	// 5.题目
	// 创建矩形和圆形实例
	rect := object.Rectangle{Width: 5, Height: 3}
	circle := object.Circle{Radius: 4}
	// 使用格式化输出保留两位小数
	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())
	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())
	// 6.题目
	// 创建员工实例并初始化
	newEmployee := object.Employee{
		Person: object.Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "1001",
	}
	// 调用方法输出信息
	newEmployee.PrintInfo()
	// 7.题目
	// 初始化通道和等待组
	ch := make(chan int)
	// 启动生产者协程
	wg.Add(1)
	go channel.Producer(ch, &wg) // 注意传递WaitGroup指针

	// 启动消费者协程
	wg.Add(1)
	go channel.Consumer(ch, &wg)
	// 等待所有协程完成
	wg.Wait()
}
