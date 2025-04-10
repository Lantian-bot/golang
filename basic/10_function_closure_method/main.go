package main

import "fmt"

type A struct {
	i int
}

// 定义方法
func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var function func(int) int

func main() {
	a := A{1}
	function = a.add

	fmt.Println("after call function, a.i = ", function(1))
	fmt.Println("a.i = ", a.i)
}
