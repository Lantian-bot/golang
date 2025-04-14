package main

import "fmt"

func main() {
	hash := map[string]int{
		"aba":     1,
		"fe":      2,
		"zi":      3,
		"counter": 4,
	}

	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}
