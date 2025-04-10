package main

import "fmt"

type Gender byte

func (g *Gender) isMale() bool {
	return *g == Male
}
func (g *Gender) isFemale() bool {
	return *g == Female
}

const (
	Male Gender = iota
	Female
)

func main() {
	var gender = Female
	fmt.Println(gender)
	fmt.Println(gender.isMale())
	fmt.Println(gender.isFemale())
}
