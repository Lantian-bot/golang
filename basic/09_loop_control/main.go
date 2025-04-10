package main

import "fmt"

func main() {
	gotoPreset := false

preset:
	a := 5

process:
	if a > 0 {
		a--
		fmt.Println("当前a的值为：", a)
		goto process
	} else if a <= 0 {
		// elseProcess:
		if !gotoPreset {
			gotoPreset = true
			goto preset
		} else {
			goto post
		}
	}

post:
	fmt.Println("main将结束，当前a的值为：", a)
}
