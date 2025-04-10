package main

import "fmt"

var b int = 10

func main() {
	var b int = 4
	fmt.Println("local variable, b = ", b)
	if b := 3; b == 3 {
		fmt.Println("if statement, b = ", b)
		b--
		fmt.Println("if statement, b = ", b)
	}
	fmt.Println("local variable, b = ", b)
}
