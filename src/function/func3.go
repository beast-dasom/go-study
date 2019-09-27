package main

import "fmt"

func main() {
	var hello func(a int, b int) int = sum
	world := sum

	fmt.Println(hello(1, 2))
	fmt.Println(world(3, 4))
}

func sum(a int, b int) int {
	return a + b
}
