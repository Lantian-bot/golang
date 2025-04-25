package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:

			fmt.Printf("receive %d from channel 2\n", y)
		case z := <-ch3:

			fmt.Printf("receive %d from channel 3\n", z)
		}
	}
}

/*
这是Go语言的通道接收操作符，箭头方向表示数据流动方向
	<-ch 表示从通道 ch 接收数据
	ch <- value 表示向通道 ch 发送数据
*/
