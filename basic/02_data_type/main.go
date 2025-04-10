package main

import "fmt"

// 1.整型
// 十六进制
var a uint8 = 0xF
var b uint8 = 0xf

// 八进制
var c uint8 = 017
var d uint8 = 0o17
var e uint8 = 0o17

// 二进制
var f uint8 = 0b1111
var g uint8 = 0b1111

// 十进制
var h uint8 = 15

// 2.浮点型
var float1 float32 = 10
var float2 = 10.0

func main() {
	//println(a, b, c, d, e, f, g, h)
	////println(float1==float2)
	//println(float1 == float32(float2))
	////复数
	//var c1 complex64
	//c1 = 1.10 + 0.1i
	//c2 := 1.10 + 0.1i
	//c3 := complex(1.10, 0.1)
	//// c2与c3是等价的
	//fmt.Println(c1 == complex64(c2))
	//fmt.Println(c1 == complex64(c3))
	//fmt.Println(complex128(c1) == c2) // 虚实部有点区别 不等
	////实部和虚部
	//x := real(c2)
	//y := imag(c2)
	//fmt.Println(x)
	//fmt.Println(y)
	// 3.byte类型
	//var s string = "Hello, world!"
	//var bytes []byte = []byte(s)
	//fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	//fmt.Println(string(bytes) == s)
	//4.rune 符文类型
	//var r1 rune = 'a'
	//var r2 rune = '世'
	//fmt.Println(r1, r2)
	//var s string = "abc，你好，世界！"
	//var runes []rune = []rune(s)
	//fmt.Println(runes)
	//fmt.Println(len(runes))
	//5.字符串 - string
	//	var s1 string = "Hello\nworld!\n"
	//	var s2 string = `Hello
	//world!
	//`
	//	fmt.Println(s1 == s2)
	//6.byte、rune 与 string 之间的联系
	//var s string = "Go语言"
	//var bytes []byte = []byte(s)
	//var runes []rune = []rune(s)
	//fmt.Println("string length: ", len(s))
	//fmt.Println("bytes length: ", len(bytes))
	//fmt.Println("runes length: ", len(runes))
	//fmt.Println("string sub: ", s[0:7])
	//fmt.Println("bytes sub: ", string(bytes[0:7]))
	//fmt.Println("runes sub: ", string(runes[0:3]))
	//零值
	var digit int
	var s1 string
	var b bool
	fmt.Println(digit)
	fmt.Println(s1)
	fmt.Println(b)

}
