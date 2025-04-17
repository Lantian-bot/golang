package main

import "fmt"

// 1.全局变量
var s1 string = "Hello"
var zero int
var b1 = true

var m map[string]int
var arr [2]byte
var slice []byte
var p *int

var (
	i  int = 123
	b2 bool
	s2 = "test"
)

var (
	group1      = 2
	group2 byte = 2
)

// 2.局部变量
func method1() {
	// 方式1，类型推导，用得最多
	a := 1
	// 方式2，完整的变量声明写法
	var b int = 2
	// 方式3，仅声明变量，但是不赋值，
	var c int
	fmt.Println(a, b, c)
}

// 方式4，直接在返回值中声明
func method2() (a int, b string) {
	// 这种方式必须声明return关键字
	// 并且同样不需要使用，并且也不用必须给这种变量赋值
	return 1, "test"
}

func method3() (a int, b string) {
	a = 1
	b = "test"
	return
}

func method4() (a int, b string) {
	return 2, "b"
}

// 3.多个变量的声明
var a, b, c int = 1, 2, 3

var e, f, g int

var h, k, j = 1, 2, "test"

func method() {
	var k, l, m int = 1, 2, 3
	var n, o, p int
	q, r, s := 1, 2, "test"
	fmt.Println(k, l, m, n, o, p, q, r, s)
}

func main() {
	// 1.全局变量
	//fmt.Println(s1)
	//fmt.Println(zero)
	//fmt.Println(b1)
	//fmt.Println(m)
	//fmt.Println(arr)
	//fmt.Println(slice)
	//fmt.Println(p)
	//fmt.Println(i)
	//fmt.Println(b2)
	//fmt.Println(s2)
	//fmt.Println(group1)
	//fmt.Println(group2)
	////m["1"] = 1
	////fmt.Println(m)
	//m = make(map[string]int, 0)
	//m["1"] = 1
	//m["2"] = 2
	//m["3"] = 3
	//fmt.Println(m)
	//slice = append(slice, 1)
	//fmt.Println(slice)
	// 2.局部变量
	//method1()
	//a1, b1 := method2()
	//a2, b2 := method3()
	//a3, b4 := method4()
	//fmt.Println(a1, b1)
	//fmt.Println(a2, b2)
	//fmt.Println(a3, b4)
	// 3.多个变量的声明
	fmt.Println(a, b, c)
	fmt.Println(e, f, g)
	fmt.Println(h, k, j)
	method()

}
