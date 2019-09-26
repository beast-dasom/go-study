package main

import "fmt"

func main() {
	a := map[string]int{"Hello": 10, "world": 20}

	b := map[string]int{
		"hello": 10,
		"world": 20,
	}

	fmt.Println(a)
	fmt.Println(b)
}
