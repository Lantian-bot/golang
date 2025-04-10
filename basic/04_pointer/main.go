package main

import (
	"fmt"
	"unsafe"
)

var i int
var p1 *int = &i
var gs string = "Hello"
var p2 *string = &gs

//var p3 *string ="aa"

func main() {
	//i := 1
	//s := "Hello"
	//// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	//// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	//p1 = &i
	//p2 = &s
	//// var p3 **string = &p2
	//p3 := &p2
	//// 零值
	//var p4 *byte
	//fmt.Println(p1)
	//fmt.Println(p2)
	//fmt.Println(p3)
	//fmt.Println(p4)
	//fmt.Println(&p4)
	//
	////2.使用指针访问值
	//fmt.Println(*p1)
	//fmt.Println(*p2)
	//fmt.Println(**p3)
	////fmt.Println(*p4) // invalid memory address or nil pointer dereference

	// 3.修改指针指向的值
	//a := 2
	//var p *int
	////fmt.Println(&a) // 0xc00000a108
	//p = &a
	////fmt.Println(p, &a) // 0xc00000a108 0xc00000a108
	//
	//var pp **int
	//pp = &p
	//fmt.Println(pp, p) // 0xc000074068 0xc00000a108
	//**pp = 3
	//fmt.Println(pp, *pp, p) // 0xc000074068 0xc00000a108 0xc00000a108
	//fmt.Println(**pp, *p)   // 3 3
	//fmt.Println(a, &a)      // 3 0xc00000a108

	//4.指针、unsafe.Pointer 和 uintptr
	a := "Hello, world!"
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
	*c = 12
	fmt.Println(*c)

}
